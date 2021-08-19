package helper

import (
	"github.com/go-sql-driver/mysql"
	"strconv"
	"unicode"
)

func ErrorResponse(e error) (int, string) {

	if mqErr, ok := e.(*mysql.MySQLError); ok {
		return int(mqErr.Number), mqErr.Message
	}

	err := e.Error() //the description of the error

	if len(err) < 6 { //if its too small return 0
		return 0, ""
	}
	i := 6 //Skip the part "Error "

	for ; len(err) > i && unicode.IsDigit(rune(err[i])); i++ {
	} // Raising i until we reach the end of err or we reach the end of error code

	n, e := strconv.Atoi(string(err[6:i])) //convert it to int

	if e != nil {
		return 0, "" //something went wrong
	}

	return n, string(err[i+2 : len(string(err))]) //return the error code
}

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message interface{} `json:"msg,omitempty"`
	//Error   string      `json:"error,omitempty"`
}
