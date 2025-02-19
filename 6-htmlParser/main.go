package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var url = "https://english.iut.ac.ir"

// var htmlDoc = `<!DOCTYPE html>
// <html>
//   <body>
//     <h1>My First Heading</h1>
//       <p>My first paragraph.</p>
//       <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
//       <img src="xxx.jpg" width="104" height="142">
//   </body>
// </html>`

func main() {

	htmlDoc, err := getPage(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not get html: %s\n", err.Error())
		os.Exit(-1)
	}

	doc, err := html.Parse(bytes.NewReader([]byte(htmlDoc)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not Parse html: %s\n", err.Error())
		os.Exit(-1)
	}

	wordCount, imgCount := countWandImg(doc)

	fmt.Printf("Word count: %d, Image Count: %d\n", wordCount, imgCount)

}

func countWandImg(doc *html.Node) (int, int) {
	var wordCount, imgCount int

	visiting(doc, &wordCount, &imgCount)

	return wordCount, imgCount
}

func visiting(n *html.Node, pWordC, pImgC *int) {

	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" && len(strings.TrimSpace(n.Data)) != 0 {
		fmt.Println("++text++:", n.Data) ///TODO
		*pWordC += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		fmt.Println("**img**:")
		*pImgC++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visiting(c, pWordC, pImgC)
	}
}

func getPage(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		content, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		os.WriteFile("index.html", content, 0644)
		return content, nil
	}

	defer res.Body.Close()
	return nil, fmt.Errorf("can not get page: %s", res.Status)
}
