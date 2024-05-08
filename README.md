# Employee CRUD

## APIs

### Create Employee
```bash
curl --location 'localhost:8000/employees' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Test1",
    "salary": 1000.00,
    "Position": "L1"
}'
```

### Get Employee
```bash
curl --location 'localhost:8000/employees/1'
```

### Update Employee
```bash
curl --location --request PUT 'localhost:8000/employees/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Test2",
    "salary": 1200.58,
    "Position": "L2"
}'
```

### Delete Employee
```bash
curl --location --request DELETE 'localhost:8000/employees/1'
```

### Get All Employees
```bash
curl --location 'localhost:8000/employees?pageNumber=1&pageSize=3'
```

**Note: pageNumber and pageSize are optional parameters with default values 1 and 10**


## How to run the application
```bash
$ cd employees
$ go run main.go
```

## How to run the tests
```bash
$ cd employees
$ go test ./...
```
