package aws

import (
	"github.com/buff91/test/configs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

func createConfig() {
	var arn string
	var externalID string

	config := configs.GetConfig()
	conf := aws.Config{Region: aws.String(config.Con)}
	sess := session.Must(session.NewSession())


	if arn != "" {
		var creds *credentials.Credentials
		if externalID != "" {
			creds = stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {
				p.ExternalID = &externalID
			})
		} else {
			creds = stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {

			})
		}
		conf.Credentials = creds
	}
	return conf
}