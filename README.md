# KFC Sales Forecasting Service

A daily sales forecasting service that predicts per-hour, per-product, per-store demand for the next day using historical sales averages. Built as a take-home assignment.  

Clarification: you may notice that it seems the same chart is being shown for different days/stores, but this is simply due to my fake seed data not being diverse (I chose not to spend time on this). I assure you the charts reflect the true sales data and the resulting forecasts.

---

## How to Run

**Requirements:** Docker and Docker Compose.

```bash
docker compose up --build
```

The app will be available at [http://localhost:5000](http://localhost:5000).

On first boot, the backend runs DB migrations, seeds historical sales data, and backfills forecasts for the seeded date range (Jan 1st-9th 2026).  
Subsequent boots reuse existing data (seeds and forecasts are idempotent).

To reset everything:
```bash
docker compose down -v && docker compose up --build
```

---

## Technologies

| Layer | Technology |
|---|---|
| Backend | Go (Gin, GORM) |
| Frontend | React + TypeScript + Vite |
| Database | PostgreSQL |
| Styling | Tailwind CSS + shadcn/ui |
| Reverse proxy | nginx |
| Containerization | Docker Compose |
`
---

## Architecture

```
nginx
  ├── /api/* → backend (Go)
  │     ├── API server (Gin)
  │     ├── Forecast worker (goroutine, runs daily at configured hour)
  │     └── PostgreSQL via GORM
  └── /* → static React bundle (built into nginx image)
```

The frontend is **not** a runtime service — it's compiled during the Docker build and served as static files by nginx.

The backend runs the API server and forecast scheduler as goroutines in a single binary (see __Design Considerations__ for details).

### API

- `GET /api/stores` — list all stores
- `GET /api/forecasts?store_id={id}&date={YYYY-MM-DD}` — hourly forecasts for a store on a given date

### Configuration

Controlled via environment variables in a `.env` file which is injected via Docker Compose:

| Variable | Default | Description |
|---|---|---|
| `LOOKBACK_DAYS` | `7` | Days of sales history to average over |
| `GENERATION_INTERVAL_DAYS` | `1` | How often to generate forecasts |
| `GENERATION_HOUR` | `0` | Hour of day (0–23) to run the forecast job |

---

## Design Considerations

### High-level Architecture
To demonstrate awareness of industry best practices I chose to use docker compose with a database service, a backend service, and an nginx service (as opposed to a simple Go app that serves static files). I wanted to create a system that's close to production systems with good separation, making my system modular and thus easier to test and modify.

### Forecast algorithm
For each `(store, product, hour-of-day)` tuple, the predicted quantity is the ceiling of the average quantity sold across the last `LOOKBACK_DAYS` days. Ceiling rounding is intentional — for food prep, underestimating is worse than overestimating.

In the current implementation, the average is computed only over sale rows that exist in the database (the SQL `AVG` does not take into account hours for which there's no sale data).  
This means that technically we assume that there exist database records for "zero-sale hours" (hours with 0 as the quantity sold). 
The fake seed data I created does not include such rows, I simply didn't choose to devote time for this, but in a real system I would either include such rows (could be helpful for other analytics or varying operating hours), or, if we wanted to save on storage / costs and potentially speed up the computation, we would add either SQL or Go logic to account for these hours -- we would have to account for potential differences in operating hours (e.g. due to weekends or holidays).

**Known limitation:** The simple mean conflates day-of-week patterns (e.g. Saturday vs Tuesday traffic). A weighted or seasonal model would improve accuracy; the spec asks for a plain average so that's what's implemented.

### Sales modeled as a fact table
`sales(store_id, product_id, sold_at, quantity)` — one row per product sold. In a real production system, sales would likely be derived from an OLTP `orders + order_items` model via ETL. This shape is used here because the only consumer is the analytics/forecasting workload and I preferred simplicity where reasonable to make it easier for reviewers and save time.

### Single binary for Forecast Generator and API Server 
Forecast generator and API server run as goroutines in one Go process. 
In a real system I consider using a CronJob (K8s), or, if the forecast logic was simple (like with simple averages), I would consider using timescaledb's "continuous aggregates" which maintains materialized views. The assignment spec asked for a generation "service" so I assumed reviewers want to see Go code.

### No service layer between DAL and API
The handlers call repository functions directly. In a production backend, a service layer would likely be needed for some business logic.
At this scope, with a few simple endpoints and no complex business rules, the extra layer would be pure boilerplate so I chose to keep it simple with less code.

## React App State Management
I chose to use react query which has response caching OOTB and is a common pattern.
It also shares the same cached responses across all consuming componenets, acting as a form of global state. In a more complex app I might use React Context or Redux. However for a quick assignment, React Query was a good choice both in terms of less boilerplate code and enough feature-richness OOTB.

---

## Limitations & Shortcuts

Given the ~4–5 hour scope, several things were intentionally left out or simplified:

- **Test coverage is minimal.** The forecast generation logic has unit tests. API handlers and the DAL layer have no automated tests — in a real project I'd add integration tests (e.g. using testcontainers) and expand unit test coverage significantly.
- **No authentication or authorization.** All endpoints are public.
- **No sales write API.** Sales data is seeded on startup. A real system would have ingestion endpoints or an ETL pipeline.
- **Single-day forecast horizon.** The scheduler generates forecasts for the next day only; multi-day lookahead is out of scope per the spec.
- **No error boundary / retry logic in the frontend.** Basic error states are handled but not extensively.
