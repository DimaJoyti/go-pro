# 🐳 Container Orchestrator (Mini Kubernetes)

Simplified container orchestrator with pod scheduling, service networking, and resource management.

## 🎯 Features

- ✅ Container lifecycle management
- ✅ Pod scheduling
- ✅ Service networking
- ✅ Health checks and auto-restart
- ✅ Resource limits (CPU, memory)
- ✅ Volume management
- ✅ kubectl-like CLI
- ✅ YAML configuration

## 🏗️ Architecture

```
container-orchestrator/
├── cmd/
│   ├── orchestrator/
│   └── ctl/
├── internal/
│   ├── scheduler/
│   ├── runtime/
│   ├── network/
│   └── storage/
└── README.md
```

## 📖 Usage

```bash
# Deploy pod
./ctl apply -f pod.yaml

# List pods
./ctl get pods

# View logs
./ctl logs mypod

# Execute command
./ctl exec mypod -- /bin/sh
```

## 🎓 Learning Objectives

- Container management
- Pod scheduling
- Service networking
- Resource management
- Health monitoring

---

**Status**: Planned | **Difficulty**: Advanced | **Time**: 30-40 hours

