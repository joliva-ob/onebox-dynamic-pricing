package dataservice



import (

	"encoding/json"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"

)


type OrderElkType struct {
	Doc SaleElkType
}

// ELK Sale struct
type SaleElkType struct{

	Code string `json:"code"`
	Token string `json:"token"`
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
	var orders []*OrderElkType
	args := make(map[string]interface{})
	args["size"] = config.Elasticsearch_limit_items
	from := page*config.Elasticsearch_limit_items
	args["from"] = from

	query := strings.Replace(config.Sales_elk_filter_eventId,"!eventId!","2627",-1)

	// Elasticsearch Search
	out, err := elk_conn.Search(config.Sales_elk_index, "", args, query)
	if len(out.Hits.Hits) > 0	{

		for i := 0; i < out.Hits.Len(); i++ {

			order := new(OrderElkType)
			err := json.Unmarshal(*out.Hits.Hits[i].Source, &order)
			if err != nil {
				log.Errorf("Error occurred while unmarshalling elasticsearch sale: %v", err)
			}
			orders = append(orders, order)
			log.Debugf("order code: %v\n", order.Doc.Code)
		}
	}
	if err != nil {
		log.Errorf("Error occurred while trying to retrieve elasticsearch sales: %v", err)
	}

	return sales
}

