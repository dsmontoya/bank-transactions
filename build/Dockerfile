ARG BUILD_TARGET
FROM golang:1.23.0-bullseye as build
ARG BUILD_TARGET

WORKDIR /app

COPY . .

RUN make setup-docker-go

RUN make build-go-$BUILD_TARGET


FROM gcr.io/distroless/base-debian11:debug-nonroot
ARG BUILD_TARGET

COPY deployments/$BUILD_TARGET-entrypoint.sh /entrypoint.sh

COPY assets /assets

COPY --from=build /app/generated/bin/cmd/$BUILD_TARGET/$BUILD_TARGET /app
