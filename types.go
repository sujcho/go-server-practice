package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Jeffail/gabs"
)

//type for input json from MyInput.json
type Input struct {
	Index  string                 `json:"_index"`
	Type   string                 `json:"_type"`
	Id     string                 `json:"_id"`
	Score  int64                  `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}
type Inputs []Input

type Entry struct {
	Key   string `json:"key"`
	Count int    `json:"doc_count"`
}

type Entries []Entry

func PrintInputs(inputs Inputs) {
	for _, input := range inputs {
		data, _ := json.Marshal(input.Source)
		fmt.Println(string(data))
	}
}

//returns go
func ReadInputs() Inputs {

	raw, err := ioutil.ReadFile("MyInputs.json")
	if err != nil {
		fmt.Print("not found")
	}

	var all Inputs
	json.Unmarshal(raw, &all)
	return all
}

func ParseAgg(raw []byte) []byte {

	jsonParsed, err := gabs.ParseJSON(raw)
	if err != nil {
		panic(err)
	}
	// S is shorthand for Search
	//	fmt.Println(jsonParsed.Path("aggregations.agg_bucket").String())

	return []byte(jsonParsed.Path("aggregations.agg_bucket").String())

}

func ParseBucket(data []byte) Entries {

	var listBucket Entries

	jsonParsed, _ := gabs.ParseJSON(data)
	children, _ := jsonParsed.S("buckets").Children()
	for _, child := range children {
		j, err := json.Marshal(child.Data())
		if err != nil {
			panic(err)
		}
		//	fmt.Println(child.Data())
		var entry Entry
		if err := json.Unmarshal(j, &entry); err != nil {
			panic(err)
		}
		listBucket = append(listBucket, entry)
	}
	return listBucket
}
