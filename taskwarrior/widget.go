package taskwarrior

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.Widget
	app    *tview.Application
	filter string
	pages  *tview.Pages
	tasks  []task
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 📝 Taskwarrior ", "taskwarrior", true),

		app:    app,
		filter: Config.UString("wtf.mods.taskwarrior.filter"),
		tasks:  []task{},
		pages:  pages,
	}

	// widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.UpdateRefreshedAt()
	widget.load()
	widget.display()
}

func (widget *Widget) load() {
	cmd := exec.Command("task", widget.filter, "export")
	tasksJSON := wtf.ExecuteCommand(cmd)
	tasks := []task{}
	err := json.Unmarshal(tasksJSON, tasks)
	if err != nil {
		panic("Trouble reading JSON from task")
	}

}

func (widget *Widget) display() {
	str := ""

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", str)
}
