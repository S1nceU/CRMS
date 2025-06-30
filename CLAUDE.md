# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Customer Relationship Management System (CRMS) built as a monorepo using Nx. The project consists of:

- **Backend API** (`apps/api/`): Go-based REST API using Gin framework with MySQL database
- **Frontend** (`apps/frontend/`): React application with TypeScript, Vite, and Tailwind CSS  
- **E2E Tests** (`apps/frontend-e2e/`): Playwright end-to-end tests

## Architecture

### Backend (Go API)
- **Framework**: Gin (HTTP router)
- **Database**: MySQL with GORM ORM
- **Architecture**: Clean Architecture pattern with distinct layers:
  - `domain/`: Domain entities (User, Customer, Citizenship, History)
  - `module/*/repository/`: Data access layer
  - `module/*/service/`: Business logic layer  
  - `module/*/delivery/http/`: HTTP handlers
- **Documentation**: Swagger/OpenAPI integration
- **Configuration**: YAML-based config in `apps/api/config/config.yaml`

### Frontend (React)
- **Framework**: React 19 with TypeScript
- **Build Tool**: Vite
- **Styling**: Tailwind CSS
- **Testing**: Vitest for unit tests

## Development Commands

### Prerequisites
- Go 1.21.3+ (for backend development)
- Node.js (for frontend and tooling)
- MySQL database

### Setup
```bash
# Install dependencies
npm install

# Note: Go is required for the @nx-go plugin to work properly
```

### Backend (API)
```bash
# Build the API
npx nx build api

# Run the API in development mode
npx nx serve api

# Run API tests
npx nx test api

# Lint Go code
npx nx lint api

# Clean up Go modules
npx nx tidy api
```

### Frontend
```bash
# Serve frontend in development
npx nx serve frontend

# Build frontend for production
npx nx build frontend

# Run frontend tests
npx nx test frontend

# Lint frontend code
npx nx lint frontend

# Type check
npx nx typecheck frontend
```

### E2E Testing
```bash
# Run Playwright e2e tests
npx nx e2e frontend-e2e
```

### Database Configuration
- Default configuration is in `apps/api/config/config.yaml`
- Database connection: MySQL on localhost:3306
- Default database name: `crms`
- SQL schema script: `apps/api/config/SQLscript.sql`

### API Documentation
- Swagger UI available when running the API server
- Swagger spec: `apps/api/docs/swagger.json`
- Generated from Go code comments using swaggo

## Key Files and Patterns

### Configuration
- `nx.json`: Nx workspace configuration with plugin setup
- `apps/api/config/config.yaml`: API server and database configuration
- `apps/api/go.mod`: Go module dependencies

### Module Structure (Backend)
Each business domain follows the same pattern:
```
module/{domain}/
├── delivery/http/     # HTTP handlers
├── repository/        # Data access
└── service/          # Business logic
```

### Testing
- Frontend: Vitest configuration in `vitest.workspace.ts`
- E2E: Playwright configuration in `apps/frontend-e2e/playwright.config.ts`
- Backend: Go standard testing

## Notes
- The project uses Nx for monorepo management with separate plugins for Go and React
- Go runtime is required for the @nx-go plugin to function properly
- Static assets are served from `apps/api/static/` including built frontend files
- CORS is configured in `apps/api/route/cors.go`