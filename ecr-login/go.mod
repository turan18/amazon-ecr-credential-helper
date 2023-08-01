module github.com/awslabs/amazon-ecr-credential-helper/ecr-login

require (
	github.com/aws/aws-sdk-go-v2 v1.20.0
	github.com/aws/aws-sdk-go-v2/config v1.18.25
	github.com/aws/aws-sdk-go-v2/credentials v1.13.24
	github.com/aws/aws-sdk-go-v2/service/ecr v1.18.11
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.17.0
	github.com/aws/smithy-go v1.14.0
	github.com/docker/docker-credential-helpers v0.7.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sirupsen/logrus v1.9.2
	github.com/stretchr/testify v1.8.2
)

go 1.13
