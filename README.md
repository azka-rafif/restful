# REST-API for Run Tracker

The Run Tracker API is a RESTful service for accessing and manipulating data related to runs. Built with Go, it offers an intuitive way for clients to create, read, update, and delete data about runs in a database.

## Getting Started

These instructions will guide you on how to run the Run Tracker API on your local machine.

### Prerequisites

You need to have Go installed on your machine. This project is developed using Go 1.17.4. While it may work with other versions, it's advisable to use the version that the project was developed with.

### Running the Project

1. Download or clone this repository to your local machine.

2. Navigate to the root directory of the project.

3. Start the application by running the command:
   ```
   make dev
   ```
   This command will start the program. If a database file does not already exist, create one inside the `/docs` folder. It will then start serving on the port 8080.

You should now have the API running locally on your machine. You can interact with it using any HTTP client, or by using the provided Swagger UI.

## API Documentation

The API documentation for this project is available through Swagger UI. Once the project is running, you can access the Swagger UI at the following address:

http://localhost:8080/swagger/

In the Swagger UI, you will be able to see all available endpoints, their expected input and output, and you can even test the endpoints directly in the UI.
