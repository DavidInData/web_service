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

type Ccvi_results struct {
    community_area_or_zip string
    community_area_name string
    ccvi_score string
}

type Waive_results struct {
    permit_id string
    permit_code string
    community_area string
}

type Loan_results struct {
    permit_id string
    permit_code string
    permit_type string
    income string
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
	http.HandleFunc("/covid", handler2)
	http.HandleFunc("/waive", handler3)
	http.HandleFunc("/loan", handler4)

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

	fmt.Fprintf(w, "WEEKLY COVID CASES REPORT\n")
	fmt.Fprintf(w, "\n")
	defer rows.Close()
	for rows.Next() {
	    
	    err = rows.Scan(&covid_results.week, &covid_results.zipcode, &covid_results.cases, &covid_results.percent)
	    if err != nil {
	        panic(err)
	    }
	    fmt.Fprintf(w, "Week of: %s\n Zip Code: %s\n Cases Weekly: %s\n Precent Tested Weekly %s\n", 
	    	covid_results.week, covid_results.zipcode, covid_results.cases, covid_results.percent)
	    fmt.Fprintf(w, "\n")
	    fmt.Fprintf(w, "\n")
	}
	err = rows.Err()
	if err != nil {
	    panic(err)
	}

	// Body := list()
	
}



func handler2(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)


	rows, err := db.Query(`select community_area_or_zip, community_area_name, ccvi_score
		from ccvi_details
		where ccvi_category = 'HIGH';`)
	if err != nil {
	    panic(err)
	}

	ccvi_results := Ccvi_results{}

	fmt.Fprintf(w, "WEEKLY HIGH CCVI REPORT\n")
	fmt.Fprintf(w, "\n")
	defer rows.Close()
	for rows.Next() {
	    
	    err = rows.Scan(&ccvi_results.community_area_or_zip, &ccvi_results.community_area_name, &ccvi_results.ccvi_score)
	    if err != nil {
	        panic(err)
	    }
	    fmt.Fprintf(w, "Community Area/Zip: %s\n Community Area Name: %s\n CCVI Score: %s\n", 
	    	ccvi_results.community_area_or_zip, ccvi_results.community_area_name, ccvi_results.ccvi_score)
	    fmt.Fprintf(w, "\n")
	    fmt.Fprintf(w, "\n")
	}
	err = rows.Err()
	if err != nil {
	    panic(err)
	}

	// Body := list()
	
}


func handler3(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)


	rows, err := db.Query(`select permit_id, permit_code, community_area from building_permits
		where community_area in (select community_area from community_area_unemployment
		order by unemployment DESC
		LIMIT 5)
		or community_area in (select community_area from community_area_unemployment
		order by below_poverty_level DESC
		LIMIT 5);`)
	if err != nil {
	    panic(err)
	}

	waive_results := Waive_results{}

	fmt.Fprintf(w, "WAIVE BUIDLING PERMITS WITH HIGH POVERTY AND/OR HIGH UNEMPLOYMENT REPORT\n")
	fmt.Fprintf(w, "\n")
	defer rows.Close()
	for rows.Next() {
	    
	    err = rows.Scan(&waive_results.permit_id, &waive_results.permit_code, &waive_results.community_area)
	    if err != nil {
	        panic(err)
	    }
	    fmt.Fprintf(w, "Permit ID: %s\n Permit Code: %s\n Community Area: %s\n", 
	    	waive_results.permit_id, waive_results.permit_code, waive_results.community_area)
	    fmt.Fprintf(w, "\n")
	    fmt.Fprintf(w, "\n")
	}
	err = rows.Err()
	if err != nil {
	    panic(err)
	}

	// Body := list()
	
}

func handler4(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)


	rows, err := db.Query(`select bp.permit_id, bp.permit_code, bp.permit_type, cau.per_capita_income from building_permits bp, community_area_unemployment cau
		where bp.community_area = cau.community_area
		and bp.permit_type = 'PERMIT - NEW CONSTRUCTION'
		and cast (per_capita_income as integer) < 30000;`)
	if err != nil {
	    panic(err)
	}

	loan_results := Loan_results{}

	fmt.Fprintf(w, "SMALL BUSINESS THAT QUALITY FOR SPECIAL LOAN\n")
	fmt.Fprintf(w, "\n")
	defer rows.Close()
	for rows.Next() {
	    
	    err = rows.Scan(&loan_results.permit_id, &loan_results.permit_code, &loan_results.permit_type, &loan_results.income)
	    if err != nil {
	        panic(err)
	    }
	    fmt.Fprintf(w, "Permit ID: %s\n Permit Code: %s\n Permit Type: %s\n Per Capita Income: %s\n", 
	    	loan_results.permit_id, loan_results.permit_code, loan_results.community_area, loan_results.income)
	    fmt.Fprintf(w, "\n")
	    fmt.Fprintf(w, "\n")
	}
	err = rows.Err()
	if err != nil {
	    panic(err)
	}

	// Body := list()
	
}


///////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////

func handler(w http.ResponseWriter, r *http.Request) {
	name := "reports"
	fmt.Fprintf(w, "MSDS 432 - Foundations of Data Engineering!!\n")
	fmt.Fprintf(w, "Run the %s!", name)
	//fmt.Fprintf(w, timeline.Id)

}

/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////

