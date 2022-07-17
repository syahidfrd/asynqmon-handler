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
    -p 3000:3000 \
    syahidfrd/asynqmon-handler
```

Here's the available flags:

| Flag                     | Default      | Description                |
|--------------------------|--------------|----------------------------|
| `-auth-username`(string) | `admin`      | Basic auth username        |
| `-auth-password`(string) | `admin`      | Basic auth password        |
| `-redis-addr`(string)    | `:6379`      | Address of redis server    |

Next, go to [localhost:3000](http://localhost:3000) and see Asynqmon dashboard: