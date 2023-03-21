FROM golang:1.19-alpine AS build

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
COPY . ./
# COPY dev.env ./
RUN go mod download
RUN echo $PORT
# RUN export PORT=8080


RUN go build -o /movie-suggestions-api

EXPOSE ${PORT}

CMD [ "/movie-suggestions-api" ]