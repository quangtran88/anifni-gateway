FROM golang:1.19
WORKDIR /app
COPY . .
RUN ls
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/main ./
ENTRYPOINT ["/out/main"]