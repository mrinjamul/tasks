package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

// Task : structure
type Task struct {
	Text     string
	Priority int
	position int
	Done     bool
}

// DatabaseFile is the db file
var DatabaseFile string = GetHomeDir() + "/.config/tasks/tasks.json"
var (
	Appname   = "tasks"
	Version   = "1.0.0"
	GitCommit = "0000000000000000000000000000000000000000"
)

// GetHomeDir returns homedir
func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory.")
	}
	return home
}

// CreateDatabase create folder and the file to save data
func CreateDatabase() error {
	dbLoc := GetHomeDir() + "/.config/tasks/"

	err := os.MkdirAll(dbLoc, 0755)
	if err != nil {
		return err
	}
	if _, err := os.Stat(DatabaseFile); os.IsNotExist(err) {
		err = ioutil.WriteFile(DatabaseFile, nil, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveTasks : save data
func SaveTasks(filename string, tasks []Task) error {
	b, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(DatabaseFile, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadTasks : read data
func ReadTasks(filename string) ([]Task, error) {
	b, err := ioutil.ReadFile(DatabaseFile)
	if err != nil {
		return []Task{}, err
	}
	var tasks []Task
	if err := json.Unmarshal(b, &tasks); err != nil {
		return []Task{}, err
	}
	for i := range tasks {
		tasks[i].position = i + 1
	}
	return tasks, nil
}

// SetPriority : sets priority to todo
func (i *Task) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

// PrettyP : prettify priority list
func (i *Task) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}

	return " "
}

// Label : index lists
func (i *Task) Label() string {
	return strconv.Itoa(i.position) + "."
}

// ByPri implements sort.Interface for []Item base on
// the priority & position field.
type ByPri []Task

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[j].Done
	}

	if s[i].Priority != s[j].Priority {
		return s[i].Priority < s[j].Priority
	}
	return s[i].position < s[j].position
}

// PrettyDone : prettify
func (i *Task) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

// RemoveTask removes todo from list
func RemoveTask(slice []Task, s int) []Task {
	return append(slice[:s], slice[s+1:]...)
}

// ConfirmPrompt will prompt to user for yes or no
func ConfirmPrompt(message string) bool {
	var response string
	fmt.Print(message + " (yes/no) :")
	fmt.Scanln(&response)

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		return false
	}
}

// SortSlice sort arrays
func SortSlice(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool { return slice[i] > slice[j] })
	return slice
}

// RemoveDuplicate removes duplicate from slice
func RemoveDuplicate(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
