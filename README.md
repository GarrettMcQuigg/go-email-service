﻿## go-email-service

Email service written in Golang that uses SMTP to send emails.

### Running Locally

1. Change directories so that you're inside the server directory.

```bash
cd server
```

2. Copy the contents of the example configuration file into a new file which is leveraged by the server.

```bash
cp infra/local/local.yaml configuration.yaml
```

3. Ensure you have an app password configured for your gmail account.

4. Download dependencies required to run the server.

```bash
go mod download
```

5. Run the server

```bash
go run cmd/server/main.go
```
