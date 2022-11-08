package main

import (
	"encoding/xml"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/romankravchuk/learn-go/link"
)

var (
	urlFlag  string
	maxDepth int
	filename string
)

func init() {
	flag.StringVar(&urlFlag, "url", "https://gophercises.com", "the url that you want to build sitemap for.")
	flag.StringVar(&filename, "filename", "map.xml", "the xml file to which the sitemap will be written.")
	flag.IntVar(&maxDepth, "depth", 0, "the maximum number of links deep to traverse")
}

/*
	TODO:
	// * GET the webpage
	// * parse all the links on the page
	// * build proper urls with our links
	// * filter out any links w/ a diff domain
	// * find all pages (BFS)
	* print out XML
*/

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	flag.Parse()

	pages := bfs(urlFlag, maxDepth)
	toXml := urlset{
		Urls:  make([]loc, len(pages)),
		Xmlns: xmlns,
	}
	for i, page := range pages {
		toXml.Urls[i] = loc{page}
	}

	if data, err := xml.MarshalIndent(toXml, "", " "); err != nil {
		exit(err)
	} else {
		data = []byte(xml.Header + string(data))
		if err := os.WriteFile(filename, data, 0644); err != nil {
			exit(err)
		}
	}
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nq[link] = struct{}{}
				}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		exit(err)
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func exit(err error) {
	log.Fatalf("ERROR | %v", err)
}
