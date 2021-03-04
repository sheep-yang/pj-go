package controller

import (
	"github.com/qiangyang/sample-controller/pkg/controller/admin"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, admin.Add)
}
