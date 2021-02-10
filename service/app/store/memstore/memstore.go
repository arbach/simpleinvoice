package memstore

import (
	"context"
	"time"

	"github.com/arbach/simpleinvoice/models"
	"github.com/pkg/errors"
)

type MemStore struct {
	invoices []models.Invoice
}

func New() *MemStore {
	store := &MemStore{}
	return store
}

func (store *MemStore) GetInvoice(ctx context.Context, id int) (models.Invoice, error) {
	for i := 0; i < len(store.invoices); i++ {
		if store.invoices[i].ID == id {
			return store.invoices[i], nil
		}
	}
	return models.Invoice{}, errors.Errorf("Could not find invoice by id: %d", id)
}

func (store *MemStore) GenerateInvoice(ctx context.Context, invoiceReq models.InvoiceRequest) (models.Invoice, error) {
	invoice := models.Invoice{
		ID:             len(store.invoices) + 1,
		Description:    invoiceReq.Description,
		PaidAmount:     0,
		Amount:         invoiceReq.Amount,
		PaymentAddress: invoiceReq.PaymentAddress,
		CreatedAt:      time.Now(),
	}
	invoice.SetStatus()
	store.invoices = append(store.invoices, invoice)
	return invoice, nil
}
