# FireBucket

API for managing file uploads and downloads with Firebase storage bucket.

## Requirements
- Go version 1.22 or higher
- Air for live reloading
- SQLC for SQL code generation

## Configurations

Create .env file in project root directory with following environment variables
```
# Server
SERVER_PORT=

# Database
DB_DRIVER=
DB_HOST=
DB_PORT=
DB_NAME=
DB_USER=
DB_PASSWORD=

# Firebase
FIREBASE_STORAGE_BUCKET=
FIREBASE_ACCOUNT_TYPE=
FIREBASE_PROJECT_ID=
FIREBASE_PRIVATE_KEY_ID=
FIREBASE_PRIVATE_KEY=
FIREBASE_CLIENT_EMAIL=
FIREBASE_CLIENT_ID=
FIREBASE_AUTH_URI=
FIREBASE_TOKEN_URI=
FIREBASE_AUTH_PROVIDER_CERT_URL=
FIREBASE_CLIENT_CERT_URL=
FIREBASE_UNIVERSE_DOMAIN=
```

## API Documentation
```
localhost:{port}/swagger/index.html
```
