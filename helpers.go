package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func constructQueryString(search string, path string, org string, repo string) string {
	if org == "" && repo == "" {
		fmt.Println("Error: --org or --repo required.")
		os.Exit(1)
	}

	queryString := search
	if path != "" {
		queryString = queryString + " path:" + path
	}
	if org != "" {
		queryString = queryString + " org:" + org
	}
	if repo != "" {
		queryString = queryString + " repo:" + repo
	}

	return queryString
}

func checkOutput(output string) {
	if !isValidOutput(output) {
		fmt.Println("Error: --output must be 'json', 'text', or 'csv'.")
		os.Exit(1)
	}

}

func isValidOutput(output string) bool {
	switch output {
	case
		"json",
		"text",
		"csv":
		return true
	}
	return false
}

func constructFinalResponseJson(sb string) []string {

	var obj2 CodeResponse
	json.Unmarshal([]byte(sb), &obj2)

	var finalResp []string
	// Get the length of our array.
	obj2len := len(obj2.Items)
	// Loop over our Response and add to our array.
	for i := range obj2.Items {
		// Check if we are at the last item.
		if i == obj2len-1 {
			// Don't include trailing comma if this is last object.
			finalResp = append(finalResp, `{"Repo":"`+obj2.Items[i].Repo.Url+`","Path":"`+obj2.Items[i].Path+`"}`)
		} else {
			finalResp = append(finalResp, `{"Repo":"`+obj2.Items[i].Repo.Url+`","Path":"`+obj2.Items[i].Path+`"},`)
		}
	}

	return finalResp
}

func constructFinalResponseText(sb string) string {

	var obj2 CodeResponse
	json.Unmarshal([]byte(sb), &obj2)

	var finalResp string
	// Loop over our Response and add to our array.
	for i := range obj2.Items {
		finalResp = finalResp + "\n" + "Repo: " + obj2.Items[i].Repo.Url + "\tPath: " + obj2.Items[i].Path
	}

	return finalResp
}

func constructFinalResponseCsv(sb string) string {

	var obj2 CodeResponse
	json.Unmarshal([]byte(sb), &obj2)

	// Initialize our string with the field headers.
	finalResp := "Repo,Path"
	// Loop over our Response and add to our array.
	for i := range obj2.Items {
		finalResp = finalResp + "\n" + obj2.Items[i].Repo.Url + "," + obj2.Items[i].Path
	}

	return finalResp
}
