package CreateCall

import (
	"context"
	"net/http"

	"google.golang.org/api/googleapi"
)

type SpreadsheetsCreateCall struct {
	s           *Service
	spreadsheet *Spreadsheet
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Create: Creates a spreadsheet, returning the newly created
// spreadsheet.
func (r *SpreadsheetsService) Create(spreadsheet *Spreadsheet) *SpreadsheetsCreateCall {
	c := &SpreadsheetsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.spreadsheet = spreadsheet
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpreadsheetsCreateCall) Fields(s ...googleapi.Field) *SpreadsheetsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpreadsheetsCreateCall) Context(ctx context.Context) *SpreadsheetsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpreadsheetsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}
