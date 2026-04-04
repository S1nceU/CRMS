# CRMS Backend (apps/api)

此文件專注於 CRMS 後端系統，說明目前實作的架構分層、設計理念、模組責任、請求流程與維運方式。  
目標是讓新加入的後端工程師可以在最短時間內理解「為什麼這樣設計」與「如何安全擴充」。

## 1. 系統定位

CRMS Backend 是一個以 Go + Gin + GORM 實作的 API 服務，負責：

- 使用者登入與 JWT 驗證
- 客戶資料管理（Customer）
- 住宿/消費歷史紀錄管理（History）
- 國籍資料查詢（Citizenship）
- 提供前端靜態檔與 SPA fallback（`/static` + `templates/index.html`）

## 2. 技術棧與核心依賴

- 語言/Runtime: Go 1.23
- Web Framework: Gin
- ORM: GORM + MySQL Driver
- 設定管理: Viper
- 認證: `github.com/golang-jwt/jwt/v5`
- API 文件: Swagger（文件檔存在 `docs/`，路由目前預設未開啟）

## 3. 設計理念

### 3.1 分層與責任分離

每個業務模組採三層結構：

- `delivery/http`: HTTP 轉換層（解 request / 回 response / status code）
- `service`: 業務規則與驗證（商業邏輯集中）
- `repository`: 資料存取（GORM + SQL 行為）

搭配 `domain` 介面做依賴反轉，避免 service 直接依賴具體 repository 實作。

### 3.2 以「模組」為邊界，而非以「技術」為邊界

模組切分為：

- `user`
- `customer`
- `history`
- `citizenship`

每個模組內保有完整 delivery/service/repository，降低跨檔案跳轉成本，也讓功能擴充時維持高內聚。

### 3.3 明確的進入點與組裝點

`main.go` 是唯一組裝點：

- 讀設定
- 建立 DB 連線與 migration
- 建立 repository $\rightarrow$ service $\rightarrow$ handler
- 掛載 public/protected 路由
- 啟動與graceful shutdown

這種做法讓依賴注入路徑可追蹤、可測試、可替換。

## 4. 目錄結構

```text
apps/api
├─ main.go                    # 程式入口、DI 組裝、HTTP Server lifecycle
├─ config.yaml                # 服務設定
├─ config/
│  ├─ config.go               # Viper 載入與熱更新
│  └─ sqlscript.go            # 首次啟動國籍資料匯入
├─ domain/                    # Repository/Service 介面定義
├─ model/
│  ├─ *.go                    # GORM model
│  └─ dto/                    # Request/Response DTO
├─ module/
│  ├─ user/
│  ├─ customer/
│  ├─ history/
│  └─ citizenship/
│     ├─ delivery/http/
│     ├─ service/
│     └─ repository/
├─ route/
│  ├─ auth.go                 # JWT middleware
│  ├─ cors.go                 # CORS middleware
│  └─ route.go                # 靜態檔與 SPA fallback
├─ docs/                      # swagger 產物
├─ static/                    # 前端 build 輸出(assets/favicon)
└─ templates/index.html       # SPA 入口頁
```

## 5. 啟動流程與生命週期

## 5.1 初始化流程（`init()`）

1. `config.Init()` 載入 `config.yaml`
2. 依設定組裝 MySQL DSN
3. 建立 GORM 連線
4. AutoMigrate:
   - `citizenships`
   - `customers`
   - `histories`
   - `users`
5. 若 `citizenships` 表不存在，執行 `SQLscript.sql` 匯入國籍種子資料

## 5.2 啟動流程（`main()`）

1. 設定 Gin mode（`config.Val.Mode`）
2. 建立 router + CORS middleware
3. 組裝所有 repository/service/handler
4. 建立路由群組：
   - `public`：`/api/userLogin`, `/api/userAuthentication`, `/api/userLogout`
   - `protected`：其餘 API，套用 `AuthMiddleware`
5. 掛載靜態檔與 `NoRoute -> templates/index.html`
6. 啟動 HTTP server
7. 監聽中斷信號，執行 graceful shutdown + 關閉 DB

## 6. 請求處理模型

以「建立客戶」為例：

1. `delivery/http` 將 JSON 綁定成 DTO
2. DTO 轉換成 `model.Customer`
3. `service` 做業務驗證（必要欄位、性別、生日、國籍等）
4. `repository` 用 GORM 寫入 DB
5. `delivery/http` 回傳結果與 HTTP code

### 依賴方向（重要）

`delivery -> service -> repository -> database`

不允許反向依賴，避免邏輯外溢到 handler 或 repository。

## 7. 認證與授權設計

## 7.1 Token 來源

`AuthMiddleware` 依序嘗試：

1. Cookie `token`
2. Header `Authorization: Bearer <token>`

