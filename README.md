# CRMS (Customer Relationship Management System)

CRMS 是一個以前後端分離架構實作的客戶關係管理系統，採用 Nx Monorepo 管理。  
系統提供登入驗證、客戶資料管理、歷史紀錄管理、國籍資料查詢，並支援將前端建置產物整合進 Go API 服務中統一部署。

## 核心功能

- 使用者驗證
  - `JWT` 驗證（`Authorization: Bearer <token>`）
  - `HttpOnly Cookie` 驗證（Cookie 名稱：`token`）
  - 登入、驗證、登出流程完整
- 客戶管理（Customer）
  - 新增 / 修改 / 刪除 / 查詢客戶
  - 依姓名、身分證字號、電話搜尋
  - 前端有欄位驗證（生日、身分證格式、電話、車牌、國籍等）
- 歷史紀錄管理（History）
  - 新增 / 修改 / 刪除 / 查詢紀錄
  - 依日期、日期區間、客戶條件搜尋
  - 與客戶資料關聯顯示
- 國籍管理（Citizenship）
  - 查詢全部國籍、依 ID / 國名查詢
  - 啟動時自動建立表並匯入國籍資料（首次）
- 前後端整合部署
  - 前端 build 後可複製到 `apps/api/static` 與 `apps/api/templates`
  - 由 Gin 統一提供靜態資源與 SPA 入口

## 技術棧

- Monorepo: Nx
- Frontend: React 19, TypeScript, Vite, Tailwind CSS
- Backend: Go 1.23, Gin, GORM
- Database: MySQL
- Auth: JWT (HS256) + Cookie / Bearer 雙模式
- API 文件: Swagger 檔案已存在於 `apps/api/docs`（目前程式預設未啟用路由）

## 系統架構

### Frontend (`apps/frontend`)

- 單頁應用（SPA），由 `AuthContext` 管理登入狀態
- `apiService` 統一封裝所有 API 呼叫與 response normalization
- 主要頁面：
  - `LoginPage`
  - `Dashboard`
  - `CustomerManagement`
  - `HistoryManagement`
- 開發環境透過 Vite Proxy 將 `/api` 轉發到 `http://localhost:8080`

### Backend (`apps/api`)

- 採模組化分層：
  - `module/*/delivery/http`
  - `module/*/service`
  - `module/*/repository`
  - `domain`（介面）
  - `model`（資料模型）
- 啟動流程：
  - 讀取 `config.yaml`
  - 連接 MySQL
  - AutoMigrate（`citizenships`, `customers`, `histories`, `users`）
  - 首次啟動時匯入國籍資料
- 路由：
  - `/api/user*`：公開
  - 其他 `/api/*`：需通過 `AuthMiddleware`

## 專案結構

```text
.
├─ apps
│  ├─ api
│  │  ├─ module
│  │  │  ├─ user
│  │  │  ├─ customer
│  │  │  ├─ history
│  │  │  └─ citizenship
│  │  ├─ config.yaml
│  │  ├─ main.go
│  │  └─ docs
│  ├─ frontend
│  │  ├─ src/components
│  │  ├─ src/contexts
│  │  └─ src/services
│  └─ frontend-e2e
├─ scripts
│  └─ copy-frontend-to-api.mjs
├─ package.json
└─ nx.json
```

## 快速開始

### 1. 環境需求

- Node.js 20+
- pnpm 10+
- Go 1.23+
- MySQL 8+

### 2. 安裝依賴

```bash
pnpm install
```

### 3. 建立資料庫

請先建立資料庫（預設名稱 `crms`），並依需求調整 `apps/api/config.yaml`：

- `DATABASE.USERNAME`
- `DATABASE.PASSWORD`
- `DATABASE.SERVER`
- `DATABASE.PORT`
- `DATABASE.DATABASE`
- `FRONTEND_ORIGIN`
- `TOKEN_SECRET`

### 4. 建立初始登入帳號（必要）

目前程式會自動建立 `users` 資料表，但不會自動建立管理員帳號。  
請先在 MySQL 執行：

```sql
INSERT INTO users (Id, Username, Password)
VALUES (UUID(), 'admin', 'password');
```

### 5. 啟動 Backend

```bash
cd apps/api
go mod tidy
go run .
```

預設服務位址：`http://localhost:8080`

### 6. 啟動 Frontend（新終端）

```bash
pnpm --filter crms-frontend dev
```

預設前端位址：`http://localhost:4200`

## API 概覽

Base Path: `/api`（目前皆為 `POST`）

### User（公開）

- `/userLogin`
- `/userAuthentication`
- `/userLogout`

### Customer（需驗證）

- `/customerList`
- `/customerNationalId`
- `/customerCre`
- `/customerMod`
- `/customerDel`
- `/customerName`
- `/customerCitizenship`
- `/customerPhone`
- `/customerID`

### History（需驗證）

- `/historyList`
- `/historyByHistoryId`
- `/historyCre`
- `/historyMod`
- `/historyDel`
- `/historyForDuring`
- `/historyForDate`
- `/historyCustomerId`

### Citizenship（需驗證）

- `/citizenships`
- `/citizenshipId`
- `/citizenshipNation`

## 測試與開發指令

### Backend

```bash
cd apps/api
go test ./...
```

### Frontend（Nx）

```bash
pnpm exec nx test frontend
pnpm exec nx lint frontend
```

### E2E（Playwright）

```bash
pnpm exec nx e2e frontend-e2e
```

## 建置與部署

### 前端獨立建置

```bash
pnpm --filter crms-frontend build
```

### 前端整合進後端（建議部署方式）

```bash
pnpm bundle:web
```

此指令會：

- 建置 `apps/frontend`
- 複製 `dist/apps/frontend/assets` 到 `apps/api/static/assets`
- 複製 `dist/apps/frontend/favicon.ico` 到 `apps/api/static/favicon.ico`
- 複製 `dist/apps/frontend/index.html` 到 `apps/api/templates/index.html`

之後啟動 `apps/api` 即可同時提供 API 與前端畫面。

## 設定重點

- CORS 由 `apps/api/route/cors.go` 控制
- `COOKIE_SECURE=true` 時會使用 `SameSite=None`（需 HTTPS）
- 前端 API Base 預設為 `/api`
- Token 過期時間由 `TOKEN_EXPIRATION_HOURS` 控制（預設 2 小時）

## 已知事項

- `apps/api/docs/swagger.*` 已存在，但 `main.go` 目前預設未啟用 swagger 路由。
- `apps/api/config.yaml` 中 `ADMIN` 區段目前未自動寫入資料庫，需手動建立帳號。

