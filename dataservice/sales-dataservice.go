package dataservice



import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"

)



// ELK Sale struct
type SaleType struct{

	Id int `json:"id"`
	Event_id int `json:"event_id"`
	Event_name string `json:"event_name"`
	Event_date string `json:"event_date"`
	Transaction_date string `json:"transaction_date"`
	Transaction_type string `json:"transaction_type"`
	Buyer_type_code string `json:"buyer_type_code"`
	Products_number int `json:"products_number"`
	Channel_id string `json:"channel_id"`
	Products []*ProductType `json:"products"`

}


// ELK Product struct
type ProductType struct {

	Id int `json:"id"`
	Session_id int `json:"session_id"`
}


/**
 *
 * Public function to retrieve needed sales details for handle with dynamic
 * pricing processes
 *
 * http://go-database-sql.org/accessing.html
 */
func GetSales(date_from string, date_to string, page int, config configuration.Config) []*SaleType {

	var sales []*SaleType


	return sales
}