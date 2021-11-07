# Writing CRD by mixing kubeuilder and code generator

- [source](https://www.fatalerrors.org/a/writing-crd-by-mixing-kubeuilder-and-code-generator.html)

## Install Kubebuilder

```bash
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)

chmod +x kubebuilder && mv kubebuilder ~/bin/
```

## Initialize project

```bash
MODULE=example.com/foo-controller
go mod init $MODULE
kubebuilder init --domain example.com
kubebuilder edit --multigroup=true
```

## Generate Resources and Manifests

```bash
kubebuilder create api --group webapp --version v1 --kind Guestbook
```

