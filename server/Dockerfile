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

WORKDIR /app

COPY --from=builder /app/dynago-server .
COPY --from=builder /app/plugins/*.so ./plugins/

CMD [ "./dynago-server" ]