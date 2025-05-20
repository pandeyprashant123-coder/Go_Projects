package main

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	requestURL := "https://scrape-me.dreamsofcode.io"

	deadLinks := crawler(requestURL)
	fmt.Println("dead returned")

	for _, link := range deadLinks {
		fmt.Println(link)
	}
}

func crawler(rootURL string) []string {
	// originDomain := rootURL
	var visited []string
	var queue = make([]string, 0)
	var deadLinks []string
	queue = enqueue(queue, rootURL)
	for len(queue) > 0 {

		currentUrl, newQueue := dequeue(queue)
		queue = newQueue
		if slices.Contains(visited, currentUrl) {
			continue
		}
		visited = append(visited, currentUrl)
		res, err := http.Get(currentUrl)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			continue
		}
		status := res.StatusCode
		if status >= 400 && status < 600 {
			// fmt.Println("dead link:", currentUrl)
			deadLinks = append(deadLinks, currentUrl)
		}
		// fmt.Println("ok:", currentUrl)
		defer res.Body.Close()
		doc, err := html.Parse(res.Body)
		if err != nil {
			fmt.Printf("client: could not read response body: %s\n", err)
		}
		paths := processBody(doc)
		for _, p := range paths {
			normalizedURL := currentUrl + p
			if !isValidURL(normalizedURL) {
				continue
			}
			if slices.Contains(visited, normalizedURL) {
				continue
			}
			queue = append(queue, normalizedURL)
		}

	}
	return deadLinks

}

func isValidURL(URL string) bool {
	return strings.HasPrefix(URL, "http://") || strings.HasPrefix(URL, "https://")
}
func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []string) (string, []string) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []string{}
		return element, tmp
	}
	return element, queue[1:] // Slice off the element once it is dequeued.
}

func processBody(n *html.Node) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && strings.HasPrefix(a.Val, "/") {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, processBody(c)...)
	}
	return links
}
