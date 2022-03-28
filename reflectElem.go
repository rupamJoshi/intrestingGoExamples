// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

func test(t interface{}) {
	b := reflect.ValueOf(t)
	fmt.Print(b.Elem())
}

func main() {
	t := 4
	test(&t)
}
