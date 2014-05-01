package main

import (
	"github.com/eaigner/hood"
)

func (m *M) CreateProfileTable_1398922624_Up(hd *hood.Hood) {
	type Profile struct {
		Id       hood.Id
		Username string `sql:validate:"presence"`
		Email    string `sql:"size(128),notnull`
		Password string `sql:"size(255),notnull`
		// These fields are auto updated on save
		Created hood.Created
		Updated hood.Updated
	}
	hd.CreateTable(&Profile{})
}

func (m *M) CreateProfileTable_1398922624_Down(hd *hood.Hood) {

}
