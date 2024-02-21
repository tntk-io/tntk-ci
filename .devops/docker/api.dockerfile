FROM --platform=amd64 golang:1.17 as build
WORKDIR /workbench
COPY . .
RUN GOARCH=amd64 GOOS=linux go build -o /output/api ./main.go
RUN chmod +x /output/api



FROM --platform=amd64 ubuntu as runtime

COPY --from=build /output/api /usr/bin/api

RUN apt update && apt install ca-certificates -y && rm -rf /var/lib/apt/lists/*

ENV LOCALSTACK_MODE_ENABLED=false
ENV AWS_ACCESS_KEY_ID=""
ENV AWS_SECRET_ACCESS_KEY=""
ENV AWS_REGION=us-east-1

ENV S3_BUCKET_NAME=test
ENV SQS_QUEUE_NAME=unprocessed
ENV DYNAMODB_TABLE_NAME=test
ENV DB_HOST=localhost
ENV DB_PORT=15432
ENV DB_USERNAME=api
ENV DB_PASSWORD=postgres
ENV DB_NAME=api
ENV API_SOCKET=:80

CMD ["/usr/bin/api"]