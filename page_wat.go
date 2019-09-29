package main

const BodyStringWat = `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("I see you're lurking... Or you're lost?'")
	fmt.Println("Well, I hope you've enjoyed this!")
	fmt.Println("Have some weird go code..")

	// Two nil variables...
	var a *int = nil
	var b interface{} = nil

	fmt.Println()

	fmt.Println("a == nil:", a == nil)
	fmt.Println("b == nil:", b == nil)
	fmt.Println("a == b:", a == b) // a == b right?
}
`