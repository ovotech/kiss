package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	awssm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	awssts "github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	clusterTagKey   = "role.k8s.aws/cluster"
	managedByTagKey = "role.k8s.aws/managed-by"
	stackTagKey     = "serviceaccount.k8s.aws/stack"
)

type Manager struct {
	client      *awssm.Client
	rolePrefix  string
	secretPrefx string
	accountId   string
	ctx         context.Context
}

func NewManagerWithDefaultConfig(
	rolePrefix string,
	secretPrefix string,
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
		client:      awssm.NewFromConfig(cfg),
		rolePrefix:  rolePrefix,
		secretPrefx: secretPrefix,
		accountId:   *callerIdentity.Account,
		ctx:         ctx,
	}
}

func NewManagerWithWebIdToken(
	rolePrefix string,
	secretPrefix string,
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
	smClient := awssm.New(awssm.Options{Region: region, Credentials: appCreds})

	return &Manager{
		client:      smClient,
		rolePrefix:  rolePrefix,
		secretPrefx: secretPrefix,
		accountId:   accountId,
		ctx:         ctx,
	}
}
