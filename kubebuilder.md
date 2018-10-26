# Using Kubebuilder

Code generated with [Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder).

### Boilerplate Code Header

```console
echo "// Amazon VPC CNI Custom Network Controller." > hack/boilerplate.go.txt
```

### Create Project

```console
$ kubebuilder init --license none --dep
```

### Add Node Controller

```console
$ kubebuilder create api --resource=false --controller --group core --version v1 --kind Node
```
