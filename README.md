## Requirements
1. Follow the setup instructions at [mtgdataload](https://github.com/JakeDodd/mtgdataload) to setup the database
2. Clone this repo and create a .env file in the root directory
```
PORT=8000
pg_host=localhost
pg_port=5432
pg_user=<FILL>
pg_password=<FILL>
pg_dbname=<FILL>
```
3. Install Air
```
go install github.com/air-verse/air@latest
```
4. Run the application
    1. It is configured so that both the go and react portions of the application will live reload: if any changes are made to source files, the whole app rebuilds and restarts.
```
sh dev.sh
```

