package utiles

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/onlyLTY/dockerCopilot/UGREEN/internal/svc"
	MyType "github.com/onlyLTY/dockerCopilot/UGREEN/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetContainerList(ctx *svc.ServiceContext) ([]MyType.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logx.Errorf("connect to docker error: %v", err)
		return nil, err
	}
	// 获取所有容器（包括停止的容器）
	dockerContainerList, err := cli.ContainerList(context.Background(), container.ListOptions{
		All: true, // 设置为true来获取所有容器
	})
	if err != nil {
		logx.Errorf("get container list error: %v", err)
		return nil, err
	}
	var containerList []MyType.Container
	for _, dockerContainerInfo := range dockerContainerList {
		containerInfo := MyType.Container{
			Container: dockerContainerInfo,
		}
		containerList = append(containerList, containerInfo)
	}
	return containerList, nil
}

func CheckImageUpdate(ctx *svc.ServiceContext, containerListData []MyType.Container) []MyType.Container {
	for i, v := range containerListData {
		if _, ok := ctx.HubImageInfo.Data[v.ImageID]; ok {
			if ctx.HubImageInfo.Data[v.ImageID].NeedUpdate {
				containerListData[i].Update = true
			}
		}
	}
	return containerListData
}
