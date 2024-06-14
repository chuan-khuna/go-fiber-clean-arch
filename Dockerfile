FROM golang:1.22.3-bullseye

WORKDIR /app
COPY . /app

RUN go get github.com/air-verse/air
RUN go install github.com/air-verse/air

# RUN go get -u github.com/gin-gonic/gin

RUN go get -u github.com/gofiber/fiber/v2
RUN go mod download
RUN go mod tidy

CMD ["air", "-c", ".air.toml"]