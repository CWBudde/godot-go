package ffi

import (
	"runtime"
	"unsafe"

	"github.com/godot-go/godot-go/pkg/log"
)

func CallBuiltinConstructor(constructor GDExtensionPtrConstructor, base GDExtensionUninitializedTypePtr, args ...GDExtensionConstTypePtr) {
	if constructor == nil {
		log.Panic("constructor is null")
	}
	argsPtr := (*GDExtensionConstTypePtr)(unsafe.SliceData(args))
	p := runtime.Pinner{}
	p.Pin(argsPtr)
	defer p.Unpin()
	CallFunc_GDExtensionPtrConstructor(constructor, base, argsPtr)
}

func CallBuiltinMethodPtrRet[T any](method GDExtensionPtrBuiltInMethod, base GDExtensionTypePtr, args ...GDExtensionTypePtr) T {
	m := (GDExtensionPtrBuiltInMethod)(method)
	b := (GDExtensionTypePtr)(base)
	a := (*GDExtensionConstTypePtr)(unsafe.SliceData(args))
	ca := (int32)(len(args))
	var ret T
	ptr := (GDExtensionTypePtr)(unsafe.Pointer(&ret))
	p := runtime.Pinner{}
	p.Pin(a)
	p.Pin(ptr)
	defer p.Unpin()
	CallFunc_GDExtensionPtrBuiltInMethod(m, b, a, ptr, ca)
	return ret
}

func CallBuiltinMethodPtrNoRet(method GDExtensionPtrBuiltInMethod, base GDExtensionTypePtr, args ...GDExtensionTypePtr) {
	m := (GDExtensionPtrBuiltInMethod)(method)
	b := (GDExtensionTypePtr)(base)
	a := (*GDExtensionConstTypePtr)(unsafe.SliceData(args))
	ca := (int32)(len(args))
	p := runtime.Pinner{}
	p.Pin(a)
	defer p.Unpin()
	CallFunc_GDExtensionPtrBuiltInMethod(m, b, a, nil, ca)
}

func CallBuiltinOperatorPtr[T any](operator GDExtensionPtrOperatorEvaluator, left GDExtensionConstTypePtr, right GDExtensionConstTypePtr) T {
	op := (GDExtensionPtrOperatorEvaluator)(operator)
	l := (GDExtensionConstTypePtr)(left)
	r := (GDExtensionConstTypePtr)(right)
	var ret T
	ptr := (GDExtensionTypePtr)(unsafe.Pointer(&ret))
	p := runtime.Pinner{}
	p.Pin(l)
	p.Pin(r)
	p.Pin(ptr)
	defer p.Unpin()
	CallFunc_GDExtensionPtrOperatorEvaluator(op, l, r, ptr)
	return ret
}

func CallBuiltinPtrGetter[T any](getter GDExtensionPtrGetter, base GDExtensionConstTypePtr) T {
	g := (GDExtensionPtrGetter)(getter)
	b := (GDExtensionConstTypePtr)(base)
	var ret T
	ptr := (GDExtensionTypePtr)(unsafe.Pointer(&ret))
	p := runtime.Pinner{}
	p.Pin(ptr)
	defer p.Unpin()
	CallFunc_GDExtensionPtrGetter(g, b, ptr)
	return ret
}

func CallBuiltinPtrSetter[T any](setter GDExtensionPtrSetter, base GDExtensionTypePtr) T {
	g := (GDExtensionPtrSetter)(setter)
	b := (GDExtensionTypePtr)(base)
	var ret T
	ptr := (GDExtensionConstTypePtr)(unsafe.Pointer(&ret))
	p := runtime.Pinner{}
	p.Pin(ptr)
	defer p.Unpin()
	CallFunc_GDExtensionPtrSetter(g, b, ptr)
	return ret
}
