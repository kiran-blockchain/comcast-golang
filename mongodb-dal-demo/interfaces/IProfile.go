// this interface will provide the requried methods
package interfaces

import "mongodemo/entities"

type IProfile interface {
	CreateProfile(profile *entities.Profile) (*entities.Profile, error)
}
