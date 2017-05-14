# Go Package: call

  <a href="https://godoc.org/github.com/cstockton/go-call"><img src="https://img.shields.io/badge/%20docs-reference-5272B4.svg?style=flat-square"></a> [![Go Report Card](https://goreportcard.com/badge/github.com/cstockton/go-call)](https://goreportcard.com/report/github.com/cstockton/go-call)

  > Get:
  > ```bash
  > go get -u github.com/cstockton/go-call
  > ```
  >
  > Example:
  > ```Go
  > func Example() {
  > 	fn, err := call.Func(func(i *int) {
  > 		*i++
  > 	})
  > 	if err != nil { log.Fatal(err) }
  > 	i := new(int)
  > 	for *i < 100 {
  > 		fn(i)
  > 	}
  > 	fmt.Printf("Looped %v times", *i)
  >
  > 	// Output:
  > 	// Looped 100 times
  > }
  > ```


## About

Package call provides a simple interface for safely calling functions (quickly) at runtime while retaining types for the callee.


## Bugs and Patches

  Feel free to report bugs and submit pull requests.

  * bugs:
    <https://github.com/cstockton/go-call/issues>
  * patches:
    <https://github.com/cstockton/go-call/pulls>



[Go Doc]: https://godoc.org/github.com/cstockton/go-call
