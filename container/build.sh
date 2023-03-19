docker login ghcr.io
docker buildx create --name foGceTYl
docker buildx use foGceTYl
docker buildx build \
    --platform linux/arm64/v8,linux/amd64 \
    -t ghcr.io/ruriazz/gopen-api:latest \
    -f container/backend/Dockerfile \
    --push .