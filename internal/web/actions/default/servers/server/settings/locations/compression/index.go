// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package compression

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	ServerId   int64
	LocationId int64
}) {
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithLocationId(this.AdminContext(), params.LocationId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id
	this.Data["compressionConfig"] = webConfig.Compression

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId           int64
	CompressionJSON []byte

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("修改Web %d 的压缩设置", params.WebId)

	// 校验配置
	var compressionConfig = &serverconfigs.HTTPCompressionConfig{}
	err := json.Unmarshal(params.CompressionJSON, compressionConfig)
	if err != nil {
		this.Fail("配置校验失败：" + err.Error())
	}

	err = compressionConfig.Init()
	if err != nil {
		this.Fail("配置校验失败：" + err.Error())
	}

	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebCompression(this.AdminContext(), &pb.UpdateHTTPWebCompressionRequest{
		HttpWebId:       params.WebId,
		CompressionJSON: params.CompressionJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
