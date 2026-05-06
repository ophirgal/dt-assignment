# CLAUDE.md

Context for AI agents working on this project. Read this first.

## Project

KFC daily sales forecasting service. Take-home assignment, ~4–5 hour scope.

A backend service generates per-hour, per-product, per-store sales forecasts for the next day on a daily schedule. A frontend lets users browse stores and view forecasts by date. The forecasting algorithm is a simple historical average — for each `(store, product, hour-of-day)`, average `quantity` across the last N days of historical sales.

## Stack

- **Backend**: Go — single binary, stdlib-first, minimal deps
- **Frontend**: React + TypeScript + Vite (built to static assets)
- **Database**: Postgres (vanilla — **no extensions**)
- **Deployment**: `docker-compose` with three services:
  - `postgres` — DB
  - `backend` — Go binary (API + scheduler)
  - `nginx` — reverse proxy for `/api/*` to `backend`, serves built React bundle for everything else

The frontend is **not** a runtime service. It's built (in a multi-stage Dockerfile) and the static `dist/` is copied into the nginx image.

## Repo Layout

```
.
├── backend/
│   ├── cmd/server/         # main.go — boots HTTP server + scheduler as goroutines
│   ├── internal/
│   │   ├── api/            # HTTP handlers, routing
│   │   ├── forecast/       # Forecast generation logic — keep pure, heavily tested
│   │   ├── storage/        # DB access, queries, migrations, seeder
│   │   └── config/         # Config file loading
│   └── Dockerfile
├── frontend/               # React + TS + Vite
│   └── Dockerfile          # multi-stage: build → nginx
├── docker-compose.yml
├── config.yaml             # runtime config (see below)
└── CLAUDE.md
```

## Key Design Decisions (with rationale)

These choices are intentional. Don't undo them without good cause.

### Single Go binary, not two services
Forecast generator and API server run as goroutines in one process but live in cleanly separated internal packages (`forecast/`, `api/`). Splitting into two services adds Docker/compose/communication boilerplate disproportionate to scope. The package boundary is real, so extracting `forecast/` into its own binary later is mechanical.

### Vanilla Postgres, not TimescaleDB
Considered Continuous Aggregates for forecast computation. Rejected:
- The averaging logic is the interesting code under review. Burying it in a `MATERIALIZED VIEW` definition makes it harder to test and harder for reviewers to evaluate.
- Edge cases (zero-sale hours, missing days, rounding, day boundaries) are far easier to nail in pure Go with table-driven tests than inside view materialization.
- Adds an extension dependency for no real win at this scale.

Forecast logic lives in pure Go and is unit-testable in isolation.

### In-process scheduler, not external cron
A goroutine with a ticker (or `gocron` if we want named jobs) reads the interval from config and invokes the generator. Simpler than a separate cron container, easier to test, and the assignment requires the schedule interval to be configurable — easier to wire in code than as a SQL refresh policy or container env var.

### Sales as a fact table, not `Sale` + `SaleItem`
The only consumer of sales data is the forecast generator (analytics workload). We have no order-level use cases — no POS, no receipts, no refunds, no basket analysis. So we go straight to the analytical shape: one row per product sold.

In a real production system this would be derived from an OLTP `orders + order_items` model via ETL. **Note this in the README** so reviewers see we're aware of the trade-off.

### Chains table kept, address fields dropped
`chains(id, name)` is essentially free and signals multi-tenancy thinking (the platform serves multiple chains, not just KFC). Stores belong to chains, products belong to chains. Address fields on stores are out of scope — the UI only needs `(id, name)`.

### No sales CRUD endpoints
Sales are seed data, populated on backend startup if the table is empty. The assignment doesn't ask for sales endpoints, and a CRUD service would be modeling a production system rather than this assignment.

## Data Model

