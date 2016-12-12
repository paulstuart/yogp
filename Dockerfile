FROM scratch
RUN mkdir /app
WORKDIR /app
ADD yogp 
CMD ["/app/yogp"]

