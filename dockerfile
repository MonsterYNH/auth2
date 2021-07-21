FROM golang:latest as builder
	
RUN mkdir /app
WORKDIR /app
COPY . .

RUN go build -o auth2 main.go

FROM ubuntu:latest
RUN mkdir /app
COPY --from=builder /app/auth2 /app
COPY ./config.yaml /app

WORKDIR /app

CMD [ "./auth2", "--config.file", "config.yaml" ]