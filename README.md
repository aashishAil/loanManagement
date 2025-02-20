This is a loan management system, which allows the admin to manage the loans and the customers to apply for the loans. The system is built using Golang and Postgres as the database.
There are limits set for the customers loan amount, repayment and tenure. Right now the system only supports loans in the currency ```INR```

### Pre-requisites

1. Golang version`1.21.X`, you can download the same from here [Golang](https://go.dev/dl/)
2. Postgres version `16.X`, you can download the same from here [Postgres](https://www.postgresql.org/download/)
3. Install `go-migrate` for database migrations, you can download the same from here [go-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
4. - Additionally if you want to use it for development purpose and generate the mocks, you can install by following the instructions on this page [mockery](https://vektra.github.io/mockery/latest/)

### Setup
- Go through the values of `.env` file at the root of the folder and make sure you provide the correct POSTGRES database
  details, including the `credentials` and `database details`.
- Run the following command from the root of the project to execute the migrations, after replacing the credentials and database details:
  ```bash
  migrate -database "postgres://<username>:<password>@localhost:5432/<database_name>?sslmode=disable" -path database/migration up
  ```
- Run the following command from the root of the project to download all the dependencies:
  ```bash
  go mod download
  ```
  
### Run
- Run the following command from the root of the project to start the server:
  ```bash
  go run main.go
  ```

### API
- The Postman API documentation can be found at the following link: [API Documentation](https://api.postman.com/collections/1567444-6893f145-b89f-4f73-9f04-5907c2726be4)
- You'd need to setup the following environment variables in order to use the API documentation:
  - `baseUrl` - The base URL of the server, e.g. `http://localhost:3000`
  - `authHeader` - The name of the header for authentication `LM-AUTH`
  - `customerToken` - The auth token for the customer (Can be obtained via using the `Login` endpoint)
  - `adminToken` - The auth token for the admin (Can be obtained via using the `Login` endpoint)
- The migrations also sets up the following base users to get started, user the same:
  - Admin:
    - Email: `admin@lm.com`
    - Password: `admin@123`
  - Customer
    - Email: `customer@lm.com`
    - Password: `qwer1234`
  
### Generating Mocks
- Run the following command from the root of the project to generate the mocks:
  ```bash
  mockery --all
  ```

### Test
Run the following commands to execute the test
```bash
cd tests
go test ./...
```