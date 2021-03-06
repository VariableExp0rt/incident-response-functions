package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/sns"
	//"github.com/VariableExp0rt/incident-response-functions/remove-all-attached-permissions/openapi"
)

const (
	reason            = "IAM Users should not disable logging tooling unless assuming a role that explicitly allows them to do so"
	mitreTactic       = "Defense Evasion"
	mitreTechnique    = "Impair Defences"
	mitreSubTechnique = "T1562.008"
)

var (
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigStateFromEnv,
	}))
)

//CloudTrailCloudWatchEventDetails is a custom type for unmarshalling the event.Details *json.RawMessage
//of a specified event, this is not at all great, as CloudTrail also doesn't specify a struct for this
//will need to look at this for specific events I'm forwarding.
//we don't need many of the fields for this simple example, however we'd need
//to create a struct to unmarshal all fields needed (requestParameters, responseElements)
//if used - as long as you have an example event https://mholt.github.io/json-to-go/
type CloudTrailCloudWatchEventDetails struct {
	EventVersion string    `json:"eventVersion"`
	EventID      string    `json:"eventID"`
	EventTime    time.Time `json:"eventTime"`
	EventType    string    `json:"eventType"`
	AwsRegion    string    `json:"awsRegion"`
	EventName    string    `json:"eventName"`
	UserIdentity struct {
		UserName    string `json:"userName"`
		PrincipalID string `json:"principalId"`
		AccessKeyID string `json:"accessKeyId"`
		InvokedBy   string `json:"invokedBy"`
		Type        string `json:"type"`
		Arn         string `json:"arn"`
		AccountID   string `json:"accountId"`
	} `json:"userIdentity"`
	EventSource string `json:"eventSource"`
}

//ThreatDetails is a custom type to serialise JSON for publishing to a SNS topic
//for alerting of detected "suspicious" actions
type ThreatDetails struct {
	EventTime               time.Time            `json:"eventTime"`
	Details                 Details              `json:"details"`
	AwsAPICall              string               `json:"AwsAPICall"`
	ServicesAccessedSummary map[string]time.Time `json:"servicesAccessedSummary"`
}

//Details is a sub-structure containing the details of the alert,
//such as the Mitre category that this type of activity falls into
type Details struct {
	Reason            string `json:"reason"`
	MitreTactic       string `json:"mitreTactic"`
	MitreTechnique    string `json:"mitreTechnique"`
	MitreSubTechnique string `json:"mitreSubTechnique"`
}

//LambdaHandler is the function executed when the Lambda receives an event from
//CloudWatch Events via an EventBridge rule
func LambdaHandler(ctx context.Context, event events.CloudWatchEvent) {

	var cTrailEvent CloudTrailCloudWatchEventDetails
	if err := json.Unmarshal(event.Detail, &cTrailEvent); err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	//I've managed to find a CloudTrail schema in the EventBridge Schema Registry
	//which can be used as a binding for a CloudTrail event type (with all the bells and
	//whistles of additional methods)
	/*
		var cEvent openapi.AWSAPICallViaCloudTrail
		_ = json.Unmarshal(event.Detail, &cEvent)
	*/

	//Establish service client and pass session information, i.e. use credentials
	//of role attached to the Lambda function
	svc := iam.New(sess)

	//** CHANGE BELOW TO ASSIGN EXPLICIT DENY FOR ALL SERVICES RATHER THAN MESSING WITH LISTING + DETACHING

	//List the policies attached to the user that attempted the API call
	//as defined by the CWE rule
	response, err := svc.ListAttachedUserPolicies(&iam.ListAttachedUserPoliciesInput{
		UserName: aws.String(cTrailEvent.UserIdentity.UserName),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Printf("Error: %s", awsErr.Error())
				return
			case iam.ErrCodeInvalidInputException:
				fmt.Printf("Error: %s", awsErr.Error())
				return
			}

		}
	}

	//Range of the AttachedPolicies object and for each policy and remove all associated
	//permissions from the IAM user
	for _, v := range response.AttachedPolicies {

		//Prepare the request
		req, resp := svc.DetachUserPolicyRequest(&iam.DetachUserPolicyInput{
			PolicyArn: v.PolicyArn,
			UserName:  aws.String(cTrailEvent.UserIdentity.UserName),
		})

		//Send the request
		err := req.Send()
		if err != nil {
			fmt.Printf("Error: %s\nResponse: %s", err.(awserr.Error).OrigErr(), resp.String())
		}
	}

	/*
		section here for grabbing accessadvisor findings for recent roles this user has assumed
		and their corresponding activity reports
	*/

	resp2, err := svc.GenerateServiceLastAccessedDetails(&iam.GenerateServiceLastAccessedDetailsInput{
		Arn:         aws.String("arn:aws:iam::" + cTrailEvent.UserIdentity.AccountID + ":user/" + cTrailEvent.UserIdentity.UserName),
		Granularity: aws.String("SERVICE_LEVEL"),
	})
	if err != nil {
		fmt.Printf("Error creating last used services report: %v", err)
	}

	resp3, err := svc.GetServiceLastAccessedDetails(&iam.GetServiceLastAccessedDetailsInput{
		JobId: resp2.JobId,
	})
	if err != nil {
		fmt.Printf("Error accessing service last used report: %v", err)
	}

	var serviceList map[string]time.Time
	for _, service := range resp3.ServicesLastAccessed {

		serviceList[*service.ServiceName] = *service.LastAuthenticated
	}

	//
	j := &ThreatDetails{
		EventTime: event.Time,
		Details: Details{
			Reason:            reason,
			MitreTactic:       mitreTactic,
			MitreTechnique:    mitreTechnique,
			MitreSubTechnique: mitreSubTechnique},
		AwsAPICall:              cTrailEvent.EventName,
		ServicesAccessedSummary: serviceList,
	}

	mj, err := json.Marshal(j)
	if err != nil {
		fmt.Printf("Unable to marshal struct to JSON format: %v", err)
	}

	svc2 := sns.New(sess)

	//Prepare the JSON object for publishing to the Simple Notification Service topic
	//The design here requires that the topic ARN is included as an environment variable
	//within your function
	resp4, err := svc2.Publish(&sns.PublishInput{
		Message:          aws.String(string(mj)),
		MessageStructure: aws.String("json"),
		TopicArn:         aws.String(os.Getenv("SNS_TOPIC_ARN")),
	})
	if err != nil {
		fmt.Printf("Error: %s", err.(awserr.Error).Message())
		return
	}

	//Print details of the published message to stdout so we can track this
	//if the alert is not delivered
	fmt.Printf("MessageId: %s, SequenceNumber: %s", *resp4.MessageId, *resp4.SequenceNumber)

}

func main() {
	lambda.Start(LambdaHandler)
}
