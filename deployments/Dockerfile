FROM golang:1.22 AS builder

ARG SERVICE
WORKDIR /src
COPY . .

RUN if [ "$SERVICE" = "app" ]; then \
      cd app && go build -o /out/service ./cmd/server; \
    else \
      cd collector && go build -o /out/service ./cmd/collector; \
    fi

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /out/service /app/service
EXPOSE 8080 9090
CMD ["/app/service"]
