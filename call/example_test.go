package call_test

import (
	"fmt"
	"log"

	call "github.com/cstockton/go-call"
)

func ExampleNew() {
	caller, err := call.New(func(s string) {
		fmt.Println(s, "World")
	})
	if err != nil {
		log.Fatal(err)
	}
	caller.Call("Hello")

	// Output:
	// Hello World
}

func Example_reuse() {
	caller, err := call.New(func(i *int) {
		*i++
	})
	if err != nil {
		log.Fatal(err)
	}

	// Once you create a caller you may call it as many times as you want.
	i := new(int)
	for *i < 100 {
		caller.Call(i)
	}
	fmt.Printf("Looped %v times", *i)

	// Output:
	// Looped 100 times
}

func Example_pointers() {

	// Pointers work fine as well.
	caller, err := call.New(func(i *int, s *string) {
		*i = 2
		*s = "World"
	})
	if err != nil {
		log.Fatal(err)
	}

	arg1, arg2 := 1, "dlroW"
	caller.Call(&arg1, &arg2)
	fmt.Println(arg1, arg2)

	// Output:
	// 2 World
}
