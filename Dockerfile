FROM golang:1.23 AS backend-build

WORKDIR /backend
COPY ./backend .

RUN go mod download && go mod verify
RUN GOOS=linux CGO_ENABLED=0 go build -o ./linuxthekernel.io ./cmd/api

FROM node:18 AS frontend-build

WORKDIR /frontend
COPY ./frontend .

RUN npm install && \
    npm run build

FROM gcr.io/distroless/static-debian12 AS prod

WORKDIR /app

COPY --from=backend-build /backend/linuxthekernel.io .
COPY --from=frontend-build /frontend/dist/ ./static
COPY ./content /app/content

ENTRYPOINT ["/app/linuxthekernel.io"]
