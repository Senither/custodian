# Setup the Node build stage
FROM node:22-alpine AS node-build

WORKDIR /app
COPY . .

# Installs and builds the frontend
RUN npm install
RUN npm run build:prod

# Setup the Go build stage
FROM golang:1.24 AS go-build

WORKDIR /app

COPY --from=node-build /app/public /app/public
COPY . .

# Disables CGO to create a lean image
ARG CGO_ENABLED=0

# Installs and builds the backend
RUN go mod download && go mod verify
RUN go build -o /custodian

# Setup a lean image to run the application
FROM scratch

WORKDIR /

COPY --from=go-build /custodian /custodian
COPY --from=go-build /app/.env.example /.env

EXPOSE 8000

CMD ["/custodian"]
