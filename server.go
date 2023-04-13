package main

import (
	"context"
	"log"
	"os"

	pb "github.com/asadzeynal/file-grpc-task/gen/file/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedFileServiceServer
	fileDirPath string
}

func NewServer(path string) *Server {
	return &Server{fileDirPath: path}

}

func (s *Server) Upload(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	os.Chdir(s.fileDirPath)
	f, err := os.Create(req.GetName())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()

	_, err = f.Write(req.File)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.UploadResponse{Name: req.Name}, nil
}

func (s *Server) LS(context.Context, *pb.LSRequest) (*pb.LSResponse, error) {
	files, err := os.ReadDir(s.fileDirPath)
	if err != nil {
		return nil, err
	}

	entries := make([]*pb.FileEntry, 0)
	for i := range files {
		info, err := files[i].Info()
		if err != nil {
			//handle
			continue
		}
		entries = append(entries, &pb.FileEntry{
			Name:      info.Name(),
			UpdatedAt: timestamppb.New(info.ModTime()),
		},
		)
	}
	return &pb.LSResponse{Files: entries}, nil
}

func (s *Server) Download(ctx context.Context, req *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	f, err := os.ReadFile(req.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.DownloadResponse{File: f}, nil
}
