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
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on Activity with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Activity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Activity with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ActivityMultiError, or nil
// if none found.
func (m *Activity) ValidateAll() error {
	return m.validate(true)
}

func (m *Activity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetStringId()) < 1 {
		err := ActivityValidationError{
			field:  "StringId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ObjectId

	// no validation rules for LinkedActivityId

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for UserId

	// no validation rules for ActivityType

	if all {
		switch v := interface{}(m.GetExtraData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ActivityValidationError{
					field:  "ExtraData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ActivityValidationError{
					field:  "ExtraData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExtraData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ActivityValidationError{
				field:  "ExtraData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ActivityMultiError(errors)
	}
	return nil
}

// ActivityMultiError is an error wrapping multiple validation errors returned
// by Activity.ValidateAll() if the designated constraints aren't met.
type ActivityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActivityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActivityMultiError) AllErrors() []error { return m }

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
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Feed) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Feed with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FeedMultiError, or nil if none found.
func (m *Feed) ValidateAll() error {
	return m.validate(true)
}

func (m *Feed) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := FeedValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for UserId

	if all {
		switch v := interface{}(m.GetExtraData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FeedValidationError{
					field:  "ExtraData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FeedValidationError{
					field:  "ExtraData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExtraData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedValidationError{
				field:  "ExtraData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FeedMultiError(errors)
	}
	return nil
}

// FeedMultiError is an error wrapping multiple validation errors returned by
// Feed.ValidateAll() if the designated constraints aren't met.
type FeedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeedMultiError) AllErrors() []error { return m }

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

// Validate checks the field values on GroupingFeed with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GroupingFeed) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GroupingFeed with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GroupingFeedMultiError, or
// nil if none found.
func (m *GroupingFeed) ValidateAll() error {
	return m.validate(true)
}

func (m *GroupingFeed) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := GroupingFeedValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for UserId

	if utf8.RuneCountInString(m.GetKeyFormat()) < 1 {
		err := GroupingFeedValidationError{
			field:  "KeyFormat",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetExtraData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GroupingFeedValidationError{
					field:  "ExtraData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GroupingFeedValidationError{
					field:  "ExtraData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExtraData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GroupingFeedValidationError{
				field:  "ExtraData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GroupingFeedMultiError(errors)
	}
	return nil
}

// GroupingFeedMultiError is an error wrapping multiple validation errors
// returned by GroupingFeed.ValidateAll() if the designated constraints aren't met.
type GroupingFeedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GroupingFeedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GroupingFeedMultiError) AllErrors() []error { return m }

// GroupingFeedValidationError is the validation error returned by
// GroupingFeed.Validate if the designated constraints aren't met.
type GroupingFeedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GroupingFeedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GroupingFeedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GroupingFeedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GroupingFeedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GroupingFeedValidationError) ErrorName() string { return "GroupingFeedValidationError" }

// Error satisfies the builtin error interface
func (e GroupingFeedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGroupingFeed.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GroupingFeedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GroupingFeedValidationError{}

// Validate checks the field values on Collection with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Collection) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Collection with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CollectionMultiError, or
// nil if none found.
func (m *Collection) ValidateAll() error {
	return m.validate(true)
}

func (m *Collection) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := CollectionValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := DeletingType_name[int32(m.GetDeletingType())]; !ok {
		err := CollectionValidationError{
			field:  "DeletingType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CollectionMultiError(errors)
	}
	return nil
}

// CollectionMultiError is an error wrapping multiple validation errors
// returned by Collection.ValidateAll() if the designated constraints aren't met.
type CollectionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CollectionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CollectionMultiError) AllErrors() []error { return m }

// CollectionValidationError is the validation error returned by
// Collection.Validate if the designated constraints aren't met.
type CollectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectionValidationError) ErrorName() string { return "CollectionValidationError" }

// Error satisfies the builtin error interface
func (e CollectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectionValidationError{}

// Validate checks the field values on Object with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Object) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Object with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ObjectMultiError, or nil if none found.
func (m *Object) ValidateAll() error {
	return m.validate(true)
}

func (m *Object) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := ObjectValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ObjectValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ObjectValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ObjectValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ObjectMultiError(errors)
	}
	return nil
}

// ObjectMultiError is an error wrapping multiple validation errors returned by
// Object.ValidateAll() if the designated constraints aren't met.
type ObjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ObjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ObjectMultiError) AllErrors() []error { return m }

// ObjectValidationError is the validation error returned by Object.Validate if
// the designated constraints aren't met.
type ObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ObjectValidationError) ErrorName() string { return "ObjectValidationError" }

// Error satisfies the builtin error interface
func (e ObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ObjectValidationError{}
