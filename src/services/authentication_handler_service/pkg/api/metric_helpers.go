package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
)

// ExtractIdOperationAndInstrument extracts an account id from a request and increments the necessary metrics
func (s *Server) ExtractIdOperationAndInstrument(r *http.Request, operation string) (uint32, error) {
	var status = constants.SUCCESS
	authnID, err := helper.ExtractIDFromRequest(r)
	if err != nil {
		status = constants.FAILURE
	}

	s.metrics.ExtractIdOperationCounter.WithLabelValues(operation, status).Inc()
	return authnID, err
}

func (s *Server) RemoteOperationAndInstrument(f func() error, operationType string, took *time.Duration, ctx opentracing.SpanContext) error {
	// we start a child span for the rpc operation
	authnSvcRpcSpan := s.tracer.StartSpan(fmt.Sprintf("AUTHENTICATION_SERVICE_%s_RPC", operationType) , opentracing.ChildOf(ctx))
	defer authnSvcRpcSpan.Finish()

	var status = constants.SUCCESS
	err := f()
	if err != nil {
		status = constants.FAILURE
	}

	s.metrics.RemoteOperationStatusCounter.WithLabelValues(operationType, status).Inc()
	s.metrics.RemoteOperationsLatencyCounter.WithLabelValues(operationType, status).Observe(took.Seconds())
	return err
}

func (s *Server) RemoteOperationAndInstrumentWithResult(
	f func() (interface{}, error),
	operationType string,
	took *time.Duration) (interface{}, error) {

	var status = constants.SUCCESS
	result, err := f()
	if err != nil {
		status = constants.FAILURE
	}

	s.metrics.RemoteOperationStatusCounter.WithLabelValues(operationType, status).Inc()
	s.metrics.RemoteOperationsLatencyCounter.WithLabelValues(operationType, status).Observe(took.Seconds())
	return result, err
}

func (s *Server) DecodeRequestAndInstrument(w http.ResponseWriter, r *http.Request, obj interface{}, operationType string) error {
	var status = constants.SUCCESS

	err := helper.DecodeJSONBody(w, r, &obj)
	if err != nil {
		status = constants.FAILURE
	}

	s.metrics.DecodeRequestStatusCounter.WithLabelValues(operationType, status).Inc()
	return err
}
