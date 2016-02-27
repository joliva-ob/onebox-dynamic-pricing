package dataservice



import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"

)



// ELK Sale struct
type SaleElkType struct{

	Code int `json:"code"`
	Token int `json:"token"`
	Products []*ProductElkType `json:"products"`
}


// ELK Product struct
type ProductElkType struct {

	EventId int `json:"eventId"`
	SessionId int `json:"sessionId"`
}


/**
 *
 * Public function to retrieve needed sales details for handle with dynamic
 * pricing processes
 *
 * http://go-database-sql.org/accessing.html
 */
func GetSales(date_from string, date_to string, page int, config configuration.Config) []*SaleElkType {

	var sales []*SaleElkType


	return sales
}