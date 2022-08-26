Requirement 1 & 2 - List of Every Microservice

	Miroservices: 
		- Retrieve/store Community Area Unemployment Data
		- Retreive/store Building Permits
		- Retreive/store Taxi Trips
		- Retreive/store Covid Details
		- Retreive/store CCVI Details

		-


Option 1. We need to execute the go app within the local machine.
	--> Needed to install PostgreSQL on local machine.
	--> Go Application will store Chicago Data into PostgreSQL database in local machine

Option 2. Build Image & Containerize Go App and PostgreSQL. Build Image & Excecute from Docker.
	--> Prereq: Docker installed
	--> docker-compose up
	--> Uses docker-compompse up  (uses YAML to build image of both database and application)
	--> Docker ps  (view on-going containers/microservies)

Option 3: Go Application and PostGreSQL image/container in Google Cloud Platform
	--> Prereq: GCP Account. GitHub Repo
	--> Upload source code to GitHub
	--> Google Cloud Trigger to retrieve updated source code
	--> Google Cloud will build and run images

