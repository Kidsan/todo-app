package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	todoapp "github.com/kidsan/todo-app"
	"github.com/kidsan/todo-app/mock"
)

func Test_newCreateTodoCommand(t *testing.T) {
	type args struct {
		client  todoapp.TodoClient
		cliArgs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				client:  &mock.TodoClient{},
				cliArgs: []string{"-n", "test todo name", "-d", "test todo description", "-t", "test task"},
			},
			want: `id: 0
name: test todo name
description: test todo description
tasks:
- id: 0
  name: test task
`,
		},
		{
			name: "multiple tasks",
			args: args{
				client:  &mock.TodoClient{},
				cliArgs: []string{"-n", "test todo name", "-d", "test todo description", "-t", "test task", "-t", "test task two"},
			},
			want: `id: 0
name: test todo name
description: test todo description
tasks:
- id: 0
  name: test task
- id: 0
  name: test task two
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := newCreateTodoCommand(tt.args.client)

			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.SetArgs(tt.args.cliArgs)
			cmd.Execute()
			out, err := ioutil.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			if string(out) != tt.want {
				t.Fatalf("expected \"%s\" got \"%s\"", tt.want, string(out))
			}
		})
	}
}
