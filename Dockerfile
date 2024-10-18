FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o /main /app/cmd/main.go

FROM public.ecr.aws/lambda/go:latest

COPY --from=builder /main /var/task/main

CMD ["main"]

