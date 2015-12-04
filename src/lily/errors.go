//
// Copyright (c) Telefonica I+D. All rights reserved.
//
package lily
import "net/http"

var (
	HTTP_400_MESSAGE = "Bad request"
	HTTP_404_MESSAGE = "Page not found"
	HTTP_405_MESSAGE = "Method not allowed"
	HTTP_500_MESSAGE = "Ups!!! We f*cked up somewhere. Maybe is better this way. This website is boring anyway."
)

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

type Http400 struct{
	HttpError
}

func NewHttp400(err string) *Http400 {
	return &Http400{*NewHttpError(err, http.StatusNotFound, HTTP_400_MESSAGE)}
}

func RaiseHttp400(err string) {
	panic(NewHttp400(err))
}

type Http404 struct{
	HttpError
}

func NewHttp404() *Http404 {
	return &Http404{*NewHttpError(HTTP_404_MESSAGE, http.StatusNotFound, HTTP_404_MESSAGE)}
}

func RaiseHttp404() {
	panic(NewHttp404())
}

type Http500 struct {
	HttpError
}

func NewHttp500(err string) *Http500 {
	return &Http500{*NewHttpError(err, http.StatusInternalServerError, HTTP_500_MESSAGE)}
}

func RaiseHttp500(err string) { panic(NewHttp500(err)) }
