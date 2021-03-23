package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func writeErrorCode(w gin.ResponseWriter, statusCode string, httpCode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	errMsg := ErrorResponse{
		ErrorMessage: errorMessage,
		Status:       statusCode,
	}

	x, err := json.Marshal(&errMsg)
	if err == nil {
		_, _ = w.Write(x)
	}
}

func writeResponse(w gin.ResponseWriter, response interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	x, err := json.Marshal(&response)

	if err == nil {
		_, _ = w.Write(x)
	}
}

func validateTransactionReq(req Transaction) (string, bool) {

	if strings.Compare(reflect.TypeOf(req.Amount).String(), reflect.Float64.String()) == -1 {
		return "Invalid value in field `amount`. ", true
	}

	if strings.Compare(reflect.TypeOf(req.Type).String(), reflect.String.String()) == -1 {
		return "Invalid value in field `type`. ", true
	}

	if req.ParentID <= -1 && strings.Compare(reflect.TypeOf(req.ParentID).String(), reflect.Float32.String()) == -1 {
		return "Invalid value in field `parentID`. ", true
	}

	return "", false
}

func PutTransaction(c *gin.Context) {

	var transactionReq Transaction
	err := json.NewDecoder(c.Request.Body).Decode(&transactionReq)
	if err != nil {
		fmt.Println("Error decoding request body.")
		writeErrorCode(c.Writer, "error", 400, "Failed to decode the body")
		return
	}

	reqErrMsg, reqErr := validateTransactionReq(transactionReq)

	if reqErr {
		fmt.Println("Error, incorrect value in request body.", reqErrMsg)
		writeErrorCode(c.Writer, "error", 400, reqErrMsg)
	}

	transactionId := c.Param("transaction_id")

	transactionReq.transactionIDs = transactionId

	go putTransaction(transactionReq)

	writeResponse(c.Writer, struct {
		Status string `json:"status"`
	}{Status: "OK"})

}

func GetTransaction(c *gin.Context) {

	transactionId := c.Param("transaction_id")

	if len(transactionId) == 0 {
		writeErrorCode(c.Writer, "error", 400, "Invalid value for `transaction_id`. ")
	}

	data, isPresent := getTransaction(transactionId)

	if isPresent {
		writeResponse(c.Writer, data)
	} else {
		writeErrorCode(c.Writer, "ok", 404, "No Data for the given `transaction_id`.")
	}

}

func GetTransactionType(c *gin.Context) {

	typesValue := c.Param("types_value")

	if len(typesValue) == 0 {
		writeErrorCode(c.Writer, "error", 400, "Invalid value for `types`. ")
	}

	data, isPresent := getTransactionType(typesValue)

	fmt.Println(data, isPresent)

	if isPresent {
		writeResponse(c.Writer, struct {
			TransactionIDs []string `json:"transaction_IDs"`
		}{TransactionIDs: data})
	} else {
		writeErrorCode(c.Writer, "ok", 404, "No Data for the given `type`.")
	}

}

func GetTransactionSum(c *gin.Context) {

	transactionId := c.Param("transaction_id")

	if len(transactionId) == 0 {
		writeErrorCode(c.Writer, "error", 400, "Invalid value for `types`. ")
	}

	data, isPresent := getTransactionSum(transactionId)

	if isPresent {
		writeResponse(c.Writer, struct {
			Sum float64 `json:"sum"`
		}{Sum: data})
	} else {
		writeErrorCode(c.Writer, "ok", 404, "Error Finding sum please make sure `transaction_id` exists.")
	}

}
