FROM golang:latest

# if left blank app will run with dev settings
# to build production image run:
# $ docker build ./api --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env
ENV GOPATH=/go
ENV GOBIN=/go/api/bin
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

RUN go get -u github.com/pressly/goose/cmd/goose

RUN mkdir -p /go/api/bin
COPY api /go/api
WORKDIR /go/api

RUN go get ./vendor/database

RUN go get .
RUN go build

RUN ls -la /go/api
RUN ls -la /go/api/bin

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	go get -u github.com/pilu/fresh && \
	bin/fresh; \
	fi

EXPOSE 8080