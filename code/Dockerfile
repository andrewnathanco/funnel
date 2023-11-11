# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git gcc libc-dev
WORKDIR /go/src/app
COPY . .
RUN go mod edit -module mural
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /go/bin/mural main.go

# final stage
FROM alpine:latest
ARG VERSION
ENV VERSION=$VERSION
WORKDIR /go/src/app
LABEL Name=mural Version=$VERSION
RUN apk --no-cache add ca-certificates

# setup node
RUN apk add --update npm
RUN npm init -y
RUN npm install tailwindcss postcss-cli autoprefixer

COPY --from=builder /go/bin/mural ./mural
COPY --from=builder /go/src/app/view ./view
COPY --from=builder /go/src/app/static ./static
COPY --from=builder /go/src/app/tailwind.config.js .

# run tailwind
RUN npx tailwindcss -i ./view/input.css -o ./static/styles.css
RUN touch .env 
EXPOSE 1323
ENTRYPOINT ./mural
