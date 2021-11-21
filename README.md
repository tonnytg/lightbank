# lightbank


Database Config:

Create a file database.env and add the following lines:

    export DB_HOST="127.0.0.1"
    export DB_PORT="5432"
    export DB_USER="postgres"
    export DB_PASSWORD="postgres"
    export DB_NAME="lightbank"
    export DB_SSLMODE="disable"


Export the configuration, and run the following command:

    go run main.go create-db

This will run migrations and create the database.
To run API of lightbank, run without arguments:

    go run main.go

after you can register your first user:

    curl -X POST -H "Content-Type: application/json" -d '{"username":"admin","email":"admin@lightbank.com","password":"admin"}' http://localhost:8888/register

Output:

    {
        "data": {
            "ID": 7,
            "Username": "admin",
            "Email": "admin@lightbank.com",
            "Accounts": [
                {
                    "ID": 7,
                    "Name": "admin's account",
                    "Balance": 0
                }
            ]
        },
        "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2Mzc1MjIwMzIsInVzZXJfaWQiOjd9.BWXKC0PqpyRcIO0C0aDAkYttqgSD6oIBfdFpiLO_pR0",
        "message": "all is fine"
    }



Calls:

- Register:

      http://localhost:8888/register

- Login:
Get jwt to authenticated user:

      curl -X POST -H "Content-Type: application/json" -d '{"username":"admin","password":"admin"}' http://localhost:8888/login

- Find:

      curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2Mzc1MjIwMzIsInVzZXJfaWQiOjd9.BWXKC0PqpyRcIO0C0aDAkYttqgSD6oIBfdFpiLO_pR0" http://localhost:8888/accounts/1

- Transaction:

    curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2Mzc1MjIwMzIsInVzZXJfaWQiOjd9.BWXKC0PqpyRcIO0C0aDAkYttqgSD6oIBfdFpiLO_pR0" -H "Content-Type: application/json" -d '{"UserID": 1,"From":1,"To":2,"amount":10}' http://localhost:8888/transaction