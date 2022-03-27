FROM node:17 AS node
COPY ./ui /ui

RUN cd /ui && npm ci && npm run build

FROM golang:1.18 AS go

COPY --from=node /ui /ui
COPY ./api /app

RUN cd /app && go build && mv /app/photoframe /photoframe && rm -rf /app

FROM alpine
COPY --from=go /ui /ui
COPY --from=go /photoframe /photoframe

RUN apk add gcompat

ENV PHOTOS_PATH=/photos
ENV UI_PATH=/ui
ENV LISTEN=:80

ENTRYPOINT "/photoframe"