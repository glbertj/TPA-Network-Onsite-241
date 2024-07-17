package utils

import "errors"

var UserNotFound = errors.New("user Not Found")
var InvalidPassword = errors.New("invalid Password")
var TokenInvalid = errors.New("token Invalid")
var ClaimInvalid = errors.New("claim Invalid")
var ParseTokenError = errors.New("error parsing token")
var UserAlreadyExist = errors.New("user already exist")
var SamePassword = errors.New("password is same as old password")
