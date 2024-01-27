package container

import (
	"context"
	"github.com/google/uuid"
	"github.com/onlyLTY/dockerCopilot/UGREEN/internal/svc"
	"github.com/onlyLTY/dockerCopilot/UGREEN/internal/types"
	"github.com/onlyLTY/dockerCopilot/UGREEN/internal/utiles"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.ContainerUpdateReq) (resp *types.Resp, err error) {
	resp = &types.Resp{}
	taskID := uuid.New().String()
	go func() {
		// Catch any panic and log the error
		defer func() {
			if r := recover(); r != nil {
				logx.Errorf("Recovered from panic in UpdateContainer: %v", r)
			}
		}()
		imageNameAndTag := req.ImageNameAndTag
		err := utiles.UpdateContainer(l.svcCtx, req.Id, req.ContainerName, imageNameAndTag, req.DelOldContainer, taskID)
		if err != nil {
			logx.Errorf("Error in UpdateContainer: %v", err)
		}
	}()
	resp.Code = 200
	resp.Msg = "success"
	resp.Data = map[string]string{"taskID": taskID}
	return resp, nil
}
