// Package stackError A have stack information with error library
package stackError

import (
	"runtime"
)

var pkgName string = "github.com/lingdor/stackError"

func init() {
	pc := make([]uintptr, 1)
	runtime.Callers(1, pc)
	frm := runtime.CallersFrames(pc)
	frmval, _ := frm.Next()
	//if ok {
	newName := funcNameToPkgName(frmval.Function)
	if newName != "" {
		pkgName = newName
	}
	//}
}
