# Time Tracker For Users

## Go Version 
This project was created with go version go1.22

## Local Deployment
1. Clone Repository.
2. Install swagger with the ```go install github.com/swaggo/swag/cmd/swag@latest``` command
3. Add to the .env file next variables:
```DB_PASSWORD=Qwerty12
DB_USER={your_db_user}
DB_NAME={your_db_name}
DB_HOST={your_db_host}
DB_PORT={your_db_port}
SERVER_HOST={your_server_host}
SERVER_PORT={your_server_port}
```

## Swagger
Swagger documentation is available at /swagger/index.html path