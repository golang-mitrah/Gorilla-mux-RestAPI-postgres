# Gorilla mux RestAPI postgres

REST API using Gorilla-mux in GO

1. Clone the git repo - `git clone git@github.com:golang-mitrah/Gorilla-mux-RestAPI-postgres.git`

2. Use `database/db_script.sql` file to crete a SQL table needed in your PostGres server for this demo project 
   
3. Change database connection details in `database/database.go` file's line #27
```
var dbname, username, password, host, port = "DB_XX", "USER_XX", "PWD_XX", "localhost", "5432"
```

4. Run below in your project path
   ```
   go get .
   go run .
   ```

5. Use the `Gorilla_mux_RestAPI_postgres.postman_collection.json` to import the API request into your [PostMan](https://www.postman.com/) tool to try out the API endpoints available

## **Contributors**

REST API using Gorilla MUX Router is authored by **[GoLang Mitrah](https://www.MitrahSoft.com/)** and everyone is welcome to contribute. 

## **Problems**

If you experience any problems with REST API using Gorilla MUX please:

* [submit a ticket to our issue tracker](https://github.com/golang-mitrah/Gorilla-mux-RestAPI-postgres/issues)
* fix the error yourself and send us a pull request

## **Social Media**

You'll find us on [Twitter](https://twitter.com/MitrahSoft) and [Facebook](http://www.facebook.com/MitrahSoft) 
