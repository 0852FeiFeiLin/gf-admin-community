// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/entity"
)

type (
	IFdInvoiceDetail interface {
		CreateInvoiceDetail(ctx context.Context, info model.FdInvoiceDetailRegister) (*entity.FdInvoiceDetail, error)
		GetInvoiceDetailById(ctx context.Context, id int64) (*entity.FdInvoiceDetail, error)
		UpdateInvoiceDetail(ctx context.Context, info entity.FdInvoiceDetail) (bool, error)
		GetInvoiceDetailList(ctx context.Context, info *model.SearchParams, isExport bool) (*model.FdInvoiceDetailListRes, error)
	}
)

var (
	localFdInvoiceDetail IFdInvoiceDetail
)

func FdInvoiceDetail() IFdInvoiceDetail {
	if localFdInvoiceDetail == nil {
		panic("implement not found for interface IFdInvoiceDetail, forgot register?")
	}
	return localFdInvoiceDetail
}

func RegisterFdInvoiceDetail(i IFdInvoiceDetail) {
	localFdInvoiceDetail = i
}
