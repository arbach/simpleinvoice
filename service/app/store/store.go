package store

import (
	"context"

	"github.com/arbach/simpleinvoice/models"
)

type Store interface {
	GetInvoice(ctx context.Context, id int) (models.Invoice, error)
	GenerateInvoice(ctx context.Context, invoiceReq models.InvoiceRequest) (models.Invoice, error)
}
