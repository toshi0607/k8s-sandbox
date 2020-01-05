package controller

import (
	"github.com/toshi0607/k8s-sandbox/sample-controller-operatorsdk/pkg/controller/foo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, foo.Add)
}
