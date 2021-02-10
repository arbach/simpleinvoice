package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/arbach/simpleinvoice/app/services"
	"github.com/arbach/simpleinvoice/app/store"
	"github.com/arbach/simpleinvoice/common"
	"github.com/arbach/simpleinvoice/models"
)

type Handler struct {
	store   store.Store
	service *services.Service
}

func New(store store.Store, service *services.Service) *Handler {
	h := &Handler{
		store:   store,
		service: service,
	}
	return h
}

func (handler *Handler) GetInvoice(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(fmt.Sprintf("Error parsing invoice id: %s", err.Error()))
		common.RespondWithError(r, w, http.StatusBadRequest, err.Error())
		return
	}

	invoice, err := handler.store.GetInvoice(r.Context(), id)
	if err != nil {
		log.Error(fmt.Sprintf("Error getting invoice: %s", err.Error()))
		common.RespondWithError(r, w, http.StatusNotFound, err.Error())
		return
	}
	invoice.PaidAmount, err = handler.service.GetBalanceInEther(invoice.PaymentAddress)
	if err != nil {
		log.Error(fmt.Sprintf("Error getting invoice balance: %s", err.Error()))
		common.RespondWithError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, invoice)
	return
}

func (handler *Handler) GenerateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoiceReq models.InvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&invoiceReq)
	if err != nil {
		log.Error(fmt.Sprintf("Error decoding invoice request: %s", err.Error()))
		common.RespondWithError(r, w, http.StatusBadRequest, fmt.Sprintf("Could not unmarshal json: %s", err.Error()))
		return
	}

	invoiceReq.PaymentAddress, err = handler.service.GenerateAddress()
	if err != nil {
		log.Error(fmt.Sprintf("Error generating address: %s", err.Error()))
		common.RespondWithError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	invoice, err := handler.store.GenerateInvoice(r.Context(), invoiceReq)
	if err != nil {
		log.Error(fmt.Sprintf("Error generating invoice: %s", err.Error()))
		common.RespondWithError(r, w, http.StatusInternalServerError, fmt.Sprintf("Could not generate invoice: %s", err.Error()))
		return
	}

	common.RespondWithJSON(w, http.StatusCreated, invoice)
	return
}
