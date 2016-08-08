FROM iron/base
WORKDIR /app
# copy binary into image
COPY blob /app/
EXPOSE 3232
ENTRYPOINT ["./blob"]