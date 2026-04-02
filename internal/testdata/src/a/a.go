package a

import (
	HTTP "net/http"

	API "externaldep"
)

type APIError struct{} // want `identifier "APIError" should use camel-case abbreviations: "ApiError"`

type ApiError struct{}

type HTTPResult string // want `identifier "HTTPResult" should use camel-case abbreviations: "HttpResult"`

type ReaderAPI interface { // want `identifier "ReaderAPI" should use camel-case abbreviations: "ReaderApi"`
	ReadURL(apiURL string) HTTPResult // want `identifier "ReadURL" should use camel-case abbreviations: "ReadUrl"` `identifier "apiURL" should use camel-case abbreviations: "apiUrl"`
}

type UserRecord struct {
	UserID string // want `identifier "UserID" should use camel-case abbreviations: "UserId"`
}

type Cache[TID any] struct{} // want `identifier "TID" should use camel-case abbreviations: "Tid"`

func ParseAPI(userID string, apiURL string) (httpURL string, HTTPResult string) { // want `identifier "ParseAPI" should use camel-case abbreviations: "ParseApi"` `identifier "userID" should use camel-case abbreviations: "userId"` `identifier "apiURL" should use camel-case abbreviations: "apiUrl"` `identifier "httpURL" should use camel-case abbreviations: "httpUrl"` `identifier "HTTPResult" should use camel-case abbreviations: "HttpResult"`
	localAPI := APIError{} // want `identifier "localAPI" should use camel-case abbreviations: "localApi"`
	_ = localAPI
	_ = API.APIError{}
	_ = HTTP.Client{}

	return "", ""
}

func (userID UserRecord) LookupURL(apiURL string) string { // want `identifier "userID" should use camel-case abbreviations: "userId"` `identifier "LookupURL" should use camel-case abbreviations: "LookupUrl"` `identifier "apiURL" should use camel-case abbreviations: "apiUrl"`
	return apiURL
}

func RangeValues(values map[string]string) {
	for userID, apiURL := range values { // want `identifier "userID" should use camel-case abbreviations: "userId"` `identifier "apiURL" should use camel-case abbreviations: "apiUrl"`
		_, _ = userID, apiURL
	}
}

var _ = struct {
	UserID string // want `identifier "UserID" should use camel-case abbreviations: "UserId"`
}{}

func WrapAPI[TID any](value TID) TID { // want `identifier "WrapAPI" should use camel-case abbreviations: "WrapApi"` `identifier "TID" should use camel-case abbreviations: "Tid"`
	return value
}

const DefaultURL = "https://example.test" // want `identifier "DefaultURL" should use camel-case abbreviations: "DefaultUrl"`

const (
	ServiceURL = "https://service.test" // want `identifier "ServiceURL" should use camel-case abbreviations: "ServiceUrl"`
)

var globalHTTPClient = 1 // want `identifier "globalHTTPClient" should use camel-case abbreviations: "globalHttpClient"`

var (
	apiURLValue = 1 // want `identifier "apiURLValue" should use camel-case abbreviations: "apiUrlValue"`
)
