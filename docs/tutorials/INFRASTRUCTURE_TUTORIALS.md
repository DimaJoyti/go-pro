# 🏗️ Infrastructure & DevOps Tutorials - Master Index

Complete guide to infrastructure, databases, messaging, and CI/CD for Go applications.

---

## 📚 Available Tutorials

### 🗄️ Databases

#### [PostgreSQL with Go](./postgresql-tutorial.md)
**Duration:** 6-8 hours | **Difficulty:** Intermediate

Master PostgreSQL integration in Go applications.

**What You'll Learn:**
- Connect to PostgreSQL with `database/sql` and `pgx`
- Implement CRUD operations
- Use prepared statements and transactions
- Handle NULL values and complex types
- Implement connection pooling
- Repository pattern for clean architecture
- Database migrations
- Testing database code

**Key Topics:**
- Basic CRUD operations
- Advanced queries and joins
- Full-text search
- Transaction management
- Connection pooling configuration
- Repository pattern implementation
- Migration strategies
- Testing with test databases

**Prerequisites:**
- Go basics (Tutorial 1-5)
- SQL fundamentals
- Understanding of database concepts

---

#### [Redis with Go](./redis-tutorial.md)
**Duration:** 5-6 hours | **Difficulty:** Intermediate

Master Redis for caching, sessions, and real-time features.

**What You'll Learn:**
- Connect to Redis with go-redis
- Use Redis data structures (strings, hashes, lists, sets, sorted sets)
- Implement caching patterns
- Pub/Sub messaging
- Distributed locks
- Rate limiting
- Leaderboards
- Session management

**Key Topics:**
- String operations and counters
- Hash operations for user profiles
- Lists for activity feeds and queues
- Sets for tags and followers
- Sorted sets for leaderboards
- Cache-aside pattern
- Write-through cache
- Pub/Sub messaging
- Distributed locks with Redlock

**Prerequisites:**
- Go basics (Tutorial 1-5)
- Understanding of caching concepts
- Basic Redis knowledge

---

### 📨 Messaging

#### [Apache Kafka with Go](./kafka-tutorial.md)
**Duration:** 7-8 hours | **Difficulty:** Advanced

Master event streaming with Apache Kafka for scalable data pipelines.

**What You'll Learn:**
- Set up Kafka producers and consumers
- Implement event-driven architectures
- Handle message serialization (JSON, Avro, Protobuf)
- Consumer groups for scalability
- Error handling and retries
- Exactly-once semantics
- Monitoring Kafka applications
- Build real-time data pipelines

**Key Topics:**
- Producer basics and async production
- Consumer basics and consumer groups
- Message serialization formats
- Error handling and dead letter queues
- Exactly-once processing
- Kafka Streams
- Schema Registry integration
- Monitoring with Prometheus

**Prerequisites:**
- Go intermediate (Tutorial 6-10)
- Understanding of distributed systems
- Event-driven architecture concepts

---

### ☁️ Cloud Platforms

#### [AWS for Go Applications](./aws-tutorial.md)
**Duration:** 8-10 hours | **Difficulty:** Advanced

Master AWS services for deploying and scaling Go applications.

**What You'll Learn:**
- Use AWS SDK for Go (v2)
- S3 for file storage
- DynamoDB for NoSQL data
- Lambda functions in Go
- SQS and SNS for messaging
- Secrets Manager for credentials
- CloudWatch for monitoring
- ECS/Fargate deployment

**Key Topics:**
- AWS SDK v2 setup and configuration
- S3 operations (upload, download, presigned URLs)
- DynamoDB CRUD and queries
- Lambda function development
- SQS message queues
- SNS notifications
- Secrets Manager integration
- CloudWatch logs and metrics
- IAM roles and policies

**Prerequisites:**
- Go advanced (Tutorial 11-15)
- AWS account
- Basic AWS knowledge
- Docker basics

---

#### [Google Cloud Platform for Go](./gcp-tutorial.md)
**Duration:** 8-10 hours | **Difficulty:** Advanced

Master GCP services for deploying and scaling Go applications.

