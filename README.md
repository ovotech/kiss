# kiss 😘

> Kaluza Infrastructure Secret Service

AWS-based secrets management for Kubernetes.

Leverages users' Kubernetes OIDC authentication tokens for AWS Secrets Manager secrets management.

## Why?

We're using [`external-secrets`](https://github.com/external-secrets/kubernetes-external-secrets) to synchronize AWS Secret Manager secrets in our k8s AWS account with k8s-native `Secret` resources.

We wanted to let our cluster users manage secrets logically scoped to their namespaces without hooking them up with direct access to the AWS account.

Since we were already using OIDC tokens for user auth/z to the cluster with `kubectl` (using [`kubelogin`](https://github.com/int128/kubelogin)), we figured we could use those same tokens for auth/z against an intermediary secrets management service that simply wrapped around AWS Secrets Manager.

## Synergy with `external-secrets`

Our service creates AWS Secrets Manager secrets with the following naming convention:

```
k8s-secret_secret-namespace_secret-name
```

We then have `external-secrets` annotations on our namespaces to only allow `ExternalSecret` resources in that namespace to sync with AWS Secrets Manager secrets logically scoped to the namespace. For example:

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: security
  annotations:
    externalsecrets.kubernetes-client.io/permitted-key-name: "k8s-secret_security_.*"
```

For example, a member of the `security` namespace can create a secret with our service and then use it as such:

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

Other commands include `create`, `update`, `delete` and `list` which do pretty much what you'd expect with AWS Secret Manager secrets logically scoped to the user's k8s namespace. Check the relevant `-help` for more information.
