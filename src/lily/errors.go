//
// Copyright (c) Telefonica I+D. All rights reserved.
//
package lily

var (
	HTTP_500_MESSAGE = "Ups!!! We f*cked up somewhere. Maybe is better this way. This website is boring anyway."
)

type IHttpError interface {
	error
	Message() string
	Status()  int
}

type Http500 error

func RaiseHttp500(err error) { panic(Http500(err)) }

func (self *Http500) Message() string { return HTTP_500_MESSAGE }

func (self *Http500) Status() int { return 500 }

type Http404 error

func RaiseHttp404(err error) { panic(Http404(err)) }

func (self *Http404) Message() string { return self.Error() }

func (self *Http404) Status() int { return 404 }

type HttpError struct {
	error
	status int
}

func NewHttpError(err error, status int) *HttpError {
	return &HttpError{err, status}
}

func RaiseHttpError(err error, status int) { panic(NewHttpError(err, status)) }

func (self *HttpError) Message() string { return self.Error() }

func (self *HttpError) Status() int { return self.status }