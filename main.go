package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"math/rand"
)

func main() {
	const (
		url = "https://www.color-hex.com/color-palettes/popular"
	)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	randomNum := rand.Intn(80)

	var f func(*html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "div" {
			for _, a := range node.Attr {
				if a.Key == "class" {
					if a.Val == "palettecolordivcon" {
						if count == randomNum {
							for c := node.FirstChild; c != nil; c = c.NextSibling {
								for _, attr := range c.Attr {
									if attr.Key == "style" {
										fmt.Println(attr.Val[17:])
									}
								}
							}
						}
						count++
					}
					break
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
}

