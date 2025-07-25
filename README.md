tenant-micro-builerplate/
├── README.md                    # Project overview, architecture diagram
├── infra/                       # CDK or Terraform files to create AWS infra
│   ├── main.ts / main.tf        # Entry point to define stacks
│   └── resources/               # Cognito, DynamoDB, API Gateway, etc.
├── shared/                      # Common code (e.g., auth middleware, utils)
│   └── tenant-context.ts
├── services/
│   ├── auth-service/            # Lambda code to handle signup/login + token validation
│   │   └── handler.ts
│   ├── tenant-service/          # Register/store tenant info
│   ├── product-service/         # Add/view/update tenant-specific products
│   └── order-service/           # Place/view tenant orders
├── frontend/                    # React/Vue frontend for admin/customer
│   └── src/
└── .github/
    └── workflows/               # GitHub Actions CI/CD for lint, test, deploy
        └── deploy.yml