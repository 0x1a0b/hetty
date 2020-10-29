// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CloseProjectResult struct {
	Success bool `json:"success"`
}

type DeleteProjectResult struct {
	Success bool `json:"success"`
}

type HTTPHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type HTTPRequestLog struct {
	ID        int64            `json:"id"`
	URL       string           `json:"url"`
	Method    HTTPMethod       `json:"method"`
	Proto     string           `json:"proto"`
	Headers   []HTTPHeader     `json:"headers"`
	Body      *string          `json:"body"`
	Timestamp time.Time        `json:"timestamp"`
	Response  *HTTPResponseLog `json:"response"`
}

type HTTPRequestLogFilter struct {
	OnlyInScope bool `json:"onlyInScope"`
}

type HTTPRequestLogFilterInput struct {
	OnlyInScope *bool `json:"onlyInScope"`
}

type HTTPResponseLog struct {
	RequestID    int64        `json:"requestId"`
	Proto        string       `json:"proto"`
	StatusCode   int          `json:"statusCode"`
	StatusReason string       `json:"statusReason"`
	Body         *string      `json:"body"`
	Headers      []HTTPHeader `json:"headers"`
}

type Project struct {
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}

type ScopeHeader struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type ScopeHeaderInput struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type ScopeRule struct {
	URL    *string      `json:"url"`
	Header *ScopeHeader `json:"header"`
	Body   *string      `json:"body"`
}

type ScopeRuleInput struct {
	URL    *string           `json:"url"`
	Header *ScopeHeaderInput `json:"header"`
	Body   *string           `json:"body"`
}

type HTTPMethod string

const (
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodHead    HTTPMethod = "HEAD"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodConnect HTTPMethod = "CONNECT"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodTrace   HTTPMethod = "TRACE"
	HTTPMethodPatch   HTTPMethod = "PATCH"
)

var AllHTTPMethod = []HTTPMethod{
	HTTPMethodGet,
	HTTPMethodHead,
	HTTPMethodPost,
	HTTPMethodPut,
	HTTPMethodDelete,
	HTTPMethodConnect,
	HTTPMethodOptions,
	HTTPMethodTrace,
	HTTPMethodPatch,
}

func (e HTTPMethod) IsValid() bool {
	switch e {
	case HTTPMethodGet, HTTPMethodHead, HTTPMethodPost, HTTPMethodPut, HTTPMethodDelete, HTTPMethodConnect, HTTPMethodOptions, HTTPMethodTrace, HTTPMethodPatch:
		return true
	}
	return false
}

func (e HTTPMethod) String() string {
	return string(e)
}

func (e *HTTPMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HTTPMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HttpMethod", str)
	}
	return nil
}

func (e HTTPMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
