# [Asynqmon](https://github.com/hibiken/asynqmon) handler with basic auth

### Docker image

```bash
# Pull the latest image
docker pull syahidfrd/asynqmon-handler

# Or specify the image by tag
docker pull syahidfrd/asynqmon-handler[:tag]
```

### Run the binary

```
docker run --rm \
    --name asynqmon-handler \
    -e ASYNQMON_USER=admin \
    -e ASYNQMON_PASSWORD=secure123 \
    -p 3000:3000 \
    syahidfrd/asynqmon-handler
```

Here's the available env:

| Env                 | Default | Description       |
|---------------------|---------|-------------------|
| `ASYNQMON_USER`     | `admin` | Asynqmon user     |
| `ASYNQMON_PASSWORD` | `admin` | Asynqmon password |
| `REDIS_ADDR`        | `:6379` | Redis address     |

Next, go to [localhost:3000](http://localhost:3000) and see Asynqmon dashboard