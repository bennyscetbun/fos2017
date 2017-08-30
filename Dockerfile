FROM golang:1.9
EXPOSE 5042

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install

ENV PORT 5042
ENV PHOTO_PATH "/var/lib/fos2017/photo"

RUN mkdir -p /var/lib/fos2017/photo

CMD ["go-wrapper", "run"]