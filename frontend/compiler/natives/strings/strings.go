// +build js

package strings

import (
	"github.com/glycerine/dynamic-go/frontend/js"
)

func IndexByte(s string, c byte) int {
	return js.InternalObject(s).Call("indexOf", js.Global.Get("String").Call("fromCharCode", c)).Int()
}
