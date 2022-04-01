FROM golang:alpine as build
WORKDIR /usr/app
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/todo-api/

FROM gcr.io/distroless/static as release
WORKDIR /usr/app
COPY --from=0 /usr/app/todo-api .
COPY app.env app.env
USER nonroot
CMD ["./todo-api"]

FROM release
