#!/bin/bash

echo "🚀 Starting CRMS Frontend..."
echo "📍 Current directory: $(pwd)"

# Method 1: Try with Nx but disable daemon and Go plugin
echo "🔄 Attempting to start with Nx (bypassing Go plugin)..."
cd apps/frontend
NX_DAEMON=false npx vite --port 4200 --host 0.0.0.0

# If that fails, try direct vite
# echo "🔄 Fallback: Starting with Vite directly..."
# npx vite --config vite.config.ts --port 4200 --host 0.0.0.0