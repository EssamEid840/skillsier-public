### 1 How & Why? Create Kubernetes Secret with Client Credentials, I need that to be Dynamic
# Make Sure from client-id=myapp
* We need to add Client Id per each Microservice/Logical service
* Folow the Steps in the file 04-keycloak-client-secret-creation.md
* Go to Credentials tab in Keycloak, copy the client secret, then:

```
# Replace YOUR_CLIENT_SECRET with actual value from Keycloak UI
kubectl create secret generic keycloak-client-myapp -n keycloak \
  --from-literal=client-id=myapp \
  --from-literal=client-secret=YOUR_CLIENT_SECRET_HERE

# Or if you want to script it, manually paste the secret value:
kubectl create secret generic keycloak-client-myapp -n keycloak \
  --from-literal=client-id=myapp \
  --from-literal=client-secret='<paste-secret-here>' \
  --dry-run=client -o yaml | kubectl apply -f -

```

### 3 Keycloak Nodeport is missing and why, and make sure from the remaining variables

```
# Export as environment variables
<!-- export KEYCLOAK_URL="http://173.212.218.251:30080" -->
export KEYCLOAK_URL="https://keycloak.skillsier.com"
export KEYCLOAK_REALM="skillsier"
export KEYCLOAK_CLIENT_ID=$(get_client_id)
export KEYCLOAK_CLIENT_SECRET=$(get_client_secret)
export KEYCLOAK_REDIRECT_URL="http://localhost:8080/callback"

```


### 3 Local Dapr is missing

```
# Create Dapr components directory if it doesn't exist
echo "Setting up Dapr components..."

mkdir -p ~/.dapr/components

# This would be done if you need local Dapr components
# For now, we're using direct Kafka connection

echo "✓ Dapr components directory created"

```