<div align="center">
<a href="https://github.com/fahmirizalbudi/countreez" target="blank">
<img src="https://raw.githubusercontent.com/JjagoKoding/icon/9553c54df9c958f7ee00506742dfd446b57ae3e1/countreez.svg" width="300" alt="Logo" />
</a>

<br />
<br />

![](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![](https://img.shields.io/badge/Gin-00A393?style=for-the-badge&logo=gin&logoColor=white)
![](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![](https://img.shields.io/badge/redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)

</div>

<br/>

## Countreez

Countreez is a dedicated API solution crafted for comprehensive country data management and discovery. Built in Go and Gin Web Framework, this project was created to learn. This project uses PostgreSQL as the database and Redis for caching. Key features include:

## Features

- **Unlimited Requests:** Handle unlimited API requests without restrictions.
- **High Performance:** Optimized response times using Redis caching.
- **Countries Management:** View and search country data; manage global information easily.

## Tech Stack

- **Go**: A statically typed programming language designed for building scalable and high-performance server-side applications.
- **Gin**: A high-performance web framework for Go, designed for building RESTful APIs and web applications.
- **PostgreSQL**: A powerful, open-source relational database system for storing and managing structured data.
- **Redis**: An in-memory data structure store, used as a database, cache, and message broker to improve API performance.

## Getting Started

To get a local copy of this project up and running, follow these steps.

### Prerequisites

- **Go** (v1.24.x or higher).
- **PostgreSQL** (or another supported SQL database).
- **Redis** (latest stable version).

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/fahmirizalbudi/countreez.git
   cd countreez
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Start the development server:**

   ```bash
   go run main.go
   ```

## Usage

### Running the Website

- **Development mode:** `go run main.go`.
- **Production mode:** `go build main.go && ./main`.

> Use [http://localhost:8080](http://localhost:8080) to test the api in your Postman.

## License

All rights reserved. This project is for educational purposes only and cannot be used or distributed without permission.
