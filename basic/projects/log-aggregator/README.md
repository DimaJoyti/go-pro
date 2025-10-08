# 📊 Log Aggregation System

Collect, parse, and analyze logs from multiple sources with real-time streaming and full-text search.

## 🎯 Features

- ✅ Multi-source log collection
- ✅ Real-time log streaming
- ✅ Full-text search (Elasticsearch)
- ✅ Log parsing and structuring
- ✅ Pattern-based alerting
- ✅ Web UI for visualization
- ✅ Retention policies
- ✅ Export capabilities

## 🏗️ Architecture

```
log-aggregator/
├── cmd/
│   ├── collector/
│   ├── processor/
│   └── api/
├── internal/
│   ├── collector/
│   ├── parser/
│   ├── storage/
│   └── search/
└── web/
    └── dashboard/
```

## 📖 Usage

```bash
# Start collector
./collector --source /var/log/app.log

# Query logs
./logctl search "error" --since 1h

# Stream logs
./logctl tail --follow
```

## 🎓 Learning Objectives

- Log parsing
- Real-time streaming
- Full-text search
- Data aggregation
- WebSocket communication

---

**Status**: Planned | **Difficulty**: Intermediate | **Time**: 12-16 hours

