package builder

type Response struct {
	User    User
	Account Account
	Address Address
}

type User struct {
	Email string
}

type Account struct {
	Balance float64
}

type Address struct {
	City string
}

type ResponseBuilder struct {
	Response Response
}

func NewBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		Response: Response{},
	}
}

func (rb *ResponseBuilder) User() *UserBuilder {
	return &UserBuilder{*rb}
}

func (rb *ResponseBuilder) Account() *AccountBuilder {
	return &AccountBuilder{*rb}
}

func (rb *ResponseBuilder) Address() *AddressBuilder {
	return &AddressBuilder{*rb}
}

func (rb *ResponseBuilder) Build() Response {
	return rb.Response
}

type UserBuilder struct {
	ResponseBuilder
}

func (ub *UserBuilder) Email(email string) *UserBuilder {
	ub.Response.User.Email = email
	return ub
}

type AccountBuilder struct {
	ResponseBuilder
}

func (ab *AccountBuilder) Balance(bal float64) *AccountBuilder {
	ab.Response.Account.Balance = bal
	return ab
}

type AddressBuilder struct {
	ResponseBuilder
}

func (addrb *AddressBuilder) City(city string) *AddressBuilder {
	addrb.Response.Address.City = city
	return addrb
}
