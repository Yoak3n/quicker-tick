package app

import (
	"context"
	"fmt"
	"quicker-tick/internal/model"
	"quicker-tick/internal/od"
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
func (a *App) AddTask(object string, date string) string {
	err := od.UnmashalRawString(object, date)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) GetTasks(dates []string) map[string][]*model.TaskView {
	view := od.GetTasks(dates)
	return view
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
