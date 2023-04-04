package models

type Plano struct {
	Plano   string  `json:"plano"`
	Duracao string  `json:"duracao"`
	Valor   float32 `json:"valor"`
}

type Assinante struct {
	Nome       string `json:"nome"`
	SobreNome  string `json:"sobrenome"`
	Plano      string `json:"plano"`
	Validade   string `json:"validade"`
	DataInicio string `json:"datainicio"`
}
