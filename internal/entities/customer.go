package entities

type Customer struct {
	Id       string
	Name     string
	Email    string
	password string
	CPF_CNPJ string
}

func (c *Customer) GetPassword() string {
	return c.password
}

func (c *Customer) SetPassword(pass string) {
	c.password = pass
}
