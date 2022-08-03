package service

import (
	"context"

	"github.com/zhaoyunxing/protobuf/search"
)

type SearchService struct {
	search.UnimplementedSearchServer
}

func (u *SearchService) Search(context.Context, *search.SearchRequest) (*search.SearchReply, error) {
	return &search.SearchReply{Message: "Hello search service"}, nil
}

type ModelService struct {
	search.UnimplementedSearchServer
}

func (u *ModelService) Search(context.Context, *search.SearchRequest) (*search.SearchReply, error) {
	return &search.SearchReply{Message: "Hello model service"}, nil
}
