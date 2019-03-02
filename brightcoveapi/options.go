package brightcoveapi

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////////////
// STRUCTS

// ClientOption interface which modifies the client on creation
type ClientOption interface {
	apply(*Client) error
}

////////////////////////////////////////////////////////////////////
// WithTimeout

type withTimeout time.Duration

func (this withTimeout) apply(client *Client) error {
	client.http.Timeout = time.Duration(this)
	return nil
}

// Provide connection timeout parameter to remote server
func WithTimeout(timeout time.Duration) ClientOption {
	return withTimeout(timeout)
}

////////////////////////////////////////////////////////////////////
// WithDebug

type withDebug bool

func (this withDebug) apply(client *Client) error {
	client.debug = bool(this)
	return nil
}

// Provide connection timeout parameter to remote server
func WithDebug(debug bool) ClientOption {
	return withDebug(debug)
}

////////////////////////////////////////////////////////////////////
// WithAccountId

type withAccountId string

func (this withAccountId) apply(client *Client) error {
	if len(this) > 0 {
		client.credentials.AccountId = string(this)
	}
	return nil
}

// Provide connection timeout parameter to remote server
func WithAccountId(account_id string) ClientOption {
	return withAccountId(account_id)
}

////////////////////////////////////////////////////////////////////
// WithLimit, WithOffset

type withLimitOffset struct {
	limit, offset uint32
}

func (this withLimitOffset) apply(client *Client) error {
	if this.limit > 0 {
		client.options.Set("limit", fmt.Sprint(this.limit))
	} else {
		client.options.Del("limit")
	}
	if this.offset > 0 {
		client.options.Set("offset", fmt.Sprint(this.offset))
	} else {
		client.options.Del("offset")
	}
	return nil
}

func WithLimit(limit uint32) ClientOption {
	return withLimitOffset{limit, 0}
}

func WithOffset(offset uint32) ClientOption {
	return withLimitOffset{0, offset}
}

func WithLimitOffset(limit, offset uint32) ClientOption {
	return withLimitOffset{limit, offset}
}

////////////////////////////////////////////////////////////////////
// WithSort

type withSort string

func (this withSort) apply(client *Client) error {
	if this != "" {
		client.options.Set("sort", string(this))
	} else {
		client.options.Del("sort")
	}
	return nil
}

func WithSort(sort string) ClientOption {
	return withSort(sort)
}

////////////////////////////////////////////////////////////////////
// WithQuery

type withQuery string

func (this withQuery) apply(client *Client) error {
	if this != "" {
		client.options.Set("q", string(this))
	} else {
		client.options.Del("q")
	}
	return nil
}

func WithQuery(q string) ClientOption {
	return withQuery(q)
}
