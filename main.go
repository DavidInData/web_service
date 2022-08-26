package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Covid_results struct {
    week string
    zipcode string
    cases string
    percent string
}

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
	PORT := "8081"
	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/ccvi", handler1)
	//http.HandleFunc("/ccvi", handler2)
	//http.HandleFunc("/ccvi", handler3)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	log.Print("Navigate to Cloud Run services and find the URL of your service")
	log.Print("Use the browser and navigate to your service URL to to check your service has started")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)


	fmt.Println("Ready to serve at", PORT)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
		

	}

}





func handler1(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)


	rows, err := db.Query(`select to_date(LEFT(week_end,10), 'YYYY-MM-DD'), zip_code, cases_weekly, percent_tested_positive_weekly
		from covid_details
		where to_date(LEFT(week_end,10), 'YYYY-MM-DD') > now() - interval '7 day';`)
	if err != nil {
	    panic(err)
	}

	covid_results := Covid_results{}
	defer rows.Close()
	for rows.Next() {
	    
	    err = rows.Scan(&covid_results.week, &covid_results.zipcode, &covid_results.cases, &covid_results.percent)
	    if err != nil {
	        panic(err)
	    }
	    fmt.Println(covid_results)
	}
	err = rows.Err()
	if err != nil {
	    panic(err)
	}

	// Body := list()
	fmt.Fprintf(w, "Week of: %s\n Zip Code: %s\n Cases Weekly: %s\n Precent Tested Weekly %d\n: ", 
		covid_results.week, covid_results.zipcode, covid_results.cases, covid_results.percent)
}



///////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////

func handler(w http.ResponseWriter, r *http.Request) {
	name := "The REPORTS"
	fmt.Fprintf(w, "MSDS 432 - Foundations of Data Engineering.. feel free to run %s!\n", name)
	//fmt.Fprintf(w, timeline.Id)

}

/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////

