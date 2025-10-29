#!/usr/bin/env bash
set -euo pipefail
ORG="${1:-YOUR_ORG}"
REPO="${2:-portsage}"
git init
git add .
git commit -m "chore: bootstrap PortSage Windows skeleton"
git branch -M main
git remote add origin git@github.com:${ORG}/${REPO}.git
git push -u origin main
echo "Done. Pushed to git@github.com:${ORG}/${REPO}.git"
