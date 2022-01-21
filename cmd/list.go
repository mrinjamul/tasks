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
	"sort"
	"text/tabwriter"

	"github.com/charmbracelet/lipgloss"
	"github.com/mrinjamul/tasks/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list all tasks",
	Long:    `List all tasks`,
	Run:     listRun,
}

var (
	// make looks good
	// width       = 96
	// columnWidth = 80

	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	morehigh  = lipgloss.AdaptiveColor{Light: "#FFCB6B", Dark: "#FFCB6B"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	// List.
	list = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		MarginRight(2)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			MarginRight(2).
			Render

	listItem          = lipgloss.NewStyle().PaddingLeft(2).Render
	listItemHighlight = lipgloss.NewStyle().Foreground(highlight).PaddingLeft(2).Render
	listItemMoreHigh  = lipgloss.NewStyle().Foreground(morehigh).PaddingLeft(2).Render

	checkMark = lipgloss.NewStyle().SetString("✓").
			Foreground(special).
			PaddingRight(1).
			String()

	listDone = func(s string) string {
		return checkMark + lipgloss.NewStyle().
			Strikethrough(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
			Render(s)
	}
)

// Main func
func listRun(cmd *cobra.Command, args []string) {

	// Check if database exists or create
	if _, err := os.Stat(todo.DatabaseFile); os.IsNotExist(err) {
		todo.CreateDatabase()
	}
	// Read data from file
	tasks, err := todo.ReadTasks(todo.DatabaseFile)

	if err != nil {
		file := []byte("[]")
		err = ioutil.WriteFile(todo.DatabaseFile, file, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
	if len(tasks) == 0 {
		fmt.Println("Empty Todo list")
	}

	sort.Sort(todo.ByPri(tasks))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	vertical := lipgloss.JoinVertical(lipgloss.Left, listHeader("TODO:"))
	for _, i := range tasks {
		if allOpt || i.Done == doneOpt {
			if i.Done {
				vertical = lipgloss.JoinVertical(lipgloss.Left, vertical, listDone(i.Label()+i.Text))
			} else if i.Priority == 1 {
				vertical = lipgloss.JoinVertical(lipgloss.Left, vertical, listItemMoreHigh(i.Label()+i.Text))
			} else if i.Priority == 3 {
				vertical = lipgloss.JoinVertical(lipgloss.Left, vertical, listItem(i.Label()+i.Text))
			} else {
				vertical = lipgloss.JoinVertical(lipgloss.Left, vertical, listItemHighlight(i.Label()+i.Text))
			}
		}
	}
	lists := lipgloss.JoinHorizontal(lipgloss.Top,
		list.Render(vertical),
	)
	fmt.Fprintln(w, lists)
	w.Flush()
}

// flag variables
var (
	doneOpt bool
	allOpt  bool
)

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "Show 'Done' Tasks")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "Show All Tasks")
}
