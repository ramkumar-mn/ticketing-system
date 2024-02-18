# Ticket Management System

This is a simple ticket management system that allows users to create, view, and delete tickets. The system is built
using Golang and gRPC.

## Project Structure

The project is structured as follows:

- `booking`: This directory contains the proto code, gRPC service definition and the generated code.
- `booking-api`: This directory contains the gRPC server and implementation of the service definition.
- `booking-client`: This directory contains the gRPC client that communicates with the server.

## Running the Project

Execute the following commands

    ```shell
    go mod tidy
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative booking/booking.proto
    go run booking-api/main.go
    go run booking-client/main.go
    ```

Note that execute server and client in separate terminal to see the communication between them.

## Assumptions made

- Each user will have a unique email address.
- Max seats per section is 100.
- Each user will book only one ticket.
- Each booking will have a unique booking id.
- In-memory datastore is used for simplicity and to avoid external dependencies.

## Future Improvements

- Use a persistent datastore like relational/non-relational databases.
- Allow user to book multiple tickets.
- Improvise seat allocation algorithm.
- Allow cancellation of specific tickets/booking.
- Design efficient database schema.
- Add user signup and login functionality.
- Add authentication and authorization to validate calling user.
