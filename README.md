# Todo App

This is a basic todo tracking application that allows a user manage todos. A todo has a name and a description and can have tasks associated with it. Each task just has a name property.

## Usage

Using the todoctl command we can manage resources in the api. Configuration can be set via env vars or via an app.env config file.

```shell
$ make build
$ ./todoctl get todos

$ ./todoctl create todo -n "this is my new todo" -d "stuff I need to do today" -t "go to the shop" -t "cook some dinner"
id: 1
name: this is my new todo
description: generic stuff I need to do today
tasks:
- id: 1
  name: go to the shop
- id: 2
  name: cook some dinner

$ ./todoctl update todo 1 -t "add an extra task"
id: 1
name: this is my new todo
description: generic stuff I need to do today
tasks:
- id: 1
  name: go to the shop
- id: 2
  name: cook some dinner
- id: 3
  name: add an extra task

$ ./todoctl update task 3 -n "DONE: add an extra task"
id: 3
name: 'DONE: add an extra task'

$ ./todoctl get todo
id: 1
name: this is my new todo
description: generic stuff I need to do today
tasks:
- id: 1
  name: go to the shop
- id: 2
  name: cook some dinner
- id: 3
  name: Done: add an extra task

$ ./todoctl delete task 3
deleted

$ ./todoctl delete todo 1
deleted
```