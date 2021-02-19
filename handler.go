package main

import (
	"file_server/proto/file_server"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
}

func (Server) Upload(req file_server.FileServer_UploadServer) (err error) {
	os.File
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (Server) Download(req *file_server.DownloadRequset, stream file_server.FileServer_DownloadServer) error {
	
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
