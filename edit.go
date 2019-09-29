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

var editTemplate = template.Must(template.New("edit.html").Funcs(map[string]interface{}{
	"printHTML": func(t string) template.HTML {
		return template.HTML(t)
	},
}).ParseFiles("edit.html"))

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

	snip := SnippetWat()

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
	data := map[string]interface{}{
		"Analytics":   true,
		"Snippet":     snip,
		"GoVersion":   runtime.Version(),
		"CurrentPage": id,
	}
	if err := editTemplate.Execute(w, data); err != nil {
		s.log.Errorf("editTemplate.Execute(w, %+v): %v", data, err)
		return
	}
}
