# Cloud Cost Resource Analyzer

A **cloud-native cost analytics platform** designed to ingest multi-project GCP metrics, process cost signals, detect anomalies, and expose actionable insights via a dashboard API.

Built with a **microservice architecture** on Google Cloud Run, the system demonstrates event-driven design, cost data processing, and scalable backend engineering patterns.

---

## 🚀 Problem Statement

Managing cloud costs across multiple GCP projects becomes increasingly complex as infrastructure scales. Native dashboards often lack centralized visibility, anomaly detection, and project-level insights.

**This project provides:**

- Multi-project onboarding  
- Centralized cost analysis  
- Anomaly detection  
- Dashboard-ready APIs  
- Event-driven cost processing pipeline  

---

## 🏗 Architecture Overview

The system follows a **microservice + event-driven architecture** deployed on Google Cloud.

> 📌 Insert architecture diagram image below this section.

### Core Flow

1. Projects are onboarded via the Upload Service.  
2. Metrics and cost signals are published via Pub/Sub.  
3. Cost Processor consumes events and processes cost data.  
4. Metrics Ingestor loads processed metrics into storage.  
5. Analysis Engine exposes aggregated data via REST APIs.  
6. Frontend Dashboard consumes analysis APIs.  

---

## 🧩 System Components

### 1️⃣ Upload Service
- Handles project onboarding  
- Publishes project events  
- Exposes Prometheus-style metrics endpoint  
- Deployed on Cloud Run  

### 2️⃣ Cost Processor
- Consumes Pub/Sub messages  
- Processes cost-related signals  
- Performs transformation & aggregation logic  

### 3️⃣ Metrics Ingestor
- Loads processed metrics into storage  
- Prepares data for analytical querying  

### 4️⃣ Analysis Engine
- Exposes REST APIs for:
  - `/dashboard`
  - `/anomalies`
  - `/insights`
  - `/projects`
  - `/add-project`
- Serves as the main backend interface for frontend  

### 5️⃣ Frontend Dashboard
- Displays:
  - Cost trends  
  - Request metrics  
  - Anomaly summaries  
  - Project insights  
- Connects directly to Analysis Engine APIs  

---

## 🛠 Tech Stack

**Backend**
- Python (FastAPI)  
- Google Cloud Run  
- Google Pub/Sub  
- PostgreSQL  

**Frontend**
- React  

**Infrastructure**
- Google Cloud Platform  
- Cloud Run (serverless deployment)  
- Pub/Sub (event-driven messaging)  
- Managed PostgreSQL  

---

## 📊 Key Features

- Multi-project onboarding  
- Event-driven cost processing  
- REST-based dashboard APIs  
- Cost anomaly detection  
- Modular microservice design  
- Prometheus-style observability endpoints  

---

## 🔍 Observability

All services expose **Prometheus-compatible metrics endpoints** for:

- Request count  
- Average latency  
- Error rates  

This enables integration with Prometheus + Grafana for production-grade monitoring.

---

## 🔄 Event-Driven Design

The system leverages **Google Pub/Sub** for asynchronous processing:

- Decoupled service communication  
- Scalable cost processing  
- Extensible pipeline for real-time analytics  

---

## 🧠 Design Considerations

- Microservice separation for modular scaling  
- Event-driven pipeline for loose coupling  
- API-first backend for frontend flexibility  
- Cloud-native deployment model  
- Stateless services with managed storage backend  

---

## 📦 Future Improvements

- Real-time multi-project metric ingestion  
- Prometheus + Grafana monitoring stack  
- Horizontal autoscaling policies  
- Multi-tenant RBAC  
- Cost forecasting models  
- Kubernetes migration for large-scale deployments  

---

# 🛠 Setup & Deployment Guide

## 🔐 Prerequisites

- Google Cloud account  
- gcloud CLI installed and authenticated  
- Terraform installed  
- Node.js (v18+)  
- Python 3.10+  

Authenticate with GCP:

```bash
gcloud auth login
gcloud config set project <YOUR_GCP_PROJECT_ID>