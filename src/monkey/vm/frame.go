package vm

import (
	"monkey/code"
	"monkey/object"
)

type Frame struct {
	fn *object.CompiledFunction
	ip int
	// The pointer to restore to, aka frame pointer
	// https://en.wikipedia.org/wiki/Call_stack
	basePointer int
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	f := &Frame{fn: fn, ip: -1, basePointer: basePointer}
	return f
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
