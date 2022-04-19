package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"log"
	"time"
)

type Origin struct {
	Cdate          time.Time
	DrcnoteCnoteNo string
	MrcnoteNo      string
	CnoteQty       int
}

type Destination struct {
	PRegional     string
	PBranch       string
	POrigin       string
	PRegionalDest string
	HawbNo        string
}

func main() {
	//_, err := gorm.Open(oracle.Open("system/SysPassword1@127.0.0.1:1521/XEPDB1"), &gorm.Config{})
	//db, err := sql.Open("godror", `user="jneapi" password="oracle" connectString="localhost:1521/XEPDB1"`)
	db, err := sql.Open("godror", `user="jneapi" password="oracle" connectString="10.16.2.116:1524/myorion"`)
	if err != nil {
		fmt.Println("Connection failed", err)
	} else {
		fmt.Println("")
		log.Println("Connection established")
		fmt.Println("")
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("select CDATE, DRCNOTE_CNOTE_NO, MRCNOTE_NO, CNOTE_QTY FROM REP_ORIGIN_215 where DRCNOTE_CNOTE_NO = '0125882100793836'")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {

		var origin Origin
		// for each row, scan the result into our tag composite object
		err = results.Scan(&origin.Cdate, &origin.DrcnoteCnoteNo, &origin.MrcnoteNo, &origin.CnoteQty)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		//log.Printf(origin.Drcn qoteCnoteNo)
		fmt.Println("DB 101 - Table REP_ORIGIN_215")
		fmt.Println("=============================")
		fmt.Println(origin.Cdate, origin.DrcnoteCnoteNo, origin.MrcnoteNo, origin.CnoteQty)
		fmt.Println("")

	}

	resultsDestination, err := db.Query("select P_REGIONAL, P_BRANCH, P_ORIGIN, P_REGIONAL_DEST, HAWB_NO from REP_DESTINATION_215 where HAWB_NO = '8000121632274'")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for resultsDestination.Next() {

		var destination Destination
		// for each row, scan the result into our tag composite object
		err = resultsDestination.Scan(&destination.PRegional, &destination.PBranch, &destination.POrigin, &destination.PRegionalDest, &destination.HawbNo)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		//log.Printf(origin.Drcn qoteCnoteNo)
		fmt.Println("DB 101 - Table REP_DESTINATION_215")
		fmt.Println("=============================")
		fmt.Println(destination.PRegional, destination.PBranch, destination.POrigin, destination.PRegionalDest, destination.HawbNo)
		fmt.Println("")
	}
}
