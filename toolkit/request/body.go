package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	requestIsNil             = "request_is_nil"
	requestBodyIsNil         = "request_body_is_nil"
	targetIsNil              = "target_is_nil"
	readAllBodyError         = "read_all_body_error"
	decodeBodyUnmarshalError = "decode_body_unmarshal_error"
)

func Unmarshal(request *http.Request, target interface{}) error {
	if request == nil {
		return errors.New(requestIsNil)
	}

	if request.Body == nil {
		return errors.New(requestBodyIsNil)
	}

	if target == nil {
		return errors.New(targetIsNil)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("code:%s error:%s", readAllBodyError, err.Error()))
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return errors.New(fmt.Sprintf("code:%s error:%s", decodeBodyUnmarshalError, err.Error()))
	}

	return nil
}
