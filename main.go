package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type CodeResponse struct {
	Items []Item
	Count int `json:"total_count"`
}

type Item struct {
	Path string     `json:"path"`
	Repo Repository `json:"repository"`
}

type Repository struct {
	Url string `json:"html_url"`
}

func main() {

	// Get our flags
	var pathFlag string
	var orgFlag string
	var repoFlag string
	var outputFlag string
	var vFlag bool
	var vvFlag bool
	var vvvFlag bool
	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&pathFlag, "path", "", "Path in repo.")
	mySet.StringVar(&orgFlag, "org", "", "GitHub organization to search within.")
	mySet.StringVar(&repoFlag, "repo", "", "GitHub repo to search within.")
	mySet.StringVar(&outputFlag, "output", "json", "Format of output 'text' or 'json'.")
	mySet.BoolVar(&vFlag, "v", false, "Logs non-sensitive values for debugging, also logs raw github output unfiltered.")
	mySet.BoolVar(&vvFlag, "vv", false, "Everything from -v plus !! sensitive values !!")
	mySet.BoolVar(&vvvFlag, "vvv", false, "Everything from -vv plus raw github output unfiltered.")
	mySet.Parse(os.Args[2:])

	// Get our required search string from Args.
	searchArg := os.Args[1]

	// Validate the output flag.
	checkOutput(outputFlag)

	// Construct query string
	queryString := constructQueryString(searchArg, pathFlag, orgFlag, repoFlag)

	// Set our GitHub user and token from os environment.
	user := os.Getenv("GH_USER")
	token := os.Getenv("GH_TOKEN")

	// Print flags
	if vFlag || vvFlag || vvvFlag {
		fmt.Println("search:", searchArg)
		fmt.Println("path:", pathFlag)
		fmt.Println("org:", orgFlag)
		fmt.Println("repo:", repoFlag)
		fmt.Println("output:", outputFlag)
		fmt.Println("queryString:", queryString)
	}
	if vvFlag || vvvFlag {
		fmt.Println("user:", user)
		fmt.Println("token:", token)
	}

	// Instantiate our http client.
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	// Set our request, but not make it yet.
	req, err := http.NewRequest("GET", "https://api.github.com/search/code", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Add auth to request.
	req.SetBasicAuth(user, token)
	// Add headers to request.
	req.Header.Add("Accept", "application/vnd.github.v3.text-match+json")

	// Add query string parameteres to request.
	q := req.URL.Query()
	q.Add("q", queryString)
	req.URL.RawQuery = q.Encode()

	// Print the request for debugging.
	if vFlag || vvFlag || vvvFlag {
		fmt.Println("Request:", req.URL)
	}

	// Make the call.
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// Get the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Convert the body to string.
	sb := string(body)

	if vvvFlag {
		fmt.Println(sb)
	}

	// Output our specified format.
	if outputFlag == "json" {
		finalResp := constructFinalResponseJson(sb)
		fmt.Println(finalResp)
	} else if outputFlag == "text" {
		finalResp := constructFinalResponseText(sb)
		fmt.Println(finalResp)
	} else if outputFlag == "csv" {
		finalResp := constructFinalResponseCsv(sb)
		fmt.Println(finalResp)
	}

}
