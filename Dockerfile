FROM iron/go
WORKDIR /app
ADD main /app/
ENTRYPOINT ["./main"]
