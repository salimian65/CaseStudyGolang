package csvservice

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"test01/datalayer"
	"test01/models"
	"time"
)

func ProcessCsvFile(db datalayer.SQLHandler, filePath string, sizeForChunking int) error {
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	//-------------------------------------------------------------------

	//-------------------------------------------------------------------
	csvrecords := csv.NewReader(csvFile)
	csvLines, err := csvrecords.ReadAll()
	fmt.Println(len(csvLines))
	var promotions []models.Promotion
	for _, line := range csvLines {
		priceFloat, err := strconv.ParseFloat(line[1], 64)
		promotion := models.Promotion{
			Id:    line[0],
			Price: priceFloat,
			Time:  line[2],
		}

		promotions = append(promotions, promotion)
		if err != nil {
			panic(err.Error())
		}

		// if i == 100 {
		// 	break
		// }
	}
	start := time.Now()
	err = db.BulkInsert(sizeForChunking, promotions)
	timeTrack(start, "BulkInsert")
	return err
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
