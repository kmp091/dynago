FROM golang as builder

WORKDIR /app
COPY . /app

WORKDIR /app/plugins

RUN chmod +x ./go-plugin-build.sh
RUN ./go-plugin-build.sh

WORKDIR /app

RUN CGO_ENABLED=1 GOOS=linux go build -o ./dynago-server .

# actual binary runner
FROM debian

ARG redis_host=redis
ARG redis_port=6379
ENV REDIS_HOST ${redis_host}
ENV REDIS_PORT ${redis_port}

WORKDIR /app

COPY --from=builder /app/dynago-server .
COPY --from=builder /app/plugins/*.so ./plugins/

CMD ./dynago-server -redis-host=${REDIS_HOST} -redis-port=${REDIS_PORT}