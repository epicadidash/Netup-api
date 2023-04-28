
FROM golang:1.19
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8000
CMD ["go","run","."]