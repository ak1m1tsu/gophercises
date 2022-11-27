package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/romankravchuk/learn-go/lib/hn"
)

func main() {
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "stories", 30, "")

	tpl := template.Must(template.ParseFiles("./index.html"))

	http.HandleFunc("/", handler(numStories, tpl))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	sc := storyCache{
		duration:   6 * time.Second,
		numStories: numStories,
	}

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			temp := storyCache{
				numStories: numStories,
				duration:   6 * time.Second,
			}
			temp.stories()
			sc.mutex.Lock()
			sc.cache = temp.cache
			sc.expiration = temp.expiration
			sc.mutex.Unlock()
			<-ticker.C
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := sc.stories()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := templateData{
			Stories: stories,
			Time:    time.Since(start),
		}
		if err := tpl.Execute(w, data); err != nil {
			http.Error(w, "failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

type storyCache struct {
	numStories int
	cache      []item
	expiration time.Time
	duration   time.Duration
	mutex      sync.Mutex
}

func (sc *storyCache) stories() ([]item, error) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	if time.Since(sc.expiration) < 0 {
		return sc.cache, nil
	}
	stories, err := getTopStories(sc.numStories)
	if err != nil {
		return nil, err
	}
	sc.expiration = time.Now().Add(sc.duration)
	sc.cache = stories
	return sc.cache, nil
}

func getTopStories(numStories int) ([]item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("failed to load top stories")
	}
	var results []result
	at := 0
	for len(results) < numStories {
		need := (numStories - len(results)) * 5 / 4
		results = append(results, getResults(&client, ids[at:at+need])...)
		at += need
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].idx < results[j].idx
	})
	return resultsToStories(results)[:numStories], nil
}

type result struct {
	idx  int
	item item
	err  error
}

func getResults(client *hn.Client, ids []int) []result {
	var results []result
	resultCh := make(chan result)
	for i := 0; i < len(ids); i++ {
		go func(idx, id int) {
			hnItem, err := client.GetItem(id)
			if err != nil {
				resultCh <- result{idx: idx, err: err}
			}
			resultCh <- result{idx: idx, item: parseHNItem(hnItem)}
		}(i, ids[i])
	}
	for i := 0; i < len(ids); i++ {
		results = append(results, <-resultCh)
	}
	return results
}

func resultsToStories(results []result) []item {
	var stories []item
	for _, res := range results {
		if res.err != nil {
			continue
		}
		if isStoryLink(res.item) {
			stories = append(stories, res.item)
		}
	}
	return stories
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
