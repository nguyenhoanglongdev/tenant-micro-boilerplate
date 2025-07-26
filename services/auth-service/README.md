# Auth Service

The **Auth Service** acts as a middleware layer between client applications and AWS Cognito (User Pool). It handles standard authentication flows as well as tenant-specific logic in a multi-tenant SaaS context.

---

##  Features

-  **Authentication**
  - User registration
  - User login
  - Token validation
  - Token refresh
-  **Tenant Determination**
  - Identify tenant based on email domain, custom header, or subdomain
  - Inject tenant context for downstream services
-  **Pluggable Server Design**
  - Supports deployment as AWS Lambda for cost-efficiency and scalability
  - Can switch to a dedicated HTTP server without changing core logic

---
auth-service/
├── cmd/ # Entry points
│ ├── lambda/ # AWS Lambda handler (e.g., API Gateway proxy)
│ └── server/ # Optional HTTP server entry point (Gin or other)
├── internal/
│ ├── utils/ #
│ ├── config/ # App configuration, environment parsing
│ ├── handler/ # HTTP handlers for auth endpoints
│ ├── provider/
│ │ └── cognito/ # AWS Cognito integration
│ │ └── userpool.go # Pluggable user pool abstraction
│ ├── router/ # Gin router and middleware setup
│ └── service/ # Business logic layer
└── go.mod