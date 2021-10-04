package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var total = 0
func visit(links []string, node *html.Node) []string {

	//if total >= 20 {
	//	return links
	//}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
				fmt.Println(a.Val)
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		visit(links, child)
	}

	return links
}

func main()  {
	url := "https://baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("get a html file failed")
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	//fmt.Println(string(b))
	doc, err := html.Parse(strings.NewReader(string(b)))
	//fmt.Println(doc)
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}