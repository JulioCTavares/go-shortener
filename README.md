# go-shortener

URL shortener em Go, similar ao bit.ly.

## Stack

- **Go** — `net/http` nativo
- **PostgreSQL** — armazenamento das URLs
- **Redis** — cache de redirects
- **golang-migrate** — migrations SQL
- **Docker Compose** — infraestrutura local

## Estrutura

```
go-shortener/
├── cmd/api/          → main.go (bootstrap)
├── internal/
│   ├── handler/      → rotas HTTP
│   └── shortener/    → regras de negócio
├── pkg/config/       → conexões e variáveis de ambiente
├── migrations/       → arquivos SQL
└── docker-compose.yml
```

## Como rodar

**1. Suba os serviços:**
```bash
docker-compose up -d
```

**2. Configure o `.env` na raiz:**
```env
PG_DATABASE_USER=postgres
PG_DATABASE_PASSWORD=docker
PG_DATABASE_DB=go-shortener
PG_DATABASE_PORT=5432
PG_DATABASE_HOST=localhost
REDIS_DATABASE_HOST=localhost
REDIS_DATABASE_PORT=6379
```

**3. Rode a API** (migrations são aplicadas automaticamente):
```bash
go run ./cmd/api/main.go
```

## Rotas

| Método | Rota       | Descrição                       |
|--------|------------|---------------------------------|
| GET    | /health    | Health check                    |
| POST   | /shorten   | Cria uma URL curta              |
| GET    | /{code}    | Redireciona para a URL original |

## Exemplos

**Criar URL curta:**
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.google.com/search?q=golang"}'

# {"short_url":"aB3xK9z"}
```

**Redirecionar:**
```bash
curl -L http://localhost:8080/aB3xK9z
```

## Schema

| coluna            | tipo          | descrição                  |
|-------------------|---------------|----------------------------|
| `id`              | `uuid`        | PK gerado automaticamente  |
| `code`            | `varchar(10)` | código único da URL curta  |
| `original_url`    | `text`        | URL original               |
| `is_active`       | `boolean`     | se a URL está ativa        |
| `click_count`     | `bigint`      | total de acessos           |
| `last_accessed_at`| `timestamptz` | último acesso              |
| `expires_at`      | `timestamptz` | expiração (nullable)       |
| `created_at`      | `timestamptz` | data de criação            |
