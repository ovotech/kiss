# kiss 😘

> Kaluza Infrastructure Secret Service

AWS-based secrets management for Kubernetes.

Leverages users' Kubernetes OIDC authentication tokens for AWS Secrets Manager secrets management. Can also manage AWS IAM policies for secrets if you're using the [AWS provider for the k8s Secrets Store CSI driver](https://aws.amazon.com/blogs/security/how-to-use-aws-secrets-configuration-provider-with-kubernetes-secrets-store-csi-driver/).

You may be interested in this project if you:
* use AWS Secrets Manager as your k8s secrets store
* use OIDC tokens for auth/z'ing users against your k8s cluster
* want your users to manage their own secrets without access to your AWS account

## Why?

We're using [`external-secrets`](https://github.com/external-secrets/kubernetes-external-secrets) to synchronize AWS Secret Manager secrets in our k8s AWS account with k8s-native `Secret` resources.

We wanted to let our cluster users manage secrets logically scoped to their namespaces without hooking them up with direct access to the AWS account.

Since we were already using OIDC tokens for user auth/z to the cluster with `kubectl` (using [`kubelogin`](https://github.com/int128/kubelogin)), we figured we could use those same tokens for auth/z against an intermediary secrets management service that simply wrapped around AWS Secrets Manager.

## Synergy with `external-secrets`

Our service creates AWS Secrets Manager secrets with the following naming convention:

**Breaking Change**
*As of version 0.1.0, the secret naming convention has been updated*
```
k8s-secret-secret-namespace-secret-name
```

We then have `external-secrets` annotations on our namespaces to only allow `ExternalSecret` resources in that namespace to sync with AWS Secrets Manager secrets logically scoped to the namespace. For example, this shows a namespace called `security` that is only allowed to read AWS Secrets Manager secrets prefixed by `k8s-secret_security_`:

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: security
  annotations:
    externalsecrets.kubernetes-client.io/permitted-key-name: "k8s-secret_security-.*"
```

A member of the `security` namespace can then create a secret with `kiss` and use it as such:

```yaml
apiVersion: kubernetes-client.io/v1
kind: ExternalSecret
metadata:
  name: foo
  namespace: security
spec:
  backendType: secretsManager
  data:
    - key: k8s-secret_security_foo
      name: foo
```

## Synergy with AWS provider for k8s Secrets Store CSI driver

If you're using the [AWS Secrets & Configuration Provider with your Kubernetes Secrets Store CSI driver](https://aws.amazon.com/blogs/security/how-to-use-aws-secrets-configuration-provider-with-kubernetes-secrets-store-csi-driver/) you can use `kiss` to:
* automatically create AWS IAM policies with read permissions when creating secrets
* attach AWS IAM policies to IAM Roles for Service Accounts, allowing the relevant `ServiceAccount` to read your secret

Combined with our [`iam-service-account-controller`](https://github.com/ovotech/iam-service-account-controller), this can allow your users to safely manage IAM roles and secret reading policies directly from their k8s namespaces, with no AWS permissions.

Check the `-policy` flags and the `bind` command for more details.

## Builds

We release binaries for the client. Please check the releases.

We don't provide any public images for the server, but it should be trivial to build your own. Check our `./.github/workflows/build-release.yaml` if you're lost.

## Tokens

Your OIDC tokens will vary depending on your setup. Kaluza's payloads look like this:

```json
{
  "cognito:groups": ["kaluza:default", "kaluza:security"],
  "iss": "https://cognito-idp.eu-west-1.amazonaws.com/eu-west-1_AbCdEf",
  "email": "user.name@kaluza.com",
  "...": "..."
}
```

The `cognito:groups` is a list of namespaces the token grants access to, prefixed with our organization name. We'll need to tell the server how to extract these things.

## Running locally

### Server

```shell
$ cd server

$ go run cmd/main.go
	-jwks-url="https://cognito-idp.eu-west-1.amazonaws.com/eu-west-1_AbCdEf/.well-known/jwks.json"
	-namespaces-key="cognito:groups"
	-namespaces-regex="kaluza:([1-9a-z-]{1,63})"
	-identifier-key="email"
	-token-path=""
```

Note we're telling the server where and how to extract the namespaces from the token as well as which field contains the user identifier, for auditing purposes. We also need to pass the relevant JWKS file to validate user tokens.

The `-token-path` is a path to a Web ID token for _AWS auth/z_ (not to be confused with our users' tokens!). When running locally, we can set an empty `-token-path` to use default AWS authentication instead. Naturally, this means your local environment should be configured for AWS access with necessary permissions.

Check the `-help` for more information.

### Client

The `ping` command is a convenient way to test the connection and user auth/z against the server.

```shell
$ cd client

$ go run cmd/main.go ping
	-token-path="~/.kube/cache/oidc-login/c19ad9a81c5044b90e02259679c9f8037acdb23d970de3425f9377a1a7242da7"
	-namespace="security"
	-server-addr="localhost:10000"
	-secure=false
Successfully sent ping


$ go run cmd/main.go ping
	-token-path="~/.kube/cache/oidc-login/c19ad9a81c5044b90e02259679c9f8037acdb23d970de3425f9377a1a7242da7"
	-namespace="kube-system"
	-server-addr="localhost:10000"
	-secure=false
2021/07/02 13:29:37 [ERROR] Error ocurred while sending ping: rpc error: code = PermissionDenied desc = user 'user.name@kaluza.com' is not authorized for namespace 'kube-system'
exit status 1
```

Other commands include `create`, `create-from`, `update`, `delete` and `list` which do pretty much what you'd expect with AWS Secret Manager secrets logically scoped to the user's k8s namespace. Check the relevant `-help` for more information.


### Creating a secret from a file

Sometimes it may be necessary to create a secret which contains newlines, to support this, you can use the `create-from` command, specifying the filepath in the `-value` argument.

**Notes**
* The `-token-path` is optional for all commands. Without it `kiss` will attempt to find the first file in the default `~/.kube/cache/oidc-login` directory
* On a Mac you can't use the `~/` path when specifying the token path - instead use `$HOME/`
* When you create a secret it will be base64 encoded, so you should specify it's value in plaintext
* A secret name cannot contain underscores as this will result in the "Error occurred while creating secret" error
* The `-interactive` flag can be used to help choose which token to load if there are multiple in your ~/.kube/cache/oidc-login` directory.
