package helper

import "strconv"

func ConvertStringToInt32(value string) int32 {

	convertValue, err := strconv.ParseInt(value, 10, 32)

	if err != nil {
		panic(err)
	}

	convertValueInt32 := int32(convertValue)

	return convertValueInt32

}

func ConvertStringToInt(value string) int {

	convertValue, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return convertValue

}

func ConvertStringToBool(value string) bool {

	convertValue, err := strconv.ParseBool(value)

	if err != nil {
		panic(err)
	}

	return convertValue

}

func ConvertIntToString(value int) string {

	convertValue := strconv.Itoa(value)

	return convertValue

}
