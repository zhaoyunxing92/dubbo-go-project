package client

import (
	"context"

	"github.com/dubbogo/triple/pkg/common/constant"

	"github.com/zhaoyunxing/protobuf/search"

	"github.com/dubbogo/grpc-go"
	"github.com/dubbogo/triple/pkg/common"
	"github.com/dubbogo/triple/pkg/triple"
)

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchClient interface {
	Search(ctx context.Context, in *search.SearchRequest, opts ...grpc.CallOption) (*search.SearchReply, common.ErrorWithAttachment)
}

type searchClient struct {
	cc *triple.TripleConn
}

type SearchClientImpl struct {
	Search func(ctx context.Context, in *search.SearchRequest) (*search.SearchReply, error)
}

func (c *SearchClientImpl) GetDubboStub(cc *triple.TripleConn) SearchClient {
	return NewSearchClient(cc)
}

func (c *SearchClientImpl) XXX_InterfaceName() string {
	return "org.apache.dubbo.Search"
}

func NewSearchClient(cc *triple.TripleConn) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) Search(ctx context.Context, in *search.SearchRequest, opts ...grpc.CallOption) (*search.SearchReply, common.ErrorWithAttachment) {
	out := new(search.SearchReply)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Search", in, out)
}
