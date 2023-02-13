Diary API


Diary API is a Go language-based application that serves as an interface for a web or mobile diary application. It allows users to interact with their diary data by providing a set of endpoints to perform CRUD (Create, Read, Update, Delete) operations.

The application consists of the following components:

Main function
Import statements
Connecting to the database
Gin framework
API endpoints
Running the application
Main function
The main function is the entry point of the application, where it connects to the database, sets up the Gin framework, and maps each endpoint to a corresponding function in the controllers package.

Import statements
The import statements import the required packages and dependencies, including the db package for database connection, the controllers package for handling the application logic, and the github.com/gin-gonic/gin package for using the Gin framework.

Connecting to the database
The db.Connect() function connects to the database using the db package and opens a database session. The session is later closed by calling the db.GetDB().Close() function in the defer statement.

Gin framework
The Gin framework is used for handling HTTP requests and routing. It is initialized by calling gin.Default(), and a group of endpoints is created for version 1 of the API by calling r.Group("/api/v1").

API endpoints
The application provides a set of endpoints for each diary component, including diaries, titles, and photos. Each endpoint is mapped to a corresponding function in the controllers package, which handles the application logic.

For example, the endpoint /api/v1/diaries maps to the controllers.GetAllDiaries function, which retrieves all diary entries from the database.

Running the application
The application is run by calling r.Run(), which starts the Gin framework and listens for incoming HTTP requests on the default port 8080.
