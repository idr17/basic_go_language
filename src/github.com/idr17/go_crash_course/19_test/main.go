package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type data struct {
	Items []struct {
		Project     string `json:"project"`
		Article     string `json:"article"`
		Granularity string `json:"granularity"`
		Timestamp   string `json:"timestamp"`
		Access      string `json:"access"`
		Agent       string `json:"agent"`
		Views       int    `json:"views"`
	} `json:"items"`
}

// type Posts []Post

func main() {

	//The name of the wikipedia post
	postName := "Smithsonian_Institution"

	//The frequency of
	period := "daily"

	//When to start the selection
	startDate := "20160101"

	//When to end the selection
	endDate := "20170101"

	url := fmt.Sprintf("https://wikimedia.org/api/rest_v1/metrics/pageviews/per-article/en.wikipedia.org/all-access/all-agents/%s/%s/%s/%s", postName, period, startDate, endDate)

	// //Get from URL
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var dataResp data
	json.Unmarshal(body, &dataResp)

	fmt.Println(dataResp.Items[0])

	// Loop over structs and display the respective views.
	for p := range dataResp.Items {
		fmt.Printf("Article = %v\n", dataResp.Items[p].Article)
	}

}
