# kiss 😘

> Kaluza Infrastructure Secret Service

AWS-based secrets management for Kubernetes.

Leverages users' Kubernetes OIDC authentication tokens for AWS Secrets Manager secrets management and binding secrets to AWS IAM roles for k8s ServiceAccounts.

Synergizes well with:

- [iam-service-account-controller](https://github.com/ovotech/iam-service-account-controller)
- [IAM roles for service accounts](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html)
- [AWS Secrets & Configuration Provider](https://aws.amazon.com/about-aws/whats-new/2021/04/aws-secrets-manager-delivers-provider-kubernetes-secrets-store-csi-driver/)
- [OIDC identity provider authentication for Amazon EKS](https://aws.amazon.com/blogs/containers/introducing-oidc-identity-provider-authentication-amazon-eks/)

and mostly useless without them.
