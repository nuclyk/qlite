package qlite

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	cases := []testCase{
		{NewQuery().Select().String(), "SELECT *"},
		{NewQuery().Select().Distinct().String(), "SELECT DISTINCT *"},
		{NewQuery().Select("id").String(), "SELECT id"},
		{NewQuery().Select("id").Distinct().String(), "SELECT DISTINCT id"},
		{NewQuery().Select("id", "name", "age").String(), "SELECT id, name, age"},
	}

	for _, test := range cases {
		result := test.input
		if result != test.expected {
			t.Errorf(`----
Inputs: %v
Expecting: %v
Actual: %v
Fail
---`, test.input, test.expected, result)
		}
	}
}

func TestFrom(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	cases := []testCase{
		{NewQuery().Select().From("users").String(), "SELECT * FROM users"},
		{NewQuery().Select("id").From("users").String(), "SELECT id FROM users"},
		{NewQuery().Select("id", "name", "age").From("users").String(), "SELECT id, name, age FROM users"},
	}

	for _, test := range cases {
		result := test.input
		if result != test.expected {
			t.Errorf(`----
Inputs: %v
Expecting: %v
Actual: %v
Fail
---`, test.input, test.expected, result)
		}
	}
}

func TestWhere(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	cases := []testCase{
		{NewQuery().Select().From("users").Where("id = ?", "1").String(),
			"SELECT * FROM users WHERE id = ?"},
		{NewQuery().Select().From("users").Where("id = ?", "1").Where("name = ?", "John").String(),
			"SELECT * FROM users WHERE id = ? AND name = ?"},
		{NewQuery().Select().From("users").Where("id = ?", "1").OrWhere("name = ?", "John").String(),
			"SELECT * FROM users WHERE id = ? OR name = ?"},
	}

	for _, test := range cases {
		result := test.input
		if result != test.expected {
			t.Errorf(`----
Inputs: %v
Expecting: %v
Actual: %v
Fail
---`, test.input, test.expected, result)
		}
	}
}

func TestGroupBy(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	cases := []testCase{
		{NewQuery().Select().From("users").GroupBy("users").String(),
			"SELECT * FROM users GROUP BY users"},
		{NewQuery().Select().From("users").GroupBy("users, name, age").String(),
			"SELECT * FROM users GROUP BY users, name, age"},
		{NewQuery().Select().From("users").Where("id = ?", "1").GroupBy("users").String(),
			"SELECT * FROM users WHERE id = ? GROUP BY users"},
	}

	for _, test := range cases {
		result := test.input
		if result != test.expected {
			t.Errorf(`----
Inputs: %v
Expecting: %v
Actual: %v
Fail
---`, test.input, test.expected, result)
		}
	}
}

func TestHaving(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	cases := []testCase{
		{NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").String(),
			"SELECT * FROM users GROUP BY users HAVING age > ?"},
		{NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").
			Having("phone = ?", "0").String(),
			"SELECT * FROM users GROUP BY users HAVING age > ? AND phone = ?"},
		{NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").
			OrHaving("phone = ?", "0").String(),
			"SELECT * FROM users GROUP BY users HAVING age > ? OR phone = ?"},
	}

	for _, test := range cases {
		result := test.input
		if result != test.expected {
			t.Errorf(`----
Inputs: %v
Expecting: %v
Actual: %v
Fail
---`, test.input, test.expected, result)
		}
	}
}

func TestOrderBy(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	cases := []testCase{
		{NewQuery().Select().From("users").OrderBy("name", ASC).String(),
			"SELECT * FROM users ORDER BY name ASC"},
		{NewQuery().Select().From("users").Where("id = ?", "1").OrderBy("name", ASC).String(),
			"SELECT * FROM users WHERE id = ? ORDER BY name ASC"},
	}

	for _, test := range cases {
		result := test.input
		if result != test.expected {
			t.Errorf(`----
Inputs: %v
Expecting: %v
Actual: %v
Fail
---`, test.input, test.expected, result)
		}
	}
}

func TestLimit(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	cases := []testCase{
		{NewQuery().Select().From("users").OrderBy("name", ASC).Limit(10).String(),
			"SELECT * FROM users ORDER BY name ASC LIMIT 10"},
		{NewQuery().Select().From("users").Limit(10).String(),
			"SELECT * FROM users LIMIT 10"},
	}

	for _, test := range cases {
		result := test.input
		if result != test.expected {
			t.Errorf(`----
Inputs: %v
Expecting: %v
Actual: %v
Fail
---`, test.input, test.expected, result)
		}
	}
}

func TestGetValues(t *testing.T) {
	type testCase struct {
		input    []any
		expected []any
	}

	cases := []testCase{
		{NewQuery().Select().From("users").Where("id = ?", "1").GetValues(),
			[]any{"1"}},
		{NewQuery().Select().From("users").Where("id = ?", "1").Where("name = ?", "John").GetValues(),
			[]any{"1", "John"}},
		{NewQuery().Select().From("users").Where("id = ?", "1").OrWhere("name = ?", "John").GetValues(),
			[]any{"1", "John"}},
		{NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").GetValues(),
			[]any{"20"}},
		{NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").
			Having("phone = ?", "0").GetValues(),
			[]any{"20", "0"}},
		{NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").
			OrHaving("phone = ?", "0").GetValues(),
			[]any{"20", "0"}},
	}

	for _, test := range cases {
		result := test.input
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf(`----
Inputs: %v
Expecting: %v
Actual: %v
Fail
---`, test.input, test.expected, result)
		}
	}
}

// Example for Where function.
func ExampleQuery_Where() {
	q := NewQuery().Select().From("users").Where("id = ?", "1")
	fmt.Println(q.String())

	// Output: "SELECT * FROM users WHERE id = ?"
}
