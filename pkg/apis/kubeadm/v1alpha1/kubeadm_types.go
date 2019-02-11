/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterConfiguration contains cluster-wide configuration for a kubeadm cluster
type ClusterConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// Etcd holds configuration for etcd.
	Etcd Etcd `json:"etcd"`

	// Networking holds configuration for the networking topology of the cluster.
	Networking Networking `json:"networking"`

	// KubernetesVersion is the target version of the control plane.
	KubernetesVersion string `json:"kubernetesVersion"`

	// ControlPlaneEndpoint sets a stable IP address or DNS name for the control plane; it
	// can be a valid IP address or a RFC-1123 DNS subdomain, both with optional TCP port.
	// In case the ControlPlaneEndpoint is not specified, the AdvertiseAddress + BindPort
	// are used; in case the ControlPlaneEndpoint is specified but without a TCP port,
	// the BindPort is used.
	// Possible usages are:
	// e.g. In a cluster with more than one control plane instances, this field should be
	// assigned the address of the external load balancer in front of the
	// control plane instances.
	// e.g.  in environments with enforced node recycling, the ControlPlaneEndpoint
	// could be used for assigning a stable DNS to the control plane.
	ControlPlaneEndpoint string `json:"controlPlaneEndpoint"`

	// APIServer contains extra settings for the API server control plane component
	APIServer APIServer `json:"apiServer,omitempty"`

	// ControllerManager contains extra settings for the controller manager control plane component
	ControllerManager ControlPlaneComponent `json:"controllerManager,omitempty"`

	// Scheduler contains extra settings for the scheduler control plane component
	Scheduler ControlPlaneComponent `json:"scheduler,omitempty"`

	// DNS defines the options for the DNS add-on installed in the cluster.
	DNS DNS `json:"dns"`

	// CertificatesDir specifies where to store or look for all required certificates.
	CertificatesDir string `json:"certificatesDir"`

	// ImageRepository sets the container registry to pull images from.
	// If empty, `k8s.gcr.io` will be used by default; in case of kubernetes version is a CI build (kubernetes version starts with `ci/` or `ci-cross/`)
	// `gcr.io/kubernetes-ci-images` will be used as a default for control plane components and for kube-proxy, while `k8s.gcr.io`
	// will be used for all the other images.
	ImageRepository string `json:"imageRepository"`

	// UseHyperKubeImage controls if hyperkube should be used for Kubernetes components instead of their respective separate images
	UseHyperKubeImage bool `json:"useHyperKubeImage,omitempty"`

	// FeatureGates enabled by the user.
	FeatureGates map[string]bool `json:"featureGates,omitempty"`

	// The cluster name
	ClusterName string `json:"clusterName,omitempty"`
}

// ControlPlaneComponent holds settings common to control plane component of the cluster
type ControlPlaneComponent struct {
	// ExtraArgs is an extra set of flags to pass to the control plane component.
	// TODO: This is temporary and ideally we would like to switch all components to
	// use ComponentConfig + ConfigMaps.
	ExtraArgs map[string]string `json:"extraArgs,omitempty"`

	// ExtraVolumes is an extra set of host volumes, mounted to the control plane component.
	ExtraVolumes []HostPathMount `json:"extraVolumes,omitempty"`
}

// APIServer holds settings necessary for API server deployments in the cluster
type APIServer struct {
	ControlPlaneComponent `json:",inline"`

	// CertSANs sets extra Subject Alternative Names for the API Server signing cert.
	CertSANs []string `json:"certSANs,omitempty"`

	// TimeoutForControlPlane controls the timeout that we use for API server to appear
	TimeoutForControlPlane *metav1.Duration `json:"timeoutForControlPlane,omitempty"`
}

// DNSAddOnType defines string identifying DNS add-on types
type DNSAddOnType string

const (
	// CoreDNS add-on type
	CoreDNS DNSAddOnType = "CoreDNS"

	// KubeDNS add-on type
	KubeDNS DNSAddOnType = "kube-dns"
)

