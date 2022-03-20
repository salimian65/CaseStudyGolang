package csvservice

import (
	"encoding/csv"
	"fmt"
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
	csvLines2, err := csvLines.ReadAll()
	fmt.Println(len(csvLines2))
	var promotions []models.Promotion
	for _, line := range csvLines2 {
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

		// if i == 9 {
		// 	break
		// }
	}

	err = db.BenchmarkBulkCreate(500, promotions)

	//number := 0
	// for {
	// 	record, err := csvLines.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		return err
	// 	}
	// 	number += 1
	// 	priceFloat, err := strconv.ParseFloat(record[1], 64)

	// 	promotion := models.Promotion{
	// 		Id:    record[0],
	// 		Price: priceFloat,
	// 		Time:  record[2],
	// 	}

	// 	err = db.AddPromotion(promotion)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	fmt.Println(promotion.Id + " " + fmt.Sprintf("%f", promotion.Price) + " " + promotion.Time)

	// 	// if number == 5 {
	// 	// 	break
	// 	// }
	// }

	return nil
}
