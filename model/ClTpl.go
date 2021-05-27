package model

import (
	"go-studio/core/errorx"
	"go-studio/core/jsontime"
	"xorm.io/xorm"
)

type ClTplModel struct {
	db *xorm.Engine
}

type ClTpl struct {
	Id        int64             `xorm:"not null pk autoincr BIGINT(20)"`
	UserId    string            `xorm:"not null default '' comment('用户id') index VARCHAR(36)"`
	Name      string            `xorm:"not null default '' comment('模板名称') VARCHAR(255)"`
	CreatedAt jsontime.JsonTime `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	UpdatedAt jsontime.JsonTime `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func NewClTplModel(db *xorm.Engine) *ClTplModel{

	return
}

func (m *ClTplModel) Insert(data *ClTpl) error {
	affect, err := m.db.InsertOne(data)
	if err != nil {
		return err
	}
	if affect != 1 {
		return errorx.DBUpdateNotAffected
	}

	return nil
}

func (m *ClTplModel) FindOne(id int64) (*ClTpl, error) {
	var resp ClTpl
	has, err := m.db.ID(id).Get(&resp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errorx.DBDataNotFound
	}
	return &resp, nil

}

func (m *ClTplModel) Update(data *ClTpl) error {
	affect, err := m.db.ID(data.Id).Update(data)
	if err != nil {
		return err
	}

	if affect != 1 {
		return errorx.DBUpdateNotAffected
	}
	return nil
}

func (m *ClTplModel) Delete(id int64) error {
	affect, err := m.db.ID(id).Delete(ClTpl{})
	if err != nil {
		return err
	}
	if affect != 1 {
		return errorx.DBUpdateNotAffected
	}
	return nil
}
