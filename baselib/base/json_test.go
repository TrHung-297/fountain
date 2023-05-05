package base

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestEmpty(t *testing.T) {
	mapData := make(map[string]interface{})

	jsonString := JSONDebugDataString(mapData)
	if jsonString == "" {
		t.Errorf("JSON string is empty while expected %q", string(emptyBytes))
	}

	if jsonString == "{}" {
		t.Log("Oke")
	}
}

func TestNotEmpty(t *testing.T) {
	stringExpected := `{"Key":"Value"}`
	mapData := make(map[string]interface{})

	err := json.Unmarshal([]byte(stringExpected), &mapData)
	if err != nil {
		t.Errorf("Can not create data test")
	}

	jsonString := JSONDebugDataString(mapData)
	if jsonString == "" {
		t.Errorf("JSON string is empty while expected %q", stringExpected)
	}

	if jsonString != stringExpected {
		t.Errorf("JSON string is %q while expected %q", jsonString, stringExpected)
	}
}

type structTest struct {
	Data  string `json:"data,omitempty"`
	Value string `json:"value,omitempty"`
}

func TestStruct(t *testing.T) {
	structTest := &structTest{
		Data:  fmt.Sprintf("%d", rand.Int63()),
		Value: time.Now().String(),
	}

	stringExpected, err := json.Marshal(structTest)
	if err != nil {
		t.Errorf("Can not create expect result for data test")
	}

	jsonData := JSONDebugData(structTest)
	if len(jsonData) == 0 {
		t.Errorf("JSON string is empty while expected result have len: %d", len(stringExpected))
	}

	if len(jsonData) != len(stringExpected) {
		t.Errorf("JSON data is not equal to expected result: %d vs %d", len(jsonData), len(stringExpected))
	}

	jsonString := JSONDebugDataString(structTest)
	if jsonString == "" {
		t.Errorf("JSON string is empty while expected %q", stringExpected)
	}

	if jsonString != string(stringExpected) {
		t.Errorf("JSON string is %q while expected %q", jsonString, stringExpected)
	}
}

func BenchmarkParser(b *testing.B) {
	structTest := &structTest{
		Data:  fmt.Sprintf("%d", rand.Int63()),
		Value: time.Now().String(),
	}

	for i := 0; i < b.N; i++ {
		JSONDebugDataString(structTest)
	}
}
