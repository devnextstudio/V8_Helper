package helper

import (
	"fmt"
	"net/url"
)

func ConvertQuery(encryptString string) map[string]string {

	m := make(map[string]string)

	crypto := GetCrypto()

	decryptoString, err := crypto.Decrypt(encryptString)

	if err != nil {
		fmt.Println(err)
	}

	params, err := url.ParseQuery(decryptoString)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for key, value := range params {
		m[key] = value[0]
	}

	return m

}

func ConvertData(encryptString string) string {

	crypto := GetCrypto()

	decryptoString, err := crypto.Decrypt(encryptString)

	if err != nil {
		fmt.Println(err)
	}

	return decryptoString

}
