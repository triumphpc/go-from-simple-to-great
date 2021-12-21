//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.6: Return values
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Entrypoint
func main() {
	fmt.Println(math.Floor(5.6))
	fmt.Println(strings.Title("Some text"))

	links, err := getLinks("https://golang.org/")
	if err != nil {
		log.Fatal(err)
	}

	// output links
	for _, v := range links {
		fmt.Println(v)
	}

	// count of images
	w, i, err := CountWordsAndImages("https://google.com/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(w, i)
}

//Collect all links from response body and return it as an array of strings
func getLinks(link string) ([]string, error) {
	// example get links
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	var links []string
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

// CountWordsAndImages get HTTP-request and return count of words and images.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	fmt.Println(doc)
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		words += len(strings.Split(n.Data, " "))
		return
	}

	if n.Data == "img" {
		images++
	} else {
		if n.FirstChild != nil {
			w, i := countWordsAndImages(n.FirstChild)
			words += w
			images += i
		}
		if n.NextSibling != nil {
			w, i := countWordsAndImages(n.NextSibling)
			words += w
			images += i
		}
	}
	return
}
