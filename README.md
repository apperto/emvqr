# emvqr
## Generate EVM compliant QR codes

Example code

``` go
testData := QRData{
  MerchantIdentifiers: map[int]map[int]string{
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
log.Println(qrstring, err)
```

Output: `000201010211454300245a32e2946e310600071234560111Test string5204343453030325802AR5907Facundo6013San Francisco63044917 <nil>`
