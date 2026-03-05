package superstruct

// Package superstruct provides a simple way to validate structs.
//
// This package is designed to be used with the NewStruct function,
// which returns a new Struct instance.
import (
	"reflect"
	"testing"
	"errors"
)

// ErrFieldMissing is returned when a required field is missing from the struct.
var ErrFieldMissing = errors.New("field is missing")

// ErrInvalidType is returned when a field has an invalid type.
var ErrInvalidType = errors.New("invalid type")

// TestStruct_Validate tests the Validate function with a valid struct.
func TestStruct_Validate(t *testing.T) {
	type User struct {
		Name  string
		Email string
	}

	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	s := NewStruct("User")
	s.AddField("Name", reflect.TypeOf((string)(nil)))
	s.AddField("Email", reflect.TypeOf((string)(nil)))

	if err := s.Validate(user); err != nil {
		t.Errorf("Validate() returned error: %w", err)
	}
}

// TestStruct_Validate_MissingField tests the Validate function with a struct that is missing a field.
func TestStruct_Validate_MissingField(t *testing.T) {
	type User struct {
		Name  string
	}

	user := User{
		Name: "John Doe",
	}

	s := NewStruct("User")
	s.AddField("Name", reflect.TypeOf((string)(nil)))
	s.AddField("Email", reflect.TypeOf((string)(nil)))

	if err := s.Validate(user); err == nil {
		t.Errorf("Validate() did not return an error for missing field: %w", err)
	} else if !errors.Is(err, ErrFieldMissing) {
		t.Errorf("Validate() returned an unexpected error: %w", err)
	}
}

// TestStruct_Validate_InvalidType tests the Validate function with a struct that has a field with an invalid type.
func TestStruct_Validate_InvalidType(t *testing.T) {
	type User struct {
		Name  int
	}

	user := User{
		Name: 1,
	}

	s := NewStruct("User")
	s.AddField("Name", reflect.TypeOf((string)(nil)))

	if err := s.Validate(user); err == nil {
		t.Errorf("Validate() did not return an error for invalid type: %w", err)
	} else if !errors.Is(err, ErrInvalidType) {
		t.Errorf("Validate() returned an unexpected error: %w", err)
	}
}

func BenchmarkStruct_Validate(b *testing.B) {
	type User struct {
		Name  string
		Email string
	}

	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	s := NewStruct("User")
	s.AddField("Name", reflect.TypeOf((string)(nil)))
	s.AddField("Email", reflect.TypeOf((string)(nil)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Validate(user)
	}
}
