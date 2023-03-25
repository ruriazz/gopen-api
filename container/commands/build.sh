docker login ghcr.io
docker buildx create --name k5r1_csw
docker buildx use k5r1_csw
docker buildx build \
    --platform linux/arm64/v8,linux/amd64 \
    -t ghcr.io/ruriazz/gopen-api-commands:latest \
    -f container/commands/Dockerfile \
    --push .