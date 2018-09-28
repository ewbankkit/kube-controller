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

### Add Service Controller

```console
$ kubebuilder create api --resource=false --controller --group core --version v1 --kind Service
Writing scaffold for you to edit...
pkg/controller/service/service_controller.go
pkg/controller/service/service_controller_test.go
Running make...
go generate ./pkg/... ./cmd/...
go fmt ./pkg/... ./cmd/...
go vet ./pkg/... ./cmd/...
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all
CRD manifests generated under '/home/kit/wrk/src/github.com/ewbankkit/kube-controller/config/crds'
RBAC manifests generated under '/home/kit/wrk/src/github.com/ewbankkit/kube-controller/config/rbac'
go test ./pkg/... ./cmd/... -coverprofile cover.out
?   	github.com/ewbankkit/kube-controller/pkg/apis	[no test files]
?   	github.com/ewbankkit/kube-controller/pkg/controller	[no test files]
2018/09/27 18:08:52 fork/exec /usr/local/kubebuilder/bin/etcd: no such file or directory
FAIL	github.com/ewbankkit/kube-controller/pkg/controller/service	0.006s
?   	github.com/ewbankkit/kube-controller/cmd/manager	[no test files]
Makefile:9: recipe for target 'test' failed
make: *** [test] Error 1
2018/09/27 18:08:52 exit status 2
```

### Add Node Controller

```console
$ kubebuilder create api --resource=false --controller --group core --version v1 --kind Node
...
```
