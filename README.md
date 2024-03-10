# Receipt Processor 

## Dependencies
- Go 1.20

## Execution
### Start server
To start the server, run the following command. This starts the service on localhost:8080.
`go run internal/main.go`
### Test
- Sample request: `curl -X POST localhost:8080/receipts/process -H 'Content-Type: application/json' -d '{"retailer": "M&M Corner Market", "purchaseDate": "2022-03-20", "purchaseTime": "14:33", "items": [{"shortDescription": "Gatorade","price": "2.25"},{"shortDescription": "Gatorade", "price": "2.25"},{"shortDescription": "Gatorade", "price": "2.25" },{"shortDescription": "Gatorade","price": "2.25"}], "total":"9.00"}'`. It returns the UUID generated for this receipt.
- Check points: `curl -X GET localhost:8080/receipts/<ENTER_UUID_HERE>/points`. This should return 109.

## Implementation
- The REST API handler is written in Go with Gin framework to serve the APIs.
- SyncMap has been used for the in-memory database as a simple solution to provide support for concurrent use.
- Points are calculated based on the given examples and return errors in case of unexpected inputs. 

## Code details
- `internal/main.go` - Gin REST API server.
- `internal/api/service.go` - logic to compute points for a given receipt.
- `internal/api/service_test.go` - unit tests for point calculation logic.
- `internal/models/models.go` - Request and Response structures for JSON parsing. 
- `internal/handler/processor.go` - API handlers to process and fetch points for a receipt.
