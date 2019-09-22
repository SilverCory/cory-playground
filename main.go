// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"net/http"
	"os"
)

var log = newStdLogger()

func main() {
	s, err := newServer(func(s *server) error {
		s.db = &inMemStore{}
		_ = s.db.PutSnippet(context.Background(), "", &snippet{Body:[]byte(`package main

import "fmt"

func main() {
	fmt.Println("Hey I'm Cory, a Golang developer.")
	fmt.Println("Looks like you've stumbled upon a WIP site of mine")
	fmt.Println("I'm not sure how, I wrote this message Sept 2019, if it's been a few months")
	fmt.Println("then this may be all you see!")
}`)})

		_ = s.db.PutSnippet(context.Background(), "about", &snippet{Body:[]byte(`package main

import "fmt"

func main() {

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}`)})
		s.log = log
		return nil
	})
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	if len(os.Args) > 1 && os.Args[1] == "test" {
		s.test()
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on :%v ...", port)
	log.Fatalf("Error listening on :%v: %v", port, http.ListenAndServe(":"+port, s))
}
