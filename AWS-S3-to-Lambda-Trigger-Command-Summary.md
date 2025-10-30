# AWS S3-to-Lambda Trigger: Command Summary

This file summarizes all the commands used to connect an **S3 bucket event** to an **AWS Lambda function**.

---

## 1. Updating the Go Code

### Step 1: Add Dependency
```bash
go get github.com/aws/aws-lambda-go/events
```
**Use Case:** Installs the `events` package that provides Go structs (like `events.S3Event`) matching AWS event JSON payloads.

---

## 2. Re-Building and Re-Deploying the Function

### Rebuild the Binary (PowerShell)
```powershell
# 1. Set target OS to Linux
$env:GOOS = "linux"

# 2. Set target architecture
$env:GOARCH = "amd64"

# 3. Build executable
go build -tags lambda.norpc -o bootstrap main.go
```
**Use Case:** Cross-compiles Go code into a Linux executable named **bootstrap** for Lambda.

### Recreate Deployment Package
```powershell
# 1. Delete old zip (optional)
Remove-Item -Path .\deployment.zip -ErrorAction SilentlyContinue

# 2. Create new zip file
Compress-Archive -Path .\bootstrap -DestinationPath .\deployment.zip
```
**Use Case:** Packages the new binary into `deployment.zip` (required for Lambda deployment).

### Update Function Code
```bash
aws lambda update-function-code --function-name hello-lambda-go --zip-file fileb://deployment.zip --profile lambda-admin
```
**Use Case:** Uploads the new deployment package, replacing the old "Hello World" code with S3 event handling logic.

---

## 3. Granting S3 Invoke Permissions

### Command (PowerShell, multi-line)
```powershell
aws lambda add-permission `
  --function-name hello-lambda-go `
  --statement-id "S3-Invoke-Permission-1" `
  --action "lambda:InvokeFunction" `
  --principal s3.amazonaws.com `
  --source-arn "arn:aws:s3:::jyotishmoyfirstbucket" `
  --profile lambda-admin
```

**Use Case:** Updates the Lambda's resource-based policy to allow S3 to invoke it.  
It grants permission for `s3.amazonaws.com` to trigger the function `hello-lambda-go`, but **only** from the specified S3 bucket (`--source-arn`).

---

## 4. Configuring the S3 Bucket

### Command
```bash
aws s3api put-bucket-notification-configuration --bucket "jyotishmoyfirstbucket" --notification-configuration file://notification.json --profile lambda-admin
```
**Use Case:** Configures the S3 bucket to send event notifications to the Lambda function.

### notification.json
```json
{
  "LambdaFunctionConfigurations": [
    {
      "LambdaFunctionArn": "arn:aws:lambda:ap-south-1:135808959611:function:hello-lambda-go",
      "Events": ["s3:ObjectCreated:*"]
    }
  ]
}
```
**Explanation:** This tells AWS, “Whenever a new object is created in the bucket, trigger this Lambda function.”

---

## 5. Testing and Verification

### Test Commands (PowerShell)
```powershell
# 1. Create a local test file
Set-Content -Path "test-trigger.txt" -Value "This is a test for Lambda!"

# 2. Upload the test file to S3
aws s3 cp .\test-trigger.txt s3://jyotishmoyfirstbucket/test-trigger.txt --profile lambda-admin
```

**Use Case:** Uploading the file triggers the Lambda function automatically.

### Verification
In the **AWS Console**:
- Go to **Lambda → Monitor → View CloudWatch Logs**
- Check the latest log stream for your function

If you see output like `"File upload detected!"`, it confirms your **S3-to-Lambda trigger pipeline** works correctly.