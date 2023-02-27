FROM golang:alpine AS builder
RUN apk update && apk add --virtual build-dependencies build-base gcc wget git
ENV GIN_MODE=release
ENV PORT=8080
WORKDIR /goapp
RUN go install github.com/cosmtrek/air@latest
COPY ./golang ./goapp/
RUN cd ./goapp && go mod tidy
RUN cd ./goapp/ && go build
EXPOSE 8080
CMD ["air"]
#CMD ["bee run -gendoc=true -downdoc=true"]