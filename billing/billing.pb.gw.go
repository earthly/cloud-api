// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: billing.proto

/*
Package billing is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package billing

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_Billing_GetBillingPlan_0(ctx context.Context, marshaler runtime.Marshaler, client BillingClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBillingPlanRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["org_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "org_id")
	}

	protoReq.OrgId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "org_id", err)
	}

	msg, err := client.GetBillingPlan(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Billing_GetBillingPlan_0(ctx context.Context, marshaler runtime.Marshaler, server BillingServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBillingPlanRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["org_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "org_id")
	}

	protoReq.OrgId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "org_id", err)
	}

	msg, err := server.GetBillingPlan(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterBillingHandlerServer registers the http handlers for service Billing to "mux".
// UnaryRPC     :call BillingServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBillingHandlerFromEndpoint instead.
func RegisterBillingHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BillingServer) error {

	mux.Handle("GET", pattern_Billing_GetBillingPlan_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/api.public.billing.Billing/GetBillingPlan", runtime.WithHTTPPathPattern("/api/v0/billing/plans/{org_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Billing_GetBillingPlan_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Billing_GetBillingPlan_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterBillingHandlerFromEndpoint is same as RegisterBillingHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBillingHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterBillingHandler(ctx, mux, conn)
}

// RegisterBillingHandler registers the http handlers for service Billing to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBillingHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBillingHandlerClient(ctx, mux, NewBillingClient(conn))
}

// RegisterBillingHandlerClient registers the http handlers for service Billing
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BillingClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BillingClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BillingClient" to call the correct interceptors.
func RegisterBillingHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BillingClient) error {

	mux.Handle("GET", pattern_Billing_GetBillingPlan_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.public.billing.Billing/GetBillingPlan", runtime.WithHTTPPathPattern("/api/v0/billing/plans/{org_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Billing_GetBillingPlan_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Billing_GetBillingPlan_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Billing_GetBillingPlan_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "v0", "billing", "plans", "org_id"}, ""))
)

var (
	forward_Billing_GetBillingPlan_0 = runtime.ForwardResponseMessage
)
