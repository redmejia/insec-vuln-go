package models

type Register struct {
	Business string `json:"bus_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	BusinessID int    `json:"bus_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type BusinessInformation struct {
	BusinessID int    `json:"bus_id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Business   string `json:"bus_name"`
}

type Deal struct {
	DealID             int     `json:"deal_id"`
	BusinessID         int     `json:"bus_id"`
	Business           string  `json:"bus_name"`
	ProductName        string  `json:"pro_name"`
	ProductDescription string  `json:"pro_desc"`
	CreatedAt          string  `json:"created_at"`
	Price              float64 `json:"price"`
}

// all deals
type Deals struct {
	BusinessID         int     `json:"bus_id"`
	Email              string  `json:"email"`
	Business           string  `json:"bus_name"`
	ProductName        string  `json:"pro_name"`
	ProductDescription string  `json:"pro_desc"`
	CreatedAt          string  `json:"created_at"`
	Price              float64 `json:"price"`
}
type DealList struct {
	Deals []Deals `json:"deals"`
}
type DealsInformation struct {
	Deals []Deal `json:"deals"`
}
