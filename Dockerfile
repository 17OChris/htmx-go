FROM node:16-alpine as nodeBuilder
RUN apk add make
COPY ./ ./
RUN make css-minify


# STAGE 1: building the executable
FROM golang:1.19-alpine AS build
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates
RUN apk add make
 
# add a user here because addgroup and adduser are not available in scratch
RUN addgroup -S myserver \
    && adduser -S -u 10000 -g myserver myserver
 
WORKDIR /src
# COPY ./go.mod ./go.sum ./
# RUN go mod download
 
COPY ./ ./
 
# Run tests
# RUN make test-api
 
# Build the executable
RUN make build-app
 
# STAGE 2: build the container to run
FROM scratch AS final
LABEL maintainer="me"
COPY --from=build /app /app
# this works - but we just want the src folder
# COPY --from=build /src /
COPY --from=build /src/src /src
COPY --from=nodeBuilder assets/output.css /assets/output.css
 
# copy ca certs
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
 
# copy users from builder (use from=0 for illustration purposes)
COPY --from=build /etc/passwd /etc/passwd
 
USER myserver

EXPOSE 80
 
ENTRYPOINT ["/app"]