package zero

import (
	"testing"
)

func TestString(t *testing.T) {
	if String("string_col") != "COALESCE(string_col, '') AS string_col" {
		t.Error("String test failed")
	}
}

func TestInt(t *testing.T) {
	if Int("int_col") != "COALESCE(int_col, 0) AS int_col" {
		t.Error("Int test failed")
	}
}

func TestFloat(t *testing.T) {
	if Float("float_col") != "COALESCE(float_col, 0.0) AS float_col" {
		t.Error("Float test failed")
	}
}

func TestBool(t *testing.T) {
	if Bool("bool_col") != "COALESCE(bool_col, FALSE) AS bool_col" {
		t.Error("Bool test failed")
	}
}

func TestStringAs(t *testing.T) {
	if StringAs("string_col", "a_str") != "COALESCE(string_col, '') AS a_str" {
		t.Error("String test failed")
	}
}

func TestIntAs(t *testing.T) {
	if IntAs("int_col", "an_int") != "COALESCE(int_col, 0) AS an_int" {
		t.Error("IntAs test failed")
	}
}

func TestFloatAs(t *testing.T) {
	if FloatAs("float_col", "a_float") != "COALESCE(float_col, 0.0) AS a_float" {
		t.Error("Float test failed")
	}
}

func TestBoolAs(t *testing.T) {
	if BoolAs("bool_col", "a_bool") != "COALESCE(bool_col, FALSE) AS a_bool" {
		t.Error("Bool test failed")
	}
}

func TestTime(t *testing.T) {
	dbTimeMap := map[string]string{
		"sqlite":   "CAST(COALESCE(time_col, '0001-01-01T00:00:00Z') as text) AS time_col",
		"mysql":    "COALESCE(time_col, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS time_col",
		"postgres": "COALESCE(time_col, (TIMESTAMP WITH TIME ZONE '0001-01-01 00:00:00+00') AT TIME ZONE 'UTC') AS time_col",
	}
	for _, db := range []string{"sqlite", "mysql", "postgres"} {
		zr := New(db)
		if zr.Time("time_col") != dbTimeMap[db] {
			t.Errorf("test Time for database %v failed", db)
		}
	}
}
func TestTimeAs(t *testing.T) {
	dbTimeMap := map[string]string{
		"sqlite":   "CAST(COALESCE(time_col, '0001-01-01T00:00:00Z') as text) AS a_time",
		"mysql":    "COALESCE(time_col, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS a_time",
		"postgres": "COALESCE(time_col, (TIMESTAMP WITH TIME ZONE '0001-01-01 00:00:00+00') AT TIME ZONE 'UTC') AS a_time",
	}
	for _, db := range []string{"sqlite", "mysql", "postgres"} {
		zr := New(db)
		if zr.TimeAs("time_col", "a_time") != dbTimeMap[db] {
			t.Errorf("test TimeAs for database %v failed", db)
		}
	}
}
