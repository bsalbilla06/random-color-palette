package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	const url = "https://www.color-hex.com/random-color-palette"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	html := doc.LastChild
	body := html.LastChild

	container := getFirstDiv(body)
	content := getFirstDiv(container)
	row := getRowDiv(content)
	randomPalette := getFirstDiv(row)
	anchor := getAnchor(randomPalette.FirstChild)
	name := anchor.LastChild.Data[1:]
	fmt.Println(name)

	count := 0
	paletteContainer := getFirstDiv(anchor)
	for c := paletteContainer.FirstChild; count < 10; c = c.NextSibling {
		if count%2 != 0 {
			fmt.Println(c.Attr[1].Val[17:])
		}
		count++
	}
}

func getFirstDiv(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "div" {
			return c
		}
	}
	return nil
}

func getAnchor(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "a" {
			return c
		}
	}
	return nil
}

func getRowDiv(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "div" {
			for _, attr := range c.Attr {
				if attr.Key == "class" && attr.Val == "row" {
					return c
				}
			}
		}
	}
	return nil
}
