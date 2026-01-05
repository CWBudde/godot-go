package gdclassimpl

import (
	"runtime"
	"unsafe"

	. "github.com/godot-go/godot-go/pkg/builtin"
)

var (
	nullptr = unsafe.Pointer(nil)
	pnr     runtime.Pinner
)

// UnpinGDClassImplPins releases pins held for the extension lifetime during shutdown.
func UnpinGDClassImplPins() {
	pnr.Unpin()
}

func (cx *ObjectImpl) ToGoString() string {
	if cx == nil || cx.Owner == nil {
		return ""
	}
	gdstr := cx.ToString()
	defer gdstr.Destroy()
	return gdstr.ToUtf8()
}

func GetInputSingleton() Input {
	owner := GetSingleton("Input")
	if owner == nil {
		return nil
	}
	return NewInputWithGodotOwnerObject((*GodotObject)(unsafe.Pointer(owner)))
}
