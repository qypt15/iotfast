// ==========================================================================
// 物联网快速开发自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-07-02 23:41:34
// 生成路径: iotfast/internal/app/service/mqtt_status.go
// 生成人：dwx
// ==========================================================================

package service

import (
	"context"
	"iotfast/api/v1/mqtt"
	"iotfast/internal/app/mqtt/dao"
	"iotfast/internal/app/mqtt/model/entity"
	systemConsts "iotfast/internal/app/system/consts"
	"iotfast/library/libErr"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

//type mqttStatus struct {
//}
//var MqttStatus = new(mqttStatus)
type IMqttStatus interface {
	List(ctx context.Context, req *mqtt.MqttStatusSearchReq) (total, page int, list []*entity.MqttStatus, err error)
	Get(ctx context.Context, id int) (info *entity.MqttStatus, err error)
	Add(ctx context.Context, req *mqtt.MqttStatusAddReq) (err error)
	Edit(ctx context.Context, req *mqtt.MqttStatusEditReq) error
	DeleteByIds(ctx context.Context, ids []int) (err error)
	ChangeStatus(ctx context.Context, req *mqtt.MqttStatusStatusReq) error
}
type mqttStatusImpl struct {
}

var mqttStatusService = mqttStatusImpl{}

func MqttStatus() IMqttStatus {
	return &mqttStatusService
}

// List 获取任务列表
func (s *mqttStatusImpl) List(ctx context.Context, req *mqtt.MqttStatusSearchReq) (total, page int, list []*entity.MqttStatus, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = systemConsts.PageSize
	}
	m := dao.MqttStatus.Ctx(ctx)
	if req.Name != "" {
		m = m.Where(dao.MqttStatus.Columns().Name+" like ?", "%"+req.Name+"%")
	}
	if req.ClientId != "" {
		m = m.Where(dao.MqttStatus.Columns().ClientId+" = ?", req.ClientId)
	}
	if req.Status != "" {
		m = m.Where(dao.MqttStatus.Columns().Status+" = ?", req.Status)
	}
	if req.UserName != "" {
		m = m.Where(dao.MqttStatus.Columns().UserName+" = ?", req.UserName)
	}
	if req.Topic != "" {
		m = m.Where(dao.MqttStatus.Columns().Topic+" = ?", req.Topic)
	}
	err = g.Try(func() {
		total, err = m.Count()
		libErr.ErrPrint(ctx, err, "获取MqttStatus列表失败")
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取总行数失败")
			return
		}
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		err = m.Page(page, req.PageSize).Order(order).Scan(&list)
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取数据失败")
		}
	})
	return
}

// Get 通过id获取
func (s *mqttStatusImpl) Get(ctx context.Context, id int) (info *entity.MqttStatus, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.MqttStatus.Ctx(ctx).Where(dao.MqttStatus.Columns().Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *mqttStatusImpl) Add(ctx context.Context, req *mqtt.MqttStatusAddReq) (err error) {
	_, err = dao.MqttStatus.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *mqttStatusImpl) Edit(ctx context.Context, req *mqtt.MqttStatusEditReq) error {
	_, err := dao.MqttStatus.Ctx(ctx).FieldsEx(dao.MqttStatus.Columns().Id, dao.MqttStatus.Columns().CreatedAt).Where(dao.MqttStatus.Columns().Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *mqttStatusImpl) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.MqttStatus.Ctx(ctx).Delete(dao.MqttStatus.Columns().Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("删除失败")
	}
	return
}

// ChangeStatus 修改状态
func (s *mqttStatusImpl) ChangeStatus(ctx context.Context, req *mqtt.MqttStatusStatusReq) error {
	_, err := dao.MqttStatus.Ctx(ctx).WherePri(req.Id).Update(g.Map{
		dao.MqttStatus.Columns().Status: req.Status,
	})
	return err
}