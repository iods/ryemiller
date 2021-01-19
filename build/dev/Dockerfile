# Use a tag (not latest)
ARG GOLANG_VER=1.15
FROM golang:$GOLANG_VER as builder

# Default labels & meta information
LABEL \
    maintaner="Rye Miller" \
    version="0.0"

# Set Environment Vars
ENV PORT=8080

WORKDIR /app
# COPY go.mod .
# RUN go mod download

COPY . .

RUN \
    cd cmd/ryemillerio \
    && go build .

CMD ["./cmd/ryemillerio/ryemillerio"]