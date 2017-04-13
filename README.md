# go server works with elastic search

## Requirement
1. Reads data from ElasticSearch summarizes the data and puts it in a different table in the database.
2. Based on HTTP GET request of how many records and destination are required - send a reply from the summary table of the number of logs.

## Tools, Languages, Libraries used
- Go:  1.8.1
- Elastic Search: 5.3.0
- Libraries:
  - "gopkg.in/olivere/elastic.v5" : Go client for Elastic Search.
    **(Used only for bulk request, Just tried out this library and didn't used it later on.)**
  - "github.com/gorilla/mux" : URL router
  - "github.com/Jeffail/gabs": Easy Json Builder.
  

## How to run & What it does
1. Run Elastic Search at localhost:9200 (default address)
2. Place the input file in the same folder (with Go executable).
3. Execute the program.
4. Go to localhost:8080.
  - at  / : It will automatically load inputs, send a INDEX request to Elastic Search. 
        Them, itt will show a list of id of Indexed documents.
  - at /aggr : It will perform aggregation of some fields over indexed documents and save results of each field into different tables in DB.
**This program performs [Term Aggregation](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-terms-aggregation.html) over some fields.**
 
 
 ## Source  code
 - main.go : main.
 - router.go : Defining routes for my server (/ and /aggr).
 - handler.go : Handlers for routes for my server.
 - bulkindex.go : Bulk indexing multiple documents.
 - mapping.go : Defining types of fields for inpupts.
 - aggregation.go : Perfoming term affregation over some field and save results to data base
 - database.go : Initializing and defining database structure
 - types.go : defining all type struct and related actions. Parsing json.
 
 
 


