# 📈 Time Series Database

Custom time-series database with compression, query language, aggregation functions, and Grafana integration.

## 🎯 Features

- ✅ Efficient time-series storage
- ✅ Data compression (Gorilla algorithm)
- ✅ Custom query language
- ✅ Aggregation functions
- ✅ Retention policies
- ✅ Grafana integration
- ✅ High write throughput
- ✅ Fast queries

## 🏗️ Architecture

```
timeseries-db/
├── cmd/
│   ├── server/
│   └── cli/
├── internal/
│   ├── storage/
│   │   ├── engine.go
│   │   ├── compression.go
│   │   └── index.go
│   ├── query/
│   │   ├── parser.go
│   │   └── executor.go
│   └── api/
└── README.md
```

## 📖 Query Language

```sql
SELECT mean(temperature) 
FROM sensors 
WHERE location = 'room1' 
AND time > now() - 1h 
GROUP BY time(5m)
```

## 🎓 Learning Objectives

- Time-series storage
- Data compression
- Query optimization
- Indexing strategies
- Aggregation functions

---

**Status**: Planned | **Difficulty**: Advanced | **Time**: 25-35 hours