// DNS defines the DNS addon that should be used in the cluster
type DNS struct {
	// Type defines the DNS add-on to be used
	Type DNSAddOnType `json:"type"`

	// ImageMeta allows to customize the image used for the DNS component
	ImageMeta `json:",inline"`
}

// ImageMeta allows to customize the image used for components that are not
// originated from the Kubernetes/Kubernetes release process
type ImageMeta struct {
	// ImageRepository sets the container registry to pull images from.
	// if not set, the ImageRepository defined in ClusterConfiguration will be used instead.
	ImageRepository string `json:"imageRepository,omitempty"`

	// ImageTag allows to specify a tag for the image.
	// In case this value is set, kubeadm does not change automatically the version of the above components during upgrades.
	ImageTag string `json:"imageTag,omitempty"`

	//TODO: evaluate if we need also a ImageName based on user feedbacks
}

// Networking contains elements describing cluster's networking configuration
type Networking struct {
	// ServiceSubnet is the subnet used by k8s services. Defaults to "10.96.0.0/12".
	ServiceSubnet string `json:"serviceSubnet"`
	// PodSubnet is the subnet used by pods.
	PodSubnet string `json:"podSubnet"`
	// DNSDomain is the dns domain used by k8s services. Defaults to "cluster.local".
	DNSDomain string `json:"dnsDomain"`
}

// Etcd contains elements describing Etcd configuration.
type Etcd struct {

	// Local provides configuration knobs for configuring the local etcd instance
	// Local and External are mutually exclusive
	Local *LocalEtcd `json:"local,omitempty"`

	// External describes how to connect to an external etcd cluster
	// Local and External are mutually exclusive
	External *ExternalEtcd `json:"external,omitempty"`
}

// LocalEtcd describes that kubeadm should run an etcd cluster locally
type LocalEtcd struct {
	// ImageMeta allows to customize the container used for etcd
	ImageMeta `json:",inline"`

	// DataDir is the directory etcd will place its data.
	// Defaults to "/var/lib/etcd".
	DataDir string `json:"dataDir"`

	// ExtraArgs are extra arguments provided to the etcd binary
	// when run inside a static pod.
	ExtraArgs map[string]string `json:"extraArgs,omitempty"`

	// ServerCertSANs sets extra Subject Alternative Names for the etcd server signing cert.
	ServerCertSANs []string `json:"serverCertSANs,omitempty"`
	// PeerCertSANs sets extra Subject Alternative Names for the etcd peer signing cert.
	PeerCertSANs []string `json:"peerCertSANs,omitempty"`
}

// ExternalEtcd describes an external etcd cluster.
// Kubeadm has no knowledge of where certificate files live and they must be supplied.
type ExternalEtcd struct {
	// Endpoints of etcd members. Required for ExternalEtcd.
	Endpoints []string `json:"endpoints"`

	// CAFile is an SSL Certificate Authority file used to secure etcd communication.
	// Required if using a TLS connection.
	CAFile string `json:"caFile"`

	// CertFile is an SSL certification file used to secure etcd communication.
	// Required if using a TLS connection.
	CertFile string `json:"certFile"`

	// KeyFile is an SSL key file used to secure etcd communication.
	// Required if using a TLS connection.
	KeyFile string `json:"keyFile"`
}

type HostPathMount struct {
	// Name of the volume inside the pod template.
	Name string `json:"name"`
	// HostPath is the path in the host that will be mounted inside
	// the pod.
	HostPath string `json:"hostPath"`
	// MountPath is the path inside the pod where hostPath will be mounted.
	MountPath string `json:"mountPath"`
	// ReadOnly controls write access to the volume
	ReadOnly bool `json:"readOnly,omitempty"`
	// PathType is the type of the HostPath.
	PathType v1.HostPathType `json:"pathType,omitempty"`
}
