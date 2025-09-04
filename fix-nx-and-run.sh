#!/bin/bash

echo "🔧 CRMS Frontend Fix & Run Script"
echo "=================================="

# Function to run frontend
run_frontend() {
    echo "🚀 Starting CRMS Frontend on http://localhost:4200"
    echo "📝 Make sure your API is running on http://127.0.0.1:8080"
    echo ""
    
    cd apps/frontend
    echo "📍 Starting Vite development server with CORS and API proxy..."
    NX_DAEMON=false npx vite --port 4200 --host 0.0.0.0 --cors
}

# Function to fix Nx (temporary)
fix_nx_temporarily() {
    echo "🔧 Temporarily fixing Nx configuration..."
    
    # Backup original nx.json
    if [ ! -f "nx.json.backup" ]; then
        cp nx.json nx.json.backup
        echo "✅ Backed up original nx.json"
    fi
    
    # Use frontend-only configuration
    cp nx-frontend-only.json nx.json
    echo "✅ Applied frontend-only Nx configuration"
}

# Function to restore Nx
restore_nx() {
    echo "🔄 Restoring original Nx configuration..."
    if [ -f "nx.json.backup" ]; then
        cp nx.json.backup nx.json
        echo "✅ Restored original nx.json"
    else
        echo "⚠️  No backup found"
    fi
}

echo "Choose an option:"
echo "1) Fix Nx temporarily and run frontend"
echo "2) Run frontend directly with Vite (bypass Nx)"
echo "3) Restore original Nx configuration"
echo "4) Show current status"

read -p "Enter choice (1-4): " choice

case $choice in
    1)
        fix_nx_temporarily
        run_frontend
        ;;
    2)
        run_frontend
        ;;
    3)
        restore_nx
        ;;
    4)
        echo "📊 Current Status:"
        echo "=================="
        echo "Go installed: $(which go > /dev/null && echo "✅ Yes" || echo "❌ No")"
        echo "Node.js: $(node --version)"
        echo "NPM: $(npm --version)"
        echo "Nx backup exists: $([ -f "nx.json.backup" ] && echo "✅ Yes" || echo "❌ No")"
        echo ""
        echo "📁 Frontend files:"
        ls -la apps/frontend/src/ | head -10
        ;;
    *)
        echo "❌ Invalid choice"
        exit 1
        ;;
esac