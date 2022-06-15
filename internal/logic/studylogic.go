package logic

import (
	"context"

	"study/internal/svc"
	"study/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudyLogic {
	return &StudyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudyLogic) Study(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
