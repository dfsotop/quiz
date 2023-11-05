# Simple Quizz

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-00C5D9?style=for-the-badge&logo=grpc&logoColor=white)
![Cobra](https://img.shields.io/badge/Cobra-ef9020?style=for-the-badge&logo=cobra&logoColor=white)

This is a simple Quiz project implemented in Go, using gRPC as the communication framework. 

In this example the server provides an API for registering answers to quiz questions and comparing them with historical data. The client, on the other hand, sends requests to the server and handles the responses. This allows users to interact with the Quiz Service from the command line.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or higher
- gRPC

### Installing

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/quiz-service.git
    ```
2. Navigate to the project directory:
    ```bash
    cd quiz-service
    ```
3. Download the dependencies:
    ```bash
    go mod download
    ```

### Running the tests

To run the tests, use the following command:

```bash
go test ./...
```

# Project components

## Server

The server for this service is implemented in Go. It uses gRPC for communication and provides an API for registering answers to quiz questions and comparing them with historical data.

The server maintains an in-memory array to store the historical data. This data is used when comparing new answers to previous ones.

To run the server, navigate to the server directory and build the application:

```bash
cd server
go build
```

This will create an executable file in the server directory. You can run this file to start the server:

```bash
./server
```

## Client

The client for this service is a command-line interface (CLI) built using the [Cobra](https://github.com/spf13/cobra) library. It connects to the server, retrieves the quiz questions, and sends the user's answers back to the server.

After the answers are sent, the client receives the quiz feedback and then, prints these results to the command line, allowing the user to see how well they did on the quiz as well as how they compare to other quizzers.

To use the client, navigate to the client directory and build the application:

```bash
cd client
go build
```
This will create an executable file in the client directory. To start the quiz, execute:

```bash
./client quiz
```

# Maintainance
If you whish to modify the contract, include your changes in the file ./proto/quiz.proto and then generate the binding files. 

Please Consider that you may also have to modify the server and client code to include your changes.

## Generate binding files

### Prerequisites

You need to have the following tools installed:

- Protocol Buffers compiler (`protoc`)
- Go plugin for Protocol Buffers (`protoc-gen-go`)
- Go plugin for gRPC (`protoc-gen-go-grpc`)

### Generating the Files

Once you have the necessary tools, navigate to the project's root folder in your terminal and execute the following command:

```bash
protoc --go_out=. --go-grpc_out=.  proto/quiz.proto
```