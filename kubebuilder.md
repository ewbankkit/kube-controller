# Using Kubebuilder

Code generated with [Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder).

### Boilerplate Code Header

```console
echo "// AWS Target Group Controller." > hack/boilerplate.go.txt
```

### Create Project

```console
$ kubebuilder init --license none --dep
dep ensure
Running make...
make
go generate ./pkg/... ./cmd/...
go fmt ./pkg/... ./cmd/...
go vet ./pkg/... ./cmd/...
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all
CRD manifests generated under '/home/kit/wrk/src/github.com/ewbankkit/kube-controller/config/crds'
RBAC manifests generated under '/home/kit/wrk/src/github.com/ewbankkit/kube-controller/config/rbac'
go test ./pkg/... ./cmd/... -coverprofile cover.out
?   	github.com/ewbankkit/kube-controller/pkg/apis	[no test files]
?   	github.com/ewbankkit/kube-controller/pkg/controller	[no test files]
?   	github.com/ewbankkit/kube-controller/cmd/manager	[no test files]
go build -o bin/manager github.com/ewbankkit/kube-controller/cmd/manager
Next: Define a resource with:
$ kubebuilder create api
```
