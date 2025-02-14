// Package handlers for the RESTful Server
//
// # Documentation for REST API
//
// Schemes: http
// BasePath: /
// Version: 0.1.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

// User defines the structure for a Full user Record
//
// swagger:model
type User struct {
	// the ID for the user
	// in: body
	//
	// required: false
	// min: 1
	ID int64 `json:"id"`

	// The username of the user
	// in: body
	//
	// required: true
	Name string `json:"username"`

	// The Password of the User
	//
	// required: true
	Password string `json:"password"`

	// The Last Login time of the User
	//
	// required: true
	// min: 0
	LastLogin int64 `json:"lastlogin"`

	// Is the User Admin or not
	//
	// required: true
	Admin int `json:"admin"`

	// Is the User Logged In or Not
	//
	// required: true
	Active int `json:"active"`
}

// swagger:parameters deleteID
type idParamWrapper struct {
	// The user id to be delete
	// in: path
	// required: true
	ID int `json:"id"`
}

// A User
// swagger:parameters getUserInfo loggedInfo
type UserInputWrapper struct {
	// A list of users
	// in: body
	Body User
}
