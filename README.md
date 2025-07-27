<p align="center">
  <img src="https://github.com/fahmirizalbudi/countreez/blob/main/countreez.png" alt="Countreez Logo" width="300"/>
</p>

<br/>

# 🌎 Countreez

Countreez – a simple REST API to search and count countries.

---

## 🛠️ Tech Stack

- Go (Golang)
- PostgreSQL
- Redis

---

## ⚙️ Instalation

1. Clone repository:
   ```bash
   git clone https://github.com/fahmirizalbudi/countreez.git
   cd book-api
   ```

3. Install dependency:
   ```bash
   go mod tidy
   ```

4. Run API:
   ```bash
   go run main.go
   ```

---

## 🔗 Endpoint

| Method | Endpoint                    | Description             |
|--------|-----------------------------|-------------------------|
| GET    | `/api/countries`                | Retrieve all country       |

Example JSON Body:
```json
{
    "id": 1,
    "name": "Afghanistan",
    "iso2": "AF.",
    "iso3": "AFG",
    "capita;": "Kabul",
    "region": "Asia",
    "language": "fa-AF,ps,uz-AF,tk",
}
```

---

## 🎯 Purpose

This project was built to practice building RESTful APIs using Go, PostgreSQL, and Redis.
