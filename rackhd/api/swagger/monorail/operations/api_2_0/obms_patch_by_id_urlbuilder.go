package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	"strings"
)

// ObmsPatchByIDURL generates an URL for the obms patch by Id operation
type ObmsPatchByIDURL struct {
	Identifier string

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *ObmsPatchByIDURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/obms/{identifier}"

	identifier := o.Identifier
	if identifier != "" {
		_path = strings.Replace(_path, "{identifier}", identifier, -1)
	} else {
		return nil, errors.New("Identifier is required on ObmsPatchByIDURL")
	}
	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ObmsPatchByIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ObmsPatchByIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ObmsPatchByIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ObmsPatchByIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ObmsPatchByIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ObmsPatchByIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
