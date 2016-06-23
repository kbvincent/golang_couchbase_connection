package main

import (
	"fmt"
	"github.com/kbvincent/couchbaseTest/db"

	"github.com/couchbase/gocb"
	"sync"
)

var doThisOnce = func() {
	// connection to pong
	var database = db.ConnectDb()

	// getting specific document from database
	assetID := "test"
	var res map[string]interface{}
	_,docErr := database.Get(assetID, &res)
	if docErr != nil {
		fmt.Println("error occurred")
		fmt.Println(docErr)
	} else {
		fmt.Println(" ")
		fmt.Println("Single Document Retreived: ")
		fmt.Println(res["name"])
		fmt.Println(" ")
	}

	myViewQuery := gocb.NewViewQuery("listall", "listall")
	rows, errs := database.ExecuteViewQuery(myViewQuery)
	if errs != nil {
		fmt.Println("listall error")
		fmt.Println(errs)
	} else {
		fmt.Println(" ")
		fmt.Println("Design Doc Query Documents Retreived: ")
		fmt.Println(rows)
		fmt.Println(" ")
	}

	// n1ql query running select *
	myN1qlQuery := gocb.NewN1qlQuery("SELECT * FROM kvincent")
	rows, err := database.ExecuteN1qlQuery(myN1qlQuery, nil)
	if err != nil {
		fmt.Println("n1ql error:")
		fmt.Println(err)
	} else {
		fmt.Println(" ")
		fmt.Println("N1Ql Query Documents Retreived: ")
		fmt.Println(rows)
		fmt.Println(" ")
	}
}

func main() {
	var once sync.Once
	once.Do(doThisOnce) //do stuff in here only once ever
}