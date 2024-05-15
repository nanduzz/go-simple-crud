## Go-Simple-Crud

### Description
This is a simple CRUD application using Go and MongoDB. This application is a simple REST API that can be used to create, read, update, and delete data from a Mongo database.

### But Why?
This is a simple project that I created to learn more about Go without using Obejct Oriented Programming. Since I'm mainly a Java programmer, I wanted to learn more about how to structure a Go project and how to use the language without using classes and objects.

### Feedbacks
I would love to hear your feedback on this project. If you have any suggestions or improvements, please let me know. This is a learning project for me and I would love to learn more from you.

### Starting the Mongo Database
Having docker and docker-compose installed, you can start the mongo database by running the following command:
```bash
docker-compose up -d
```
this will start the mongo database on port 27017. You can change the port by modifying the docker-compose.yml file.

### Running the Application
To run the application, you can use the following command:
```bash
go run main.go
```

### Testing the Application
You can run unit tests by running the following command:
```bash
go test ./...
```

