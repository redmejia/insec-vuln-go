package database

import (
	"fmt"
	"time"

	"github.com/redmejia/models"
)

func (db *DbModel) RegisterNewUser(user models.Register) (bool, int) {

	var busId int

	tx, err := db.DB.Begin()
	if err != nil {
		db.ErrorLog.Fatal("tx : ", err)
		return false, 0
	}

	row := tx.QueryRow(`INSERT INTO vuln_register_users (bus_name, name, email)
						VALUES($1, $2, $3) RETURNING bus_id`, user.Business, user.Name, user.Email)

	err = row.Scan(&busId)
	if err != nil {
		db.ErrorLog.Fatal("error bus_id : ", err)
		return false, 0
	}

	_, err = tx.Exec(`INSERT INTO vuln_login_users (bus_id, email, password)
				VALUES($1, $2, $3)`, busId, user.Email, user.Password)

	if err != nil {
		db.ErrorLog.Println("exec :", err)
		tx.Rollback()
		return false, 0
	}

	err = tx.Commit()
	if err != nil {
		db.ErrorLog.Fatal("commit error : ", err)
		return false, 0
	}

	db.InfoLog.Println("user was created")
	return true, busId
}

func (db *DbModel) NewLogIn(user models.Login) models.Login {

	row := db.DB.QueryRow(`SELECT bus_id, email, password FROM vuln_login_users where email = $1`, user.Email)

	var userInfo models.Login
	err := row.Scan(&userInfo.BusinessID, &userInfo.Email, &userInfo.Password)
	if err != nil {
		db.ErrorLog.Println(err)
		return models.Login{}
	}

	return userInfo

}

func (db *DbModel) CreateNewDeal(deal models.Deal) bool {

	_, err := db.DB.Exec(`INSERT INTO vuln_users_new_deal (bus_id, bus_name, pro_name, pro_description, created_at, price)
							VALUES ($1, $2, $3, $4, $5, $6)`, deal.BusinessID, deal.Business, deal.ProductName, deal.ProductDescription, time.Now(), deal.Price)
	if err != nil {
		db.ErrorLog.Fatal(err)
		return false
	}

	db.InfoLog.Println("created")
	return true

}

func (db *DbModel) GetDealByBusinessName(busName string) *models.DealsInformation {
	// quote base BAD
	// sql = 'SELECT * FROM Users WHERE Name ="' + uName + '" AND Pass ="' + uPass + '"'
	query := fmt.Sprintf("SELECT * FROM vuln_users_new_deal WHERE bus_name = '%s'", busName)
	db.InfoLog.Println(query)
	// rows, err := db.DB.Query(`SELECT * FROM vuln_users_new_deal WHERE bus_name = $1`, busName)
	rows, err := db.DB.Query(query)

	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	var infoD models.DealsInformation

	for rows.Next() {

		var deal models.Deal

		rows.Scan(
			&deal.DealID,
			&deal.BusinessID,
			&deal.Business,
			&deal.ProductName,
			&deal.ProductDescription,
			&deal.CreatedAt,
			&deal.Price,
		)

		infoD.Deals = append(infoD.Deals, deal)
	}

	return &infoD // deal info

}

func (db *DbModel) GetBusinessInfoById(busID string) *models.BusinessInformation {

	txtSQL := "SELECT * FROM vuln_register_users WHERE bus_id = " + busID
	row := db.DB.QueryRow(txtSQL)

	var businessInfo models.BusinessInformation
	err := row.Scan(
		&businessInfo.BusinessID,
		&businessInfo.Business, // business name
		&businessInfo.Name,
		&businessInfo.Email,
	)

	if err != nil {
		db.ErrorLog.Fatal("this ", err)
		return nil
	}

	return &businessInfo

}

func (db *DbModel) UpdateBusinessInfo(business *models.BusinessInformation) bool {

	tx, err := db.DB.Begin()
	if err != nil {
		db.ErrorLog.Fatal(err)
		return false
	}

	_, err = tx.Exec(`UPDATE vuln_register_users
				SET email = $1
				WHERE bus_id = $2
	`, business.Email, business.BusinessID)

	if err != nil {
		db.ErrorLog.Fatal(err)
		tx.Rollback()
	}

	_, err = tx.Exec(`UPDATE vuln_login_users
					SET email = $1
				    	WHERE bus_id = $2
	`, business.Email, business.BusinessID)

	if err != nil {
		db.ErrorLog.Fatal(err)
		tx.Rollback()
	}

	err = tx.Commit()
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	return true

}

func (db *DbModel) GetDealByDealID(dealId int) *models.Deal {

	row := db.DB.QueryRow(`SELECT 
			pro_name, 
			pro_description, 
			price 
		FROM vuln_users_new_deal 
		WHERE deal_id = $1`, dealId)

	var deal models.Deal

	err := row.Scan(
		&deal.ProductName,
		&deal.ProductDescription,
		&deal.Price,
	)

	if err != nil {
		db.ErrorLog.Fatal(err)
		return nil
	}

	return &deal

}

func (db *DbModel) GetAllDeals() models.DealList {

	rows, err := db.DB.Query(`SELECT 
	ru.bus_id,
	ru.email,
	d.bus_name,
	d.pro_name,
	d.pro_description,
	d.created_at, 
	d.price
FROM vuln_register_users AS ru
JOIN vuln_users_new_deal AS d ON ru.bus_id = d.bus_id`)
	if err != nil {
		db.ErrorLog.Fatal(err)
		return models.DealList{}
	}

	var deals models.DealList

	for rows.Next() {
		var deal models.Deals
		rows.Scan(
			&deal.BusinessID,
			&deal.Email,
			&deal.Business,
			&deal.ProductName,
			&deal.ProductDescription,
			&deal.CreatedAt,
			&deal.Price,
		)
		deals.Deals = append(deals.Deals, deal)
	}

	return deals
}

func (db *DbModel) UpdateDeal(deal *models.Deal) bool {

	_, err := db.DB.Exec(`UPDATE vuln_users_new_deal
		SET pro_name = $1, pro_description = $2, price = $3
		WHERE deal_id = $4
	`,
		deal.ProductName,
		deal.ProductDescription,
		deal.Price,
		deal.DealID)

	if err != nil {
		db.ErrorLog.Fatal(err)
		return false
	}

	return true
}
