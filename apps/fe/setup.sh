# ============================================
# SETUP SCRIPT
# ============================================
# File: setup.sh
#!/bin/bash

echo "🚀 Setting up Skillsier Monorepo..."

# Check if pnpm is installed
if ! command -v pnpm &> /dev/null; then
    echo "📦 Installing pnpm..."
    npm install -g pnpm@9.15.0
fi

# Install dependencies
echo "📦 Installing dependencies..."
pnpm install

# Setup husky
echo "🪝 Setting up git hooks..."
pnpm prepare

# Generate necessary files
echo "📝 Generating config files..."
pnpm turbo gen

echo "✅ Setup complete!"
echo ""
echo "Next steps:"
echo "  1. Copy .env.example files and configure"
echo "  2. Run 'pnpm dev:web' for web development"
echo "  3. Run 'pnpm dev:mobile' for mobile development"
