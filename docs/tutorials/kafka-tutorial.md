# 📨 Apache Kafka with Go - Complete Tutorial

Master event streaming with Apache Kafka in Go applications for building scalable, real-time data pipelines.

## 🎯 Learning Objectives

By the end of this tutorial, you will be able to:
- Set up Kafka producers and consumers in Go
- Implement event-driven architectures
- Handle message serialization (JSON, Avro, Protobuf)
- Implement consumer groups for scalability
- Handle errors and retries
- Monitor Kafka applications
- Implement exactly-once semantics
- Build real-time data pipelines

---

## 📚 Table of Contents

1. [Setup and Configuration](#setup-and-configuration)
2. [Producer Basics](#producer-basics)
3. [Consumer Basics](#consumer-basics)
4. [Consumer Groups](#consumer-groups)
5. [Message Serialization](#message-serialization)
6. [Error Handling](#error-handling)
7. [Advanced Patterns](#advanced-patterns)
8. [Monitoring](#monitoring)
9. [Best Practices](#best-practices)

---

## 1. Setup and Configuration

### Install Kafka Client

```bash
# Using Confluent's Kafka Go client (recommended)
go get github.com/confluentinc/confluent-kafka-go/v2/kafka

# Or Shopify's Sarama
go get github.com/IBM/sarama
```

### Docker Compose for Local Kafka

```yaml
# docker-compose.yml
version: '3.8'

services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000
    ports:
      - "2181:2181"

  kafka:
    image: confluentinc/cp-kafka:latest
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
      - "9093:9093"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092,PLAINTEXT_INTERNAL://kafka:9093
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_INTERNAL:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT_INTERNAL
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: "true"

  kafka-ui:
    image: provectuslabs/kafka-ui:latest
    depends_on:
      - kafka
    ports:
      - "8080:8080"
    environment:
      KAFKA_CLUSTERS_0_NAME: local
      KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS: kafka:9093
```

Start Kafka:
```bash
docker-compose up -d
```

---

## 2. Producer Basics

### Simple Producer

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
    producer *kafka.Producer
}

func NewKafkaProducer(brokers string) (*KafkaProducer, error) {
    config := &kafka.ConfigMap{
        "bootstrap.servers": brokers,
        "client.id":         "go-pro-producer",
        "acks":              "all",
    }
    
    producer, err := kafka.NewProducer(config)
    if err != nil {
        return nil, fmt.Errorf("creating producer: %w", err)
    }
    
    return &KafkaProducer{producer: producer}, nil
}

func (kp *KafkaProducer) ProduceMessage(topic, key, value string) error {
    deliveryChan := make(chan kafka.Event)
    
    err := kp.producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{
            Topic:     &topic,
            Partition: kafka.PartitionAny,
        },
        Key:   []byte(key),
        Value: []byte(value),
    }, deliveryChan)
    
    if err != nil {
        return fmt.Errorf("producing message: %w", err)
    }
    
    // Wait for delivery report
    e := <-deliveryChan
    m := e.(*kafka.Message)
    
    if m.TopicPartition.Error != nil {
        return fmt.Errorf("delivery failed: %w", m.TopicPartition.Error)
    }
    
    log.Printf("Message delivered to %v", m.TopicPartition)
    close(deliveryChan)
    
    return nil
}

func (kp *KafkaProducer) Close() {
    kp.producer.Close()
}

func main() {
    producer, err := NewKafkaProducer("localhost:9092")
    if err != nil {
        log.Fatal(err)
    }
    defer producer.Close()
    
    // Produce a message
    err = producer.ProduceMessage("user-events", "user-123", `{"action":"login","timestamp":"2025-01-15T10:00:00Z"}`)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Message sent successfully!")
}
```

### Async Producer with Goroutines

```go
func (kp *KafkaProducer) ProduceAsync(topic, key, value string) {
    go func() {
        err := kp.producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{
                Topic:     &topic,
                Partition: kafka.PartitionAny,
            },
            Key:   []byte(key),
            Value: []byte(value),
        }, nil)
        
        if err != nil {
            log.Printf("Failed to produce message: %v", err)
        }
    }()
}

// Handle delivery reports in background
func (kp *KafkaProducer) HandleDeliveryReports() {
    go func() {
        for e := range kp.producer.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    log.Printf("Delivery failed: %v", ev.TopicPartition.Error)
                } else {
                    log.Printf("Delivered message to %v", ev.TopicPartition)
                }
            }
        }
    }()
}
```

### Batch Producer

```go
type Event struct {
    UserID    string    `json:"user_id"`
    Action    string    `json:"action"`
    Timestamp time.Time `json:"timestamp"`
}

func (kp *KafkaProducer) ProduceBatch(topic string, events []Event) error {
    for _, event := range events {
        data, err := json.Marshal(event)
        if err != nil {
            return fmt.Errorf("marshaling event: %w", err)
        }
        
        err = kp.producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{
                Topic:     &topic,
                Partition: kafka.PartitionAny,
            },
            Key:   []byte(event.UserID),
            Value: data,
        }, nil)
        
        if err != nil {
            return fmt.Errorf("producing message: %w", err)
        }
    }
    
    // Flush to ensure all messages are sent
    kp.producer.Flush(15 * 1000) // 15 seconds timeout
    
    return nil
}
```

---

## 3. Consumer Basics

### Simple Consumer

```go
package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
    consumer *kafka.Consumer
}

