/*Package cmd ...
Copyright © 2020 Injamul Mohammad Mollah

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/mrinjamul/tasks/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new task",
	Long:    `Add will create a new task to the list`,
	Run:     addRun,
}

var (
	addMark = lipgloss.NewStyle().SetString("[+]").
		Foreground(lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}).
		PaddingRight(1).
		String()
	newlist = func(s string) string {
		return addMark + lipgloss.NewStyle().
			Foreground(highlight).
			Render(s)
	}
)

// Main func
func addRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Too short argument")
		fmt.Println("Usage: tasks add [task]")
		os.Exit(1)
	}
	tasks, err := todo.ReadTasks(todo.DatabaseFile)
	if err != nil {
		file := []byte("[]")
		err = ioutil.WriteFile(todo.DatabaseFile, file, 0644)
		if err != nil {
			fmt.Println("Error: Unable to create database file")
			os.Exit(1)
		}
	}
	if _, err := os.Stat(todo.DatabaseFile); os.IsNotExist(err) {
		file := []byte("[]")
		err = ioutil.WriteFile(todo.DatabaseFile, file, 0644)
		if err != nil {
			fmt.Println("Error: Unable to create database file")
			os.Exit(1)
		}
	}
	var todoName string
	for _, x := range args {
		todoName += x + " "
	}
	task := todo.Task{Text: todoName}
	task.SetPriority(priority)
	tasks = append(tasks, task)
	err = todo.SaveTasks(todo.DatabaseFile, tasks)
	// fmt.Println("[+] New task added")
	li := lipgloss.JoinHorizontal(lipgloss.Left,
		newlist("New task added "+todoName),
	)
	fmt.Println(li)
	if err != nil {
		fmt.Println(err)
	}
}

// flag variables
var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
