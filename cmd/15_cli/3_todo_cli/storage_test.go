package main

import (
	"os"
	"strings"
	"testing"
)

type TestData struct {
	Name string
	Age  int
}

/*
The test suite includes:
- Test cases for adding items to storage and verifying JSON content
- Test cases for loading items from storage and validating data integrity
- Error handling tests for non-existent files and invalid paths
- Proper cleanup using temporary files to avoid affecting actual storage
- Type safety verification using generics with a TestData struct

*/

func TestStorage_Save(t *testing.T) {
	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "storage_test_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Initialize storage with test data
	storage := NewStorage[TestData](tmpFile.Name())
	testItem := TestData{Name: "John", Age: 30}

	// Test adding item
	err = storage.Save(testItem)
	if err != nil {
		t.Errorf("Failed to add item: %v", err)
	}

	// Verify file content
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	// Normalize JSON by removing whitespace
	expectedJSON := `{"Name":"John","Age":30}`
	actualJSON := string(content)
	expectedJSON = strings.ReplaceAll(strings.ReplaceAll(expectedJSON, " ", ""), "\n", "")
	actualJSON = strings.ReplaceAll(strings.ReplaceAll(actualJSON, " ", ""), "\n", "")
	if actualJSON != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(content))
	}
}

func TestStorage_Load(t *testing.T) {
	// Create a temporary file with test data
	tmpFile, err := os.CreateTemp("", "storage_test_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// Write test data to file
	testJSON := `{"Name":"Jane","Age":25}`
	if err := os.WriteFile(tmpFile.Name(), []byte(testJSON), 0644); err != nil {
		t.Fatal(err)
	}

	// Initialize storage and load data
	storage := NewStorage[TestData](tmpFile.Name())
	var loadedData TestData
	err = storage.Load(&loadedData)
	if err != nil {
		t.Errorf("Failed to load data: %v", err)
	}

	// Verify loaded data
	expectedData := TestData{Name: "Jane", Age: 25}
	if loadedData != expectedData {
		t.Errorf("Expected %+v, got %+v", expectedData, loadedData)
	}
}

func TestStorage_LoadNonExistentFile(t *testing.T) {
	// Test loading from non-existent file
	storage := NewStorage[TestData]("non_existent.json")
	var data TestData
	err := storage.Load(&data)
	if err == nil {
		t.Error("Expected error when loading from non-existent file")
	}
}

func TestStorage_SaveInvalidPath(t *testing.T) {
	// Test adding to invalid path
	storage := NewStorage[TestData]("/invalid/path/file.json")
	testItem := TestData{Name: "John", Age: 30}

	err := storage.Save(testItem)
	if err == nil {
		t.Error("Expected error when adding to invalid path")
	}
}
