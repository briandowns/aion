Bring up supporting containers (nsq and mysql)

```bash
docker-compose up -d
```

First run you need to setup the db table structure for dev.

```bash
./aion_start -db-setup
```

After the db is setup you can run the script without args to start the bin

```bash
./aion_start
```

Run an ubuntu host with shared volume and links to required containers.

```bash
docker-compose run --service-ports aion bash
```

Ports are dynamically assigned use docker ps to find mapped port for 9898.

The container has curl, mysql-client and netcat-openbsd installed for testing purposes

```bash
# example
docker-compose run aion curl -d "hello world" http://nsqd:4151/put?topic=test
OK
```
