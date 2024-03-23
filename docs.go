// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/docs/docs.go */

package docs

import (
	"strings"
	"github.com/utopiagio/docs/apireference"
	"github.com/utopiagio/docs/overview"
	"github.com/utopiagio/docs/reference"
)


func Page(docRequest string) (packgName string, docName string, content []string) {
	packg, doc, found := strings.Cut(docRequest, ".")
	if found {
		switch packg {
			case "api":
				packgName = "API Reference"
				docName = doc
				content = apireference.Page(doc)
			case "ref":
				packgName = "Reference"
				docName = doc
				content = reference.Page(doc)
			case "ove":
				packgName = "Overview"
				docName = doc
				content = overview.Page(doc)
			default:
				packgName = "Page Not Found"
				docName = ""
				content = PageNotFound
		}
	}
	return
}

var PageNotFound []string = []string{""}
