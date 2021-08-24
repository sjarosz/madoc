# MaDOC

MaDOC or "My Doctor" is a sample application developed by Sqoop Data to demo and test various Identity & Access Management use cases.

## How to Use?

To summarize, the following is available within MaDOC:

* Three different types of users - `ADMIN`, `PATIENT`, `DOCTOR`;
* Create Staff Users or Register Patients;
* Staff Can Create Appointments;
* Doctors Can Create Health Records;

The application publishes a REST API with the following API endpoints:

**`/users`** [Actions Available: GET, POST]

**`/users/{userID}`** [Actions Available: GET, PUT]

**`/appointments`** [Actions Available: GET, POST]

**`/appointments/{userID}`** [Actions Available: GET]

**`/appointments?username={username}&apptId={apptId}`** [Actions Available: GET]

**`/healthrecords`** [Actions Available: GET, POST]

**`/healthrecords/{healthRecordId}`** [Actions Available: GET]

Please note MaDOC is a Dockerized application, and therefore, Docker is a prerequisite.

Run `docker-compose up --build` to bring the server up and running. 

## License

[Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/)

