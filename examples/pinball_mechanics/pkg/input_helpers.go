package pkg

import (
	"unsafe"

	. "github.com/godot-go/godot-go/pkg/builtin"
	. "github.com/godot-go/godot-go/pkg/core"
	. "github.com/godot-go/godot-go/pkg/ffi"
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
	. "github.com/godot-go/godot-go/pkg/gdutilfunc"
)

// getGDClassInstance looks up a custom GDExtension class instance by its Object.
// This is needed because ObjectCastTo doesn't work for custom GDExtension classes.
func getGDClassInstance[T GDClass](obj Object) T {
	var zero T
	if obj == nil {
		return zero
	}
	owner := obj.GetGodotObjectOwner()
	if owner == nil {
		return zero
	}
	id := CallFunc_GDExtensionInterfaceObjectGetInstanceId(
		(GDExtensionConstObjectPtr)(unsafe.Pointer(owner)),
	)
	inst, ok := Internal.GDClassInstances.Get(id)
	if !ok {
		return zero
	}
	result, ok := inst.(T)
	if !ok {
		return zero
	}
	return result
}

func nodePath(path string) NodePath {
	str := NewStringWithUtf8Chars(path)
	defer str.Destroy()
	return NewNodePathWithString(str)
}

func printLine(text string) {
	v := NewVariantGoString(text)
	defer v.Destroy()
	Print(v)
}

func inputSingleton() Input {
	return GetInputSingleton()
}
