# 🌐 Google Cloud Platform for Go - Complete Tutorial

Master GCP services for deploying and scaling Go applications in production.

## 🎯 Learning Objectives

By the end of this tutorial, you will be able to:
- Deploy Go applications on Cloud Run and GKE
- Use Google Cloud Storage for file storage
- Implement Firestore for NoSQL data
- Use Cloud Pub/Sub for messaging
- Set up Cloud Monitoring and Logging
- Implement Cloud Functions in Go
- Use Secret Manager for credentials
- Deploy with Cloud Build

---

## 📚 Table of Contents

1. [GCP SDK Setup](#gcp-sdk-setup)
2. [Cloud Storage](#cloud-storage)
3. [Firestore](#firestore)
4. [Cloud Pub/Sub](#cloud-pubsub)
5. [Cloud Functions](#cloud-functions)
6. [Secret Manager](#secret-manager)
7. [Cloud Run Deployment](#cloud-run-deployment)
8. [Monitoring and Logging](#monitoring-and-logging)

---

## 1. GCP SDK Setup

### Install GCP Libraries

```bash
go get cloud.google.com/go/storage
go get cloud.google.com/go/firestore
go get cloud.google.com/go/pubsub
go get cloud.google.com/go/secretmanager/apiv1
go get cloud.google.com/go/logging
```

### Initialize GCP Client

```go
package gcp

import (
    "context"
    "fmt"
    
    "google.golang.org/api/option"
)

type GCPConfig struct {
    ProjectID      string
    CredentialsFile string
}

func NewGCPConfig(projectID, credentialsFile string) *GCPConfig {
    return &GCPConfig{
        ProjectID:      projectID,
        CredentialsFile: credentialsFile,
    }
}

func (cfg *GCPConfig) ClientOptions() []option.ClientOption {
    if cfg.CredentialsFile != "" {
        return []option.ClientOption{
            option.WithCredentialsFile(cfg.CredentialsFile),
        }
    }
    return nil
}
```

---

## 2. Cloud Storage

### Storage Client

```go
package storage

import (
    "context"
    "fmt"
    "io"
    "time"
    
    "cloud.google.com/go/storage"
    "google.golang.org/api/option"
)

type GCSStorage struct {
    client *storage.Client
    bucket string
}

func NewGCSStorage(ctx context.Context, projectID, bucket string, opts ...option.ClientOption) (*GCSStorage, error) {
    client, err := storage.NewClient(ctx, opts...)
    if err != nil {
        return nil, fmt.Errorf("creating storage client: %w", err)
    }
    
    return &GCSStorage{
        client: client,
        bucket: bucket,
    }, nil
}

func (gcs *GCSStorage) Close() error {
    return gcs.client.Close()
}

// Upload file
func (gcs *GCSStorage) Upload(ctx context.Context, objectName string, content io.Reader, contentType string) error {
    obj := gcs.client.Bucket(gcs.bucket).Object(objectName)
    writer := obj.NewWriter(ctx)
    writer.ContentType = contentType
    
    if _, err := io.Copy(writer, content); err != nil {
        return fmt.Errorf("copying content: %w", err)
    }
    
    if err := writer.Close(); err != nil {
        return fmt.Errorf("closing writer: %w", err)
    }
    
    return nil
}

// Download file
func (gcs *GCSStorage) Download(ctx context.Context, objectName string) ([]byte, error) {
    obj := gcs.client.Bucket(gcs.bucket).Object(objectName)
    reader, err := obj.NewReader(ctx)
    if err != nil {
        return nil, fmt.Errorf("creating reader: %w", err)
    }
    defer reader.Close()
    
    data, err := io.ReadAll(reader)
    if err != nil {
        return nil, fmt.Errorf("reading data: %w", err)
    }
    
    return data, nil
}

// Delete file
func (gcs *GCSStorage) Delete(ctx context.Context, objectName string) error {
    obj := gcs.client.Bucket(gcs.bucket).Object(objectName)
    if err := obj.Delete(ctx); err != nil {
        return fmt.Errorf("deleting object: %w", err)
    }
    return nil
}

// List files
func (gcs *GCSStorage) List(ctx context.Context, prefix string) ([]string, error) {
    query := &storage.Query{Prefix: prefix}
    it := gcs.client.Bucket(gcs.bucket).Objects(ctx, query)
    
    var objects []string
    for {
        attrs, err := it.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("iterating objects: %w", err)
        }
        objects = append(objects, attrs.Name)
    }
    
    return objects, nil
}

// Generate signed URL
func (gcs *GCSStorage) GetSignedURL(ctx context.Context, objectName string, duration time.Duration) (string, error) {
    obj := gcs.client.Bucket(gcs.bucket).Object(objectName)
    
    url, err := obj.SignedURL(&storage.SignedURLOptions{
        Method:  "GET",
        Expires: time.Now().Add(duration),
    })
    
    if err != nil {
        return "", fmt.Errorf("generating signed URL: %w", err)
    }
    
    return url, nil
}

// Set metadata
func (gcs *GCSStorage) SetMetadata(ctx context.Context, objectName string, metadata map[string]string) error {
    obj := gcs.client.Bucket(gcs.bucket).Object(objectName)
    
    _, err := obj.Update(ctx, storage.ObjectAttrsToUpdate{
        Metadata: metadata,
    })
    
    if err != nil {
        return fmt.Errorf("updating metadata: %w", err)
    }
    
    return nil
}
```

---

## 3. Firestore

### Firestore Client

```go
package database

import (
    "context"
    "fmt"
    "time"
    
    "cloud.google.com/go/firestore"
    "google.golang.org/api/iterator"
    "google.golang.org/api/option"
)

type FirestoreDB struct {
    client *firestore.Client
}

func NewFirestoreDB(ctx context.Context, projectID string, opts ...option.ClientOption) (*FirestoreDB, error) {
    client, err := firestore.NewClient(ctx, projectID, opts...)
    if err != nil {
        return nil, fmt.Errorf("creating firestore client: %w", err)
    }
    
    return &FirestoreDB{client: client}, nil
}

func (db *FirestoreDB) Close() error {
    return db.client.Close()
}

// User model
type User struct {
    ID        string    `firestore:"id"`
    Email     string    `firestore:"email"`
    Name      string    `firestore:"name"`
    Active    bool      `firestore:"active"`
    CreatedAt time.Time `firestore:"created_at"`
    UpdatedAt time.Time `firestore:"updated_at"`
}

// Create user
func (db *FirestoreDB) CreateUser(ctx context.Context, user *User) error {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    
    _, err := db.client.Collection("users").Doc(user.ID).Set(ctx, user)
    if err != nil {
        return fmt.Errorf("creating user: %w", err)
    }
    
    return nil
}

// Get user
func (db *FirestoreDB) GetUser(ctx context.Context, userID string) (*User, error) {
    doc, err := db.client.Collection("users").Doc(userID).Get(ctx)
    if err != nil {
        return nil, fmt.Errorf("getting user: %w", err)
    }
    
    var user User
    if err := doc.DataTo(&user); err != nil {
        return nil, fmt.Errorf("parsing user data: %w", err)
    }
    
    return &user, nil
}

// Update user
func (db *FirestoreDB) UpdateUser(ctx context.Context, userID string, updates map[string]interface{}) error {
    updates["updated_at"] = time.Now()
    
    _, err := db.client.Collection("users").Doc(userID).Set(ctx, updates, firestore.MergeAll)
    if err != nil {
        return fmt.Errorf("updating user: %w", err)
    }
    
    return nil
}

// Delete user
func (db *FirestoreDB) DeleteUser(ctx context.Context, userID string) error {
    _, err := db.client.Collection("users").Doc(userID).Delete(ctx)
    if err != nil {
        return fmt.Errorf("deleting user: %w", err)
    }
    
    return nil
}

// Query users
func (db *FirestoreDB) GetActiveUsers(ctx context.Context) ([]*User, error) {
    iter := db.client.Collection("users").
        Where("active", "==", true).
        OrderBy("created_at", firestore.Desc).
        Documents(ctx)
    
    var users []*User
    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("iterating users: %w", err)
        }
        
        var user User
        if err := doc.DataTo(&user); err != nil {
            return nil, fmt.Errorf("parsing user: %w", err)
        }
        users = append(users, &user)
    }
    
    return users, nil
}

// Batch write
func (db *FirestoreDB) BatchCreateUsers(ctx context.Context, users []*User) error {
    batch := db.client.Batch()
    
    for _, user := range users {
        user.CreatedAt = time.Now()
        user.UpdatedAt = time.Now()
        
        ref := db.client.Collection("users").Doc(user.ID)
        batch.Set(ref, user)
    }
    
    _, err := batch.Commit(ctx)
    if err != nil {
        return fmt.Errorf("committing batch: %w", err)
    }
    
    return nil
}

// Transaction
func (db *FirestoreDB) TransferProgress(ctx context.Context, fromUserID, toUserID, lessonID string) error {
    return db.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
        // Read from source
        fromRef := db.client.Collection("users").Doc(fromUserID).
            Collection("progress").Doc(lessonID)
        
        fromDoc, err := tx.Get(fromRef)
        if err != nil {
            return err
        }
        
        var progress map[string]interface{}
        fromDoc.DataTo(&progress)
        
        // Delete from source
        if err := tx.Delete(fromRef); err != nil {
            return err
        }
        
        // Write to destination
        toRef := db.client.Collection("users").Doc(toUserID).
            Collection("progress").Doc(lessonID)
        
        return tx.Set(toRef, progress)
    })
}

// Real-time listener
func (db *FirestoreDB) WatchUser(ctx context.Context, userID string, callback func(*User)) error {
    snapshots := db.client.Collection("users").Doc(userID).Snapshots(ctx)
    defer snapshots.Stop()
    
    for {
        snap, err := snapshots.Next()
        if err != nil {
            return fmt.Errorf("watching user: %w", err)
        }
        
        var user User
        if err := snap.DataTo(&user); err != nil {
            return fmt.Errorf("parsing user: %w", err)
        }
        
        callback(&user)
    }
}
```

---

## 4. Cloud Pub/Sub

### Pub/Sub Client

```go
package messaging

import (
    "context"
    "fmt"
    
    "cloud.google.com/go/pubsub"
    "google.golang.org/api/option"
)

type PubSubClient struct {
    client *pubsub.Client
}

func NewPubSubClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*PubSubClient, error) {
    client, err := pubsub.NewClient(ctx, projectID, opts...)
    if err != nil {
        return nil, fmt.Errorf("creating pubsub client: %w", err)
    }
    
    return &PubSubClient{client: client}, nil
}

func (ps *PubSubClient) Close() error {
    return ps.client.Close()
}

// Publish message
func (ps *PubSubClient) Publish(ctx context.Context, topicID string, data []byte) (string, error) {
    topic := ps.client.Topic(topicID)
    defer topic.Stop()
    
    result := topic.Publish(ctx, &pubsub.Message{
        Data: data,
    })
    
    messageID, err := result.Get(ctx)
    if err != nil {
        return "", fmt.Errorf("publishing message: %w", err)
    }
    
    return messageID, nil
}

// Publish with attributes
func (ps *PubSubClient) PublishWithAttributes(ctx context.Context, topicID string, data []byte, attrs map[string]string) (string, error) {
    topic := ps.client.Topic(topicID)
    defer topic.Stop()
    
    result := topic.Publish(ctx, &pubsub.Message{
        Data:       data,
        Attributes: attrs,
    })
    
    messageID, err := result.Get(ctx)
    if err != nil {
        return "", fmt.Errorf("publishing message: %w", err)
    }
    
    return messageID, nil
}

// Subscribe to messages
func (ps *PubSubClient) Subscribe(ctx context.Context, subscriptionID string, handler func(context.Context, *pubsub.Message) error) error {
    sub := ps.client.Subscription(subscriptionID)
    
    err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
        if err := handler(ctx, msg); err != nil {
            msg.Nack()
            return
        }
        msg.Ack()
    })
    
    if err != nil {
        return fmt.Errorf("receiving messages: %w", err)
    }
    
    return nil
}

// Create topic
func (ps *PubSubClient) CreateTopic(ctx context.Context, topicID string) error {
    topic := ps.client.Topic(topicID)
    exists, err := topic.Exists(ctx)
    if err != nil {
        return fmt.Errorf("checking topic existence: %w", err)
    }
    
    if exists {
        return nil
    }
    
    _, err = ps.client.CreateTopic(ctx, topicID)
    if err != nil {
        return fmt.Errorf("creating topic: %w", err)
    }
    
    return nil
}

// Create subscription
func (ps *PubSubClient) CreateSubscription(ctx context.Context, subscriptionID, topicID string) error {
    sub := ps.client.Subscription(subscriptionID)
    exists, err := sub.Exists(ctx)
    if err != nil {
        return fmt.Errorf("checking subscription existence: %w", err)
    }
    
    if exists {
        return nil
    }
    
    _, err = ps.client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
        Topic: ps.client.Topic(topicID),
    })
    
    if err != nil {
        return fmt.Errorf("creating subscription: %w", err)
    }
    
    return nil
}
```

### Example: Event Processing

```go
type LessonCompletedEvent struct {
    UserID     string    `json:"user_id"`
    LessonID   int64     `json:"lesson_id"`
    Score      int       `json:"score"`
    CompletedAt time.Time `json:"completed_at"`
}

func PublishLessonCompleted(ctx context.Context, ps *PubSubClient, event *LessonCompletedEvent) error {
    data, err := json.Marshal(event)
    if err != nil {
        return fmt.Errorf("marshaling event: %w", err)
    }
    
    attrs := map[string]string{
        "event_type": "lesson_completed",
        "user_id":    event.UserID,
    }
    
    _, err = ps.PublishWithAttributes(ctx, "lesson-events", data, attrs)
    return err
}

func ProcessLessonEvents(ctx context.Context, ps *PubSubClient) error {
    handler := func(ctx context.Context, msg *pubsub.Message) error {
        var event LessonCompletedEvent
        if err := json.Unmarshal(msg.Data, &event); err != nil {
            return fmt.Errorf("unmarshaling event: %w", err)
        }
        
        log.Printf("Processing lesson completion: User=%s, Lesson=%d, Score=%d",
            event.UserID, event.LessonID, event.Score)
        
        // Process event
        return nil
    }
    
    return ps.Subscribe(ctx, "lesson-events-sub", handler)
}
```

---


