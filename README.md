## Kong API gateway

---

### 1. Kong DB

Kong use database to save route, service, consumer data, this compose file spin up and postgres 13 database. Need to run migration first

```shell
docker compose up -d migration
```

After run migration start kong

```shell
docker compose up -d kong
```

Visit `http://localhost:8000` for http listen port
and `http://localhost:8001` for admin http listen port.

### 2. Kong UI

I use konga to manage kong, konga use db to store some resource like kong endpoint, user credential, to start konga, run ui_migration first then running ui

```
docker compose up ui_migration
docker compose up -d ui
```

Visit ui on `http://localhost:8002`
