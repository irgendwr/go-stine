package api

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

func scrapeTableBody(document *goquery.Document, rowCallback func(index int, tr *goquery.Selection)) error {
	tables := document.Find("table tbody")
	if tables.Length() != 1 {
		return errors.Errorf("expected 1 table, got %d instead", tables.Length())
	}

	tables.Find("tr").Each(rowCallback)
	return nil
}

// Text gets the combined text contents of each element in the set of matched
// elements, including their descendants.
func scrapeText(s *goquery.Selection) string {
	var buf bytes.Buffer
	var cb func(n *html.Node)
	cb = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(reSpace.ReplaceAllLiteralString(strings.TrimSpace(n.Data), " "))
		} else if n.Type == html.ElementNode && n.Data == "br" {
			buf.WriteRune('\n')
		}
		if n.FirstChild != nil {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				cb(c)
			}
		}
	}
	for _, n := range s.Nodes {
		cb(n)
	}

	return buf.String()
}
