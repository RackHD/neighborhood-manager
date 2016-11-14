package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	"strings"
)

// LookupsDelByIDURL generates an URL for the lookups del by Id operation
type LookupsDelByIDURL struct {
	ID string

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *LookupsDelByIDURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/lookups/{id}"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("ID is required on LookupsDelByIDURL")
	}
	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *LookupsDelByIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *LookupsDelByIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *LookupsDelByIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on LookupsDelByIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on LookupsDelByIDURL")
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
func (o *LookupsDelByIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
