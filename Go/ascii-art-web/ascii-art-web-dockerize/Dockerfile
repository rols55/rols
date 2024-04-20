# syntax=docker/dockerfile:1
FROM golang:1.19-alpine
LABEL description="Project: Ascii-Art-Web 0.0.1-beta"
LABEL vendor="Team - Jaan.Ginzul, rols55, _Parker_"  
COPY . /app
WORKDIR /app
EXPOSE 8080
RUN apk add build-base
RUN go build -o ascii . 
CMD ["./ascii"]