```sql
chains    (id, name)
stores    (id, name, chain_id → chains)
products  (id, name, price, chain_id → chains)
sales     (id, store_id → stores, product_id → products, sold_at timestamptz, quantity int)
forecasts (id, store_id → stores, product_id → products,
           forecast_date date, hour int CHECK 0..23,
           predicted_quantity numeric, generated_at timestamptz,
           UNIQUE(store_id, product_id, forecast_date, hour))
```

Indexes:
- `sales(store_id, sold_at)` — composite, supports historical window scan
- `forecasts(store_id, forecast_date)` — supports primary read path

## Configuration

All operational parameters live in `config.yaml` loaded at startup. The UI does **not** expose configuration — config is operator-controlled, per spec.

```yaml
forecast:
  interval: 24h          # how often the generator runs
  history_days: 30       # days of history to average over
  run_at: "02:00"        # wall-clock time to run daily
db:
  host: postgres
  port: 5432
  # ...
server:
  port: 8080
seed:
  enabled: true          # generate fake history on first boot
  days: 60               # how many days of history to fabricate
```

## API Surface

Read-only, scoped to forecast retrieval:

- `GET /api/stores` — list stores
- `GET /api/forecasts?store_id={id}&date={YYYY-MM-DD}` — forecast rows for store + date

Optional if time permits: `GET /api/products?store_id={id}` (depends on whether the UI needs product names client-side or gets them embedded in the forecast response).

## Watch out for

The reviewer hinted there's a subtle trap that surfaces by running with different inputs. Most likely candidates — write table-driven tests for each:

- **Zero-sale hours**: divide by total days, or only days with sales? A store closed Sundays shouldn't drag the Monday forecast.
- **Missing history**: what if some products have fewer than `history_days` of data?
- **Hour buckets**: 0–23 vs 1–24, midnight handling, day-boundary off-by-one.
- **Timezones**: store timestamps in UTC vs forecast displayed in store-local time. Day boundaries shift.
- **Rounding**: forecasts are "items to prepare" — fractional outputs need a defined rounding policy (probably ceiling — underprep is worse than overprep — but be explicit and document it).
- **Day-of-week effects**: simple mean across last N days conflates Saturdays and Tuesdays. Spec says "avg," so follow the spec literally, but note this in code comments and README.
- **Generator cold start**: no historical data yet → empty forecasts, error, or fallback? Decide and document.

Forecast logic must have unit tests covering each of these. This is where the assignment is graded.

## Out of scope

Explicit list — do not build:

- Sales CRUD endpoints
- Authentication / authorization
- API gateway / rate limiting
- Address fields on stores
- Order / OrderItem modeling
- Multi-day forecast horizon (spec says "next day" only)
- ML beyond simple averaging
- Real-time updates / websockets
- User-facing config

## Dev workflow

```bash
docker compose up --build           # bring everything up
docker compose down -v              # tear down + drop volumes
docker compose logs -f backend      # tail backend logs
```

Local backend dev (without Docker, requires running Postgres):
```bash
cd backend
go run ./cmd/server
go test ./...
```

Local frontend dev:
```bash
cd frontend
npm install
npm run dev                         # Vite dev server, proxy /api to localhost:8080
```

## Conventions

- **Go**: stdlib first; add deps only with reason. Logging via `log/slog`. Config loaded once into a typed struct, passed by value to constructors. No global state.
- **TypeScript**: strict mode on, no `any`. Functional components + hooks. Small components.
- **DB migrations**: SQL files in `backend/internal/storage/migrations/`, applied on startup (golang-migrate or similar).
- **Tests**: forecast logic gets table-driven unit tests with explicit edge-case coverage. API handlers get integration tests against a real Postgres (testcontainers or compose-managed).
- **Commits**: small, focused, clear messages.

## README checklist (separate from this file)

The user-facing `README.md` should cover:
- One-line description, quickstart (`docker compose up`)
- Architecture diagram (or ascii)
- Configuration reference
- Explicit notes on:
  - "Sales modeled as fact table; in production would be derived from OLTP orders via ETL"
  - "Forecast averaging: zero-sale hours handled as [decision]; rounding policy is [decision]"
  - Known limitations / trade-offs
