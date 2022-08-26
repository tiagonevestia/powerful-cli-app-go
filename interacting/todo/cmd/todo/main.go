package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tiagoneves.tia/powerful-cli-app-go/interacting/todo"
)

var todoFileName = ".todo.json"

func main() {

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// Parsing command line flags
	add := flag.Bool("add", false, "Add task to the ToDo list")
	task := flag.String("task", "", "Task to be added to the ToDo list")
	list := flag.Bool("list", false, "List all the tasks in the ToDo list")
	complete := flag.Int("complete", 0, "Item to be completed")
	del := flag.Int("del", 0, "Item to be deleted")
	pretty := flag.Bool("p", false, "Pretty print the ToDo list")
	only_completed := flag.Bool("c", false, "Only print completed tasks")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for how to practice the book.\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright (C) 2022 Tiago Neves\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Define an items list
	l := &todo.List{}

	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		// When developing command-line tools, its a good practice to use the standard
		// error (STDERR) output instead of the standard output (STDOUT) to display
		// erro messages as the user can easily filter them out if they desire.
		fmt.Fprintln(os.Stderr, err)
		// Another good practice is to exit your program with a return code different than 0
		// when erros occur as this is a convention that clearly indicates
		// that the program has an erro or abnormal condition.
		os.Exit(1)
	}

	// Decide what to do based on the provided flags
	switch {
	// For no extra arguments, print the list
	case *list:
		// List current to do items
		if *pretty {
			fmt.Print(l.Pretty())
		} else if *only_completed {
			fmt.Print(l.Completed())
		} else {
			fmt.Print(l)
		}
	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// Add the task
		l.Add(*task)
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		// When any arguments (excluding flags) are provided, they will be
		// used as the new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, i := range t {
			// Add the task
			l.Add(i)
			// Save the new list
			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}

	case *del > 0:
		// Delete the given item
		if err := l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

// getTask function decides where to get the description for a new
// task from: arguments or STDIN
func getTask(r io.Reader, args ...string) ([]string, error) {
	if len(args) > 0 {
		return []string{strings.Join(args, " ")}, nil
	}

	input := make([]string, 0)

	s := bufio.NewScanner(r)
	for s.Scan() {
		if len(s.Text()) == 0 {
			return nil, fmt.Errorf("Task cannot be blank")
		}
		if err := s.Err(); err != nil {
			return nil, err
		}
		if s.Text() == "exit" {
			break
		}
		if s.Text() == "exit" {
			fmt.Println("Bye..")
			os.Exit(0)
		}
		input = append(input, s.Text())
	}

	return input, nil
}
