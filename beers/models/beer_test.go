package models

import "testing"

func NewBeer(name string, price float64, description string) *CreateBeerCMD {
	return &CreateBeerCMD{
		Name:        name,
		Price:       price,
		Description: description,
	}
}
func Test_withCorrectParams(t *testing.T) {

	beer := NewBeer("Duff", 5, "Duff is a beer of the simpsons")

	err := beer.Validate()

	if err != nil {
		t.Errorf("The validation did not pass")
		t.Fail()
	}
}

func Test_shouldFailWithWrongLengthOfName(t *testing.T) {

	beer := NewBeer("Beer Beer Beer Beer Beer Beer Beer Beer Beer Beer Beer Beer", 5, "Duff is a beer of the simpsons")

	err := beer.Validate()

	if err == nil {
		t.Errorf("The max length of name is 50")
		t.Fail()
	}
}

func Test_shouldFailWithWrongLengthOfDescription(t *testing.T) {

	beer := NewBeer("Duff", 5,
		` Duff is a beer of the simpsons Duff is a beer of the simpsons Duff is a beer of the simpsons 
			Duff is a beer of the simpsons Duff is a beer of the simpsons Duff is a beer of the simpsons 
			Duff is a beer of the simpsons Duff is a beer of the simpsons Duff is a beer of the simpsons 
			Duff is a beer of the simpsons Duff is a beer of the simpsons Duff is a beer of the simpsons 
			Duff is a beer of the simpsons Duff is a beer of the simpsons Duff is a beer of the simpsons 
			Duff is a beer of the simpsons Duff is a beer of the simpsons Duff is a beer of the simpsons `,
	)

	err := beer.Validate()

	if err == nil {
		t.Errorf("The max length of description is 500")
		t.Fail()
	}
}

func Test_shouldFailWithWrongPrice(t *testing.T) {

	beer := NewBeer("Duff", -1, "Duff is a beer of the simpsons ")

	err := beer.Validate()

	if err == nil {
		t.Errorf("The price should be greater than 0")
		t.Fail()
	}
}
