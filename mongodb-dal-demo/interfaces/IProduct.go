// this interface will provide the requried methods
package interfaces

import "mongodemo/entities"

type IProduct interface {
	Create(product *entities.Product) (*entities.Product, error)
}
