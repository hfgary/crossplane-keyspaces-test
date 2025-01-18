package awsclient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/keyspaces"
)

// ListKeyspaces fetches the list of Keyspaces
func ListKeyspaces() ([]*keyspaces.KeyspaceSummary, error) {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // specify your region here
	}))
	keyspacesSvc := keyspaces.New(awsSession)

	result, err := keyspacesSvc.ListKeyspaces(&keyspaces.ListKeyspacesInput{})
	if err != nil {
		return nil, err
	}

	return result.Keyspaces, nil
}
