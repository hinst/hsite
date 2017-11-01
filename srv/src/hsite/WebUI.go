package hsite

import (
	"hgo"
	"net/http"
	"sync"
	"sync/atomic"
)

type WebUI struct {
	URL          string
	Holder       sync.WaitGroup
	Stopped      int32
	DebugEnabled bool
}

func (this *WebUI) Create() {
	this.URL = "/hsite"
}

func (this *WebUI) Start() {
	this.SetupFiles()
}

func (this *WebUI) WrapRequestHandler(h hgo.HttpFunc) hgo.HttpFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		this.Holder.Add(1)
		defer this.Holder.Done()
		if atomic.LoadInt32(&this.Stopped) == 0 {
			h(responseWriter, request)
		}
	}
}

func (this *WebUI) Stop() {
	atomic.AddInt32(&this.Stopped, 1)
	this.Holder.Wait()
}

func (this *WebUI) SetupFiles() {
	var dir = hgo.AppDir + "/../ui/jsc"
	var url = this.URL + "/jsc/"
	http.Handle(url, http.StripPrefix(url,
		http.FileServer(http.Dir(dir))))
}
