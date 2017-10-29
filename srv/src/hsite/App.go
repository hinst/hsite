package hsite

import (
	"hgo"
	"sync"
)

type App struct {
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
}

func (this *App) Stop() {
}
