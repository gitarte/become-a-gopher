package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// URL to by scrapped
var URL = "http://www.fabpedigree.com/james/mathmen.htm"
var UserAgent = "FRIENDLY SCANNER"

func main() {
	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", UserAgent)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to variable
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	pageContent := string(bodyBytes)
	//fmt.Println(pageContent)

	/*
		=====================================================
			        FIND THE TITLE OF THE PAGE
		=====================================================
	*/
	// Find the beginning of title section
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}
	// offset the infex just after <title> so '<title>' will be excluded
	titleStartIndex += 7

	// Find the index of the closing tag
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		os.Exit(0)
	}
	pageTitle := pageContent[titleStartIndex:titleEndIndex]
	fmt.Println(pageTitle)

	split := strings.Split(pageContent, `<a href="#top">Top</a>`)
	fmt.Println(split)

	/*
		=====================================================
			        FIND ALL <li> ELEMENTS ON THE PAGE
		=====================================================
	*/
	// restore body since we had destroyed it
	response.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(pageContent)))
	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	document.Find("li").Each(
		func(index int, element *goquery.Selection) {
			a := element.Find("a")
			link, _ := a.Attr("href")
			document.Find("a").Each(
				func(index int, element *goquery.Selection) {
					href, _ := element.Attr("href")
					if href == link {
						fmt.Println(href)
					}
				})
		})
}
