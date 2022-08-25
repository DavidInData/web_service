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

type Timeline struct {
    Id int
    ccvi_score float64
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


	rows, err := db.Query(`SELECT id, ccvi_score from ccvi_details where id = 1;`)
	if err != nil {
	    panic(err)
	}

	timeline := Timeline{}
	defer rows.Close()
	for rows.Next() {
	    
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


	// Body := list()
	fmt.Fprintf(w, "%s", timeline)
}



///////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////

func handler(w http.ResponseWriter, r *http.Request) {
	name := "The REPORTS"
	fmt.Fprintf(w, "CBI data collection microservices' goroutines have started for %s!\n", name)
	//fmt.Fprintf(w, timeline.Id)

}

/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////

