package qlite_test

import (
	"fmt"
	q "nuclyk/qlite"
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	cases := []testCase{
		{q.NewQuery().Select().String(), "SELECT *"},
		{q.NewQuery().Select().Distinct().String(), "SELECT DISTINCT *"},
		{q.NewQuery().Select("id").String(), "SELECT id"},
		{q.NewQuery().Select("id").Distinct().String(), "SELECT DISTINCT id"},
		{q.NewQuery().Select("id", "name", "age").String(), "SELECT id, name, age"},
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
		{q.NewQuery().Select().From("users").String(), "SELECT * FROM users"},
		{q.NewQuery().Select("id").From("users").String(), "SELECT id FROM users"},
		{q.NewQuery().Select("id", "name", "age").From("users").String(), "SELECT id, name, age FROM users"},
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
		{q.NewQuery().Select().From("users").Where("id = ?", "1").String(),
			"SELECT * FROM users WHERE id = ?"},
		{q.NewQuery().Select().From("users").Where("id = ?", "1").Where("name = ?", "John").String(),
			"SELECT * FROM users WHERE id = ? AND name = ?"},
		{q.NewQuery().Select().From("users").Where("id = ?", "1").OrWhere("name = ?", "John").String(),
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
		{q.NewQuery().Select().From("users").GroupBy("users").String(),
			"SELECT * FROM users GROUP BY users"},
		{q.NewQuery().Select().From("users").GroupBy("users, name, age").String(),
			"SELECT * FROM users GROUP BY users, name, age"},
		{q.NewQuery().Select().From("users").Where("id = ?", "1").GroupBy("users").String(),
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
		{q.NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").String(),
			"SELECT * FROM users GROUP BY users HAVING age > ?"},
		{q.NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").
			Having("phone = ?", "0").String(),
			"SELECT * FROM users GROUP BY users HAVING age > ? AND phone = ?"},
		{q.NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").
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
		{q.NewQuery().Select().From("users").OrderBy("name", q.ASC).String(),
			"SELECT * FROM users ORDER BY name ASC"},
		{q.NewQuery().Select().From("users").Where("id = ?", "1").OrderBy("name", q.ASC).String(),
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
		{q.NewQuery().Select().From("users").OrderBy("name", q.ASC).Limit(10).String(),
			"SELECT * FROM users ORDER BY name ASC LIMIT 10"},
		{q.NewQuery().Select().From("users").Limit(10).String(),
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
		{q.NewQuery().Select().From("users").Where("id = ?", "1").GetValues(),
			[]any{"1"}},
		{q.NewQuery().Select().From("users").Where("id = ?", "1").Where("name = ?", "John").GetValues(),
			[]any{"1", "John"}},
		{q.NewQuery().Select().From("users").Where("id = ?", "1").OrWhere("name = ?", "John").GetValues(),
			[]any{"1", "John"}},
		{q.NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").GetValues(),
			[]any{"20"}},
		{q.NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").
			Having("phone = ?", "0").GetValues(),
			[]any{"20", "0"}},
		{q.NewQuery().Select().From("users").GroupBy("users").Having("age > ?", "20").
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
	q := q.NewQuery().Select().From("users").Where("id = ?", "1")
	fmt.Println(q.String())

	// Output: SELECT * FROM users WHERE id = ?
}
