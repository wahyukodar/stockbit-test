package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
)

type ErrorNo int64

func (e ErrorNo) String() string {
	return strconv.FormatInt(int64(e), 10)
}

const (
	Success ErrorNo = 0

	ReqPathError = 10001

	DatabaseError = 40001

	BusParamError        = 30001
	BusParamConvertError = 30002

	ServiceError = 50001
)

var ReturnMsg = map[ErrorNo]string{

	Success: "success",

	ReqPathError: "Request path error",

	DatabaseError: "Something wrong communication with database ",

	// Business parameters
	BusParamError:        "Business parameter error",            // Parameter error
	BusParamConvertError: "Business parameter conversion error", // BusParamConvertError

	// Server-side processing exception
	ServiceError: "Server-side processing exception",
}

func errCommon(code ErrorNo, errMsg ...string) interface{} {
	return map[string]interface{}{
		"code":    code.String(),
		"message": ReturnMsg[code] + strings.Join(errMsg, " "),
	}
}

func RPCErr(c echo.Context, err error) error {
	st := status.Convert(err)
	log.Info(st.Message())
	switch st.Code() {
	case codes.InvalidArgument:
		return c.JSON(http.StatusBadRequest, errCommon(BusParamError, ":", st.Message()))
	default:
		return c.JSON(http.StatusInternalServerError, errCommon(ServiceError))
	}
}

func HandleSuccess(c echo.Context, i ...interface{}) error {
	resp := make(map[string]interface{})
	resp["code"] = Success
	resp["message"] = ReturnMsg[Success]
	switch len(i) {
	case 1:
		resp["data"] = i
	case 2:
		resp["data"] = i[0]
		resp["pages"] = i[1]
	case 3:
		resp["data"] = i[0]
		resp["pages"] = i[1]
		resp["page"] = i[2]
	case 4:
		resp["data"] = i[0]
		resp["pages"] = i[1]
		resp["page"] = i[2]
		resp["totalResults"] = i[3]
	default:
	}
	return c.JSON(http.StatusOK, resp)
}

func HandleError(c echo.Context, errCode ErrorNo, errMsg ...string) error {
	co := int64(errCode)
	var code int
	switch {
	case co == 0:
		code = http.StatusOK
	case co < 50000:
		code = http.StatusBadRequest
	case co < 60000:
		code = http.StatusInternalServerError
	default:
		code = http.StatusInternalServerError
	}
	return c.JSON(code, map[string]interface{}{
		"code":    errCode.String(),
		"message": ReturnMsg[errCode] + strings.Join(errMsg, " "),
	})
}
