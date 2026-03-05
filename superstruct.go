package superstruct

// Package superstruct provides a simple way to define and validate structured data.
// It includes types for defining structures and fields, as well as methods for adding fields and validating data.
//
// Example usage:
//
// struct := NewStruct("my_struct")
// struct.AddField("name", reflect.TypeOf(""))
// err := struct.Validate(MyStruct{Name: "John"})
// if err != nil {
//     fmt.Println(err)
// }

import (
    "context"
    "errors"
    "fmt"
    "reflect"
)

// ErrInvalidData is returned when the data is invalid.
var ErrInvalidData = errors.New("invalid data")

// ErrFieldMissing is returned when a field is missing.
var ErrFieldMissing = errors.New("field is missing")

// ErrFieldTypeMismatch is returned when the field type does not match.
var ErrFieldTypeMismatch = errors.New("field type mismatch")

// Struct represents a structured data definition
// Struct is used to define the structure of data and validate it.
type Struct struct {
    name  string
    fields map[string]*Field
}

// NewStruct creates a new instance of Struct.
//
// NewStruct takes a name as an argument and returns a new Struct instance.
func NewStruct(name string) *Struct {
    return &Struct{
        name:  name,
        fields: make(map[string]*Field),
    }
}

// AddField adds a new field to the structure.
//
// AddField takes a name and a type as arguments and adds a new field to the structure.
// It returns the updated Struct instance.
func (s *Struct) AddField(name string, typ reflect.Type) *Struct {
    if name == "" {
        return nil
    }
    if typ == nil {
        return nil
    }
    s.fields[name] = &Field{
        name: name,
        typ:  typ,
    }
    return s
}

// Validate validates the given data against the structure.
//
// Validate takes a context and data as arguments and checks if the data matches the structure.
// It returns an error if the data is invalid.
func (s *Struct) Validate(ctx context.Context, data interface{}) error {
    if ctx == nil {
        return fmt.Errorf("context is nil: %w", ErrInvalidData)
    }
    if data == nil {
        return fmt.Errorf("data is nil: %w", ErrInvalidData)
    }
    rv := reflect.ValueOf(data)
    if rv.Kind() != reflect.Struct {
        return fmt.Errorf("expected a struct, got %s: %w", rv.Kind(), ErrInvalidData)
    }

    for fieldName, field := range s.fields {
        fv := rv.FieldByName(fieldName)
        if !fv.IsValid() {
            return fmt.Errorf("field %s is missing: %w", fieldName, ErrFieldMissing)
        }
        if fv.Type() != field.typ {
            return fmt.Errorf("field %s has incorrect type, expected %s, got %s: %w", fieldName, field.typ, fv.Type(), ErrFieldTypeMismatch)
        }
    }

    return nil
}

// Field represents a single field in the structure.
//
// Field is used to define a single field in the structure.
type Field struct {
    name string
    typ  reflect.Type
    tag  string
}

// NewField creates a new instance of Field.
//
// NewField takes a name and a type as arguments and returns a new Field instance.
func NewField(name string, typ reflect.Type) *Field {
    if name == "" {
        return nil
    }
    if typ == nil {
        return nil
    }
    return &Field{
        name: name,
        typ:  typ,
    }
}

// SetTag sets the tag for the field.
//
// SetTag takes a tag as an argument and sets it for the field.
// It returns the updated Field instance.
func (f *Field) SetTag(tag string) *Field {
    if tag == "" {
        return nil
    }
    f.tag = tag
    return f
}
