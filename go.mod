module github.com/crowdstrike/falcon-operator

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.9.2
	github.com/aws/aws-sdk-go-v2/config v1.8.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.5.0
	github.com/containerd/cgroups v1.0.2 // indirect
	github.com/containers/image/v5 v5.19.0
	github.com/crowdstrike/gofalcon v0.2.16
	github.com/go-logr/logr v0.4.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.16.0
	github.com/openshift/api v0.0.0-20210428205234-a8389931bee7
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	k8s.io/api v0.20.6
	k8s.io/apimachinery v0.20.6
	k8s.io/client-go v0.20.6
	sigs.k8s.io/controller-runtime v0.8.3
)

replace github.com/opencontainers/image-spec => github.com/opencontainers/image-spec v1.0.2-0.20210819154149-5ad6f50d6283
