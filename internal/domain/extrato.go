package domain

import "time"

type ExtratoDto struct {
	Saldo             SaldoClienteDto `json:"saldo"`
	UltimasTransacoes []TransacaoDto  `json:"ultimas_transacoes"`
}

type SaldoClienteDto struct {
	Total       int64     `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      int64     `json:"limite"`
}

type TransacaoDto struct {
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	Valor       int64     `json:"valor"`
	RealizadaEm time.Time `json:"realizada_em"`
}
