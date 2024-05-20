package pelocalhandler

import "h5travelotobackend/component/appContext"

type peLocalHandler struct {
	appCtx appContext.AppContext
}

func NewPELocalHandler(appCtx appContext.AppContext) *peLocalHandler {
	return &peLocalHandler{appCtx: appCtx}
}
