package domains

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "nics",
		Price: 1.00,
		SKU:   "abc-def-feh",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
