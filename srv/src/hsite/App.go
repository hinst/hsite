package hsite

import (
	"hgo"
	"net/http"
	"sync"
)

type App struct {
	WebUI  *WebUI
	Holder sync.WaitGroup
}

func (this *App) Run() {
	this.Start()
	this.Holder.Add(1)
	hgo.InstallShutdownReceiver(
		func() { this.Holder.Done() })
	this.Holder.Wait()
	this.Stop()
}

func (this *App) Start() {
	this.WebUI = &WebUI{}
	this.WebUI.Create()
	this.WebUI.Start()
	http.ListenAndServe(":9003", nil)
}

func (this *App) Stop() {
	this.WebUI.Stop()
}
