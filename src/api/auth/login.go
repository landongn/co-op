package login

import (
	"api/users"
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
)

func AttemptLoginForUser(r render.Render, rw http.ResponseWriter, req *http.Request, db *gorm.DB) {

	username, password := req.FormValue("username"), req.FormValue("password")
	userRecord := users.Profile{Username: username}

	if err := db.Where("username = ?", username).First(&userRecord).Error; err != nil {
		hashedPass, oldpass := []byte(userRecord.Password), []byte(password)

		if bcrypt.CompareHashAndPassword(hashedPass, oldpass) != nil {
			r.JSON(200, map[string]interface{}{
				"error":  true,
				"reason": "Password or Username Invalid. Try again?"})
			return
		}

		r.JSON(200, map[string]interface{}{
			"error":    false,
			"redirect": "index"})
	}

}
