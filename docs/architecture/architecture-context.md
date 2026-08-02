# Architecture Context

This document summarizes the architecture of DapurPintar AI and serves as a quick reference for developers and AI assistants.

---

# Architecture Style

- Modular Monolith
- Clean Architecture
- Domain Driven Design

---

# System Layers

Presentation

↓

Application

↓

Domain

↓

Repository

↓

Infrastructure

---

# Core Domain

AI Kitchen Intelligence

Responsible for:

- Recommendation
- Personalization
- Decision Support
- Ingredient Analysis

---

# Supporting Domains

- Pantry Management
- Recipe Management
- Meal Planning
- Shopping Planning
- User Profile

Future:

- Nutrition
- Household Collaboration

---

# Generic Domains

- Identity
- Notifications
- Observability
- External Integrations
- Billing

---

# Technology

Frontend

- Next.js
- React
- TypeScript
- Tailwind

Backend

- Go
- Fiber

Storage

- PostgreSQL
- Redis

AI

- AI Gateway
- OpenAI

Observability

- OpenTelemetry
- Grafana

---

# Development Principles

- Keep modules independent.
- Keep business logic inside Domain/Application.
- Infrastructure depends on Domain.
- Never reverse dependencies.
- Every module owns its own language.

---

# Repository Standards

- SQLC
- Goose
- PostgreSQL

---

# API Standards

- REST
- Versioned API
- JSON

---

# Observability

Every module should support:

- Logging
- Metrics
- Tracing

through OpenTelemetry.

---

# Current Architecture Status

Completed

- Architecture Vision
- ADR
- Domain Discovery
- Bounded Context

In Progress

- Tactical DDD

Upcoming

- Database Design
- API Design
- Authentication
- AI Architecture
- Deployment