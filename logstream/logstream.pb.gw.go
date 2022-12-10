// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: logstream.proto

/*
Package logstream is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package logstream

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

func request_LogStream_StreamLogs_0(ctx context.Context, marshaler runtime.Marshaler, client LogStreamClient, req *http.Request, pathParams map[string]string) (LogStream_StreamLogsClient, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.StreamLogs(ctx)
	if err != nil {
		grpclog.Infof("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq StreamLogRequest
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Infof("Failed to decode request: %v", err)
			return err
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Infof("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Infof("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Infof("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

func request_LogStream_GetFirebaseAuthToken_0(ctx context.Context, marshaler runtime.Marshaler, client LogStreamClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetFirebaseAuthTokenRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetFirebaseAuthToken(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_LogStream_GetFirebaseAuthToken_0(ctx context.Context, marshaler runtime.Marshaler, server LogStreamServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetFirebaseAuthTokenRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetFirebaseAuthToken(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterLogStreamHandlerServer registers the http handlers for service LogStream to "mux".
// UnaryRPC     :call LogStreamServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterLogStreamHandlerFromEndpoint instead.
func RegisterLogStreamHandlerServer(ctx context.Context, mux *runtime.ServeMux, server LogStreamServer) error {

	mux.Handle("POST", pattern_LogStream_StreamLogs_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("GET", pattern_LogStream_GetFirebaseAuthToken_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/api.public.logstream.LogStream/GetFirebaseAuthToken", runtime.WithHTTPPathPattern("/api/v0/logstream/firebase-auth-token"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_LogStream_GetFirebaseAuthToken_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogStream_GetFirebaseAuthToken_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterLogStreamHandlerFromEndpoint is same as RegisterLogStreamHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterLogStreamHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterLogStreamHandler(ctx, mux, conn)
}

// RegisterLogStreamHandler registers the http handlers for service LogStream to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterLogStreamHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterLogStreamHandlerClient(ctx, mux, NewLogStreamClient(conn))
}

// RegisterLogStreamHandlerClient registers the http handlers for service LogStream
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "LogStreamClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "LogStreamClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "LogStreamClient" to call the correct interceptors.
func RegisterLogStreamHandlerClient(ctx context.Context, mux *runtime.ServeMux, client LogStreamClient) error {

	mux.Handle("POST", pattern_LogStream_StreamLogs_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.public.logstream.LogStream/StreamLogs", runtime.WithHTTPPathPattern("/api.public.logstream.LogStream/StreamLogs"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_LogStream_StreamLogs_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogStream_StreamLogs_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_LogStream_GetFirebaseAuthToken_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.public.logstream.LogStream/GetFirebaseAuthToken", runtime.WithHTTPPathPattern("/api/v0/logstream/firebase-auth-token"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_LogStream_GetFirebaseAuthToken_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogStream_GetFirebaseAuthToken_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_LogStream_StreamLogs_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"api.public.logstream.LogStream", "StreamLogs"}, ""))

	pattern_LogStream_GetFirebaseAuthToken_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v0", "logstream", "firebase-auth-token"}, ""))
)

var (
	forward_LogStream_StreamLogs_0 = runtime.ForwardResponseStream

	forward_LogStream_GetFirebaseAuthToken_0 = runtime.ForwardResponseMessage
)