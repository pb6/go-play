FROM scratch
COPY ./envleaker /app/
EXPOSE 8080
USER 1000
ENTRYPOINT ["/app/envleaker"]
