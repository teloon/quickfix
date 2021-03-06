package quickfix

import (
	"testing"
)

func TestFieldMap_Clear(t *testing.T) {
	var fMap FieldMap
	fMap.init()

	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))

	fMap.Clear()

	if fMap.Has(1) || fMap.Has(2) {
		t.Error("All fields should be cleared")
	}
}

func TestFieldMap_SetAndGet(t *testing.T) {
	var fMap FieldMap
	fMap.init()

	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))

	var testCases = []struct {
		tag         Tag
		expectErr   bool
		expectValue string
	}{
		{tag: 1, expectValue: "hello"},
		{tag: 2, expectValue: "world"},
		{tag: 44, expectErr: true},
	}

	for _, tc := range testCases {
		var testField FIXString
		err := fMap.GetField(tc.tag, &testField)

		if tc.expectErr {
			if err == nil {
				t.Error("Expected Error")
			}
			continue
		}

		if err != nil {
			t.Error("Unexpected Error", err)
		}

		if string(testField) != tc.expectValue {
			t.Errorf("Expected %v got %v", tc.expectValue, testField)
		}
	}
}

func TestFieldMap_Length(t *testing.T) {
	var fMap FieldMap
	fMap.init()
	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))
	fMap.SetField(8, FIXString("FIX.4.4"))
	fMap.SetField(9, FIXInt(100))
	fMap.SetField(10, FIXString("100"))

	if fMap.length() != 16 {
		t.Error("Length should include all fields but beginString, bodyLength, and checkSum- got ", fMap.length())
	}
}

func TestFieldMap_Total(t *testing.T) {

	var fMap FieldMap
	fMap.init()
	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))
	fMap.SetField(8, FIXString("FIX.4.4"))
	fMap.SetField(Tag(9), FIXInt(100))
	fMap.SetField(10, FIXString("100"))

	if fMap.total() != 2116 {
		t.Error("Total should includes all fields but checkSum- got ", fMap.total())
	}
}

func TestFieldMap_TypedSetAndGet(t *testing.T) {
	var fMap FieldMap
	fMap.init()

	fMap.SetString(1, "hello")
	fMap.SetInt(2, 256)

	s, err := fMap.GetString(1)
	if err != nil {
		t.Error("Unexpected Error", err)
	} else if s != "hello" {
		t.Errorf("Expected %v got %v", "hello", s)
	}

	i, err := fMap.GetInt(2)
	if err != nil {
		t.Error("Unexpected Error", err)
	} else if i != 256 {
		t.Errorf("Expected %v got %v", 256, i)
	}

	_, err = fMap.GetInt(1)
	if err == nil {
		t.Error("Type mismatch should occur error but nil")
	}

	s, err = fMap.GetString(2)
	if err != nil {
		t.Error("Type mismatch should occur error but nil")
	} else if s != "256" {
		t.Errorf("Expected %v got %v", "256", s)
	}
}
