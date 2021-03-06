// +build js

package syscall

import (
	"bytes"
	"unsafe"

	"github.com/glycerine/dyg/frontend/js"
)

var warningPrinted = false
var lineBuffer []byte

func init() {
	js.Global.Set("$flushConsole", js.InternalObject(func() {
		if len(lineBuffer) != 0 {
			js.Global.Get("console").Call("log", string(lineBuffer))
			lineBuffer = nil
		}
	}))
}

func printWarning() {
	if !warningPrinted {
		println("warning: system calls not available, see https://github.com/glycerine/dyg/frontend/blob/master/doc/syscalls.md")
	}
	warningPrinted = true
}

func printToConsole(b []byte) {
	goPrintToConsole := js.Global.Get("goPrintToConsole")
	if !goPrintToConsole.IsUndefined() {
		goPrintToConsole.Invoke(js.InternalObject(b))
		return
	}

	lineBuffer = append(lineBuffer, b...)
	for {
		i := bytes.IndexByte(lineBuffer, '\n')
		if i == -1 {
			break
		}
		js.Global.Get("console").Call("log", string(lineBuffer[:i])) // don't use println, since it does not externalize multibyte characters
		lineBuffer = lineBuffer[i+1:]
	}
}

func use(p unsafe.Pointer) {
	// no-op
}
