package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

const workerCount = 10

var requestURL = flag.String("link", "", "URL to start crawling from")

func main() {
	// "https://scrape-me.dreamsofcode.io"
	flag.Parse()

	if *requestURL == "" {
		fmt.Println("Please provide a URL using --link flag")
		return
	}

	deadLinks := crawler(*requestURL)

	fmt.Println("Dead links:")
	for _, link := range deadLinks {
		fmt.Println(link)
	}
}

func crawler(rootURL string) []string {
	var wg sync.WaitGroup
	var crawlWg sync.WaitGroup

	visited := sync.Map{}
	queue := make(chan string, 100)
	deadLinkChan := make(chan string, 100)

	rootParsed, err := url.Parse(rootURL)
	if err != nil {
		fmt.Println("Invalid root URL:", err)
		return nil
	}
	rootHost := rootParsed.Host

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, queue, deadLinkChan, &visited, &crawlWg, rootHost)
	}

	crawlWg.Add(1)
	queue <- rootURL
	go func() {
		crawlWg.Wait()
		close(queue)
	}()

	go func() {
		wg.Wait()
		close(deadLinkChan)
	}()

	var deadLinks []string
	for link := range deadLinkChan {
		deadLinks = append(deadLinks, link)
	}
	return deadLinks
}

func worker(wg *sync.WaitGroup, queue chan string, deadlinks chan string, visited *sync.Map, crawlWg *sync.WaitGroup, rootHost string) {
	defer wg.Done()

	for current := range queue {
		if _, seen := visited.LoadOrStore(current, true); seen {
			crawlWg.Done()
			continue
		}

		resp, err := http.Get(current)
		if err != nil {
			deadlinks <- current
			crawlWg.Done()
			continue
		}
		if resp.StatusCode >= 400 && resp.StatusCode < 600 {
			deadlinks <- current
			resp.Body.Close()
			crawlWg.Done()
			continue
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			crawlWg.Done()
			continue
		}

		links := extractLinks(doc)
		for _, link := range links {
			absURL := resolveURL(current, link)
			if absURL == "" || !isValidLink(absURL, rootHost) {
				continue
			}
			crawlWg.Add(1)
			queue <- absURL
		}
		crawlWg.Done()
	}
}

func isValidLink(link, rootHost string) bool {
	if strings.HasPrefix(link, "mailto:") || strings.HasPrefix(link, "javascript:") {
		return false
	}
	parsed, err := url.Parse(link)
	if err != nil || parsed.Host == "" {
		return false
	}
	return parsed.Host == rootHost
}

func extractLinks(n *html.Node) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, extractLinks(c)...)
	}
	return links
}

func resolveURL(base, href string) string {
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	hrefURL, err := url.Parse(href)
	if err != nil {
		return ""
	}
	return baseURL.ResolveReference(hrefURL).String()
}
