package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	"strings"
)

// PollersIDGetURL generates an URL for the pollers Id get operation
type PollersIDGetURL struct {
	Identifier string

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *PollersIDGetURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/pollers/{identifier}"

	identifier := o.Identifier
	if identifier != "" {
		_path = strings.Replace(_path, "{identifier}", identifier, -1)
	} else {
		return nil, errors.New("Identifier is required on PollersIDGetURL")
	}
	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PollersIDGetURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PollersIDGetURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PollersIDGetURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PollersIDGetURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PollersIDGetURL")
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
func (o *PollersIDGetURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
