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
	assetID := "21st_amendment_brewery_cafe-21a_ipa"
	var res map[string]interface{}
	_,docErr := database.Get(assetID, &res)
	if docErr != nil {
		fmt.Println("error occurred")
		fmt.Println(docErr)
	} else {
		fmt.Println("Single Document Retreived: ")
		fmt.Println(fmt.Sprintf("Name: %s", res["name"]))
		fmt.Println(fmt.Sprintf("TYPE: %s", res["type"]))
		fmt.Println(fmt.Sprintf("ABV: %s", res["abv"]))
	}

	myViewQuery := gocb.NewViewQuery("beer", "brewery_beers")
	rows, errs := database.ExecuteViewQuery(myViewQuery)
	if errs != nil {
		fmt.Println("brewery_beers error")
		fmt.Println(errs)
	} else {
		fmt.Println(" ")
		fmt.Println("Design Doc Query Documents Retreived: ")
		fmt.Println(rows)
	}

	myN1qlQuery := gocb.NewN1qlQuery("SELECT name FROM `beer-sample`;")
	rows, err := database.ExecuteN1qlQuery(myN1qlQuery, nil)
	if err != nil {
		fmt.Println("n1ql error:")
		fmt.Println(err)
	} else {
		fmt.Println("N1Ql Query Documents Successfully Retreived")
		fmt.Println(rows)
	}
}

func main() {
	var once sync.Once
	once.Do(doThisOnce) //do stuff in here only once ever
}