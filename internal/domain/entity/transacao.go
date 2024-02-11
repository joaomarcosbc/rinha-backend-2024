package entity

import "time"

type Transacao struct {
	ID          int       `json:"id"`
	Tipo        string    `json:"tipo"`
	Valor       int64     `json:"valor"`
	Descricao   int       `json:"descricao"`
	ClienteID   int       `json:"cliente_id"`
	RealizadaEm time.Time `json:"realizada_em"`
}
