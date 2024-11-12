package app

import (
	"context"
	"fmt"
	"log"
	"quicker-tick/internal/controller"
	"quicker-tick/internal/model"
	"quicker-tick/internal/od"
	"quicker-tick/package/shortcut"
)

var app *App

// App struct
type App struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	app = &App{}
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
	registerTray()
}

func (a *App) BoardcastCtx(ctx *context.Context) {
	for {
		if a.Ctx != nil {
			ctx = &a.Ctx
		}
	}
}
func (a *App) AddTask(object model.TaskView) string {
	r := od.ObjectToTaskTable(&object)
	controller.CreateTask(r)
	return ""
}

func (a *App) GetTasks(dates []string) map[string][]*model.TaskView {
	view := od.GetTasks(dates)
	return view
}

func (a *App) AddAction(object model.Action) {
	record := od.ObjectToActionTable(&object)
	controller.CreateAction(record)
}

func (a *App) GetActions() []*model.Action {
	log.Println("GetActions")
	return od.GetActions()
}

func (a *App) ConvertTaskView(tasks []*model.TaskView) model.Item {
	item := od.ConvertToDoc(tasks)
	return item
}
func (a *App) ConvertTask(item model.Item) string {
	htmlString := od.ConvertToHtmlString(item, false)
	return htmlString
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RunCommand(url string) {
	log.Println("RunCommand")
	shortcut.OpenBrowser(url)
}
