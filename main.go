package main

import (
	"fmt"
	"log"
	"test01/csvservice"
	"test01/datalayer"
	"test01/restapi"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("hi")

	//------------------------------------------------------------------
	db, err := datalayer.CreateDBConnection("root:S@limian65@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		log.Fatalln(err)
	}
	//-------------------------------------------------------------------
	err = db.TruncatePromotion()
	if err != nil {
		log.Fatalln(err)
	}
	//-------------------------------------------------------------------
	err = csvservice.ProcessCsvFile("promotions.csv", *db)
	if err != nil {
		log.Fatalln(err)
	}
	//-------------------------------------------------------------------
	restapi.RunApi(":8383", *db)
	//-------------------------------------------------------------------
}
