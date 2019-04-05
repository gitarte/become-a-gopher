package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// URL to by scrapped
const (
	URL             = "http://www.fabpedigree.com/james/mathmen.htm"
	UserAgentValue  = "FRIENDLY SCANNER"
	UserAgentHeader = "User-Agent"
	Timeout         = 30
)

func main() {
	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set(UserAgentHeader, UserAgentValue)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: Timeout * time.Second,
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
	titleStartIndex += len("<title>")

	// Find the index of the closing tag
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		os.Exit(0)
	}
	pageTitle := pageContent[titleStartIndex:titleEndIndex]
	fmt.Println(pageTitle)

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

	mat := make([]string, 0)
	document.Find("li").Each(
		func(index int, element *goquery.Selection) {
			a := element.Find("a").Text()
			if len(a) > 0 {
				mat = append(mat, a)
			}
		})

	sort.Strings(mat)

	fmt.Printf("%v\n", mat)
	fmt.Println(mat[0])
	fmt.Println(mat[len(mat)-1])
}
