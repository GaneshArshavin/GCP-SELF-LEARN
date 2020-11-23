From golang:1.12

RUN apt-get update || exit 0
RUN apt-get upgrade -y
RUN apt-get install vim sudo dbus curl bc supervisor -y

RUN apt-get install -y lsb-release
RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
RUN apt-get update
RUN apt-get install -y google-cloud-sdk
RUN apt-get clean

RUN mkdir -p /go/src/github.com/carousell/gcp-self-study
WORKDIR /go/src/github.com/carousell/gcp-self-study
COPY . /go/src/github.com/carousell/gcp-self-study
RUN go install github.com/carousell/gcp-self-study/server

EXPOSE 8080
