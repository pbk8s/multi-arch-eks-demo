# Build: 1st stage
#
FROM golang:1.18-alpine as builder 
ARG BUILDPLATFORM    
ARG TARGETPLATFORM
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"  
WORKDIR /app
COPY go.mod .
COPY hello.go .
RUN go build -o /hello && \
    apk add --update --no-cache file && \
    file /hello   

#
# Release: 2nd stage
#
FROM alpine
WORKDIR /
COPY --from=builder /hello /hello
CMD [ "/hello" ]
