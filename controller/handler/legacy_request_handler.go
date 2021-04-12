package handler

import (
	"github.com/faasflow/runtime"
	"github.com/faasflow/runtime/controller/util"

	"github.com/faasflow/sdk/executor"
)

func LegacyRequestHandler(response *runtime.Response, request *runtime.Request, ex executor.Executor) error {
	var handler func(response *runtime.Response, request *runtime.Request, ex executor.Executor) error

	id := ""
	switch {
	case util.IsDagExportRequest(request.RawQuery):
		handler = GetDagHandler

	case util.GetPauseRequestID(request.RawQuery) != "":
		id = util.GetPauseRequestID(request.RawQuery)
		handler = PauseFlowHandler

	case util.GetStopRequestID(request.RawQuery) != "":
		id = util.GetStopRequestID(request.RawQuery)
		handler = StopFlowHandler

	case util.GetResumeRequestID(request.RawQuery) != "":
		id = util.GetResumeRequestID(request.RawQuery)
		handler = ResumeFlowHandler

	case util.GetStateRequestID(request.RawQuery) != "":
		id = util.GetStateRequestID(request.RawQuery)
		handler = FlowStateHandler

	default:
		id = request.GetHeader(util.RequestIdHeader)
		if id == "" {
			handler = ExecuteFlowHandler
		} else {
			handler = PartialExecuteFlowHandler
		}
	}

	request.RequestID = id
	err := handler(response, request, ex)
	return err
}
