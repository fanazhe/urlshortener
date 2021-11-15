# Fontend generation start
FROM node:14.18-alpine as frontend

WORKDIR /frontend-build

ENV NODE_ENV production

COPY ./web/package.json ./web/yarn.lock ./
RUN yarn

COPY ./web/ .

RUN yarn build
# Frontend generated

# Backend generation start
FROM golang:1.17.3-alpine3.13 as backend

WORKDIR /backend-build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# -ldflags="-w -s" means omit DWARF symbol table and the symbol table and debug information
RUN CGO_ENABLED=0 GOOS=linux go build --tags "release" -ldflags="-w -s" -o urlshortener ./main.go
# Backend generated

# Use redis docker to run
FROM alpine:3.14.3

# Copy frontend asset
RUN mkdir -p /app/web/dist
COPY --from=frontend /frontend-build/dist /app/web/dist

# Copy backend executable file
COPY --from=backend /backend-build/urlshortener /app/
RUN mkdir -p /app/config
COPY config/config.docker.json /app/config/config.json

WORKDIR /app

EXPOSE 8082

CMD ["/app/urlshortener"]
