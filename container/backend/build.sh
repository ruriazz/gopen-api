docker login ghcr.io
docker buildx create --name yHyfmBiH
docker buildx use yHyfmBiH
docker buildx build \
    --platform linux/arm64/v8,linux/amd64 \
    -t ghcr.io/ruriazz/gopen-api:latest \
    -f container/backend/Dockerfile \
    --push .