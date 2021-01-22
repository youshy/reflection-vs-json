package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func BenchmarkStructToMap(b *testing.B) {
	t := Test{
		String: "value string",
		Number: 33245,
	}
	_ = structToMap(t)
}

func BenchmarkJsonToMap(b *testing.B) {
	t := Test{
		String: "value string",
		Number: 33245,
	}

	_ = jsonToMap(t)
}

func TestStructToMap(t *testing.T) {
	s := Test{
		String: "value string",
		Number: 33245,
	}

	v := map[string]interface{}{
		"this-is-string-field": "value string",
		"this-is-number-field": 33245,
	}

	r := structToMap(s)

	fmt.Printf("Type test for %+v\n%T\t%T\n", r, r["this-is-string-field"], r["this-is-number-field"])
	assert.Equal(t, v, r)
}

func TestJsonToMap(t *testing.T) {
	s := Test{
		String: "value string",
		Number: 33245,
	}

	v := map[string]interface{}{
		"this-is-string-field": "value string",
		"this-is-number-field": 33245,
	}

	r := jsonToMap(s)

	fmt.Printf("Type test for %+v\n%T\t%T\n", r, r["this-is-string-field"], r["this-is-number-field"])
	// NotEqual because marshal/unmarshal combo yields float64 on int field
	assert.NotEqual(t, v, r)
}
