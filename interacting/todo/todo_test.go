package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/tiagoneves.tia/powerful-cli-app-go/interacting/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %s, got %s", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	l.Add("New Task")
	if l[0].Done {
		t.Errorf("New task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	task := []string{"New Task", "Another Task"}

	for _, name := range task {
		l.Add(name)
	}

	if l[0].Task != task[0] {
		t.Errorf("Expected %s, got %s", task[0], l[0].Task)
	}

	l.Delete(1)

	if len(l) != 1 {
		t.Errorf("Expected 1 item, got %d", len(l))
	}

	if l[0].Task != task[1] {
		t.Errorf("Expected %s, got %s", task[0], l[0].Task)
	}
}

func TestSaveGet(t *testing.T) {

	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %s, got %s", taskName, l1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Expected %s, got %s", l1[0].Task, l2[0].Task)
	}

}
