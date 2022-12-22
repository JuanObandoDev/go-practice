package main

import "fmt"

// taskList struct
type taskList struct {
	tasks []*task // slice of pointers to task struct
}

// add method
func (t *taskList) add(task *task) {
	t.tasks = append(t.tasks, task)
}

// delete method
func (t *taskList) delete(task *task) {
	for i, v := range t.tasks {
		if v == task {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
		}
	}
}

// print method
func (t *taskList) print() {
	fmt.Println("FROM TASK LIST:")
	for _, task := range t.tasks {
		fmt.Printf("name: %s\ndescription: %s\ndone: %t\n", task.name, task.description, task.done)
	}
}

// just done tasks
func (t *taskList) doneTasks() {
	fmt.Println("DONE TASKS:")
	for _, task := range t.tasks {
		if task.done {
			fmt.Printf("name: %s\ndescription: %s\ndone: %t\n", task.name, task.description, task.done)
		}
	}
}

// task struct
type task struct {
	name        string
	description string
	done        bool
}

// markDone method
func (t *task) markDone() {
	t.done = true
}

// updateDescription method
func (t *task) updateDescription(description string) {
	t.description = description
}

// updateName method
func (t *task) updateName(name string) {
	t.name = name
}

// print method
func (t *task) print() {
	fmt.Printf("name: %s\ndescription: %s\ndone: %t\n", t.name, t.description, t.done)
}

func main() {
	// create new task
	myTask := task{name: "My Task", description: "My Task Description", done: false}

	// print task
	myTask.print()

	// mark task as done
	myTask.markDone()

	// print task
	myTask.print()

	// update task description
	myTask.updateDescription("My Task Description Updated")

	// print task
	myTask.print()

	// update task name
	myTask.updateName("My Task Updated")

	// print task
	myTask.print()

	// another way
	myTask2 := task{"My Task 2", "My Task Description 2", false}
	myTask2.print()
	// myTask2.markDone()
	myTask2.updateDescription("My Task Description 2 Updated")
	myTask2.updateName("My Task 2 Updated")
	myTask2.print()

	// create new task list
	myTaskList := taskList{tasks: []*task{}}

	// add task to task list
	myTaskList.add(&myTask)
	myTaskList.add(&myTask2)

	// print task list
	myTaskList.print()

	// print done tasks
	myTaskList.doneTasks()

	// delete task from task list
	myTaskList.delete(&myTask)

	// print task list
	myTaskList.print()

	// creating maps
	taskMap := make(map[string]*taskList)
	taskMap["myTaskList"] = &myTaskList

	myTask3 := task{"My Task 3", "My Task Description 3", false}
	myTask4 := task{"My Task 4", "My Task Description 4", false}
	myTaskList2 := taskList{tasks: []*task{&myTask3, &myTask4}}
	taskMap["myTaskList2"] = &myTaskList2

	taskMap["myTaskList"].print()
	taskMap["myTaskList2"].print()
}
