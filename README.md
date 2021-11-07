# Writing CRD by mixing kubeuilder and code generator

- [source](https://www.fatalerrors.org/a/writing-crd-by-mixing-kubeuilder-and-code-generator.html)

## Install Kubebuilder

```bash
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)

chmod +x kubebuilder && mv kubebuilder ~/bin/
```

## Kubebuilder

### Initialize project

```bash
MODULE=example.com/foo-controller
go mod init $MODULE
kubebuilder init --domain example.com
kubebuilder edit --multigroup=true
```

### Generate Resources and Manifests

```bash
kubebuilder create api --group webapp --version v1 --kind Guestbook

make manifests
```


## Code-generator

### Prepare script

### Download code-generator

check version in `go.mod`, for example `k8s.io/client-go v0.22.1`

```bash
K8S_VERSION=v0.22.1
go get k8s.io/code-generator@$K8S_VERSION
go mod vendor
```

### Update dependant version

seems unnecessary

### Generating code


