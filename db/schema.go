package db

import (
	"github.com/eaigner/hood"
)

type Profile struct {
	Id       hood.Id
	Username string `sql:validate:"presence"`
	Email    string `sql:"size(128),notnull`
	Password string `sql:"size(255),notnull`
	Created  hood.Created
	Updated  hood.Updated
}
