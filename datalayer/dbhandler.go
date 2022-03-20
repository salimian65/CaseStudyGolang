package datalayer

import (
	"fmt"
	"strings"
	"test01/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type SQLHandler struct {
	db *gorm.DB
}

func CreateDBConnection(connString string) (*SQLHandler, error) {

	db, err := gorm.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Promotion{})
	return &SQLHandler{
		db: db,
	}, nil
}

func (handler *SQLHandler) GetPromotions() ([]models.Promotion, error) {

	rows, err := handler.db.DB().Query("select * from Promotions")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	promotions := []models.Promotion{}

	for rows.Next() {
		p := models.Promotion{}
		err = rows.Scan(&p.Id, &p.Price, &p.Time)
		if err != nil {
			return nil, err
		}
		promotions = append(promotions, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return promotions, nil
}

func (handler *SQLHandler) GetPromitionById(id string) (models.Promotion, error) {
	row := handler.db.DB().QueryRow("select * from promotions where Id=?", id)
	p := models.Promotion{}
	err := row.Scan(&p.Id, &p.Price, &p.Time)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (handler *SQLHandler) AddPromotion(promotion models.Promotion) error {
	_, err := handler.db.DB().Exec("INSERT INTO  promotions VALUES (?,?,?)", promotion.Id, promotion.Price, promotion.Time)
	return err
}

func (handler *SQLHandler) UpdatePromotion(promotion models.Promotion) error {
	_, err := handler.db.DB().Exec("update promotions set 'price'=? , time=? where 'Id'=?", promotion.Price, promotion.Time, promotion.Id)
	return err
}

func (handler *SQLHandler) DeleteAllPromotions() error {
	_, err := handler.db.DB().Exec("delete from promotions")
	return err
}

func (handler *SQLHandler) DeletePromotion(promotion models.Promotion) error {
	_, err := handler.db.DB().Exec("delete from promotions where Id=?", promotion.Id)
	return err
}

func (handler *SQLHandler) TruncatePromotions() error {
	_, err := handler.db.DB().Exec("TRUNCATE table promotions")
	return err
}

func (handler *SQLHandler) BulkInsert(size int, promotions []models.Promotion) error {
	tx := handler.db.Begin()
	chunkList := chunk(promotions, size)
	for _, chunk := range chunkList {
		valueStrings := []string{}
		valueArgs := []interface{}{}
		for _, promotion := range chunk {
			valueStrings = append(valueStrings, "(?, ?, ?)")
			valueArgs = append(valueArgs, promotion.Id)
			valueArgs = append(valueArgs, promotion.Price)
			valueArgs = append(valueArgs, promotion.Time)
		}

		stmt := fmt.Sprintf("INSERT INTO `mydb`.`promotions` (`Id`,`Price`,`Time`) VALUES %s", strings.Join(valueStrings, ","))
		err := tx.Exec(stmt, valueArgs...).Error
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
		}
	}
	err := tx.Commit().Error
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func chunk(list []models.Promotion, size int) [][]models.Promotion {

	var divided [][]models.Promotion

	chunkSize := (len(list) + size - 1) / size

	for i := 0; i < len(list); i += chunkSize {
		end := i + chunkSize

		if end > len(list) {
			end = len(list)
		}

		divided = append(divided, list[i:end])
	}
	return divided
}
