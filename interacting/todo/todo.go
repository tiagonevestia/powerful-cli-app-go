package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// item struct represents a ToDo item
type item struct {
	Task        string
	Done        bool
	CreateAt    time.Time
	CompletedAt time.Time
}

type List []item

// Add creates a new todo item and appends ir to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreateAt:    time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete method marks a ToDo item as completed by setting
// the Done field to true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	i--
	ls[i].Done = true
	ls[i].CompletedAt = time.Now()

	return nil
}

// Delete method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	i--
	*l = append(ls[:i], ls[i+1:]...)

	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

// String prints out a formatted list
// Implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X "
		}
		// Adjust the item number k to print numbers stating from 1 instead of 0
		formatted += fmt.Sprintf("%d)%s %s\n", k+1, prefix, t.Task)
	}
	return formatted
}

func (l *List) Pretty() string {
	formatted := ""

	for k, t := range *l {
		prefix := "No"
		if t.Done {
			prefix = "Yes!"
		}
		// Adjust the item number k to print numbers stating from 1 instead of 0
		formatted += fmt.Sprintf("# ID: %d\n# Is completed: %s\n# Description: %s\n# Created in: %s\n\n", k+1, prefix, t.Task, t.CreateAt.Format("2006-01-02 15:04:05"))
	}
	return formatted
}

func (l *List) Completed() string {
	formatted := ""

	for k, t := range *l {
		if t.Done {
			// Adjust the item number k to print numbers stating from 1 instead of 0
			formatted += fmt.Sprintf("%d) %s\n", k+1, t.Task)
		}
	}
	return formatted
}
