package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	awsiam "github.com/aws/aws-sdk-go-v2/service/iam"
	awssts "github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	clusterTagKey   = "role.k8s.aws/cluster"
	managedByTagKey = "role.k8s.aws/managed-by"
	stackTagKey     = "serviceaccount.k8s.aws/stack"
)

type Manager struct {
	client     *awsiam.Client
	rolePrefix string
	accountId  string
	ctx        context.Context
}

func NewManagerWithDefaultConfig(
	rolePrefix string,
	region string,
) *Manager {
	ctx := context.Background()

	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(region))
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config: %v", err)
	}

	stsClient := awssts.NewFromConfig(cfg)
	callerIdentity, err := stsClient.GetCallerIdentity(
		ctx,
		&awssts.GetCallerIdentityInput{},
	)
	if err != nil {
		log.Fatalf("Unable to get account identifer from AWS STS: %v", err)
	}

	return &Manager{
		client:    awsiam.NewFromConfig(cfg),
		accountId: *callerIdentity.Account,
		ctx:       ctx,
	}
}

func NewManagerWithWebIdToken(
	rolePrefix string,
	region string,
	roleARN string,
	tokenPath string,
) *Manager {
	ctx := context.Background()

	// get creds
	stsClient := awssts.New(awssts.Options{Region: region})
	appCreds := aws.NewCredentialsCache(
		stscreds.NewWebIdentityRoleProvider(
			stsClient,
			roleARN,
			stscreds.IdentityTokenFile(tokenPath),
			func(o *stscreds.WebIdentityRoleOptions) {
				o.RoleSessionName = "kaluza-infrastructure-secret-service"
			},
		),
	)

	// get account id for manager
	acctSTSClient := awssts.New(awssts.Options{Region: region, Credentials: appCreds})
	callerIdentity, err := acctSTSClient.GetCallerIdentity(
		ctx,
		&awssts.GetCallerIdentityInput{},
	)
	if err != nil {
		log.Fatalf("Unable to get account identifer from AWS STS: %v", err)
	}
	accountId := *callerIdentity.Account

	// get iam client for manager
	iamClient := awsiam.New(awsiam.Options{Region: region, Credentials: appCreds})

	return &Manager{
		client:    iamClient,
		accountId: accountId,
		ctx:       ctx,
	}
}
