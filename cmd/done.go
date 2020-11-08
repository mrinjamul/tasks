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
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/mrinjamul/tasks/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark Task as Done",
	Long:    `Mark Task as Done`,
	Run:     doneRun,
}

// Main func
func doneRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Too short argument")
		fmt.Println("Usage: tasks done [task id]")
		os.Exit(1)
	}
	tasks, err := todo.ReadTasks(todo.DatabaseFile)
	i, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println(args[0], "is not a valid index\ninvalid syntax")
		os.Exit(1)
	}
	if i > 0 && i <= len(tasks) {
		tasks[i-1].Done = true
		fmt.Printf("%q %v\n", tasks[i-1].Text, "marked done")

		sort.Sort(todo.ByPri(tasks))
		todo.SaveTasks(todo.DatabaseFile, tasks)
	} else {
		log.Println(i, "doesn't match any tasks")
	}
}

// flag variables

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
