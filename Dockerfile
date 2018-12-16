FROM golang:stretch AS builder
RUN cd /; wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt-get update
RUN dpkg -i /google-chrome-stable_current_amd64.deb 2>/dev/null || true
RUN apt-get install -qq -y -f

ENV GOAPP github.com/noonien/resume
ENV GODIR $GOPATH/src/$GOAPP
WORKDIR $GODIR
COPY . .
RUN go build -mod=vendor -o resume-srv && (./resume-srv >/dev/null &) &&\
    go run -mod=vendor genpdf/main.go
RUN go install -mod=vendor $GODIR/vendor/github.com/gobuffalo/packr/packr
RUN packr
RUN go install -mod=vendor


FROM bitnami/minideb
COPY --from=builder /go/bin/resume /
CMD /resume