**What You'll Learn:**
- Use Google Cloud client libraries
- Cloud Storage for files
- Firestore for NoSQL data
- Cloud Pub/Sub for messaging
- Cloud Functions in Go
- Secret Manager
- Cloud Run deployment
- Cloud Monitoring and Logging

**Key Topics:**
- GCP SDK setup
- Cloud Storage operations
- Firestore CRUD and real-time listeners
- Pub/Sub publish and subscribe
- Cloud Functions development
- Secret Manager integration
- Cloud Run deployment
- Stackdriver monitoring
- Cloud Build CI/CD

**Prerequisites:**
- Go advanced (Tutorial 11-15)
- GCP account
- Basic GCP knowledge
- Docker basics

---

### 🏗️ Infrastructure as Code

#### [Terraform for Go Applications](./terraform-tutorial.md)
**Duration:** 8-10 hours | **Difficulty:** Advanced

Master Infrastructure as Code with Terraform for Go deployments.

**What You'll Learn:**
- Write Terraform configurations
- Provision AWS infrastructure (VPC, ECS, RDS, ElastiCache)
- Provision GCP infrastructure (Cloud Run, Cloud SQL)
- Kubernetes deployment
- Modules for reusability
- State management
- CI/CD integration
- Multi-cloud deployments

**Key Topics:**
- Terraform basics and HCL syntax
- AWS infrastructure (VPC, subnets, security groups)
- ECS/Fargate for containers
- RDS PostgreSQL setup
- ElastiCache Redis setup
- GCP Cloud Run deployment
- Kubernetes manifests
- Terraform modules
- Remote state with S3/GCS
- Workspaces for environments

**Prerequisites:**
- Go advanced (Tutorial 11-15)
- Cloud platform knowledge (AWS or GCP)
- Docker and Kubernetes basics
- Infrastructure concepts

---

### 🚀 CI/CD

#### [GitHub Actions for Go](./github-actions-tutorial.md)
**Duration:** 6-8 hours | **Difficulty:** Intermediate-Advanced

Master CI/CD with GitHub Actions for automated testing and deployment.

**What You'll Learn:**
- Set up CI/CD pipelines
- Automate testing and quality checks
- Build and publish Docker images
- Deploy to AWS, GCP, and Kubernetes
- Security scanning
- Release automation
- Matrix builds
- Caching strategies

**Key Topics:**
- Basic workflow syntax
- Matrix builds for multiple Go versions
- Testing with services (PostgreSQL, Redis)
- Code quality checks (golangci-lint, staticcheck)
- Docker build and push
- AWS ECS deployment
- GCP Cloud Run deployment
- Kubernetes deployment
- Security scanning (Gosec, Trivy)
- Release automation
- Secrets management

**Prerequisites:**
- Go basics (Tutorial 1-10)
- Git and GitHub knowledge
- Docker basics
- YAML syntax

---

## 🎯 Learning Paths

### Path 1: Database Specialist (2-3 weeks)
**Focus:** Master database integration

**Week 1:** PostgreSQL Tutorial
- Complete all PostgreSQL exercises
- Build a CRUD application
- Implement repository pattern

**Week 2:** Redis Tutorial
- Complete all Redis exercises
- Implement caching layer
- Build real-time features

**Week 3:** Integration Project
- Combine PostgreSQL and Redis
- Build a scalable application
- Implement best practices

---

### Path 2: Cloud Engineer (3-4 weeks)
**Focus:** Master cloud deployments

**Week 1:** Terraform Tutorial
- Learn IaC fundamentals
- Provision AWS infrastructure
- Provision GCP infrastructure

**Week 2:** AWS Tutorial
- Master AWS SDK
- Deploy to ECS/Fargate
- Implement AWS services

**Week 3:** GCP Tutorial
- Master GCP client libraries
- Deploy to Cloud Run
- Implement GCP services

**Week 4:** Multi-Cloud Project
- Deploy to both AWS and GCP
- Implement disaster recovery
- Cost optimization

---

### Path 3: DevOps Engineer (2-3 weeks)
**Focus:** Master CI/CD and automation

**Week 1:** GitHub Actions Tutorial
- Set up CI/CD pipelines
- Automate testing
- Security scanning

