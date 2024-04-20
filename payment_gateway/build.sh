RUN_NAME="payment_gateway"
mkdir -p output/bin output/conf output/data
go mod tidy
go build -o output/bin/${RUN_NAME}