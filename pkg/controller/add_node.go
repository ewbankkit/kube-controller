// Amazon VPC CNI Custom Network Controller.

package controller

import (
	"github.com/ewbankkit/kube-controller/pkg/controller/node"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, node.Add)
}
