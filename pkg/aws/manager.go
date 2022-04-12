package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	awsiam "github.com/aws/aws-sdk-go-v2/service/iam"
	awssm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	awssts "github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	managedByTagKey   = "security.kaluza.com/managed-by"
	managedByTagValue = "kiss"
	namespaceTagKey   = "security.kaluza.com/secret-namespace"
	nameTagKey        = "security.kaluza.com/secret-name"
)

type Manager struct {
	smclient     *awssm.Client
	iamclient    *awsiam.Client
	rolePrefix   string
	secretPrefix string
	accountId    string
	ctx          context.Context
}

func NewManagerWithDefaultConfig(
	rolePrefix string,
	secretPrefix string,
	region string,
	statsdTraceUrl string,
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
		smclient:     awssm.NewFromConfig(cfg),
		iamclient:    awsiam.NewFromConfig(cfg),
		rolePrefix:   rolePrefix,
		secretPrefix: secretPrefix,
		accountId:    *callerIdentity.Account,
		ctx:          ctx,
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
	iamClient := awsiam.New(awsiam.Options{Region: region, Credentials: appCreds})

	return &Manager{
		smclient:     smClient,
		iamclient:    iamClient,
		rolePrefix:   rolePrefix,
		secretPrefix: secretPrefix,
		accountId:    accountId,
		ctx:          ctx,
	}
}
