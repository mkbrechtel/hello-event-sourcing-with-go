package event

import (
	"testing"

	"github.com/gofrs/uuid/v5"
)

func TestNewEvent(t *testing.T) {
	type TestDetails struct {
		Name string
		Age  int
	}

	details := TestDetails{
		Name: "Test",
		Age:  30,
	}

	event := NewEvent(details, "test", 1)

	if event.ID == uuid.Nil {
		t.Error("Expected non-nil UUID")
	}

	if event.Type != "TestDetails" {
		t.Errorf("Expected type to be 'TestDetails', got '%s'", event.Type)
	}

	if event.Details.Name != "Test" {
		t.Errorf("Expected Details.Name to be 'Test', got '%s'", event.Details.Name)
	}

	if event.Details.Age != 30 {
		t.Errorf("Expected Details.Age to be 30, got %d", event.Details.Age)
	}
}

func TestGetType(t *testing.T) {
	type TestStruct struct{}

	// Test with struct
	typeName := getType(TestStruct{})
	if typeName != "TestStruct" {
		t.Errorf("Expected 'TestStruct', got '%s'", typeName)
	}

	// Test with pointer
	typeNamePtr := getType(&TestStruct{})
	if typeNamePtr != "TestStruct" {
		t.Errorf("Expected 'TestStruct', got '%s'", typeNamePtr)
	}
}
