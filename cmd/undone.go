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

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:     "undone",
	Aliases: []string{"undo"},
	Short:   "Mark Task as UnDone",
	Long:    `Mark Task as UnDone`,
	Run:     undoneRun,
}

// Main func
func undoneRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Too short argument")
		fmt.Println("Usage: tasks undone [task id]")
		os.Exit(1)
	}
	tasks, err := todo.ReadTasks(todo.DatabaseFile)
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid index\ninvalid syntax")
	}
	if i > 0 && i <= len(tasks) {
		tasks[i-1].Done = false
		fmt.Printf("%q %v\n", tasks[i-1].Text, "marked undone")

		sort.Sort(todo.ByPri(tasks))
		todo.SaveTasks(todo.DatabaseFile, tasks)
	} else {
		log.Println(i, "doesn't match any tasks")
	}
}

// flag variables

func init() {
	rootCmd.AddCommand(undoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
