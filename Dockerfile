FROM golang:1.23-alpine AS builder

ARG APP_NAME=stress-test
ARG BUILD_DIR=/build

WORKDIR ${BUILD_DIR}

LABEL maintainer="DevOps Team" \
      description="Stress Test Application" \
      version="1.0"

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ${APP_NAME} .

FROM alpine:latest AS final

ARG APP_NAME=stress-test
ARG APP_DIR=/app
ENV APP_USER=appuser

RUN apk --no-cache add ca-certificates && \
    adduser -D -h ${APP_DIR} ${APP_USER}

WORKDIR ${APP_DIR}

COPY --from=builder /build/${APP_NAME} .

USER ${APP_USER}

ENTRYPOINT ["/app/stress-test"]