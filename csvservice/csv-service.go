package csvservice

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"test01/datalayer"
	"test01/models"
)

func ProcessCsvFile(filePath string, db datalayer.SQLHandler) error {
	csvFile, err := os.Open("promotions.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	//-------------------------------------------------------------------

	//-------------------------------------------------------------------
	csvLines := csv.NewReader(csvFile)
	number := 0
	for {
		record, err := csvLines.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		number += 1
		priceFloat, err := strconv.ParseFloat(record[1], 64)

		promotion := models.Promotion{
			Id:    record[0],
			Price: priceFloat,
			Time:  record[2],
		}

		err = db.AddPromotion(promotion)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(promotion.Id + " " + fmt.Sprintf("%f", promotion.Price) + " " + promotion.Time)

		if number == 5 {
			break
		}
	}

	return nil
}
