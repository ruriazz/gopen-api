docker login ghcr.io
docker buildx create --name foGceTYldev
docker buildx use foGceTYldev
docker buildx build \
    --platform linux/amd64 \
    -t ghcr.io/ruriazz/gopen-api:dev \
    -f container/backend/Dockerfile \
    --push .
docker buildx build \
    --platform linux/amd64 \
    -t ghcr.io/ruriazz/gopen-api-commands:dev \
    -f container/commands/Dockerfile \
    --push .