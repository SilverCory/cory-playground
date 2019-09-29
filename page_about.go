package main

func SnippetAbout() *snippet {
	const BodyStringAbout = `
package main

import "fmt"

func main() {

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
`
	const HTMLStringAbout = `<div class="vh-100" style="background:red;"></div>`

	return &snippet{Body: []byte(BodyStringAbout), HTML: []byte(HTMLStringAbout)}
}
