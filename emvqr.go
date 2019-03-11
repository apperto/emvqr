package emvqr

import (
	"errors"
	"fmt"

	"github.com/apperto/emvqr/crc16"
)

// Reference: https://www.mastercardlabs.com/masterpass-qr/

type QRData struct {
	MerchantIdentifiers map[int]map[int]string // Merchat identifiers form 16 to 51
	MCC                 string                 // TAG 52. Merchant Category Code
	Currency            string                 // TAG 53
	Country             string                 // TAG 58. ISO 3166-1 alpha 2
	MerchantName        string                 // TAG 59
	City                string                 // TAG 60
	Checksum            string                 // TAG 63
}

func (data QRData) GenerateString() (string, error) {

	if len(data.MCC) != 4 {
		return "", errors.New("MCC must be 4 digits")
	}

	if len(data.Currency) != 3 {
		return "", errors.New("Currency must be 3 numerical digits (ISO 4217)")
	}

	if len(data.Country) != 2 {
		return "", errors.New("Country must be 2 letters")
	}

	qr := "000201010211"
	for identifier, idValue := range data.MerchantIdentifiers {
		customData := ""
		for tag, value := range idValue {
			customData += fmt.Sprintf("%02d", tag)
			customData += fmt.Sprintf("%02d", len(value))
			customData += value
		}
		if len(customData) > 99 {
			return "", errors.New("custom identifier too long (max length: 99 characters)")
		}
		customData = fmt.Sprintf("%02d", identifier) + fmt.Sprintf("%02d", len(customData)) + customData
		qr += customData
	}

	qr += "52" + fmt.Sprintf("%02d", len(data.MCC)) + data.MCC
	qr += "5303" + data.Currency
	qr += "5802" + data.Country
	qr += "59" + fmt.Sprintf("%02d", len(data.MerchantName)) + data.MerchantName
	qr += "60" + fmt.Sprintf("%02d", len(data.City)) + data.City
	qr += "6304"
	qr += crc16.Checksum([]byte(qr))

	return qr, nil
}
