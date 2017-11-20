FROM golang:1.9.2-stretch

LABEL maintainer="user@email.com"

RUN apt-get update && apt-get -y upgrade 

RUN go get github.com/JohanAanesen/CloudTech_oblig3/cmd/currency

RUN go install github.com/JohanAanesen/CloudTech_oblig3/cmd/currency

WORKDIR /src/github.com/JihanAanesen/CloudTech_oblig3/cmd/currency

CMD ["/currency"] 
