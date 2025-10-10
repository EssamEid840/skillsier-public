#!/bin/bash

# ============================================
# SKILLSIER ENVIRONMENT SETUP SCRIPT
# ============================================

set -e

echo "🔧 Setting up Skillsier environment files..."
echo ""

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Production Keycloak URL
KEYCLOAK_PROD_URL="https://keycloak.skillsier.com"
KEYCLOAK_REALM="skillsier"

# Determine API URL based on environment
read -p "Are you setting up for (1) Local development or (2) Kubernetes? [1/2]: " env_choice

if [ "$env_choice" = "2" ]; then
  API_URL="http://users-be-service.default.svc.cluster.local:8080/api"
  echo -e "${BLUE}Using Kubernetes internal service URL${NC}"
else
  API_URL="http://localhost:8080/api"
  echo -e "${BLUE}Using local development URL${NC}"
fi

# ============================================
# WEB ENVIRONMENT (.env.local)
# ============================================

echo ""
echo -e "${GREEN}Creating Web environment file...${NC}"

cat > apps/web/.env.local << EOF
# API Configuration
NEXT_PUBLIC_API_URL=${API_URL}

# Keycloak Configuration (Production)
NEXT_PUBLIC_KEYCLOAK_URL=${KEYCLOAK_PROD_URL}
NEXT_PUBLIC_KEYCLOAK_REALM=${KEYCLOAK_REALM}
NEXT_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-fe

# Server-side Keycloak Configuration
KEYCLOAK_ISSUER_URL=${KEYCLOAK_PROD_URL}/realms/${KEYCLOAK_REALM}
KEYCLOAK_CLIENT_SECRET=your-client-secret-here

# Management Client (for registration APIs)
KEYCLOAK_MGMT_CLIENT_ID=skillsier-bff
KEYCLOAK_MGMT_CLIENT_SECRET=your-mgmt-client-secret-here

# App Configuration
NEXT_PUBLIC_APP_URL=http://localhost:3000
NODE_ENV=development
EOF

echo -e "${GREEN}✓ Web environment created: apps/web/.env.local${NC}"

# ============================================
# MOBILE ENVIRONMENT (.env)
# ============================================

echo ""
echo -e "${GREEN}Creating Mobile environment file...${NC}"

cat > apps/mobile/.env << EOF
# API Configuration
EXPO_PUBLIC_API_URL=${API_URL}

# Keycloak Configuration (Production)
EXPO_PUBLIC_KEYCLOAK_URL=${KEYCLOAK_PROD_URL}
EXPO_PUBLIC_KEYCLOAK_REALM=${KEYCLOAK_REALM}
EXPO_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-mobile

# App Configuration
EXPO_PUBLIC_APP_NAME=Skillsier
EOF

echo -e "${GREEN}✓ Mobile environment created: apps/mobile/.env${NC}"

# ============================================
# INSTRUCTIONS
# ============================================

echo ""
echo -e "${YELLOW}⚠️  IMPORTANT: Update the following values in your .env files:${NC}"
echo ""
echo "📝 apps/web/.env.local:"
echo "   - KEYCLOAK_CLIENT_SECRET (get from Keycloak Admin > Clients > skillsier-fe > Credentials)"
echo "   - KEYCLOAK_MGMT_CLIENT_SECRET (get from Keycloak Admin > Clients > skillsier-bff > Credentials)"
echo ""
echo "🔐 To get Keycloak client secrets:"
echo "   1. Go to ${KEYCLOAK_PROD_URL}/admin/"
echo "   2. Select realm: ${KEYCLOAK_REALM}"
echo "   3. Navigate to Clients > skillsier-fe > Credentials tab"
echo "   4. Copy the 'Client secret' value"
echo "   5. Repeat for skillsier-bff client"
echo ""
echo -e "${GREEN}✅ Environment setup complete!${NC}"
echo ""
echo "Next steps:"
echo "  1. Update the client secrets in apps/web/.env.local"
echo "  2. Run: pnpm install"
echo "  3. Run: pnpm dev:web (for web) or pnpm dev:mobile (for mobile)"