FROM iron/go
WORKDIR /app
ADD myapp /app/
ENTRYPOINT ["./myapp"]
