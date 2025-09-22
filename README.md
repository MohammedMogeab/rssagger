# rssagger

A minimal RSS aggregator written in Go. It exposes a small REST API to:

- Create users and authenticate via API key (Bearer token)
- Create and list RSS feeds
- Follow/unfollow feeds per user
- Fetch posts from followed feeds
- Continuously scrape feeds in the background

It uses Chi for routing, Postgres for storage, and sqlc for type‑safe DB access.

## Requirements

- Go 1.24+
- PostgreSQL 13+

## Setup

1) Create a Postgres database and run the migrations in `sql/schema` in order:

- `001_users.sql`
- `002_users_apikey.sql`
- `003_feeds.sql`
- `004_feedsfollow.sql`
- `005_feeds_fetch.sql`
- `006_posts.sql`

You can run them with your preferred tool (psql, a GUI, or goose). Example with `psql`:

```bash
psql "$DB_URL" -f sql/schema/001_users.sql
psql "$DB_URL" -f sql/schema/002_users_apikey.sql
psql "$DB_URL" -f sql/schema/003_feeds.sql
psql "$DB_URL" -f sql/schema/004_feedsfollow.sql
psql "$DB_URL" -f sql/schema/005_feeds_fetch.sql
psql "$DB_URL" -f sql/schema/006_posts.sql
```

2) Create a `.env` file in the project root:

```env
PORT=8080
DB_URL=postgres://user:pass@localhost:5432/rssagger?sslmode=disable
```

3) Install deps and run:

```bash
go run .
# or
go build -o rssagger
./rssagger
```

Server starts on `:$PORT` and mounts API under `/v1`.

## Authentication

- Create a user to receive an API key.
- Send the API key on protected routes using the `Authorization: Bearer <api_key>` header.

## API

Base URL: `http://localhost:<PORT>/v1`

- Health: `GET /healthz`
  - Returns `{}` (currently implemented as an empty JSON response).

- Test error: `GET /error`
  - Returns `{ "error": "something is wrong" }` with status 400.

- Create user: `POST /users`
  - Body: `{ "name": "Alice" }`
  - Response: user object including `api_key`.

- Get current user: `GET /users` (auth required)
  - Header: `Authorization: Bearer <api_key>`
  - Response: user object.

- Create feed: `POST /feeds` (auth required)
  - Body: `{ "name": "Go Blog", "url": "https://blog.golang.org/index.atom" }`
  - Response: feed object.

- List feeds: `GET /feeds`
  - Response: array of feeds.

- Follow a feed: `POST /feeds/follow` (auth required)
  - Body: `{ "feed_id": "<feed UUID>" }`
  - Response: follow relationship object.

- List my follows: `GET /feeds/follow` (auth required)
  - Response: array of follow relationships.

- Unfollow: `DELETE /feeds/follow/{feedfollowId}` (auth required)
  - Path param: `feedfollowId` is the follow relationship UUID.

- Get my posts: `GET /posts` (auth required)
  - Returns posts for feeds you follow, ordered by `published_at` desc.
  - Optional query: `?limit=10` (defaults to 4, max 100).

### Example cURL

```bash
# Create user
curl -sX POST http://localhost:8080/v1/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Alice"}'

# Set API key
API_KEY=... # from the response

# Create a feed
curl -sX POST http://localhost:8080/v1/feeds \
  -H "Authorization: Bearer $API_KEY" \
  -H 'Content-Type: application/json' \
  -d '{"name":"Go Blog","url":"https://blog.golang.org/index.atom"}'

# List feeds
curl -s http://localhost:8080/v1/feeds

# Follow a feed
curl -sX POST http://localhost:8080/v1/feeds/follow \
  -H "Authorization: Bearer $API_KEY" \
  -H 'Content-Type: application/json' \
  -d '{"feed_id":"<FEED_UUID>"}'

# Get my posts
curl -s http://localhost:8080/v1/posts \
  -H "Authorization: Bearer $API_KEY"
```

## Background Scraper

The server starts a background worker that:

- Periodically selects the next feeds to fetch (by `last_fetch_at` nulls‑first)
- Fetches each feed concurrently
- Parses items from the RSS/Atom feed
- Inserts posts (skipping duplicates by unique URL)
- Marks the feed as fetched (`last_fetch_at = now()`)

Configurable via env:

- `SCRAPER_CONCURRENCY` (default 3)
- `SCRAPER_INTERVAL_SECONDS` (default 60)

## Project Structure

- `main.go` – server setup, routes, middleware
- `middleware.go` – API key auth wrapper
- `UserHandler.go`, `handler_feed.go` – HTTP handlers
- `json.go` – JSON helpers
- `models.go` – response models and mappers
- `rss.go` – RSS fetch/parse helpers
- `scrapper.go` – background scraping logic
- `internal/database` – sqlc‑generated DB layer
- `sql/schema` – SQL migrations
- `sql/queries` – SQL queries (sqlc input)

## Development Notes

- CORS is enabled for local development; see `main.go` for origins.
- DB access code is generated with `sqlc` using `sqlc.yaml`.
- Errors are returned as JSON `{ "error": "..." }` with appropriate HTTP status codes.

## License

No license specified. Add one if needed before distributing.
