package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

type A struct {
	a [4]byte
}

type B struct {
	b [4]byte
	A
}

func (a *A) getself() *A {
	return a
}

func getnewa() *A {
	b := B{[4]byte{80, 80, 80, 80}, A{[4]byte{81, 81, 81, 81}}}
	return &A{b.a}
}

func main() {
	f, _ := os.Create(os.Args[1])
	defer f.Close()

	go func() {
		defer fmt.Print(".")
		time.Sleep(1 * time.Second)
	}()

	a := getnewa()
	runtime.GC()

	debug.WriteHeapDump(f.Fd())
	fmt.Printf("%v\n", a)
}
