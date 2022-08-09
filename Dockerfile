FROM golang:1.18-alpine3.16 as BUILD
WORKDIR /app
COPY . .
RUN go build

FROM alpine:3.16 as RUNTIME
WORKDIR /app
COPY --from=BUILD /app/simpleapi simpleapi
CMD [ "./simpleapi" ]