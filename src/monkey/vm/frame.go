package vm

import (
	"monkey/code"
	"monkey/object"
)

type Frame struct {
	cl *object.Closure
	ip int
	// The pointer to restore to, aka frame pointer
	// https://en.wikipedia.org/wiki/Call_stack
	basePointer int
}

func NewFrame(cl *object.Closure, basePointer int) *Frame {
	f := &Frame{cl: cl, ip: -1, basePointer: basePointer}
	return f
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
