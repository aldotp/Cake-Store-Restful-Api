FROM golang:1.19-alpine

RUN apk update && apk add libc-dev && apk add gcc && apk add make 
WORKDIR /app

COPY . . 
RUN go mod tidy 
RUN go mod download

# CMD ["go", "run", "main.go"]
CMD go test -v ./... && go run main.go 


