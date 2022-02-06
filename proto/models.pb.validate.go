// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/models.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Activity with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Activity) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetInternalID()) < 1 {
		return ActivityValidationError{
			field:  "InternalID",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetExternalID()) < 1 {
		return ActivityValidationError{
			field:  "ExternalID",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Time

	// no validation rules for UserID

	// no validation rules for ActivityType

	// no validation rules for ExtraData

	return nil
}

// ActivityValidationError is the validation error returned by
// Activity.Validate if the designated constraints aren't met.
type ActivityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActivityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActivityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActivityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActivityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActivityValidationError) ErrorName() string { return "ActivityValidationError" }

// Error satisfies the builtin error interface
func (e ActivityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActivity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActivityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActivityValidationError{}

// Validate checks the field values on Feed with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Feed) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetInternalID()) < 1 {
		return FeedValidationError{
			field:  "InternalID",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for UserID

	// no validation rules for ExtraData

	return nil
}

// FeedValidationError is the validation error returned by Feed.Validate if the
// designated constraints aren't met.
type FeedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedValidationError) ErrorName() string { return "FeedValidationError" }

// Error satisfies the builtin error interface
func (e FeedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeed.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedValidationError{}