package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	Controllers         []ClusterMember `json:"controllers"`
	Workers             []ClusterMember `json:"workers"`
	DefaultSSHKeySecret string          `json:"defaultsshkeysecret"`
	Version             string          `json:"version"`
}

type ClusterMember struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

type ControllerStatus struct {
	Ready     bool   `json:"ready"`
	AdminConf string `json:"adminconf"`
}

type WorkerStatus struct {
	Ready bool `json:"ready"`
}

type ClusterStatusPKI struct {
	CACert         string `json:"cacert"`
	CAKey          string `json:"cakey"`
	SAPrivate      string `json:"saprivate"`
	SAPublic       string `json:"sapublic"`
	FrontProxyKey  string `json:"frontproxykey"`
	FrontProxyCert string `json:"frontproxycert"`
	EtcdCert       string `json:"etcdcert"`
	EtcdKey        string `json:"etcdkey"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	Controllers              map[string]ControllerStatus `json:"controllers"`
	Workers                  map[string]WorkerStatus     `json:"workers"`
	KubeadmConfig            *ClusterConfiguration       `json:"kubeadmconfig"`
	JoinToken                string                      `json:"jointoken"`
	DiscoveryTokenCACertHash string                      `json:"discoverytokencacerthash"`
	AdminConf                string                      `json:"adminconf"`
	PKI                      ClusterStatusPKI            `json:"pki"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cluster is the Schema for the clusters API
// +k8s:openapi-gen=true
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
