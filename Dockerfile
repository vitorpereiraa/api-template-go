FROM golang:1.21.0-alpine

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-template-go

CMD ["/api-template-go"]