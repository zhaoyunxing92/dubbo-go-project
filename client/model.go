package client

import (
	"context"

	"github.com/dubbogo/triple/pkg/common/constant"

	"github.com/zhaoyunxing/protobuf/search"

	"github.com/dubbogo/grpc-go"
	"github.com/dubbogo/triple/pkg/common"
	"github.com/dubbogo/triple/pkg/triple"
)

// ModelClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelClient interface {
	Search(ctx context.Context, in *search.SearchRequest, opts ...grpc.CallOption) (*search.SearchReply, common.ErrorWithAttachment)
}

type modelClient struct {
	cc *triple.TripleConn
}

type ModelClientImpl struct {
	Search func(ctx context.Context, in *search.SearchRequest) (*search.SearchReply, error)
}

func (c *ModelClientImpl) GetDubboStub(cc *triple.TripleConn) ModelClient {
	return NewSearchClient(cc)
}

func (c *ModelClientImpl) XXX_InterfaceName() string {
	return "org.apache.dubbo.Model"
}

func NewModelClient(cc *triple.TripleConn) ModelClient {
	return &modelClient{cc}
}

func (c *modelClient) Search(ctx context.Context, in *search.SearchRequest, opts ...grpc.CallOption) (*search.SearchReply, common.ErrorWithAttachment) {
	out := new(search.SearchReply)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Search", in, out)
}
