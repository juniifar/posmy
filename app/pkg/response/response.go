package response

import (
	"encoding/json"
	"github.com/postmy/app/pkg/errors"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

// Response is response model
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

// Meta is meta model
type Meta struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// WriteResponse is function to write the response
func WriteResponse(controller *web.Controller, errs errors.IError, response interface{}) {
	customCode := 0
	customMessage := ""

	if errs != nil {
		customCode = errs.GetCustomCode()
		customMessage = errs.GetCustomMessage()
	}

	if customCode == 0 {
		customCode = http.StatusOK
		customMessage = "OK"
	} else {
		controller.Ctx.ResponseWriter.WriteHeader(errs.GetHTTPCode())
	}

	responseModel := Response{
		Meta: Meta{
			StatusCode: customCode,
			Message:    customMessage,
		},
		Data: response,
	}

	resp, _ := json.Marshal(responseModel)
	controller.Ctx.ResponseWriter.Write(resp)
}