若都不存在或驗證失敗，回傳 `401 Unauthorized`。

## 7.2 JWT 規格

- 演算法: `HS256`
- Claims: `username`, `exp`, `iss`
- Secret: `TOKEN_SECRET`
- Issuer: `TOKEN_ISSUER`
- 期限: `TOKEN_EXPIRATION_HOURS`（<=0 時 fallback 2 小時）

## 7.3 Cookie 策略

登入成功會設定 `HttpOnly` cookie：

- `Secure` 由 `COOKIE_SECURE` 決定
- 只有 `COOKIE_SECURE=true` 時才使用 `SameSite=None`（符合瀏覽器規範）

## 8. 模組職責說明

## 8.1 User

- 功能：登入、Token 驗證、登出
- 邏輯重點：
  - `Login` 讀取 DB 使用者比對密碼
  - 產生 JWT，並同時回傳 JSON token + 設定 cookie
  - `Authentication` 專責解析 JWT

## 8.2 Customer

- 功能：客戶 CRUD + 多條件查詢（姓名/身分證/電話/國籍）
- 邏輯重點：
  - service 層統一檢查必要欄位與格式條件
  - 刪除客戶時先刪關聯 `histories`（目前由 repository 手動處理）

## 8.3 History

- 功能：歷史紀錄 CRUD + 依日期/區間/客戶查詢
- 邏輯重點：
  - 建立/更新前需確認客戶存在
  - 日期查詢禁止未來日期
  - 區間查詢需 `start <= end`

## 8.4 Citizenship

- 功能：國籍主資料查詢
- 邏輯重點：
  - 系統啟動時可自動匯入種子資料
  - 前端表單依此資料建立下拉選單

## 9. 資料模型（摘要）

- `Customer`
  - PK: `Id (uuid)`
  - FK: `CitizenshipId -> Citizenship.Id`
  - 與 `History` 一對多
- `History`
  - PK: `Id (uuid)`
  - FK: `CustomerId -> Customer.Id`
- `Citizenship`
  - 唯一欄位: `Id`, `Nation`, `Alpha3`
- `User`
  - PK: `Id (uuid)`
  - 唯一欄位: `Username`

## 10. 設定檔說明（`config.yaml`）

### 10.1 Server

- `MODE`: `debug` / `release`
- `PORT`: API port
- `FRONTEND_ORIGIN`: CORS 允許來源
- `COOKIE_SECURE`: 是否啟用 Secure cookie

### 10.2 Database

- `DATABASE.USERNAME`
- `DATABASE.PASSWORD`
- `DATABASE.NETWORK`
- `DATABASE.SERVER`
- `DATABASE.PORT`
- `DATABASE.DATABASE`

### 10.3 Token

- `TOKEN_SECRET`
- `TOKEN_ISSUER`
- `TOKEN_EXPIRATION_HOURS`

## 11. API 風格與約定

- Base Path: `/api`
- 目前所有 endpoint 皆使用 `POST`（包含查詢）
- 回應多採 `{ "Message": "...", ...data }` 結構
- 錯誤處理目前是字串訊息導向（`err.Error()`）

這是既有契約，若要改為更標準 REST（GET/POST/PUT/DELETE + 錯誤碼物件），建議以版本化或兼容層逐步演進。

## 12. 本地開發

### 12.1 先決條件

- Go 1.23+
- MySQL 8+
- 已建立 `crms` database

### 12.2 啟動

```bash
cd apps/api
go mod tidy
go run .
```

預設：`http://localhost:8080`

### 12.3 測試

```bash
cd apps/api
go test ./...
```

## 13. 與前端整合部署

前端建置後可複製到後端靜態目錄，由 Gin 同源提供：

- `static/assets/*`
- `static/favicon.ico`
- `templates/index.html`

整合指令在 repo root 執行：

```bash
pnpm bundle:web
```

## 14. 已知限制與技術債

- `config.yaml` 內 `ADMIN` 區段目前未自動建立使用者，需手動建帳號
- Swagger 文件已生成，但 `main.go` 預設不啟用 swagger route
- 錯誤訊息以字串比對為主，建議演進為 typed error + 統一錯誤碼
- 查詢 API 全為 POST，若對外開放建議逐步 RESTful 化

## 15. 擴充建議（維持既有架構）

新增新模組（例如 `booking`）時，建議遵循：

1. 在 `domain/booking.go` 先定義 repository/service 介面
2. 新增 `module/booking/{repository,service,delivery/http}`
3. 在 `main.go` 組裝依賴並掛路由
4. 將驗證放 service，HTTP 轉換放 delivery
5. 以 table-driven tests 優先覆蓋 service 邏輯

這樣可以持續保持「邏輯集中、依賴單向、模組高內聚」的後端架構特性。

