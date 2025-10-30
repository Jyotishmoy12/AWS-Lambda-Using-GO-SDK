# AWS Lambda with Golang: A "Hello, World" Guide

This guide covers the complete, step-by-step process for creating, compiling, and deploying a serverless "Hello, World" function on AWS Lambda using Go (Golang).

---

## Overview
The process involves:

1. Writing the Go code.  
2. Creating an IAM Role for the function.  
3. Compiling the Go code for Linux.  
4. Packaging and deploying the function using the AWS CLI.  
5. Invoking and testing the function from the command line.

---

## 1. Go Project Setup & SDK

### Step 1: Initialize Go Module
```bash
go mod init hello-lambda
```
**Use Case:** Initializes a new Go project and creates a `go.mod` file for dependency management.

### Step 2: Install AWS Lambda SDK
```bash
go get github.com/aws/aws-lambda-go/lambda
```
**Use Case:** Installs the official `aws-lambda-go` library, providing the `lambda.Start()` function to connect your Go handler with AWS Lambda.

---

## 2. Building the Go Binary for Lambda

AWS Lambda runs on a **Linux environment**, so we must compile our Go code accordingly.

### PowerShell Commands
```powershell
# 1. Target Linux OS
$env:GOOS = "linux"

# 2. Target amd64 architecture
$env:GOARCH = "amd64"

# 3. Build executable
go build -tags lambda.norpc -o bootstrap main.go
```

**Explanation:**
- `$env:GOOS = "linux"` → Builds for Linux OS.  
- `go build ...` → Compiles `main.go`.  
- `-o bootstrap` → Output file name required by Lambda.  
- `-tags lambda.norpc` → Optional build tag for smaller, optimized binary.

---

## 3. Packaging for Deployment

Lambda requires your code in a **.zip file**.

### PowerShell Command
```powershell
Compress-Archive -Path .ootstrap -DestinationPath .\deployment.zip
```

**Alternative (macOS/Linux):**
```bash
zip deployment.zip bootstrap
```

---

## 4. AWS CLI: Configuration

### Step 1: Configure AWS CLI
```bash
aws configure --profile lambda-admin
```

**Use Case:** Sets up credentials as a *named profile*. You’ll provide:
- AWS Access Key ID  
- AWS Secret Access Key  
- Default region  
- Output format

### Step 2: Verify Configuration
```bash
aws sts get-caller-identity --profile lambda-admin
```
**Use Case:** Confirms your credentials are valid and shows your AWS account info.

---

## 5. AWS CLI: Lambda Deployment

### Command (PowerShell, single line)
```powershell
aws lambda create-function --function-name hello-lambda-go --runtime provided.al2023 --handler bootstrap --role "IAM execution ARN" --zip-file fileb://deployment.zip --profile lambda-admin
```

**Parameter Breakdown:**
- `--function-name hello-lambda-go` → Function name.  
- `--runtime provided.al2023` → Custom runtime (Amazon Linux 2023).  
- `--handler bootstrap` → Entry executable name.  
- `--role` → ARN of your IAM execution role.  
- `--zip-file fileb://deployment.zip` → Uploads the deployment package.  
- `--profile lambda-admin` → Uses the correct AWS credentials profile.

---

## 6. AWS CLI: Lambda Invocation (Testing)

### Save Output to File
```powershell
aws lambda invoke --function-name hello-lambda-go --payload file://payload.json --profile lambda-admin real-output.json
```

**Explanation:**
- `--payload file://payload.json` → Reads JSON input from file.  
- `real-output.json` → File to store Lambda response.

### Print Output to Terminal
```powershell
aws lambda invoke --function-name hello-lambda-go --payload file://payload.json --profile lambda-admin -
```

**Explanation:**
- The `-` at the end prints directly to the terminal instead of saving.

### View Output File (PowerShell)
```powershell
Get-Content real-output.json
```

**Alternative (macOS/Linux):**
```bash
cat real-output.json
```
