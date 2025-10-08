# ☁️ AWS for Go Applications - Complete Tutorial

Master AWS services for deploying and scaling Go applications in production.

## 🎯 Learning Objectives

By the end of this tutorial, you will be able to:
- Deploy Go applications on AWS ECS/Fargate
- Use AWS SDK for Go (v2)
- Implement S3 for file storage
- Use DynamoDB for NoSQL data
- Set up CloudWatch for monitoring
- Implement Lambda functions in Go
- Use SQS and SNS for messaging
- Secure applications with IAM and Secrets Manager

---

## 📚 Table of Contents

1. [AWS SDK Setup](#aws-sdk-setup)
2. [S3 File Storage](#s3-file-storage)
3. [DynamoDB](#dynamodb)
4. [Lambda Functions](#lambda-functions)
5. [SQS Messaging](#sqs-messaging)
6. [SNS Notifications](#sns-notifications)
7. [Secrets Manager](#secrets-manager)
8. [CloudWatch](#cloudwatch)

---

## 1. AWS SDK Setup

### Install AWS SDK v2

```bash
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3
go get github.com/aws/aws-sdk-go-v2/service/dynamodb
go get github.com/aws/aws-sdk-go-v2/service/sqs
go get github.com/aws/aws-sdk-go-v2/service/sns
go get github.com/aws/aws-sdk-go-v2/service/secretsmanager
```

### Initialize AWS Config

```go
package aws

import (
    "context"
    "fmt"
    
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
)

type AWSConfig struct {
    Config aws.Config
}

func NewAWSConfig(ctx context.Context, region string) (*AWSConfig, error) {
    cfg, err := config.LoadDefaultConfig(ctx,
        config.WithRegion(region),
    )
    if err != nil {
        return nil, fmt.Errorf("loading AWS config: %w", err)
    }
    
    return &AWSConfig{Config: cfg}, nil
}

// With custom credentials
func NewAWSConfigWithCredentials(ctx context.Context, region, accessKey, secretKey string) (*AWSConfig, error) {
    cfg, err := config.LoadDefaultConfig(ctx,
        config.WithRegion(region),
        config.WithCredentialsProvider(
            aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
                return aws.Credentials{
                    AccessKeyID:     accessKey,
                    SecretAccessKey: secretKey,
                }, nil
            }),
        ),
    )
    if err != nil {
        return nil, fmt.Errorf("loading AWS config: %w", err)
    }
    
    return &AWSConfig{Config: cfg}, nil
}
```

---

## 2. S3 File Storage

### S3 Client

```go
package storage

import (
    "context"
    "fmt"
    "io"
    "time"
    
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Storage struct {
    client *s3.Client
    bucket string
}

func NewS3Storage(cfg aws.Config, bucket string) *S3Storage {
    return &S3Storage{
        client: s3.NewFromConfig(cfg),
        bucket: bucket,
    }
}

// Upload file
func (s *S3Storage) Upload(ctx context.Context, key string, body io.Reader, contentType string) error {
    _, err := s.client.PutObject(ctx, &s3.PutObjectInput{
        Bucket:      aws.String(s.bucket),
        Key:         aws.String(key),
        Body:        body,
        ContentType: aws.String(contentType),
    })
    
    if err != nil {
        return fmt.Errorf("uploading to S3: %w", err)
    }
    
    return nil
}

// Download file
func (s *S3Storage) Download(ctx context.Context, key string) (io.ReadCloser, error) {
    result, err := s.client.GetObject(ctx, &s3.GetObjectInput{
        Bucket: aws.String(s.bucket),
        Key:    aws.String(key),
    })
    
    if err != nil {
        return nil, fmt.Errorf("downloading from S3: %w", err)
    }
    
    return result.Body, nil
}

// Delete file
func (s *S3Storage) Delete(ctx context.Context, key string) error {
    _, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
        Bucket: aws.String(s.bucket),
        Key:    aws.String(key),
    })
    
    if err != nil {
        return fmt.Errorf("deleting from S3: %w", err)
    }
    
    return nil
}

// List files
func (s *S3Storage) List(ctx context.Context, prefix string) ([]string, error) {
    result, err := s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
        Bucket: aws.String(s.bucket),
        Prefix: aws.String(prefix),
    })
    
    if err != nil {
        return nil, fmt.Errorf("listing S3 objects: %w", err)
    }
    
    var keys []string
    for _, obj := range result.Contents {
        keys = append(keys, *obj.Key)
    }
    
    return keys, nil
}

// Generate presigned URL
func (s *S3Storage) GetPresignedURL(ctx context.Context, key string, duration time.Duration) (string, error) {
    presignClient := s3.NewPresignClient(s.client)
    
    request, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
        Bucket: aws.String(s.bucket),
        Key:    aws.String(key),
    }, func(opts *s3.PresignOptions) {
        opts.Expires = duration
    })
    
    if err != nil {
        return "", fmt.Errorf("generating presigned URL: %w", err)
    }
    
    return request.URL, nil
}

// Upload with multipart for large files
func (s *S3Storage) UploadLarge(ctx context.Context, key string, body io.Reader) error {
    uploader := manager.NewUploader(s.client)
    
    _, err := uploader.Upload(ctx, &s3.PutObjectInput{
        Bucket: aws.String(s.bucket),
        Key:    aws.String(key),
        Body:   body,
    })
    
    if err != nil {
        return fmt.Errorf("uploading large file: %w", err)
    }
    
    return nil
}
```

### Example: User Avatar Upload

```go
func UploadUserAvatar(ctx context.Context, s3 *S3Storage, userID string, file io.Reader) (string, error) {
    key := fmt.Sprintf("avatars/%s/%d.jpg", userID, time.Now().Unix())
    
    if err := s3.Upload(ctx, key, file, "image/jpeg"); err != nil {
        return "", err
    }
    
    // Generate presigned URL valid for 1 hour
    url, err := s3.GetPresignedURL(ctx, key, 1*time.Hour)
    if err != nil {
        return "", err
    }
    
    return url, nil
}
```

---

## 3. DynamoDB

### DynamoDB Client

```go
package database

import (
    "context"
    "fmt"
    
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBClient struct {
    client    *dynamodb.Client
    tableName string
}

func NewDynamoDBClient(cfg aws.Config, tableName string) *DynamoDBClient {
    return &DynamoDBClient{
        client:    dynamodb.NewFromConfig(cfg),
        tableName: tableName,
    }
}

// User model
type User struct {
    UserID    string `dynamodbav:"user_id"`
    Email     string `dynamodbav:"email"`
    Name      string `dynamodbav:"name"`
    CreatedAt int64  `dynamodbav:"created_at"`
}

// Put item
func (db *DynamoDBClient) PutUser(ctx context.Context, user *User) error {
    item, err := attributevalue.MarshalMap(user)
    if err != nil {
        return fmt.Errorf("marshaling user: %w", err)
    }
    
    _, err = db.client.PutItem(ctx, &dynamodb.PutItemInput{
        TableName: aws.String(db.tableName),
        Item:      item,
    })
    
    if err != nil {
        return fmt.Errorf("putting item: %w", err)
    }
    
    return nil
}

// Get item
func (db *DynamoDBClient) GetUser(ctx context.Context, userID string) (*User, error) {
    result, err := db.client.GetItem(ctx, &dynamodb.GetItemInput{
        TableName: aws.String(db.tableName),
        Key: map[string]types.AttributeValue{
            "user_id": &types.AttributeValueMemberS{Value: userID},
        },
    })
    
    if err != nil {
        return nil, fmt.Errorf("getting item: %w", err)
    }
    
    if result.Item == nil {
        return nil, fmt.Errorf("user not found")
    }
    
    var user User
    if err := attributevalue.UnmarshalMap(result.Item, &user); err != nil {
        return nil, fmt.Errorf("unmarshaling user: %w", err)
    }
    
    return &user, nil
}

// Query items
func (db *DynamoDBClient) QueryUsersByEmail(ctx context.Context, email string) ([]User, error) {
    result, err := db.client.Query(ctx, &dynamodb.QueryInput{
        TableName:              aws.String(db.tableName),
        IndexName:              aws.String("email-index"),
        KeyConditionExpression: aws.String("email = :email"),
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":email": &types.AttributeValueMemberS{Value: email},
        },
    })
    
    if err != nil {
        return nil, fmt.Errorf("querying items: %w", err)
    }
    
    var users []User
    if err := attributevalue.UnmarshalListOfMaps(result.Items, &users); err != nil {
        return nil, fmt.Errorf("unmarshaling users: %w", err)
    }
    
    return users, nil
}

// Update item
func (db *DynamoDBClient) UpdateUserName(ctx context.Context, userID, name string) error {
    _, err := db.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
        TableName: aws.String(db.tableName),
        Key: map[string]types.AttributeValue{
            "user_id": &types.AttributeValueMemberS{Value: userID},
        },
        UpdateExpression: aws.String("SET #name = :name"),
        ExpressionAttributeNames: map[string]string{
            "#name": "name",
        },
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":name": &types.AttributeValueMemberS{Value: name},
        },
    })
    
    if err != nil {
        return fmt.Errorf("updating item: %w", err)
    }
    
    return nil
}

// Delete item
func (db *DynamoDBClient) DeleteUser(ctx context.Context, userID string) error {
    _, err := db.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
        TableName: aws.String(db.tableName),
        Key: map[string]types.AttributeValue{
            "user_id": &types.AttributeValueMemberS{Value: userID},
        },
    })
    
    if err != nil {
        return fmt.Errorf("deleting item: %w", err)
    }
    
    return nil
}

// Batch write
func (db *DynamoDBClient) BatchPutUsers(ctx context.Context, users []User) error {
    var writeRequests []types.WriteRequest
    
    for _, user := range users {
        item, err := attributevalue.MarshalMap(user)
        if err != nil {
            return fmt.Errorf("marshaling user: %w", err)
        }
        
        writeRequests = append(writeRequests, types.WriteRequest{
            PutRequest: &types.PutRequest{
                Item: item,
            },
        })
    }
    
    _, err := db.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
        RequestItems: map[string][]types.WriteRequest{
            db.tableName: writeRequests,
        },
    })
    
    if err != nil {
        return fmt.Errorf("batch writing items: %w", err)
    }
    
    return nil
}
```

---

## 4. Lambda Functions

### Lambda Handler

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
    UserID   string `json:"user_id"`
    LessonID int64  `json:"lesson_id"`
}

type Response struct {
    Message string `json:"message"`
    Success bool   `json:"success"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    var req Request
    if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: 400,
            Body:       `{"error": "Invalid request"}`,
        }, nil
    }
    
    // Process request
    result := processLesson(ctx, req.UserID, req.LessonID)
    
    response := Response{
        Message: fmt.Sprintf("Processed lesson %d for user %s", req.LessonID, req.UserID),
        Success: result,
    }
    
    body, _ := json.Marshal(response)
    
    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
        Body: string(body),
    }, nil
}

func processLesson(ctx context.Context, userID string, lessonID int64) bool {
    // Business logic here
    return true
}

func main() {
    lambda.Start(HandleRequest)
}
```

### Build and Deploy Lambda

```bash
# Build for Lambda
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
zip function.zip bootstrap

# Deploy with AWS CLI
aws lambda create-function \
  --function-name go-pro-processor \
  --runtime provided.al2 \
  --handler bootstrap \
  --zip-file fileb://function.zip \
  --role arn:aws:iam::ACCOUNT_ID:role/lambda-role
```

---


