// AWS Target Group Controller.

package controller

import (
	"github.com/ewbankkit/kube-controller/pkg/controller/service"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, service.Add)
}
