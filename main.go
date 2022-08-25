package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kelvins/geocoder"
	_ "github.com/lib/pq"
)

// Declare my database connection
var db *sql.DB

func init() {
	var err error

	fmt.Println("Initializing the DB connection")

	//Option 4
	//Database application running on Google Cloud Platform. GCP will build, containerize and run.
	db_connection := "user=postgres dbname=chicago_business_intelligence password=root host=/cloudsql/citric-dream-359919:us-central1:mypostgres sslmode=disable port = 5432"

	db, err = sql.Open("postgres", db_connection)
	if err != nil {
		log.Fatal(fmt.Println("Couldn't Open Connection to database"))
		panic(err)
	}

}


///////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////

func main() {

	

	log.Print("starting CBI Microservices ...")

	rows, err := db.Query(`SELECT id from ccvi_id where id = 1;`)
	if err != nil {
	    panic(err)
	}

	type Timeline struct {
    Id int
    ccvi_score float64

	defer rows.Close()
	for rows.Next() {
	    timeline := Timeline{}
	    err = rows.Scan(&timeline.Id, &timeline.ccvi_score)
	    if err != nil {
	        panic(err)
	    }
	    fmt.Println(timeline)
	}
	err = rows.Err()
	if err != nil {
	    panic(err)
	}


	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	log.Print("Navigate to Cloud Run services and find the URL of your service")
	log.Print("Use the browser and navigate to your service URL to to check your service has started")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
		

	}

}


///////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("PROJECT_ID")
	if name == "" {
		name = "CBI-Project"
	}

	fmt.Fprintf(w, "CBI data collection microservices' goroutines have started for %s!\n", name)
}

/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////

