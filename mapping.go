package main

import (
	"bytes"
	"net/http"
)

func RequestMapping(client *http.Client) {
	//client := &http.Client{}

	var URL string = "http://localhost:9200/dev_nanodemoxxx1_traffic_logs/kafka/_mapping"

	var jsonStr = []byte(`{"kafka": {"properties": {
    "srcInst": {
        "type": "text",
        "fielddata": true
     },
    "reason": {
        "type":"long",
    },
    "srcIp": {
        "type": "ip",
        "fielddata": true
    },
    "httpPath": {
        "type": "text",
        "fielddata": true
     },
     "nanoMode": {
         "type": "long",
     },
     "destInst": {
         "type": "text",
         "fielddata": true
     },
     "linkProps": {
         "type": "text",
         "fielddata": true
     },
     "destComp": {
         "type": "text",
         "fielddata": true
     },
     "srcComp": {
         "type": "text",
         "fielddata": true
     },
     "destIP": {
         "type": "ip",
         "fielddata": true
     }
    }}}`)

	req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	/*
		body, err := ioutil.ReadAll(resp.Body)
		jsonParsed, err := gabs.ParseJSON(body)
		if err != nil {
			panic(err)
		}
		//value, ok := jsonParsed.Path("acknowledged").Data().(string)
	*/
}
