package main

import (
	"context"
	"encoding/json"

	"gopkg.in/olivere/elastic.v5"
)

func RequestBulkIndex() []*elastic.BulkResponseItem {

	//reads input set
	inputs := ReadInputs()

	//elastic
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200/"))
	if err != nil {
		panic(err)
	}

	bulkRequest := client.Bulk()
	bulkRequest = BuildIndexRequest(inputs, bulkRequest)

	// Do sends the bulk requests to Elasticsearch
	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		// ...
	}
	indexed := bulkResponse.Indexed()

	return indexed
}

//build multiple bulk requests
func BuildIndexRequest(entries Inputs, s *elastic.BulkService) *elastic.BulkService {
	for _, entry := range entries {
		req := WrapperBulkIndexRequest(entry.Index, entry.Type, entry.Id, entry.Source)
		s = s.Add(req)
	}
	return s
}

//!!!! handle score later!!
//wrapper index NewBulkIndexRequest
func WrapperBulkIndexRequest(theIndex string, theType string, theId string, theSource map[string]interface{}) *elastic.BulkIndexRequest {
	source, _ := json.Marshal(theSource)
	return elastic.NewBulkIndexRequest().Index(theIndex).Type(theType).Id(theId).Doc(string(source))
}
