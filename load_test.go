package goconfig_test

import (
	"fmt"
	"profiralex/goconfig"
	"testing"
)

const (
	testKeyMissing = "testKeyMissing"

	testStringKey   = "testStringKey"
	testStringValue = "testStringValue"

	testStringSliceKey = "testStringSliceKey"

	testBoolKey1       = "testBoolKey1"
	testBoolValue1     = "true"
	testBoolValueReal1 = true

	testBoolKey2       = "testBoolKey2"
	testBoolValue2     = "false"
	testBoolValueReal2 = false

	testBoolKeyInvalid   = "testBoolKeyInvalid"
	testBoolValueInvalid = "bas"

	testIntKey       = "testIntKey"
	testIntValue     = "5"
	testIntValueReal = int(5)

	testInt8Key       = "testInt8Key"
	testInt8Value     = "8"
	testInt8ValueReal = int8(8)

	testInt16Key       = "testInt16Key"
	testInt16Value     = "13"
	testInt16ValueReal = int16(13)

	testInt32Key       = "testInt32Key"
	testInt32Value     = "21"
	testInt32ValueReal = int32(21)

	testInt64Key       = "testInt64Key"
	testInt64Value     = "34"
	testInt64ValueReal = int64(34)

	testIntKeyInvalid   = "testIntKeyInvalid"
	testIntValueInvalid = "5a"

	testUintKey       = "testUintKey"
	testUintValue     = "55"
	testUintValueReal = uint(55)

	testUint8Key       = "testUint8Key"
	testUint8Value     = "89"
	testUint8ValueReal = uint8(89)

	testUint16Key       = "testUint16Key"
	testUint16Value     = "144"
	testUint16ValueReal = uint16(144)

	testUint32Key       = "testUint32Key"
	testUint32Value     = "233"
	testUint32ValueReal = uint32(233)

	testUint64Key       = "testUint64Key"
	testUint64Value     = "377"
	testUint64ValueReal = uint64(377)

	testUintKeyInvalid   = "testUintKeyInvalid"
	testUintValueInvalid = "55lkd"

	testFloat32Key       = "testFloat32Key"
	testFloat32Value     = "231243.23"
	testFloat32ValueReal = float32(231243.23)

	testFloat64Key       = "testFloat64Key"
	testFloat64Value     = "12223.53"
	testFloat64ValueReal = float64(12223.53)

	testFloat64KeyInvalid   = "testFloat64KeyInvalid"
	testFloat64ValueInvalid = "12223.53alshf"
)

var (
	testStringSliceValue     = "tesStringSliceValue1,tesStringSliceValue2,tesStringSliceValue3"
	testStringSliceValueReal = []string{"tesStringSliceValue1", "tesStringSliceValue2", "tesStringSliceValue3"}

	mockEnv = map[string]string{
		testStringKey:      testStringValue,
		testStringSliceKey: testStringSliceValue,

		testBoolKey1:       testBoolValue1,
		testBoolKey2:       testBoolValue2,
		testBoolKeyInvalid: testBoolValueInvalid,

		testIntKey:        testIntValue,
		testInt8Key:       testInt8Value,
		testInt16Key:      testInt16Value,
		testInt32Key:      testInt32Value,
		testInt64Key:      testInt64Value,
		testIntKeyInvalid: testIntValueInvalid,

		testUintKey:        testUintValue,
		testUint8Key:       testUint8Value,
		testUint16Key:      testUint16Value,
		testUint32Key:      testUint32Value,
		testUint64Key:      testUint64Value,
		testUintKeyInvalid: testUintValueInvalid,

		testFloat32Key:        testFloat32Value,
		testFloat64Key:        testFloat64Value,
		testFloat64KeyInvalid: testFloat64ValueInvalid,
	}
	mockProvider = mockProviderStruct{}
)

type mockProviderStruct struct {
}

// Lookup looks up for value in the environment
func (p *mockProviderStruct) Lookup(key string) (string, error) {
	value, ok := mockEnv[key]
	if !ok {
		return "", fmt.Errorf("Value not found")
	}

	return value, nil
}

