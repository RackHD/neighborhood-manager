package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"

	"github.com/go-openapi/swag"
)

// WorkflowsGetURL generates an URL for the workflows get operation
type WorkflowsGetURL struct {
	DollarSkip *int64
	DollarTop  *int64
	Active     *bool
	Sort       *string

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *WorkflowsGetURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/workflows"

	result.Path = _path

	qs := make(url.Values)

	var dollarSkip string
	if o.DollarSkip != nil {
		dollarSkip = swag.FormatInt64(*o.DollarSkip)
	}
	if dollarSkip != "" {
		qs.Set("$skip", dollarSkip)
	}

	var dollarTop string
	if o.DollarTop != nil {
		dollarTop = swag.FormatInt64(*o.DollarTop)
	}
	if dollarTop != "" {
		qs.Set("$top", dollarTop)
	}

	var active string
	if o.Active != nil {
		active = swag.FormatBool(*o.Active)
	}
	if active != "" {
		qs.Set("active", active)
	}

	var sort string
	if o.Sort != nil {
		sort = *o.Sort
	}
	if sort != "" {
		qs.Set("sort", sort)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *WorkflowsGetURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *WorkflowsGetURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *WorkflowsGetURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on WorkflowsGetURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on WorkflowsGetURL")
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
func (o *WorkflowsGetURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
