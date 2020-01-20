# api-books
Project in go created for studies

1. Requirements:
    - go version go1.13
    - docker and docker-compose

2. Run project:
    - docker-compose up
    - go run main.go
    
3. Endpoints:
    - ```/books```  GET, POST
    - ```/books/{bookID}```  GET, PUT, DELETE
    
4. Json example:
````
{
	"name": "Water for Elephants",
	"author": "Sara Gruen",
	"publication": "May 26, 2006"
}
````
