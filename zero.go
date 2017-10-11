package zero

import (
	"fmt"
)

var supportedDatabases = []string{"sqlite", "mysql", "postgres"}

type Nullable struct {
	Database string
}

func New(database string) *Nullable {
	isSupported := false
	for _, db := range supportedDatabases {
		if database == db {
			isSupported = true
		}
	}
	if !isSupported {
		panic("not supported database")
	}
	return &Nullable{Database: database}
}

func (n *Nullable) String(field string) string {
	return StringAs(field, field)
}

func (n *Nullable) Int(field string) string {
	return IntAs(field, field)
}

func (n *Nullable) Float(field string) string {
	return FloatAs(field, field)
}

func (n *Nullable) Bool(field string) string {
	return BoolAs(field, field)
}

func (n *Nullable) Time(field string) string {
	return TimeAs(n.Database, field, field)
}

func (n *Nullable) StringAs(field, as string) string {
	return StringAs(field, as)
}

func (n *Nullable) IntAs(field, as string) string {
	return IntAs(field, as)
}

func (n *Nullable) FloatAs(field, as string) string {
	return FloatAs(field, as)
}

func (n *Nullable) BoolAs(field, as string) string {
	return BoolAs(field, as)
}

func (n *Nullable) TimeAs(field, as string) string {
	return TimeAs(n.Database, field, as)
}

// String return converted SQL chunk for a nullable string typed field,
// here the parameter "field" stands for field name
func String(field string) string {
	return StringAs(field, field)
}

// Int return converted SQL chunk for a nullable int* typed field
func Int(field string) string {
	return FloatAs(field, field)
}

// Float return converted SQL chunk for a nullable float* typed field
func Float(field string) string {
	return FloatAs(field, field)
}

// Bool return converted SQL chunk for a nullable boolean typed field
func Bool(field string) string {
	return BoolAs(field, field)
}

// Time return converted SQL chunk for a nullable datetime typed field,
// and this function is database dependent
func Time(database, field string) string {
	return TimeAs(database, field, field)
}

// StringAs return converted SQL chunk for a nullable string typed field
// here the parameter "field" stands for field name
func StringAs(field, as string) string {
	return fmt.Sprintf("COALESCE(%v, '') AS %v", field, as)
}

// NullIntAs return converted SQL chunk for a nullable int* typed field
func IntAs(field, as string) string {
	return fmt.Sprintf("COALESCE(%v, 0) AS %v", field, as)
}

// FloatAs return converted SQL chunk for a nullable float* typed field
func FloatAs(field, as string) string {
	return fmt.Sprintf("COALESCE(%v, 0.0) AS %v", field, as)
}

// BoolAs return converted SQL chunk for a nullable boolean typed field
func BoolAs(field, as string) string {
	return fmt.Sprintf("COALESCE(%v, FALSE) AS %v", field, as)
}

// TimeAS return converted SQL chunk for a nullable datetime typed field,
// and this function is database dependent
func TimeAs(database, field, as string) string {
	switch database {
	case "mysql":
		return fmt.Sprintf("COALESCE(%v, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS %v", field, as)
	case "postgres":
		return fmt.Sprintf("COALESCE(%v, (TIMESTAMP WITH TIME ZONE '0001-01-01 00:00:00+00') AT TIME ZONE 'UTC') AS %v", field, as)
	}
	return ""
}