**Week 2:** Terraform Tutorial
- Infrastructure as Code
- Multi-environment setup
- State management

**Week 3:** Full Pipeline Project
- End-to-end CI/CD
- Automated deployments
- Monitoring and alerting

---

### Path 4: Event-Driven Architect (2 weeks)
**Focus:** Master event-driven systems

**Week 1:** Kafka Tutorial
- Event streaming basics
- Producer and consumer patterns
- Schema management

**Week 2:** Integration Project
- Build event-driven microservices
- Implement CQRS pattern
- Event sourcing

---

### Path 5: Full-Stack Infrastructure (6-8 weeks)
**Complete all tutorials in order**

**Weeks 1-2:** Databases
- PostgreSQL
- Redis

**Weeks 3-4:** Cloud Platforms
- AWS
- GCP

**Weeks 5-6:** Infrastructure & CI/CD
- Terraform
- GitHub Actions

**Weeks 7-8:** Messaging & Projects
- Kafka
- Capstone project

---

## 📊 Tutorial Comparison

| Tutorial | Duration | Difficulty | Prerequisites | Cloud | Local Dev |
|----------|----------|------------|---------------|-------|-----------|
| PostgreSQL | 6-8h | Intermediate | Go basics | ✅ | ✅ |
| Redis | 5-6h | Intermediate | Go basics | ✅ | ✅ |
| Kafka | 7-8h | Advanced | Go intermediate | ✅ | ✅ |
| AWS | 8-10h | Advanced | Go advanced | ✅ | ❌ |
| GCP | 8-10h | Advanced | Go advanced | ✅ | ❌ |
| Terraform | 8-10h | Advanced | Cloud knowledge | ✅ | ✅ |
| GitHub Actions | 6-8h | Intermediate-Advanced | Git/Docker | ✅ | ❌ |

---

## 🛠️ Setup Requirements

### Local Development

**Required:**
- Go 1.22+
- Docker and Docker Compose
- Git

**Optional:**
- PostgreSQL client (psql)
- Redis CLI
- Kafka CLI tools
- Terraform CLI
- kubectl

### Cloud Accounts

**AWS:**
- AWS account
- AWS CLI configured
- IAM user with appropriate permissions

**GCP:**
- GCP account
- gcloud CLI configured
- Service account with appropriate permissions

**GitHub:**
- GitHub account
- Repository access
- Secrets configured for deployments

---

## 📁 Project Structure

```
go-pro/
├── docs/
│   └── tutorials/
│       ├── INFRASTRUCTURE_TUTORIALS.md  # This file
│       ├── postgresql-tutorial.md
│       ├── redis-tutorial.md
│       ├── kafka-tutorial.md
│       ├── aws-tutorial.md
│       ├── gcp-tutorial.md
│       ├── terraform-tutorial.md
│       └── github-actions-tutorial.md
├── examples/
│   ├── postgresql/
│   ├── redis/
│   ├── kafka/
│   ├── aws/
│   └── gcp/
├── terraform/
│   ├── aws/
│   ├── gcp/
│   └── modules/
└── .github/
    └── workflows/
```

---

## 🎓 Certification Path

After completing all tutorials, you'll be able to:

✅ Design and implement scalable database architectures  
✅ Build event-driven microservices  
✅ Deploy applications to AWS and GCP  
✅ Implement Infrastructure as Code  
✅ Set up production-grade CI/CD pipelines  
✅ Monitor and troubleshoot distributed systems  
✅ Implement security best practices  
✅ Optimize costs and performance  

---

## 📚 Additional Resources

### Documentation
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Redis Documentation](https://redis.io/documentation)
- [Kafka Documentation](https://kafka.apache.org/documentation/)
- [AWS Documentation](https://docs.aws.amazon.com/)
- [GCP Documentation](https://cloud.google.com/docs)
- [Terraform Documentation](https://www.terraform.io/docs)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

### Community
- [Go Forum](https://forum.golangbridge.org/)
- [r/golang](https://reddit.com/r/golang)
- [Gophers Slack](https://gophers.slack.com/)

---

**Ready to start?** Choose your learning path and dive into the tutorials! 🚀

