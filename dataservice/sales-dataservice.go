package dataservice



import (

	"encoding/json"
	"strings"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

)

// Elk parent level
type OrderDocElkType struct {
	Doc OrderElkType `json:"doc"`
}

// ELK Order struct
type OrderElkType struct {

	Date DateElkType `json:"date"`
	Code string `json:"code"`
	Products []*ProductElkType `json:"products"`
	Price PriceElkType `json:"price"`
	OrderData OrderDataElkType `json:"orderData"`
}

type OrderDataElkType struct {
	ChannelType string `json:"channelType"`
	ChannelId int `json:"channelId"`
}

type PriceElkType struct {
	FinalPrice float32 `json:"finalPrice"`
	BasePrice float32 `json:"basePrice"`
}

type DateElkType struct {
	Purchased string `json:"purchased"`
}

// ELK Product struct
type ProductElkType struct {

	ProductId string `json:"id"`
	EventId int `json:"eventId"`
	SessionId int `json:"sessionId"`
	TicketData TicketDataElkType `json:"ticketData"`
}

type TicketDataElkType struct {

	PriceZoneId int `json:"priceZoneId"`
	PriceZoneName string `json:"priceZoneName"`
	SectorName string `json:"sectorName"`
	RowOrder int `json:"rowOrder"`
	NumSeat string `json:"numSeat"`
}


/**
 *
 * Public function to retrieve needed sales details for handle with dynamic
 * pricing processes
 *
 * http://go-database-sql.org/accessing.html
 */
func GetSales(dateFrom string, dateTo string, eventId int, saleId string, page int, uuid string, oauthtoken *Oauthtoken, pageSize int) []*OrderDocElkType {

	var sales []*OrderDocElkType
	args := make(map[string]interface{})
	args["size"] = pageSize
	from := page*pageSize
	args["from"] = from
	saleId = strings.ToLower(saleId)
	key := dateFrom + dateTo + strconv.Itoa(pageSize) + strconv.Itoa(from) + strconv.Itoa(eventId) + saleId + oauthtoken.UserName

	// Get the string associated with the key from the cache
	salesFromCache, found := salesCache.Get(key)
	if !found {

		// Get the query and fill placeholders properly
		query := GetQuery(dateFrom, dateTo, eventId, saleId, oauthtoken.UserName)

		// Elasticsearch Search
		out, err := elk_conn.Search(config.Sales_elk_index, "", args, query)
		if len(out.Hits.Hits) > 0 {

			for i := 0; i < out.Hits.Len(); i++ {

				sale := new(OrderDocElkType)
				json.Unmarshal(*out.Hits.Hits[i].Source, &sale)
				sales = append(sales, sale)
			}
		}
		if err != nil {
			log.Errorf("{%v} Error occurred while trying to retrieve elasticsearch sales: %v", uuid, err)
		}

		// Store the prices struct to cache
		salesCache.Set(key, sales, 0)

	} else {

		// Retrieve sales struct from cache
		sales = salesFromCache.([]*OrderDocElkType) // Cast interface{} retrieved from cache to []*PriceType
	}

	return sales
}



/*
 Get the correct query from configuration
 depending on the Url params
 eventId = -1 means there is no event id requested
  */
func GetQuery(dateFrom string, dateTo string, eventId int, saleId string, username string)  string {

	var query string
	_, restrictions := GetRestrictions( username, false )

	if eventId > 0 {

		if restrictions {
			query = config.Sales_elk_filter_event_restricted
		} else {
			query = config.Sales_elk_filter_event
			query = strings.Replace(query, EVENT_ID, strconv.Itoa(eventId), 1)
		}

	} else if saleId != "" {

		if restrictions {
			query = config.Sales_elk_filter_sale_restricted
		} else {
			query = config.Sales_elk_filter_sale
		}
		query = strings.Replace(query, SALE_ID, saleId, 1)

	} else {

		if restrictions {
			query = config.Sales_elk_filter_dates_restricted
		} else {
			query = config.Sales_elk_filter_dates
		}
		query = strings.Replace(query,START_DATE,dateFrom,1)
		query = strings.Replace(query,END_DATE,dateTo,1)
		log.Debugf("Query ELK: %v", query)
	}

	return query
}