# coach-ai-assignment

## Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop)
- [Go](https://go.dev/doc/install) (for running the app locally)
- [Make](https://www.gnu.org/software/make/) (optional, for convenience)

---

## Running the Stack

### 1. Run init

This creates a local copy of the .env file that is expected from services.

```sh
make init
```

---

### 2. Start Infrastructure

This will start **Postgres**, **Kafka**, **Prometheus**, **Pushgateway**, and **Grafana**:

```sh
make up
```

or, if you don't have `make`:

```sh
docker-compose up -d
```

---

### 3. Run the Ingestor service

This publishes example events to Kafka:

```sh
make ingestor
```

or

```sh
go run main.go start ingestor
```

---

### 4. Run the Consumer

This consumes events from Kafka topic and pushes metrics to the Prometheus Pushgateway:

```sh
make consumer t=transactions
```

or

```sh
go run main.go start consumer --topic=transactions
```

You can run multiple consumers (optionally with different `INSTANCE_ID`s by setting the environment variable).

---

## Service URLs

- **Postgres:**
  [http://localhost:5439](http://localhost:5439)
- **Kafka Broker:**
  [http://localhost:9092](http://localhost:9092)
- **Prometheus:**
  [http://localhost:9090](http://localhost:9090)
- **Pushgateway:**
  [http://localhost:9091](http://localhost:9091)
- **Grafana:**
  [http://localhost:3000](http://localhost:3000)

  Login: `admin` / `admin`

---

## Observability

- The consumer pushes metrics to the Pushgateway.
- Prometheus scrapes the Pushgateway.
- You can visualize metrics in Grafana by adding Prometheus as a data source (`http://prometheus:9090` from within Docker, or `http://localhost:9090` from your host).

---

## Links

- **Kafka UI:**

  Inspect Kafka via Web UI
  [http://localhost:8080/](http://localhost:8080/)

- **Pushgateway UI:**

  View all pushed metrics and grouping keys at
  [http://localhost:9091/metrics](http://localhost:9091/metrics)

- **Prometheus UI:**

  Check configured targets:
  [http://localhost:9090/targets](http://localhost:9090/targets)

  Query and inspect metrics at
  [http://localhost:9090/graph](http://localhost:9090/graph)

---

## Stopping Everything

```sh
make down
```

or

```sh
docker-compose down
```

---

## Troubleshooting

- If you see connection errors, ensure all containers are running (`docker-compose ps`).
- For Kafka consumer group parallelism, increase the number of partitions for your topic.
- Check logs with `docker-compose logs <service>`.

---
