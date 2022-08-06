# MaDOC

MaDOC or "My Doctor" is a sample application developed by Sqoop Data to demo and test various Identity & Access Management use cases.

## How to build and start
docker compose build
docker compose up

## To start in Kubernetes
skaffold run -f skaffold-pg.yaml; skaffold dev -f skaffold.yaml --default-repo <your-repo> 


## How to Use?

To summarize, the following is available within MaDOC:

* Three different types of users - `ADMIN`, `PATIENT`, `DOCTOR`;
* CRU Users;
* CRU Appointments;
* CRU Health Records;

## Endpoints

| Endpoint               | Query Params | Action | Description                    |
|------------------------|--------------|--------|--------------------------------|
| /                      | -            | GET    | Index                          |
| /users/{id}            | -            | GET    | Get User by ID                 |
| /users                 | -            | GET    | Get All Users                  |
| /users                 | -            | POST   | Create User                    |
| /users/{id}            | -            | PUT    | Update User                    |
| /appointments          | -            | POST   | Create Appointment             |
| /appointments          | username     | GET    | Get All Appointments           |
| /appointments/{apptId} | -            | GET    | Get Appointment by ID          |
| /appointments          | -            | PUT    | Update Appointment             |
| /healthrecords         | -            | POST   | Create Health Record           |
| /healthrecords         | username     | GET    | Get Health Records for Patient |
| /healthrecords/{hrId}  | -            | PUT    | Update Health Record           |
| /metrics               | -            | GET    | Get Prometheus Monitoring Data |

## License

[Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/)

