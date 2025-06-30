# 🚀 CRMS 啟動指南

## ✅ CORS 已修復！

### 修復內容：

1. **Vite 配置** (`vite.config.ts`)：
   - ✅ 啟用 CORS：`cors: true`
   - ✅ API 代理：`/api` → `http://localhost:8080`
   - ✅ 跨域設定：`changeOrigin: true`

2. **API 服務** (`services/api.ts`)：
   - ✅ 使用代理路徑：`/api` (透過 Vite 代理到後端)

## 🔧 啟動步驟

### 1. 啟動 API 服務器
```bash
# 在專案根目錄
npx nx serve api
```
API 會運行在：`http://127.0.0.1:8080`

### 2. 啟動前端（三種方式任選一種）

**方式 1：使用腳本**
```bash
./fix-nx-and-run.sh
# 選擇選項 2
```

**方式 2：直接使用 npm**
```bash
cd apps/frontend
npm run dev
```

**方式 3：手動 Vite**
```bash
cd apps/frontend
NX_DAEMON=false npx vite --port 4200 --host 0.0.0.0 --cors
```

### 3. 訪問應用

- **前端**: http://localhost:4200
- **API**: http://localhost:8080
- **Swagger**: http://localhost:8080/swagger/index.html

## 🔄 CORS 工作原理

```
前端 (localhost:4200) → Vite 代理 → API (localhost:8080)
     /api/* 請求      →    代理轉發    →    實際 API
```

### 優點：
- ✅ 無 CORS 問題
- ✅ 開發環境友好
- ✅ 生產環境可配置不同 API URL

## 🐛 問題排除

### 如果還有 CORS 錯誤：

1. **確認 API 正在運行**：
   ```bash
   curl http://localhost:8080/api/userLogin
   ```

2. **測試 Citizenship API**：
   ```bash
   node test-citizenship-api.js
   ```

3. **確認 Vite 代理**：
   檢查瀏覽器開發者工具的 Network 標籤，確認請求是發到 `localhost:4200/api/*`

4. **重啟前端服務器**：
   ```bash
   # Ctrl+C 停止，然後重新啟動
   npm run dev
   ```

### Citizenship 資料載入問題：

1. **檢查 API 回應**：
   - 開啟瀏覽器開發者工具
   - 查看 Network 標籤中的 `/api/citizenships` 請求
   - 確認回應包含國籍資料陣列

2. **檢查控制台訊息**：
   - 應該看到 "Loaded X citizenships from API"
   - 如果看到錯誤訊息，檢查 API 服務器狀態

## 📝 注意事項

- 前端必須通過 `http://localhost:4200` 訪問（不要用 IP）
- API 必須運行在 `localhost:8080`
- 確保兩個服務都在運行

現在 CORS 問題已經完全解決！🎉