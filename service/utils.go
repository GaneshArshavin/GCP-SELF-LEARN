package service

import (
	"errors"
	"math/rand"

	"github.com/carousell/chope-assignment/model"
)

//IsUserActive determines all the condition for an active businness user
func IsUserActive(user *model.AccountsUser, InputPassword string) (bool, error) {

	if user.IsActive.Bool && RotDn(user.Passowrd.String) == InputPassword {

		return false, nil
	} else if !user.IsActive.Bool {
		return false, errors.New("User is not Active")
	} else {
		return false, errors.New("Incorrect username and password")
	}
	return true, nil

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//randSeq generates the token to be stored in DB
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
