#!/bin/bash

# ============================================
# SKILLSIER INTERACTIVE DEV LAUNCHER
# ============================================

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

clear

echo -e "${BLUE}"
echo "╔════════════════════════════════════════╗"
echo "║     🚀 Skillsier Dev Launcher 🚀      ║"
echo "╔════════════════════════════════════════╗"
echo -e "${NC}"
echo ""
echo "Choose what to run:"
echo ""
echo -e "${GREEN}1.${NC} Web App only (Next.js)"
echo -e "${GREEN}2.${NC} Mobile App only (Expo)"
echo -e "${GREEN}3.${NC} Both Web + Mobile (recommended)"
echo -e "${GREEN}4.${NC} Type check all"
echo -e "${GREEN}5.${NC} Lint all"
echo -e "${GREEN}6.${NC} Build all"
echo -e "${RED}7.${NC} Clean & reinstall"
echo ""
read -p "Enter your choice [1-7]: " choice

case $choice in
  1)
    echo ""
    echo -e "${BLUE}🌐 Starting Web development server...${NC}"
    echo -e "${YELLOW}Available at: http://localhost:3000${NC}"
    echo ""
    pnpm dev:web
    ;;
  2)
    echo ""
    echo -e "${BLUE}📱 Starting Mobile development server...${NC}"
    echo -e "${YELLOW}Scan QR code with Expo Go app${NC}"
    echo ""
    pnpm dev:mobile
    ;;
  3)
    echo ""
    echo -e "${BLUE}🚀 Starting both Web and Mobile servers...${NC}"
    echo -e "${YELLOW}Web: http://localhost:3000${NC}"
    echo -e "${YELLOW}Mobile: Scan QR code${NC}"
    echo ""
    pnpm dev
    ;;
  4)
    echo ""
    echo -e "${BLUE}🔍 Running TypeScript type check...${NC}"
    echo ""
    pnpm type-check
    ;;
  5)
    echo ""
    echo -e "${BLUE}✨ Running linter...${NC}"
    echo ""
    pnpm lint
    ;;
  6)
    echo ""
    echo -e "${BLUE}🏗️  Building all packages...${NC}"
    echo ""
    pnpm build
    ;;
  7)
    echo ""
    echo -e "${RED}🧹 Cleaning and reinstalling...${NC}"
    echo -e "${YELLOW}This will remove all node_modules and rebuild${NC}"
    read -p "Are you sure? [y/N]: " confirm
    if [[ $confirm == [yY] ]]; then
      pnpm clean
      rm -rf node_modules
      pnpm install
      echo -e "${GREEN}✅ Clean install complete!${NC}"
    else
      echo "Cancelled"
    fi
    ;;
  *)
    echo ""
    echo -e "${RED}Invalid choice. Please run again and choose 1-7.${NC}"
    exit 1
    ;;
esac