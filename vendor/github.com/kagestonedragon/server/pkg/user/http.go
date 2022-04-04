// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kagestonedragon/server/tools/logging"
	"github.com/kagestonedragon/server/tools/tracing"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func MakeHTTPHandler(ctx context.Context, s Service) http.Handler {
	logger := logging.FromContext(ctx)
	logger = log.With(logger, "http handler", "user")
	tracer := tracing.FromContext(ctx)

	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
		// httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerBefore(opentracing.HTTPToContext(tracer, "http server", logger)),
		httptransport.ServerFinalizer(closeHTTPTracer()),
	}

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		makeAddUserEndpoint(s),
		decodePOSTAddUserRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/user").Handler(httptransport.NewServer(
		makeGetUserListEndpoint(s),
		decodeGETGetUserListRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/user/{id}").Handler(httptransport.NewServer(
		makeUpdateUserEndpoint(s),
		decodePUTUser,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/user/{id}").Handler(httptransport.NewServer(
		makeDeleteUserByIdEndpoint(s),
		decodeDELETEDeleteUserByIdRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		makeGetUserByIdEndpoint(s),
		decodeGETGetUserByIdRequest,
		encodeResponse,
		options...,
	))

	return accessControl(r)
}

func httpToContext() httptransport.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		return context.WithValue(ctx, ContextHTTPKey{}, HTTPInfo{
			Method:   req.Method,
			URL:      req.RequestURI,
			From:     req.RemoteAddr,
			Protocol: req.Proto,
		})
	}
}
func closeHTTPTracer() httptransport.ServerFinalizerFunc {
	return func(ctx context.Context, code int, r *http.Request) {
		span := stdopentracing.SpanFromContext(ctx)
		span.Finish()
	}
}

func decodePOSTAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AddUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "decode request body")
	}
	{
		if err := validate(request); err != nil {
			return nil, errors.Wrap(ErrInvalidRequest, err.Error())
		}
	}
	return request, nil
}

func decodeDELETEDeleteUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteUserByIdRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "decode request body")
	}
	vars := mux.Vars(r)

	{
		s, ok := vars["id"]
		if !ok {
			return nil, errors.WithStack(errBadRoute)
		}
		id, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return nil, errors.Wrap(errBadRoute, err.Error())
		}

		request.Id = uint64(id)
	}

	{
		if err := validate(request); err != nil {
			return nil, errors.Wrap(ErrInvalidRequest, err.Error())
		}
	}
	return request, nil
}

func decodeGETGetUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUserByIdRequest

	{
		decoder := schema.NewDecoder()
		err := decoder.Decode(&request, r.URL.Query())
		if err != nil {
			return nil, errors.Wrap(ErrInvalidArgument, err.Error())
		}
	}
	vars := mux.Vars(r)

	{
		s, ok := vars["id"]
		if !ok {
			return nil, errors.WithStack(errBadRoute)
		}
		id, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return nil, errors.Wrap(errBadRoute, err.Error())
		}

		request.Id = uint64(id)
	}

	{
		if err := validate(request); err != nil {
			return nil, errors.Wrap(ErrInvalidRequest, err.Error())
		}
	}
	return request, nil
}

func decodeGETGetUserListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUserListRequest

	{
		decoder := schema.NewDecoder()
		err := decoder.Decode(&request, r.URL.Query())
		if err != nil {
			return nil, errors.Wrap(ErrInvalidArgument, err.Error())
		}
	}
	{
		if err := validate(request); err != nil {
			return nil, errors.Wrap(ErrInvalidRequest, err.Error())
		}
	}
	return request, nil
}

func decodePUTUser(_ context.Context, r *http.Request) (interface{}, error) {
	var request User

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "decode request body")
	}

	vars := mux.Vars(r)

	{
		s, ok := vars["id"]
		if !ok {
			return nil, errors.WithStack(errBadRoute)
		}
		id, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return nil, errors.Wrap(errBadRoute, err.Error())
		}

		request.Id = uint64(id)
	}

	{
		if err := validate(request); err != nil {
			return nil, errors.Wrap(ErrInvalidRequest, err.Error())
		}
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encodeError handles error from business-layer.
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("X-Esp-Error", err.Error())
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")

	w.WriteHeader(getHTTPStatusCode(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// accessControl is CORS middleware.
func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}