package emvqr

import (
	"testing"
)

func TestGenerateString(t *testing.T) {
	testData := QRData{
		MerchantIdentifiers: map[int]map[int]string{
			50: map[int]string{
				0: "somecuit",
			},
			51: map[int]string{
				0: "somealias",
			},
			45: map[int]string{
				0: "5a32e2946e31060007123456",
				1: "Test string",
			},
		},
		MCC:          "3434",
		Currency:     "032",
		Country:      "AR",
		MerchantName: "Facundo",
		City:         "San Francisco",
	}

	qrstring, err := testData.GenerateString()
	if err != nil {
		t.Error("Error should be nil", err)
	}
	if qrstring != "00020101021150120008somecuit51130009somealias454300245a32e2946e310600071234560111Test string5204343453030325802AR5907Facundo6013San Francisco630431d0" {
		t.Error("Output doesn't match", qrstring)
	}
}