func NewKafkaConsumer(brokers, groupID string, topics []string) (*KafkaConsumer, error) {
    config := &kafka.ConfigMap{
        "bootstrap.servers": brokers,
        "group.id":          groupID,
        "auto.offset.reset": "earliest",
    }
    
    consumer, err := kafka.NewConsumer(config)
    if err != nil {
        return nil, fmt.Errorf("creating consumer: %w", err)
    }
    
    err = consumer.SubscribeTopics(topics, nil)
    if err != nil {
        return nil, fmt.Errorf("subscribing to topics: %w", err)
    }
    
    return &KafkaConsumer{consumer: consumer}, nil
}

func (kc *KafkaConsumer) Consume(handler func(*kafka.Message) error) error {
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
    
    run := true
    for run {
        select {
        case sig := <-sigchan:
            log.Printf("Caught signal %v: terminating", sig)
            run = false
            
        default:
            ev := kc.consumer.Poll(100)
            if ev == nil {
                continue
            }
            
            switch e := ev.(type) {
            case *kafka.Message:
                log.Printf("Message on %s: %s", e.TopicPartition, string(e.Value))
                
                if err := handler(e); err != nil {
                    log.Printf("Error handling message: %v", err)
                    continue
                }
                
                // Commit offset after successful processing
                _, err := kc.consumer.CommitMessage(e)
                if err != nil {
                    log.Printf("Error committing offset: %v", err)
                }
                
            case kafka.Error:
                log.Printf("Error: %v", e)
                if e.Code() == kafka.ErrAllBrokersDown {
                    run = false
                }
            }
        }
    }
    
    return nil
}

func (kc *KafkaConsumer) Close() {
    kc.consumer.Close()
}

func main() {
    consumer, err := NewKafkaConsumer(
        "localhost:9092",
        "go-pro-consumer-group",
        []string{"user-events"},
    )
    if err != nil {
        log.Fatal(err)
    }
    defer consumer.Close()
    
    // Message handler
    handler := func(msg *kafka.Message) error {
        log.Printf("Processing: %s", string(msg.Value))
        // Process message here
        return nil
    }
    
    if err := consumer.Consume(handler); err != nil {
        log.Fatal(err)
    }
}
```

### Consumer with Structured Messages

```go
type UserEvent struct {
    UserID    string    `json:"user_id"`
    Action    string    `json:"action"`
    Timestamp time.Time `json:"timestamp"`
}

