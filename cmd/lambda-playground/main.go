package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/get-ytt/ytt/pkg/cmd"
)

type HandlerFuncAdapter struct {
	RequestAccessor
	handler http.Handler
}

func New(handler http.Handler) *HandlerFuncAdapter {
	return &HandlerFuncAdapter{
		handler: handler,
	}
}

func (h *HandlerFuncAdapter) Proxy(event events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	req, err := h.ProxyEventToHTTPRequest(event)
	if err != nil {
		return events.ALBTargetGroupResponse{StatusCode: 421}, fmt.Errorf("Could not convert event to request: %v", err)
	}

	w := NewProxyResponseWriter()
	h.handler.ServeHTTP(http.ResponseWriter(w), req)

	resp, err := w.GetProxyResponse()
	if err != nil {
		return events.ALBTargetGroupResponse{StatusCode: 422}, fmt.Errorf("Error while generating response: %v", err)
	}

	return resp, nil
}

func main() {
	playgroundOpts := cmd.NewPlaygroundOptions()
	playgroundOpts.BinaryPath = "/var/task/ytt"
	playgroundOpts.CheckCookie = true
	lambda.Start(New(playgroundOpts.Server().Mux()).Proxy)
}
