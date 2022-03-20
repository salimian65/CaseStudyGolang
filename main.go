package main

import (
	"fmt"
	"log"
	"test01/csvservice"
	"test01/datalayer"
	"test01/restapi"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jasonlvhit/gocron"
)

func main() {
	fmt.Println("Strating the CaseStudyGolang")

	connectionString := "root:S@limian65@tcp(127.0.0.1:3306)/mydb"
	webApiPort := ":8384"
	sizeForChunking := 1000 // it can be 1000 for inhance the speed of insert data in mysql
	csvFileName := "promotions.csv"

	//------------------------------------------------------------------
	db, err := datalayer.CreateDBConnection(connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	//-------------------------------------------------------------------
	ProcessCsvFile(*db, csvFileName, sizeForChunking)
	//-------------------------------------------------------------------
	gocron.Every(30).Minutes().Do(ProcessCsvFile, *db, csvFileName, sizeForChunking)
	//-------------------------------------------------------------------
	go func() {
		restapi.RunApi(webApiPort, *db)
	}()
	//-------------------------------------------------------------------
	<-gocron.Start()
}

func ProcessCsvFile(db datalayer.SQLHandler, csvFileName string, sizeForChunking int) {
	fmt.Println("state Processing the the data form csv file")
	err := db.TruncatePromotions()
	if err != nil {
		log.Fatalln(err)
	}

	err = csvservice.ProcessCsvFile(db, csvFileName, sizeForChunking)
	if err != nil {
		log.Fatalln(err)
	}
}
