package interfaces

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type ProductRepoIF interface {
	Add(data *model.Product) (*model.Product, error)
	GetAll() (*model.Products, error)
	GetById(id string) (*model.Product, error)
	Update(data *model.Product) (*model.Product, error)
	Delete(id string) (*model.Product, error)
	GetByCategory(category string) (*model.Products, error)
	Sort(by string, order string) (*model.Products, error)
	GetBySlug(slug string) (*model.Product, error)
	Search(query string) (*model.Products, error)
}

type ProductServiceIF interface {
	Add(data *model.Product) *lib.Response
	GetAll() *lib.Response
	GetById(id string) *lib.Response
	Update(data *model.Product) *lib.Response
	Delete(id string) *lib.Response
	GetByCategory(category string) *lib.Response
	Sort(by string, order string) *lib.Response
	GetBySlug(slug string) *lib.Response
	Search(query string) *lib.Response
}
