package users

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	"time"
)

type Profile struct {
	Id            int64  `primaryKey:"yes"`
	Username      string `sql:"not null;unique"`
	Email         string `sql:"not null;unique"`
	Password      string `sql:"not null;unique"`
	Created       time.Time
	Updated       time.Time
	Bio           string
	VerifiedEmail bool
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateProfile(r render.Render, rw http.ResponseWriter, req *http.Request, db *gorm.DB, log *log.Logger) {
	username, email, password, passwordConfirm := req.FormValue("username"), req.FormValue("email"), req.FormValue("password"), req.FormValue("passwordConfirm")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	panicIf(err)

	var possibleMatches []Profile
	db.Where("username = ?", username).Or("email = ?", email).Find(&possibleMatches)

	if len(possibleMatches) == 0 {

		if password != passwordConfirm {
			r.JSON(200, map[string]interface{}{
				"code": "password",
				"msg":  "Passwords must match"})
			return
		}

		user := Profile{}

		user.Password = string(hashedPassword)
		user.Username = username
		user.Email = email

		user.VerifiedEmail = false
		user.Created = time.Now()
		user.Updated = time.Now()

		db.Save(&user)

		r.JSON(200, map[string]interface{}{
			"code": 200,
			"msg":  "success"})

	} else {

		var usernameMatch bool
		var emailMatch bool

		if len(possibleMatches) > 1 {
			for i := 0; i < len(possibleMatches); i++ {

				if possibleMatches[i].Username == username {
					usernameMatch = true
				}

				if possibleMatches[i].Email == email {
					emailMatch = true
				}
			}
		} else {
			if possibleMatches[0].Username == username {
				usernameMatch = true
			}

			if possibleMatches[0].Email == email {
				emailMatch = true
			}
		}

		if usernameMatch == true && emailMatch == true {
			r.JSON(200, map[string]interface{}{
				"code": "error",
				"msg": map[string]interface{}{
					"username": "Username already being used. Try another.",
					"email":    "That email address has already been used to register an account"}})
		}

		if usernameMatch == true && emailMatch == false {
			r.JSON(200, map[string]interface{}{
				"code": "name",
				"msg":  "Username already being used. Try another."})
		}

		if emailMatch == true && usernameMatch == false {
			r.JSON(200, map[string]interface{}{
				"code": "email",
				"msg":  "That email address has already been used to register an account. "})
		}

	}
}
