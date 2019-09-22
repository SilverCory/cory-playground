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
		_ = s.db.PutSnippet(context.Background(), "about", &snippet{Body:[]byte(`package main

import "fmt"

func main() {
	fmt.Println("This is some kind of buullll shieeeet")
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