func ProcessUserEvent(msg *kafka.Message) error {
    var event UserEvent
    if err := json.Unmarshal(msg.Value, &event); err != nil {
        return fmt.Errorf("unmarshaling event: %w", err)
    }
    
    log.Printf("User %s performed %s at %v", event.UserID, event.Action, event.Timestamp)
    
    // Business logic here
    switch event.Action {
    case "login":
        // Handle login
    case "logout":
        // Handle logout
    case "purchase":
        // Handle purchase
    }
    
    return nil
}
```

---

## 4. Consumer Groups

### Parallel Processing with Consumer Group

```go
type ConsumerGroup struct {
    consumers []*KafkaConsumer
    wg        sync.WaitGroup
}

func NewConsumerGroup(brokers, groupID string, topics []string, numConsumers int) (*ConsumerGroup, error) {
    cg := &ConsumerGroup{
        consumers: make([]*KafkaConsumer, numConsumers),
    }
    
    for i := 0; i < numConsumers; i++ {
        consumer, err := NewKafkaConsumer(brokers, groupID, topics)
        if err != nil {
            return nil, fmt.Errorf("creating consumer %d: %w", i, err)
        }
        cg.consumers[i] = consumer
    }
    
    return cg, nil
}

func (cg *ConsumerGroup) Start(handler func(*kafka.Message) error) {
    for i, consumer := range cg.consumers {
        cg.wg.Add(1)
        go func(id int, c *KafkaConsumer) {
            defer cg.wg.Done()
            log.Printf("Starting consumer %d", id)
            if err := c.Consume(handler); err != nil {
                log.Printf("Consumer %d error: %v", id, err)
            }
        }(i, consumer)
    }
}

func (cg *ConsumerGroup) Wait() {
    cg.wg.Wait()
}

func (cg *ConsumerGroup) Close() {
    for _, consumer := range cg.consumers {
        consumer.Close()
    }
}
```

---

## 5. Message Serialization

### JSON Serialization

```go
type LessonCompletedEvent struct {
    UserID     string    `json:"user_id"`
    LessonID   int64     `json:"lesson_id"`
    Score      int       `json:"score"`
    CompletedAt time.Time `json:"completed_at"`
}

func ProduceJSONEvent(producer *KafkaProducer, topic string, event *LessonCompletedEvent) error {
    data, err := json.Marshal(event)
    if err != nil {
        return fmt.Errorf("marshaling event: %w", err)
    }
    
    return producer.ProduceMessage(topic, event.UserID, string(data))
}

func ConsumeJSONEvent(msg *kafka.Message) (*LessonCompletedEvent, error) {
    var event LessonCompletedEvent
    if err := json.Unmarshal(msg.Value, &event); err != nil {
        return nil, fmt.Errorf("unmarshaling event: %w", err)
    }
    return &event, nil
}
```

### Protobuf Serialization

```protobuf
// events.proto
syntax = "proto3";

package events;

option go_package = "go-pro/events";

message LessonCompleted {
    string user_id = 1;
    int64 lesson_id = 2;
    int32 score = 3;
    int64 completed_at = 4;
}
```

```go
import (
    "google.golang.org/protobuf/proto"
    "go-pro/events"
)

func ProduceProtobufEvent(producer *KafkaProducer, topic string, event *events.LessonCompleted) error {
    data, err := proto.Marshal(event)
    if err != nil {
        return fmt.Errorf("marshaling protobuf: %w", err)
    }
    
    return producer.ProduceMessage(topic, event.UserId, string(data))
}

func ConsumeProtobufEvent(msg *kafka.Message) (*events.LessonCompleted, error) {
    var event events.LessonCompleted
    if err := proto.Unmarshal(msg.Value, &event); err != nil {
        return nil, fmt.Errorf("unmarshaling protobuf: %w", err)
    }
    return &event, nil
}
```

---


