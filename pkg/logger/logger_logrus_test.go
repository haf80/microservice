package logger

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

var logrusFieldsTests = []struct {
	name	string
	fields	interface{}
	message string
	expectedOk bool
}{
	{"original_fields_and_message", logrus.Fields{"msg": "everything is ok"}, "test msg", true},
	{"original_fields_null_message", logrus.Fields{"msg": "everything is ok"}, "", true},
	{"not_original_fields_and_message", map[string]interface{}{"msg": "everything is ok"}, "test msg", true},
	{"not_original_fields_null_message", map[string]interface{}{"msg": "everything is ok"}, "", true},
	{"invalid_fields_and_message", "invalid param", "", false},
}

func TestConvertToLogrusFields(t *testing.T) {
	for _, tt := range logrusFieldsTests {
		f, err := convertToLogrusFields(tt.fields)
		if tt.expectedOk && err != nil {
			t.Fatalf("%s: fields=%v error=%s", tt.name, f, err.Error())
		}
	}
}

func TestExtractLogrusParams(t *testing.T) {
	var f logrus.Fields
	var s string
	var err error
	for _, tt := range logrusFieldsTests {
		if tt.message != "" {
			f, s, err = extractLogrusParams(tt.fields, tt.message)
		} else {
			f, s, err = extractLogrusParams(tt.fields)
		}
		if tt.message != s {
			t.Fatalf("%s: passed message=%s recieved message=%s", tt.name, tt.message, s)
		}
		if tt.expectedOk && err != nil {
			t.Fatalf("%s: fields=%v message=%s error=%s", tt.name, f, s, err.Error())
		}
	}
}

func TestExtractLogrusParamsInvalidParams(t *testing.T) {
	ts := logrusFieldsTests[0]
	_, _, err := extractLogrusParams(ts.fields, ts.message, "another unhandled param")
	if err == nil {
		t.Fatal("error expected but didn't return an error")
	}
}

func TestMustExtractLogrusParams(t *testing.T) {
	ts := logrusFieldsTests[0]
	f, s := mustExtractLogrusParams(ts.fields, ts.message)
	ft := reflect.TypeOf(f).Kind()
	st := reflect.TypeOf(s).Kind()
	if ft != reflect.TypeOf(logrus.Fields{}).Kind() && st != reflect.String {
		t.Fatalf("logrus.Fields{} and string, got %v and %v", ft, st)
	}
}
