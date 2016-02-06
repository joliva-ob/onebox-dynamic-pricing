package dataservice

import (
	"flag"
	"fmt"
	"github.com/mattbaird/elastigo/lib"
	"encoding/json"
	"log"
	"os"
)

var (
	host *string = flag.String("host", "192.168.99.100", "Elasticsearch Host")
)

func testelastic() {
	c := elastigo.NewConn()
	log.SetFlags(log.LstdFlags)
	flag.Parse()

	// Trace all requests
/*	c.RequestTracer = func(method, url, body string) {
		log.Printf("Requesting %s %s", method, url)
		log.Printf("Request body: %s", body)
	}
*/
	fmt.Println("host = ", *host)
	// Set the Elasticsearch Host to Connect to
	c.Domain = *host

	// Index a document
	_, err := c.Index("testindex", "user", "docid_1", nil, `{"name":"bob"}`)
	exitIfErr(err)

	// Index a doc using a map of values
	_, err = c.Index("testindex", "user", "docid_2", nil, map[string]string{"name": "venkatesh"})
	exitIfErr(err)

	// Index a doc using Structs
	_, err = c.Index("testindex", "user", "docid_3", nil, MyUser{"wanda", 22})
	exitIfErr(err)

	// Search Using Raw json String
	searchJson := `{
	    "query" : {
	        "term" : { "Name" : "wanda" }
	    }
	}`

	// Elasticsearch Search
	out, err := c.Search("testindex", "user", nil, searchJson)
	if len(out.Hits.Hits) == 1	{
		// PArsing response search
		var myUser MyUser
		err := json.Unmarshal(*out.Hits.Hits[0].Source, &myUser)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", myUser)
		fmt.Printf("name: %v\n", myUser.Name)
		fmt.Printf("age: %v\n", myUser.Age)
		//fmt.Println("%v", out.Hits.Hits[0].Source)
	}
	exitIfErr(err)

}
func exitIfErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}

type MyUser struct {
	Name string
	Age  int
}
