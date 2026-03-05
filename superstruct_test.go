package superstruct

import (
    "testing"
)

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
    s.AddField("Name", reflect.TypeOf(""))
    s.AddField("Email", reflect.TypeOf(""))

    err := s.Validate(user)
    if err != nil {
        t.Errorf("Validate() returned error: %v", err)
    }
}

func TestStruct_Validate_MissingField(t *testing.T) {
    type User struct {
        Name  string
    }

    user := User{
        Name: "John Doe",
    }

    s := NewStruct("User")
    s.AddField("Name", reflect.TypeOf(""))
    s.AddField("Email", reflect.TypeOf(""))

    err := s.Validate(user)
    if err == nil {
        t.Errorf("Validate() did not return an error for missing field")
    }
}

func TestStruct_Validate_InvalidType(t *testing.T) {
    type User struct {
        Name  int
    }

    user := User{
        Name: 1,
    }

    s := NewStruct("User")
    s.AddField("Name", reflect.TypeOf(""))

    err := s.Validate(user)
    if err == nil {
        t.Errorf("Validate() did not return an error for invalid type")
    }
}
