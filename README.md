# Mock Telco API Pipeline

This project simulates a data ingestion pipeline for telco-like data using a custom mock API built with Go. The API is exposed publicly using Ngrok, and the data is intended to be ingested into PostgreSQL for further processing. The purpose of this project is to practice data engineering concepts end-to-end, from ingestion to storage.

---

## Project Structure

```text
mock-telco-api-pipeline/
├── ingestion/
│   └── go_api/
│       ├── main.go               # Go server for mock telco API
│       ├── handlers.go           # HTTP handlers for each endpoint
│       ├── models.go             # Data models (User, CallLog, DataUsage)
│       ├── mockdata.go           # Functions to generate mock data
│       ├── utils.go              # Helper functions like writeJSON
│       ├── .air.toml             # Air config (not committed)
│       ├── data/                 # Static mock response files (JSON, CSV, etc.)
│       └── tmp/                  # Temp dir used by Air (git-ignored)
│
├── scripts/
│   └── ingest_to_postgres.py     # Python script to fetch from API and load into PostgreSQL
│
├── requirements.txt              # Python dependencies (for scripts/)
├── .gitignore                    # Ignored files and dirs
└── README.md                     # Project overview and instructions
```

---

## Tech Stack

| Tool         | Role                                             |
| ------------ | ------------------------------------------------ |
| `Go`         | Serves mock telco data via HTTP API              |
| `Air`        | Enables live reload during Go development        |
| `Ngrok`      | Tunnels the local API publicly via HTTPS         |
| `PostgreSQL` | Relational database for storing ingested data    |
| `DBeaver`    | GUI to connect and query the PostgreSQL database |

---

## Getting Started

### 1. Run the Go API with Air (live reload)

Ensure you are inside the `go_api` directory:

```bash
cd ingestion/go_api
air
```

This will build and start your API, automatically reloading on file changes.

### 2. Expose API publicly using Ngrok

In a new terminal, run:

```bash
ngrok http 8080
```

You’ll receive a public HTTPS URL like:

```
https://1234-abcde.ngrok-free.app -> http://localhost:8080
```

Use this URL for external access (for example, testing on your phone or for ingestion scripts).

---

## Current Progress

- Built mock API with `Go`
- Configured `Air` for hot-reloading
- Installed and tested `Ngrok` tunnel
- Confirmed external device access via Ngrok
- Created project folder and Git structure
- Next: Implement Python ingestion to PostgreSQL

---

## Notes

- The API serves mock telco-like data and no real production data involved.
- `Ngrok` authtoken is stored globally at:
  ```
  C:\Users\<USERNAME>\AppData\Local\ngrok\ngrok.yml
  ```
- `.air.toml`, `.env`, and compiled binaries are intentionally ignored via `.gitignore`.

---

## Useful Commands

| Task                         | Command                         |
| ---------------------------- | ------------------------------- |
| Start API (with Air)         | `air`                           |
| Tunnel API (Ngrok)           | `ngrok http 8080`               |
| Check Ngrok Web UI           | `http://localhost:4040`         |
| Generate Python requirements | `pip freeze > requirements.txt` |

---

## To Do (Upcoming)

- [ ] Develop Python script to pull API data and insert into PostgreSQL
- [ ] Create PostgreSQL schema and test DB connection
- [ ] (Optional) Integrate with orchestration tools (Airflow, NiFi, etc.)

---

## License

This is a learning project and currently not licensed for distribution. Please treat it as personal or educational use only.
