package beam

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	port   = ":3333"
	appCtx context.Context
)

type beamRequest struct {
	Payload string
}

func handleBeam(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		runtime.LogErrorf(appCtx, "error reading request body: %s", err)
	}

	var requestData beamRequest
	err = json.Unmarshal(body, &requestData)

	if err != nil {
		runtime.LogErrorf(appCtx, "error parsing json: %s", err)
		requestData.Payload = string(body)
	}

	runtime.EventsEmit(appCtx, "beamMessage", requestData)

	runtime.LogInfof(appCtx, "parsed json: %+v", requestData)
}

func StartServer(ctx context.Context) {
	appCtx = ctx

	runtime.LogInfof(ctx, "starting http server on port %v \n", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/beam", handleBeam)

	err := http.ListenAndServe(port, mux)

	if errors.Is(err, http.ErrServerClosed) {
		runtime.LogInfo(ctx, "server closed")
	} else if err != nil {
		runtime.LogErrorf(ctx, "error starting server: %s", err)
	}
}
