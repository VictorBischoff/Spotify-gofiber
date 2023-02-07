# Golang Fiber using Postgres.
<img src="./ghsrc/embed.png"></img>
## Usage:

1. Clone the repo

```BASH
git clone git@github.com:VictorBischoff/GoFiber-Postgres.git
```

2. Enter the repo directory

```BASH
cd GoFiber-Postgres
```

3. Run postgres

```BASH
docker compose up -d
```

4. Install dependencies

```BASH
go mod tidy
```

5. run the fiber server, and access at http://localhost:8080

```BASH
go run cmd/web/*.go
```
