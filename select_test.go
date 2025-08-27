package main

import "testing"

func TestSelect(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	cases := []testCase{
		{NewQuery().Select().Build(), "SELECT *"},
		{NewQuery().Select("id").Build(), "SELECT id"},
		{NewQuery().Select("id", "name", "age").Build(), "SELECT id, name, age"},
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
		{NewQuery().Select().From("users").Build(), "SELECT * FROM users"},
		{NewQuery().Select("id").From("users").Build(), "SELECT id FROM users"},
		{NewQuery().Select("id", "name", "age").From("users").Build(), "SELECT id, name, age FROM users"},
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
		{NewQuery().Select().From("users").Where("id = ?", "1").Build(),
			"SELECT * FROM users WHERE id = ?"},
		{NewQuery().Select().From("users").Where("id = ?", "1").Where("name = ?", "John").Build(),
			"SELECT * FROM users WHERE id = ? AND name = ?"},
		{NewQuery().Select().From("users").Where("id = ?", "1").Or("name = ?", "John").Build(),
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
