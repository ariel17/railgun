# Railgun API
[![CircleCI](https://circleci.com/gh/ariel17/railgun/tree/master.svg?style=svg)](https://circleci.com/gh/ariel17/railgun/tree/master)
[![codecov](https://codecov.io/gh/ariel17/railgun/branch/master/graph/badge.svg?token=4IKZHQEA8G)](https://codecov.io/gh/ariel17/railgun)
[![Go Report Card](https://goreportcard.com/badge/github.com/ariel17/railgun)](https://goreportcard.com/report/github.com/ariel17/railgun)

## Local execution
```bash
ENVIRONMENT=production \
  DATABASE_USERNAME=root DATABASE_PASSWORD=root DATABASE_HOST=localhost DATABASE_PORT=3306 DATABASE_NAME=railgun \
  AUTH0_DOMAIN=ariel17.auth0.com AUTH0_AUDIENCE=GvJN0vB06vn5flX1ChZeEfteS3KzsCx60 \
  go run main.go
```
