package common

import (
	"database/sql"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type (
	// Route holds all the infomation needed by a route
	Route struct {
		Method     string
		Path       string
		Handler    echo.HandlerFunc
		Middleware []echo.MiddlewareFunc
	}

	// Controller is an interface that can provide routes
	Controller interface {
		Routes() []Route
	}
)

// JsonNullString is a nullable string that can be JSON Marshalled
type JsonNullString struct {
	sql.NullString
}

type JsonNullInt32 struct {
	sql.NullInt32
}

type JsonNullFloat64 struct {
	sql.NullFloat64
}

type JsonNullBool struct {
	sql.NullBool
}

func (v JsonNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal("")
	}
}

func (v *JsonNullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

func (v JsonNullInt32) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int32)
	} else {
		return json.Marshal("")
	}
}

func (v *JsonNullInt32) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int32
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int32 = *x
	} else {
		v.Valid = false
	}
	return nil
}

func (v JsonNullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	} else {
		return json.Marshal("")
	}
}

func (v *JsonNullFloat64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *float64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Float64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

func (v JsonNullBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	} else {
		return json.Marshal("")
	}
}

func (v *JsonNullBool) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *bool
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Bool = *x
	} else {
		v.Valid = false
	}
	return nil
}
