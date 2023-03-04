package interfaces

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type CartRepoIF interface {
	Add(data *model.Cart) (*model.Cart, error)
	GetAll(userID string) (*model.Carts, error)
	Delete(id string) (*model.Cart, error)

	Checkout(data *model.Cart) (*model.Cart, error)
	GetCheckout(userID string) (*model.Carts, error)

	Payment() (*model.Histories, error)
}

type CartServiceIF interface {
	Add(data *model.Cart, productID, userID string) *lib.Response
	GetAll(userID string) *lib.Response
	Delete(id string) *lib.Response 

	Checkout(data *model.Cart) *lib.Response
	GetCheckout(userID string) *lib.Response

	Payment() *lib.Response
}
