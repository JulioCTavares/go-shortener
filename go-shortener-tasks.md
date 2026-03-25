# go-shortener — plano de tasks

URL Shortener em Go · plano de um dia · até 17h

---

## Fase 1 · Setup e estrutura do projeto

- [ ] Inicializar módulo Go e estrutura de pastas (`cmd/`, `internal/`, `pkg/`)
- [ ] Levantar Postgres e Redis via docker-compose
- [ ] Configurar variáveis de ambiente com godotenv ou viper
- [ ] Criar conexão com Postgres e rodar primeira migration (tabela `urls`)
- [ ] Criar conexão com Redis e validar ping

---

## Fase 2 · Core da API HTTP

- [ ] Definir rotas: `POST /shorten`, `GET /:code`, `GET /stats/:code`
- [ ] Implementar handler de criação de URL curta (geração de hash/slug)
- [ ] Implementar handler de redirect com lookup no Postgres
- [ ] Middleware de logging estruturado (slog ou zerolog)
- [ ] Middleware de recovery (panic handler)
- [ ] Middleware de CORS e Content-Type

---

## Fase 3 · Cache com Redis + Rate Limit

- [ ] Implementar cache-aside: checar Redis antes de ir ao Postgres no redirect
- [ ] Invalidar cache ao deletar ou atualizar URL
- [ ] Rate limiter por IP com Redis (sliding window ou token bucket)
- [ ] Retornar `429` com header `Retry-After` quando limite atingido

---

## Fase 4 · Métricas assíncronas com goroutines

- [ ] Criar worker goroutine que consome um channel de eventos de acesso
- [ ] Incrementar contador de cliques no Postgres de forma não bloqueante
- [ ] Usar `sync.WaitGroup` para graceful shutdown do worker
- [ ] Endpoint `GET /stats/:code` retornando total de cliques e última visita

---

## Fase 5 · Testes e benchmark

- [ ] Testes unitários do gerador de slug/hash
- [ ] Testes de integração para `POST /shorten` e `GET /:code` com `httptest`
- [ ] Mock do Redis e Postgres nos testes com interface + fake impl
- [ ] Benchmark do handler de redirect (`go test -bench`)
- [ ] README com arquitetura, como rodar e exemplos curl

---

## Rotas

| Rota | Método | Cache | Worker |
|---|---|---|---|
| `/shorten` | POST | invalidar se existir | — |
| `/:code` | GET | leitura (Redis-first) | dispara evento |
| `/stats/:code` | GET | sem cache | — |
| `/urls/:code` | DELETE | invalidar | — |
| `/health` | GET | sem cache | — |

---

## Estrutura de pastas

```
go-shortener/
├── cmd/api/          → main.go, bootstrap
├── internal/
│   ├── handler/      → HTTP handlers
│   ├── middleware/   → logging, rate limit, recovery
│   ├── shortener/    → regra de negócio (geração de slug)
│   ├── cache/        → interface + impl Redis
│   └── store/        → interface + impl Postgres
├── pkg/config/       → env/config
├── migrations/       → SQL files
└── docker-compose.yml
```

---

## Schema do banco

### `urls`
| coluna | tipo | observação |
|---|---|---|
| `id` | `uuid` | PK, `gen_random_uuid()` |
| `code` | `varchar(10)` | UK, indexado |
| `original_url` | `text` | — |
| `title` | `varchar(255)` | nullable |
| `is_active` | `boolean` | default `true` |
| `click_count` | `bigint` | incrementado pelo worker |
| `last_accessed_at` | `timestamptz` | nullable |
| `expires_at` | `timestamptz` | nullable, índice parcial |
| `created_at` | `timestamptz` | default `NOW()` |
| `updated_at` | `timestamptz` | default `NOW()` |

### `access_logs`
| coluna | tipo | observação |
|---|---|---|
| `id` | `uuid` | PK |
| `url_id` | `uuid` | FK → `urls.id` ON DELETE CASCADE |
| `ip_address` | `inet` | IPv4 e IPv6 |
| `user_agent` | `text` | nullable |
| `country_code` | `varchar(10)` | nullable |
| `referer` | `text` | nullable |
| `accessed_at` | `timestamptz` | default `NOW()`, indexado |