func TestSuccessStringMissingNotStrict(t *testing.T) {
	type testStruct struct {
		StringField string `cfg:"testKeyMissing"`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, false)

	if len(v.StringField) != 0 {
		t.Errorf("Unexpected value assigned %s", v.StringField)
	}

	if err != nil {
		t.Errorf("Unexpected error recived %w", err)
	}
}

func TestFailStringMissingStrict(t *testing.T) {
	type testStruct struct {
		StringField string `cfg:"testKeyMissing"`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, true)

	if len(v.StringField) != 0 {
		t.Errorf("Unexpected value assigned %s", v.StringField)
	}

	if err == nil {
		t.Errorf("Expected error not recived")
	}
}

func TestSuccessStringMissingStrictDefaultValue(t *testing.T) {
	defaultValue := "abcdef"

	type testStruct struct {
		StringField string `cfg:"testKeyMissing" cfg-default:"abcdef"`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, true)

	if v.StringField != defaultValue {
		t.Errorf("Expected value %s got %s", defaultValue, v.StringField)
	}

	if err != nil {
		t.Errorf("Unexpected error received")
	}
}

func TestSuccessStringNoStrictNoDefault(t *testing.T) {
	type testStruct struct {
		StringField string `cfg:"testStringKey"`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, false)

	if v.StringField != testStringValue {
		t.Errorf("Expected value %s got %s", testStringValue, v.StringField)
	}

	if err != nil {
		t.Errorf("Unexpected error received")
	}
}

func TestSuccessStringStrictNoDefault(t *testing.T) {
	type testStruct struct {
		StringField string `cfg:"testStringKey"`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, true)

	if v.StringField != testStringValue {
		t.Errorf("Expected value %s got %s", testStringValue, v.StringField)
	}

	if err != nil {
		t.Errorf("Unexpected error received")
	}
}

func TestSuccessStringNoStrictDefault(t *testing.T) {
	type testStruct struct {
		StringField string `cfg:"testStringKey" cfg-default:"abcd"`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, false)

	if v.StringField != testStringValue {
		t.Errorf("Expected value %s got %s", testStringValue, v.StringField)
	}

	if err != nil {
		t.Errorf("Unexpected error received")
	}
}

func TestSuccessStringStrictDefault(t *testing.T) {
	type testStruct struct {
		StringField string `cfg:"testStringKey" cfg-default:"abcd"`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, true)

	if v.StringField != testStringValue {
		t.Errorf("Expected value %s got %s", testStringValue, v.StringField)
	}

	if err != nil {
		t.Errorf("Unexpected error received")
	}
}

func TestSuccessStringSliceMissingDefaultEmpty(t *testing.T) {
	type testStruct struct {
		StringSliceField []string `cfg:"testKeyMissing" cfg-default:""`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, true)

	if err != nil {
		t.Errorf("Unexpected error received")
	}

	if len(v.StringSliceField) != 0 {
		t.Errorf("Expected slice len %d got %d", 0, len(testStringSliceValueReal))
	}
}

func TestSuccessStringSliceMissingDefaultNonEmpty(t *testing.T) {
	type testStruct struct {
		StringSliceField []string `cfg:"testKeyMissing" cfg-default:"testValue"`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, true)

	if err != nil {
		t.Errorf("Unexpected error received")
	}

	if len(v.StringSliceField) != 1 {
		t.Errorf("Expected slice len %d got %d", 1, len(testStringSliceValueReal))
	}

	if v.StringSliceField[0] != "testValue" {
		t.Errorf("Expected value at index %d %s got %s", 0, "testValue", v.StringSliceField[0])
	}
}

func TestSuccessStringSlice(t *testing.T) {
	type testStruct struct {
		StringSliceField []string `cfg:"testStringSliceKey"`
	}
	v := testStruct{}

	err := goconfig.Load(&v, &mockProvider, true)

	if err != nil {
		t.Errorf("Unexpected error received")
	}

	if len(v.StringSliceField) != len(testStringSliceValueReal) {
		t.Errorf("Expected slice len %d got %d", len(v.StringSliceField), len(testStringSliceValueReal))
	}

	for i, v := range v.StringSliceField {
		if v != testStringSliceValueReal[i] {
			t.Errorf("Expected value at index %d %s got %s", i, testStringSliceValueReal[i], v)
		}
	}
}
