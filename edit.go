// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime"
	"strings"
)

const hostname = "play.golang.org"

var editTemplate = template.Must(template.ParseFiles("edit.html"))

type editData struct {
	Snippet   *snippet
	Share     bool
	Analytics bool
	GoVersion string
}

func (s *server) handleEdit(w http.ResponseWriter, r *http.Request) {
	// Redirect foo.play.golang.org to play.golang.org.
	if strings.HasSuffix(r.Host, "."+hostname) {
		http.Redirect(w, r, "https://"+hostname, http.StatusFound)
		return
	}

	snip := &snippet{Body: []byte(wat)}

	id := r.URL.Path[1:]
	serveText := false
	if strings.HasSuffix(id, ".go") {
		id = id[:len(id)-3]
		serveText = true
	}

	if err := s.db.GetSnippet(r.Context(), id, snip); err != nil {
		//snip = &snippet{Body: []byte(wat)}
	}

	if serveText {
		if r.FormValue("download") == "true" {
			w.Header().Set(
				"Content-Disposition", fmt.Sprintf(`attachment; filename="%s.go"`, id),
			)
		}
		w.Header().Set("Content-type", "text/plain; charset=utf-8")
		w.Write(snip.Body)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := &editData{
		Snippet:   snip,
		Share:     allowShare(r),
		Analytics: r.Host == hostname,
		GoVersion: runtime.Version(),
	}
	if err := editTemplate.Execute(w, data); err != nil {
		s.log.Errorf("editTemplate.Execute(w, %+v): %v", data, err)
		return
	}
}

const wat = `package main

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
