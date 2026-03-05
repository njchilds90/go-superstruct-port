package superstruct

import (
    "fmt"
    "reflect"
)

// Struct represents a structured data definition
type Struct struct {
    name  string
    fields map[string]*Field
}

// NewStruct creates a new instance of Struct
func NewStruct(name string) *Struct {
    return &Struct{
        name:  name,
        fields: make(map[string]*Field),
    }
}

// AddField adds a new field to the structure
func (s *Struct) AddField(name string, typ reflect.Type) *Struct {
    s.fields[name] = &Field{
        name: name,
        typ:  typ,
    }
    return s
}

// Validate validates the given data against the structure
func (s *Struct) Validate(data interface{}) error {
    rv := reflect.ValueOf(data)
    if rv.Kind() != reflect.Struct {
        return fmt.Errorf("expected a struct, got %s", rv.Kind())
    }

    for fieldName, field := range s.fields {
        fv := rv.FieldByName(fieldName)
        if !fv.IsValid() {
            return fmt.Errorf("field %s is missing", fieldName)
        }
        if fv.Type() != field.typ {
            return fmt.Errorf("field %s has incorrect type, expected %s, got %s", fieldName, field.typ, fv.Type())
        }
    }

    return nil
}

// Field represents a single field in the structure
type Field struct {
    name string
    typ  reflect.Type
    tag  string
}

// NewField creates a new instance of Field
func NewField(name string, typ reflect.Type) *Field {
    return &Field{
        name: name,
        typ:  typ,
    }
}

// SetTag sets the tag for the field
func (f *Field) SetTag(tag string) *Field {
    f.tag = tag
    return f
}
