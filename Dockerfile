FROM golang

RUN mkdir /app
COPY . /app

WORKDIR /app/server/plugins

RUN chmod +x ./go-plugin-build.sh
RUN ./go-plugin-build.sh

WORKDIR /app/server

RUN go build -o ./dynago-server .

CMD [ "/app/server/dynago-server" ]