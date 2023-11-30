FROM golang:1.21.4 as build

WORKDIR /annuums
# Copy dependencies list
# # Build with optional lambda.norpc tag
COPY . .
RUN go build -tags lambda.norpc -o main starter.go

# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /annuums/main /main

EXPOSE 8080
ENTRYPOINT [ "/main" ]