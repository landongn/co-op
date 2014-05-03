package users

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
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

func CreateProfile(r render.Render, rw http.ResponseWriter, req *http.Request, db *gorm.DB) {
	username, email, password := req.FormValue("username"), req.FormValue("email"), req.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	panicIf(err)

	if db.Where("username = ?", username).Or("email = ?", email).Find(&Profile{Username: username, Email: email}).RecordNotFound() {
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
		r.JSON(200, map[string]interface{}{
			"code": 403,
			"msg":  "Username or Email Address already taken."})
	}
}
