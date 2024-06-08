package gogen

import (
	"reflect"
	"testing"
)

func TestGoStructString(t *testing.T) {
	goStruct := &GoStruct{
		Comment: "This is a test struct",
		Name:    "TestStruct",
		Fields: []*GoField{
			{
				Comment: "This is a test field",
				Name:    "TestField",
				Type:    &GoFieldType{Name: "string"},
				Tags:    []*GoFieldTag{{Key: "json", Values: []string{"test_field"}}},
			},
		},
	}

	expected := `// This is a test struct
type TestStruct struct {
	// This is a test field
	TestField string ` + "`json:\"test_field\"`" + `
}`

	if goStruct.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, goStruct.String())
	}
}

func TestSplitStringByCharLimit(t *testing.T) {
	text := "This is a test string to check the splitting by character limit."
	expected := []string{"This is a test string to check the splitting", "by character limit."}
	result := SplitStringByCharLimit(text, 45)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

func TestGoFieldString(t *testing.T) {
	goField := &GoField{
		Comment: "This is a test field",
		Name:    "TestField",
		Type:    &GoFieldType{Name: "string"},
		Tags:    []*GoFieldTag{{Key: "json", Values: []string{"test_field"}}},
	}

	expected := `// This is a test field
TestField string ` + "`json:\"test_field\"`"

	if goField.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, goField.String())
	}
}

func TestGoFieldTypeString(t *testing.T) {
	tests := []struct {
		fieldType *GoFieldType
		expected  string
	}{
		{&GoFieldType{Name: "string"}, "string"},
		{&GoFieldType{Name: "int", IsPointer: true}, "*int"},
		{&GoFieldType{Name: "float64", IsSlice: true}, "[]float64"},
		{&GoFieldType{Name: "MyStruct", IsPointer: true, IsSlice: true}, "*[]MyStruct"},
	}

	for _, test := range tests {
		if test.fieldType.String() != test.expected {
			t.Errorf("Expected: %s, Got: %s", test.expected, test.fieldType.String())
		}
	}
}

func TestGoFieldTagString(t *testing.T) {
	goFieldTag := &GoFieldTag{
		Key:    "json",
		Values: []string{"test_field"},
	}

	expected := `json:"test_field"`

	if goFieldTag.String() != expected {
		t.Errorf("Expected: %s, Got: %s", expected, goFieldTag.String())
	}
}

func TestParse(t *testing.T) {
	source := `
		// TestStruct is a test struct
		type TestStruct struct {
			// TestField is a test field
			TestField string ` + "`json:\"test_field\"`" + `
		}
	`

	expected := []*GoStruct{
		{
			Comment: "TestStruct is a test struct",
			Name:    "TestStruct",
			Fields: []*GoField{
				{
					Comment: "TestField is a test field",
					Name:    "TestField",
					Type:    &GoFieldType{Name: "string"},
					Tags:    []*GoFieldTag{{Key: "json", Values: []string{"test_field"}}},
				},
			},
		},
	}

	structs, err := Parse([]byte(source))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(structs, expected) {
		t.Errorf("Expected: %+v, Got: %+v", expected, structs)
	}
}
