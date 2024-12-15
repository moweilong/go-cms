FROM debian:stable-slim

ARG APP_RELATIVE_PATH
ARG GRPC_PORT
ARG REST_PORT

# RUN apt update && \
#     apt install -y --no-install-recommends ca-certificates netbase && \
#     rm -rf /var/lib/apt/lists/ && \
#     apt autoremove -y && \
#     apt autoclean -y

COPY ./bin /app

WORKDIR /app

EXPOSE ${GRPC_PORT} ${REST_PORT}
VOLUME /data/conf

USER nonroot:nonroot

CMD ["./server", "-conf", "/data/conf"]
