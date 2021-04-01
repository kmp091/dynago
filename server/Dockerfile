FROM golang

RUN mkdir /app
COPY . /app

WORKDIR /app/plugins

RUN chmod +x ./go-plugin-build.sh
RUN ./go-plugin-build.sh

WORKDIR /app

RUN go build -o ./dynago-server .

CMD [ "/app/dynago-server" ]