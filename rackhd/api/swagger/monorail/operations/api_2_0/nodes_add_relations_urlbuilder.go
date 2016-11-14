package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	"strings"
)

// NodesAddRelationsURL generates an URL for the nodes add relations operation
type NodesAddRelationsURL struct {
	Identifier string

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *NodesAddRelationsURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/nodes/{identifier}/relations"

	identifier := o.Identifier
	if identifier != "" {
		_path = strings.Replace(_path, "{identifier}", identifier, -1)
	} else {
		return nil, errors.New("Identifier is required on NodesAddRelationsURL")
	}
	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *NodesAddRelationsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *NodesAddRelationsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *NodesAddRelationsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on NodesAddRelationsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on NodesAddRelationsURL")
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
func (o *NodesAddRelationsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
