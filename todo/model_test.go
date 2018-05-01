package todo

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	type testcase struct {
		name  string
		input CreateTodo
		want  string
	}

	testcases := []testcase{}

	testcases = append(testcases, testcase{name: "title and status blank", input: CreateTodo{Title: "", Status: ""}, want: "Todo request is missing status or title"})
	testcases = append(testcases, testcase{name: "title blank", input: CreateTodo{Title: "", Status: "xyz"}, want: "Todo request is missing status or title"})
	testcases = append(testcases, testcase{name: "status blank", input: CreateTodo{Title: "123", Status: ""}, want: "Todo request is missing status or title"})

	testcases = append(testcases, testcase{name: "status invalid", input: CreateTodo{Title: "123", Status: "xyz"}, want: "The provided status is not supported"})
	testcases = append(testcases, testcase{name: "status invalid (lowercase)", input: CreateTodo{Title: "123", Status: "new"}, want: "The provided status is not supported"})

	testcases = append(testcases, testcase{name: "all clean", input: CreateTodo{Title: "123", Status: "New"}, want: ""})

	for _, tc := range testcases {
		got := isValid(tc.input)
		if got != tc.want {
			t.Errorf(`test case "%s" failed: wanted %s but got %s`, tc.name, tc.want, got)
		}

	}
}
