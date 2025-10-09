#!/bin/bash
set -e

NAMESPACE="keycloak"
REALM_JSON_FILE="skillsier-realm.json"
ADMIN_USER="admin"
ADMIN_PASS="admin"
KEYCLOAK_URL="http://keycloak.keycloak.svc:8080"

echo "=== Importing Skillsier Realm to Keycloak ==="
echo ""

# Check if realm file exists
if [[ ! -f "$REALM_JSON_FILE" ]]; then
    echo "Error: Realm file '$REALM_JSON_FILE' not found!"
    exit 1
fi

echo "Step 1: Creating temporary pod with curl and jq..."
kubectl -n $NAMESPACE run realm-importer --rm -i --restart=Never \
  --image=curlimages/curl:8.5.0 -- sh <<EOF
set -e

echo "Step 2: Getting admin access token..."
TOKEN=\$(curl -s -X POST \
  "$KEYCLOAK_URL/realms/master/protocol/openid-connect/token" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "username=$ADMIN_USER" \
  -d "password=$ADMIN_PASS" \
  -d "grant_type=password" \
  -d "client_id=admin-cli" | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)

if [ -z "\$TOKEN" ]; then
    echo "Error: Failed to get admin token"
    exit 1
fi

echo "Step 3: Checking if realm already exists..."
REALM_EXISTS=\$(curl -s -o /dev/null -w "%{http_code}" \
  "$KEYCLOAK_URL/admin/realms/skillsier" \
  -H "Authorization: Bearer \$TOKEN")

if [ "\$REALM_EXISTS" = "200" ]; then
    echo "Warning: Realm 'skillsier' already exists. Deleting it first..."
    curl -s -X DELETE \
      "$KEYCLOAK_URL/admin/realms/skillsier" \
      -H "Authorization: Bearer \$TOKEN"
    echo "Existing realm deleted."
fi

echo "Step 4: Importing realm..."
cat > /tmp/realm.json <<'REALM_EOF'
$(cat $REALM_JSON_FILE)
REALM_EOF

RESPONSE=\$(curl -s -w "\n%{http_code}" -X POST \
  "$KEYCLOAK_URL/admin/realms" \
  -H "Authorization: Bearer \$TOKEN" \
  -H "Content-Type: application/json" \
  -d @/tmp/realm.json)

HTTP_CODE=\$(echo "\$RESPONSE" | tail -n1)
BODY=\$(echo "\$RESPONSE" | head -n-1)

if [ "\$HTTP_CODE" = "201" ]; then
    echo "✓ Realm 'skillsier' imported successfully!"
    exit 0
elif [ "\$HTTP_CODE" = "409" ]; then
    echo "✓ Realm 'skillsier' already exists"
    exit 0
else
    echo "✗ Error: Failed to import realm (HTTP \$HTTP_CODE)"
    echo "Response: \$BODY"
    exit 1
fi
EOF

echo ""
echo "=== Import Complete ==="
echo ""
echo "Next steps:"
echo "  1. Access Keycloak at https://keycloak.skillsier.com"
echo "  2. Login with admin/admin"
echo "  3. Switch to 'skillsier' realm (top-left dropdown)"
echo "  4. Test login with:"
echo "     - admin@skillsier.com / Skillsier@admin2025!"
echo "     - freelancer@skillsier.com / Freelancer@2025!"
echo "     - client@skillsier.com / Client@2025!"
echo "  5. Keycloak Account Console (End-user self-service):"   
echo "     - https://keycloak.skillsier.com/realms/skillsier/account"
echo ""