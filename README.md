# 🌐 How we wrap GinGonic with Lambda

To wrap a Gin Gonic REST API in an AWS Lambda function, we use the `gin` web framework to handle HTTP requests locally
and invoke Lambda using AWS SAM CLI when running in serverless mode.

1. **Gin Gonic for Local REST API**:
   - When running the application locally, Gin Gonic handles HTTP requests as a typical Go REST API.
   - The `gin.Default()` creates the Gin router, and you define your API routes like `/hello`, `/echo`, etc.
   - This part allows you to test your API locally without needing AWS infrastructure.

2. **Lambda Proxy Mode**:
   - When the application runs in Lambda Proxy mode, the API Gateway invokes the Lambda function instead of Gin directly.
   - We use **AWS SAM CLI** to emulate the API Gateway locally and route requests to the Lambda function.
   - The Lambda function itself is invoked in response to API requests, and the response is returned through API Gateway.

3. **Conditional Routing Based on Environment**:
   - By checking the environment variable (`ENV`), the app switches between running as a REST API with Gin locally or in Lambda Proxy mode when deployed to AWS.
   

## **Dependencies**:
We use the following dependencies to wrap gingonic server and transform the traditional rest api into a lambda
handler
(no need to install, they are already in the project)

```shell
  go get -u github.com/aws/aws-lambda-go/lambda
  go get -u github.com/awslabs/aws-lambda-go-api-proxy/gin
```

These are resources from [awslabs](https://github.com/awslabs) and [aws](https://github.com/aws) official AWS users.
Here we can find [aws-lambda-go-proxy](https://github.com/awslabs/aws-lambda-go-api-proxy)
and [aws-lambda-go](https://github.com/aws/aws-lambda-go) dependencies.

**Install SAM CLI**:

As we said, AWS SAM CLI includes a local server that emulates Lambda behavior. Make sure you have Docker installed and running,
as SAM uses containers to simulate the Lambda environment.

```shell
  brew install aws/tap/aws-sam-cli
```
In the project you can find `template.yml` file. This file is used to configure sam cli.

The next step is to execute the Makefile as indicated below, in the Makefile section:

# 🚀 Running the Makefile Correctly with Environment Variables

This guide will help you run the Makefile for your Go application using different
environment configurations (`ENV`). 
The `Makefile` allows you to compile, run, and test your Go application either locally or in Lambda Proxy mode using AWS SAM CLI.

## **Available Commands**

### 1. **Run Locally in REST API Mode**

To run your application as a REST API locally using the default `local` environment, execute the following command:

```shell
  make run-local
```
This command will:
- Set the ENV variable to local by default.
- Run the application as a standard Go API on your local machine (main.go).

### 2. **Run in Lambda Proxy Mode Using SAM CLI**

To execute the application with Lambda Proxy mode (using AWS SAM CLI), set the ENV variable to serverless:

```shell
  make ENV=serverless run-proxy
```
This command will:
- Compile the Go application for AWS Lambda.
- Use AWS SAM CLI to start an API Gateway locally, emulating how your application would run in AWS Lambda.
- Load the appropriate environment variables from env.json.

### 3. **Format Code Using goimports and gofmt**

To format your Go code and ensure it adheres to best practices, you can run:

```shell
  make format
```
This command will:
- Run goimports to automatically fix imports and format your code.
- Run gofmt to ensure your code is properly formatted according to Go's style guidelines.

## Environment Configuration

### **Default Environment: `local`**
If no environment is specified, the `ENV` variable defaults to `local`. This environment assumes you're running 
your application as a standard REST API on your local machine.

### **Custom Environment: `serverless`**
You can set the `ENV` variable to `serverless` to run your application in Lambda Proxy mode:
```shell
  make ENV=serverless run-proxy
```
This mode will invoke your API via AWS Lambda using SAM CLI and load the corresponding configuration for the serverless environment.

# Tests 🧪

Execute unit tests

```shell
  go test -v ./...
```

Create & update mocks

```shell
  mockery --all
```