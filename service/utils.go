package service

import (
	"math/rand"

	"github.com/carousell/chope-assignment/model"
)

//IsUserActive determines all the condition for an active businness user
func IsUserActive(user *model.AccountsUser, InputPassword string) bool {
	if user.IsActive.Bool && RotDn(user.Passowrd.String) == InputPassword {
		return true
	}
	return false
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
