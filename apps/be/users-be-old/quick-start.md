
```text  


# 1. Create project structure and copy all files
mkdir -p users-be && cd users-be

# 2. Initialize and download dependencies
go mod init users-be
make deps

# 3. deploy postgresql
make db.apply

# 4. Setup local environment
make setup-local
source scripts/get-secrets.sh

# 5. Run the service
make run

# 6. Test it
curl http://localhost:8080/health


```
