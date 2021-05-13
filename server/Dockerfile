FROM golang as builder

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=1 GOOS=linux go build -o ./dynago-server .

# actual binary runner
FROM golang

ARG redis_host=redis
ARG redis_port=6379
ENV REDIS_HOST ${redis_host}
ENV REDIS_PORT ${redis_port}

WORKDIR /app

COPY --from=builder /app/dynago-server .

RUN mkdir plugins

CMD ./dynago-server -redis-host=${REDIS_HOST} -redis-port=${REDIS_PORT} -redis-secret-path=${REDIS_SECRET}