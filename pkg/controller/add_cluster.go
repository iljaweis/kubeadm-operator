package controller

import (
	"github.com/iljaweis/kubeadm-operator/pkg/controller/cluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cluster.Add)
}
