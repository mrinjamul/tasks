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
	"os"
	"sort"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/mrinjamul/tasks/todo"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a task",
	Long:    `Remove will remove a task from the list by Label(index)`,
	Run:     removeRun,
}

var (
	removeMark = lipgloss.NewStyle().SetString("[-]").
			Foreground(lipgloss.AdaptiveColor{Light: "#FF5F5F", Dark: "#FF5F5F"}).
			PaddingRight(1).
			String()

	removeList = func(s string) string {
		return removeMark + lipgloss.NewStyle().
			Foreground(morehigh).
			Render(s)
	}
)

// Main func
func removeRun(cmd *cobra.Command, args []string) {
	// remove only done tasks func
	if doOpt {
		response := todo.ConfirmPrompt("Do you want to remove all done task(s)?")
		if response {
			tasks, err := todo.ReadTasks(todo.DatabaseFile)
			var undoneTasks []todo.Task
			if err != nil {
				fmt.Println("Something went wrong!")
				os.Exit(0)
			}
			for i, task := range tasks {
				if !tasks[i].Done {
					undoneTasks = append(undoneTasks, task)
				}
				if tasks[i].Done {
					text := tasks[i].Text
					// fmt.Println("- " + "\"" + strconv.Itoa(i) + ". " + text + "\"" + " has been removed")
					li := lipgloss.JoinHorizontal(lipgloss.Left,
						removeList("'"+strconv.Itoa(i)+". "+text+"'"+" has been removed"),
					)
					fmt.Println(li)
				}
			}
			sort.Sort(todo.ByPri(undoneTasks))
			todo.SaveTasks(todo.DatabaseFile, undoneTasks)
		}
	} else { // Remove one by one
		if len(args) == 0 {
			fmt.Println("Error: Too short argument")
			fmt.Println("Usage: tasks remove [task id]")
			os.Exit(1)
		}
		tasks, err := todo.ReadTasks(todo.DatabaseFile)
		if err != nil {
			fmt.Println("Failed to get todo lists")
			os.Exit(1)
		}

		// create  remove lists for batch remove purpose
		var rmList = []int{}

		for arg := 0; arg < len(args); arg++ {
			i, err := strconv.Atoi(args[arg])

			if err != nil {
				fmt.Println(args[arg] + " is not a valid index\ninvalid syntax")
				os.Exit(0)
			}
			rmList = append(rmList, i)
		}
		// sort and remove duplicate
		rmList = todo.SortSlice(rmList)
		rmList = todo.RemoveDuplicate(rmList)
		for _, i := range rmList {
			if i > 0 && i <= len(tasks) {
				text := tasks[i-1].Text
				tasks = todo.RemoveTask(tasks, i-1)
				// fmt.Println("- " + "\"" + strconv.Itoa(i) + ". " + text + "\"" + " has been removed")
				li := lipgloss.JoinHorizontal(lipgloss.Left,
					removeList("'"+strconv.Itoa(i)+". "+text+"'"+" has been removed"),
				)
				fmt.Println(li)
				sort.Sort(todo.ByPri(tasks))
				todo.SaveTasks(todo.DatabaseFile, tasks)
			} else {
				fmt.Println(i, "doesn't match any tasks")
			}
		}
	}
}

// flag variables
var (
	doOpt bool
)

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	removeCmd.Flags().BoolVarP(&doOpt, "done", "d", false, "remove only done tasks")
}
