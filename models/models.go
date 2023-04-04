package models

type Plano struct {
	Plano   string  `json:"Plano"`
	Duracao string  `json:"Duracao"`
	Valor   float32 `json:"Valor"`
}

type Assinante struct {
	Nome       string `json:"Nome"`
	SobreNome  string `json:"SobreNome"`
	Plano      string `json:"Plano"`
	Validade   string `json:"Validade"`
	DataInicio string `json:"DataInicio"`
}
