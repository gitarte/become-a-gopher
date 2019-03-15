# BUILDING APP
FROM    golang:1.12-stretch AS builder
WORKDIR /app
COPY    . /app
RUN     rm -rf /go/pkg/mod \
        && go build

# RUN
FROM    debian:stretch-slim
WORKDIR /app
RUN     apt -y update \
        && apt -y upgrade \
        && apt -y autoremove \
        && apt -y clean \
        && apt -y autoclean
COPY    --from=builder /app/httpServer /app/httpServer
CMD     ["/app/httpServer"]

