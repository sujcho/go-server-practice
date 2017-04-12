package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

type Bucket struct {
	// db tag lets you specify the column name if it differs from the struct field
	Key      string `db:"key"`
	DocCount int    `db:"doc_count"`
	Created  int64
	Id       int64 `db:"id"`
}

func newBucket(key string, docCount int) Bucket {
	return Bucket{
		Created:  time.Now().UnixNano(),
		Key:      key,
		DocCount: docCount,
	}
}

func PutToDB(data []byte, TableName string) []Bucket {
	// initialize the DbMap
	dbmap := initDb(TableName)
	defer dbmap.Db.Close()

	// delete any existing rows
	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	entries := ParseBucket(data)
	println(entries)
	for _, entry := range entries {
		p := newBucket(entry.Key, entry.Count)
		// insert rows - auto increment PKs will be set properly after the insert
		err = dbmap.Insert(&p)
		checkErr(err, "Insert failed")
	}

	// fetch all rows
	var buckets []Bucket
	var str = "select * from " + TableName + " order by id"
	_, err = dbmap.Select(&buckets, str)
	checkErr(err, "Select failed")

	return buckets

}

func initDb(TableName string) *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "./myDb.db")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Bucket{}, TableName).SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
