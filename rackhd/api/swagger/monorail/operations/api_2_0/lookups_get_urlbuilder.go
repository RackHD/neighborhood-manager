package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
)

// LookupsGetURL generates an URL for the lookups get operation
type LookupsGetURL struct {
	Q *string

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *LookupsGetURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/lookups"

	result.Path = _path

	qs := make(url.Values)

	var q string
	if o.Q != nil {
		q = *o.Q
	}
	if q != "" {
		qs.Set("q", q)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *LookupsGetURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *LookupsGetURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *LookupsGetURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on LookupsGetURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on LookupsGetURL")
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
func (o *LookupsGetURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
