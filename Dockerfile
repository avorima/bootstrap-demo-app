FROM ubuntu

WORKDIR /data

COPY bootstrap-demo-app /data/app
COPY site /data/site

CMD ["/data/app"]
