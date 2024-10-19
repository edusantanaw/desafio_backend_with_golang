package schema

type CustomerSchema struct {
	Name     string `json:"name"`
	Pass     string `json:"pass"`
	Email    string `json:"email"`
	CPF_CNPJ string `json:"cpf_cnpj"`
}
