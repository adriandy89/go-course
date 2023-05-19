package main

import "fmt"

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) toPrint() {
	fmt.Printf("Name: %s\nDescription: %s\nCompleted: %t\n", t.name, t.description, t.completed)
}

func (t *Task) toggleCompleted() {
	t.completed = !t.completed
}

func main() {
	task1 := Task{
		name:        "Name 1",
		description: "Description task 1",
		completed:   false,
	}
	task2 := Task{
		name:        "Name 2",
		description: "Description task 2",
		completed:   true,
	}
	task1.toPrint()
	task2.toPrint()
	task1.toggleCompleted()
	task1.toPrint()

	tl1 := TaskList{}
	tl1.appendTask(&task1)
	tl1.appendTask(&task2)
	tl1.appendTask(&task1)
	tl1.appendTask(&task2)
	fmt.Println(tl1.tasks[0])
	fmt.Println(len(tl1.tasks))
	tl1.removeTask(0)
	fmt.Println(tl1.tasks[0])
	fmt.Println(len(tl1.tasks))

	for _, v := range tl1.tasks {
		v.toPrint()
	}

}
