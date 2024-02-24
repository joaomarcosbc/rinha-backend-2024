package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	application "github.com/joaomarcosbc/rinha-backend-2024/internal/application/transacao"
)

type TransacaoController struct {
	transacaoService application.Service
}

func NewTransacaoController(s application.Service) *TransacaoController {
	return &TransacaoController{
		transacaoService: s,
	}
}

func (tc *TransacaoController) ObterExtrato(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if id > 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	extrato, err := tc.transacaoService.ObterExtrato(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "applicarion/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(extrato)
}
