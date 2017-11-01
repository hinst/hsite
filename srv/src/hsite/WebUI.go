package hsite

import (
	"hgo"
	"net/http"
	"sync"
	"sync/atomic"
)

type WebUI struct {
	Holder  sync.WaitGroup
	Stopped int32
}

func (this *WebUI) Start() {
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
}
