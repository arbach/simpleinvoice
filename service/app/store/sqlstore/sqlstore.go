package sqlstore

import (
	"context"
	"database/sql"

	"github.com/arbach/simpleinvoice/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type SqlStore struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *SqlStore {
	store := &SqlStore{db}
	return store
}

func (store *SqlStore) GetInvoice(ctx context.Context, id int) (models.Invoice, error) {
	var invoice models.Invoice
	if err := store.db.GetContext(ctx, &invoice, "SELECT * FROM invoices WHERE id = $1", id); err != nil {
		if sql.ErrNoRows == err {
			return invoice, errors.Errorf("Could not find invoice by id: %d", id)
		}
		return invoice, err
	}
	return invoice, nil
}

func (store *SqlStore) GenerateInvoice(ctx context.Context, invoiceReq models.InvoiceRequest) (models.Invoice, error) {
	var invoice models.Invoice
	insertQuery := `INSERT INTO invoices(amount, description, status, payment_address, paid_amount) 
					VALUES($1, $2, $3, $4, $5)
					RETURNING id, amount, description, status, payment_address, paid_amount, created_at, updated_at`
	err := store.db.QueryRowxContext(ctx, insertQuery, invoiceReq.Amount, invoiceReq.Description, "Unpaid", invoiceReq.PaymentAddress, 0.0).StructScan(&invoice)
	if err != nil {
		return invoice, err
	}

	return invoice, nil
}
