# loyalty_program_exercise

Backend that implements a loyalty program

## Table of Contents

- [Description](#description)
- [Libraries and Dependencies](#libraries-and-dependencies)
- [Start project](#start-project)

## Description

This project is a backend that implements a loyalty program. It is a REST API that allows to create, update, delete and get customers and their loyalty points. It also allows to create, update, delete and get transactions.

## Getting Started

### Prerequisites

- Go 1.21.9
- Gin 1.10.0
- Gorm 1.25.11
- Uber Go Dig 1.18.0

### Installation

1. Clone the repository

```bash
git clone

# install Dependencies
got mod tidy

# Run the project
go run main.go
```

## Libraries and Dependencies

### Uber Go Dig

Allow to easily handle dependency injection in Go.

### Swagger

Allow to easily create API documentation with and web UI interface.

## Start project

### Create docker image

This application run using docker, so in order to run the application use:

```bash
  docker build -t loyalty_program .
```

### Run docker compose file

In order to setup the application, once the docker image is created you can use the docker compose file to start the app and the database so it can be immediately use

```bash
  docker compose -d up
```

### documentation

To see the documentation of the API, you can go to the following URL:

```bash
  http://localhost:8080/swagger/index.html
```
