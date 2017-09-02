FROM golang:1.9
COPY example02.go /
RUN go build -o /example02 /example02.go
EXPOSE 8000
ENTRYPOINT ["/example02"]

