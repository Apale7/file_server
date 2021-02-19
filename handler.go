package main

import (
	"context"
	"file_server/proto/file_server"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
}

func (Server) Upload(context.Context, *file_server.UploadRequest) (*file_server.UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (Server) Download(context.Context, *file_server.DownloadRequset) (*file_server.DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
