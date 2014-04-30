package users

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"github.com/eaigner/hood"
	"github.com/martini-contrib/render"
	"net/http"
)

type Profile struct {
	Id       hood.Id
	Username string `sql:"pk" validate:"presence"`
	Email    string `sql:validate:"presence"`
	Password []byte `sql:size(255)`
	// These fields are auto updated on save
	Created hood.Created
	Updated hood.Updated
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateProfile(r render.Render, rw http.ResponseWriter, req *http.Request, db *hood.Hood) {
	username, email, password := req.FormValue("username"), req.FormValue("email"), req.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	panicIf(err)

	user := Profile{}

	user.Password = hashedPassword
	user.Username = username
	user.Email = email

	tx := db.Begin()

	id, err := tx.Save(user)
	panicIf(err)

	err = tx.Commit()
	panicIf(err)

	fmt.Println("Inserted a new user with id of: ", id)

}
