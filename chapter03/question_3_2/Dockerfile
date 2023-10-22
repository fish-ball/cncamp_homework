# Stage 1
FROM golang AS build_src
COPY main.go .
WORKDIR /go
RUN go build main.go

# Stage 2
FROM debian
COPY --from=build_src /go/main .
EXPOSE 80
CMD ./main
