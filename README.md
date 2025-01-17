# Run

To run  on server with public IP address

```  docker-compose up ```

This will build and deploy application service and apache containers, and serve the UI on port 80

Needs tweaking to run on localhost

# Design

https://docs.google.com/document/d/e/2PACX-1vS5Ic8JQ91TSn0MRJYloWqs2QNi-1cS-61l9fzqiIS-aIsd4UkvTs_7rET0DaTp9qZoyCaDCF1b-n9q/pub

# Backend Service
## Configuration

config.env contains all the env variables we can configure e.g. db settings, ethereum node url etc. The flag `USE_DB_STORAGE` can be used to toggle between database storage and in-memory storage.
If this flag is true, please make sure you have setup database first by running `scripts/db.sql` and database credentials are properly setup.

## Building 

To build project, run `go build main.go`. This will generate an executable which when run, starts the server. By default, it uses port 80 which can be configured in config.env file.

## Endpoints

`POST /invoice`:

- Accepts: `amount: float` and `description: string`, generates an ethereum address and stores these values in storage.
- Response: `{ "id": 1, "status": "Unpaid", "amount": 0.000001, "paymentAddress": "0x0d3ee89c22379694f89fc7b4bcb0454a02cb972a", "paidAmount": 0 }`

`GET /invoice?id=1`:

- Accepts: `id: int` of the invoice, fetches invoice from storage, checks balance of the invoice from blockchain and updates status based on amoun paid.
- Response: `{ "id": 1, "status": "Unpaid", "amount": 0.000001, "paymentAddress": "0x0d3ee89c22379694f89fc7b4bcb0454a02cb972a", "paidAmount": 0 }`
