FROM golang:1.20.6-alpine3.18
WORKDIR /home
COPY ./ .
ENV PORT 8080
ENV GIN_MODE release
RUN go build .
RUN chmod +x main
EXPOSE 8080
CMD ["./main"]
