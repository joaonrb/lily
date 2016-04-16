//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"fmt"
	"net/http"
)

var (
	HTTP400 = "Bad request"
	HTTP404 = "Page not found"
	HTTP405 = "Method not allowed"
	HTTP500 = "Ups!!! We f*cked up somewhere. Maybe is better this way. This website is boring anyway."
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Http errors
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type IHttpError interface {
	ToResponse() *Response
}

type HttpError struct {
	*Response
	err string
}

func (self *HttpError) ToResponse() *Response {
	return self.Response
}

func NewHttpError(err string, status int, message string) *HttpError {
	return &HttpError{Response: &Response{status, map[string]string{}, message, nil}, err: err}
}

func RaiseHttpError(err string, status int, message string) { panic(NewHttpError(err, status, message)) }

func (self *HttpError) Error() string {
	return self.err
}

type Http400 struct {
	HttpError
}

func NewHttp400(err string) *Http400 {
	return &Http400{*NewHttpError(err, http.StatusNotFound, HTTP400)}
}

func RaiseHttp400(err string) {
	panic(NewHttp400(err))
}

type Http404 struct {
	HttpError
}

func NewHttp404() *Http404 {
	return &Http404{*NewHttpError(HTTP404, http.StatusNotFound, HTTP404)}
}

func RaiseHttp404() {
	panic(NewHttp404())
}

type Http500 struct {
	HttpError
}

func NewHttp500(err string) *Http500 {
	return &Http500{*NewHttpError(err, http.StatusInternalServerError, HTTP500)}
}

func RaiseHttp500(err string) { panic(NewHttp500(err)) }

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Routing errors
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type PathAlreadyExist struct {
	path string
}

func (self *PathAlreadyExist) Error() string {
	return fmt.Sprintf("The path %s already exists in router", self.path)
}

func NewPathAlreadyExist(path string) *PathAlreadyExist {
	return &PathAlreadyExist{path}
}
