package datalayer

import (
	"database/sql"
	"test01/models"

	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	db *sql.DB
}

func CreateDBConnection(connString string) (*SQLHandler, error) {

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	return &SQLHandler{
		db: db,
	}, nil
}

func (handler *SQLHandler) GetPromotions() ([]models.Promotion, error) {

	rows, err := handler.db.Query("select * from Promotions")
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
	row := handler.db.QueryRow("select * from promotions where Id=?", id)
	p := models.Promotion{}
	err := row.Scan(&p.Id, &p.Price, &p.Time)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (handler *SQLHandler) AddPromotion(promotion models.Promotion) error {
	_, err := handler.db.Exec("INSERT INTO  promotions VALUES (?,?,?)", promotion.Id, promotion.Price, promotion.Time)
	return err
}

func (handler *SQLHandler) UpdatePromotion(promotion models.Promotion) error {
	_, err := handler.db.Exec("update promotions set 'price'=? , time=? where 'Id'=?", promotion.Price, promotion.Time, promotion.Id)
	return err
}

func (handler *SQLHandler) DeleteAllPromotions() error {
	_, err := handler.db.Exec("delete from promotions")
	return err
}

func (handler *SQLHandler) DeletePromotion(promotion models.Promotion) error {
	_, err := handler.db.Exec("delete from promotions where Id=?", promotion.Id)
	return err
}

func (handler *SQLHandler) TruncatePromotion() error {
	_, err := handler.db.Exec("TRUNCATE table promotions")
	return err
}
