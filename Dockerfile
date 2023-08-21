FROM golang:latest
LABEL authors="ruiwong"
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go build -o docker-gs-ping
EXPOSE 9092
CMD ["./docker-gs-ping"]
