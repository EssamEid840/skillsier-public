#!/bin/bash

# ============================================
# MOBILE DEVELOPMENT CLIENT SETUP
# ============================================

set -e

GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${BLUE}"
echo "╔════════════════════════════════════════╗"
echo "║   📱 Mobile Dev Client Setup 📱       ║"
echo "╔════════════════════════════════════════╗"
echo -e "${NC}"
echo ""

cd apps/mobile

# Check platform
if [[ "$OSTYPE" == "darwin"* ]]; then
  HAS_MACOS=true
else
  HAS_MACOS=false
fi

echo -e "${GREEN}Step 1: Installing Expo development client...${NC}"
npx expo install expo-dev-client

echo ""
echo -e "${GREEN}Step 2: Prebuilding native projects...${NC}"
npx expo prebuild --clean

if [ "$HAS_MACOS" = true ]; then
  echo ""
  echo -e "${GREEN}Step 3: Installing iOS CocoaPods dependencies...${NC}"
  cd ios
  pod install
  cd ..
  echo -e "${GREEN}✓ iOS dependencies installed${NC}"
else
  echo ""
  echo -e "${YELLOW}⏭️  Skipping iOS setup (not on macOS)${NC}"
fi

echo ""
echo -e "${GREEN}✅ Mobile development client setup complete!${NC}"
echo ""
echo "Next steps:"
echo ""
echo "For iOS (macOS only):"
echo -e "  ${BLUE}npx expo run:ios${NC}"
echo ""
echo "For Android:"
echo -e "  ${BLUE}npx expo run:android${NC}"
echo ""
echo "Or use Expo Go (simpler):"
echo -e "  ${BLUE}pnpm dev:mobile${NC}"
echo "  Then scan QR code with Expo Go app"