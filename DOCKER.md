Bring up supporting containers (nsq and mysql)

```bash
docker-compose up -d
```

Run an ubuntu host with shared volume and links to required containers.

```bash
docker-compose run --service-ports aion bash
```

Ports are dynamically assigned use docker ps to find mapped port for 9898.
