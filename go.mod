module github.com/banzaicloud/logging-operator

go 1.12

require (
	emperror.dev/errors v0.4.2
	github.com/MakeNowJust/heredoc v0.0.0-20171113091838-e9091a26100e
	github.com/andreyvit/diff v0.0.0-20170406064948-c7f18ee00883
	github.com/banzaicloud/k8s-objectmatcher v1.0.0
	github.com/coreos/prometheus-operator v0.33.0
	github.com/go-logr/logr v0.1.0
	github.com/goph/emperror v0.17.2
	github.com/gordonklaus/ineffassign v0.0.0-20190601041439-ed7b1b5ee0f8 // indirect
	github.com/iancoleman/orderedmap v0.0.0-20190318233801-ac98e3ecb4b0
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/pborman/uuid v1.2.0
	github.com/sergi/go-diff v1.0.0 // indirect
	golang.org/x/lint v0.0.0-20190909230951-414d861bb4ac // indirect
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	k8s.io/api v0.0.0-20190813020757-36bff7324fb7
	k8s.io/apimachinery v0.0.0-20190809020650-423f5d784010
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.2.0
	sigs.k8s.io/controller-tools v0.2.1 // indirect
)

replace (
	github.com/prometheus/prometheus => github.com/prometheus/prometheus v1.8.2-0.20190818123050-43acd0e2e93f
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	// required for test deps only
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190409022649-727a075fdec8
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
