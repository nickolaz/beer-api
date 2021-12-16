package models

import "errors"

const maxLengthName = 50
const maxLengthDescription = 500

// Beer model structure for beer
type Beer struct {
	Id          uint64  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Brewery     string  `json:"brewery"`
	Country     string  `json:"country"`
	Description string  `json:"description"`
}

// CreateBeerCMD model structure for create a Beer
type CreateBeerCMD struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Brewery     string  `json:"brewery"`
	Country     string  `json:"country"`
	Description string  `json:"description"`
}

type ErrorMsg struct {
	Msg string `json:"msg"`
}

// Validate function for validate a beer model
func (cmd *CreateBeerCMD) Validate() error {
	if len(cmd.Name) == 0 || len(cmd.Description) == 0 || cmd.Price < 0 {
		return errors.New("invalid beer data")
	}
	if len(cmd.Name) > maxLengthName || len(cmd.Description) > maxLengthDescription {
		return errors.New("invalid beer data")
	}
	return nil
}
