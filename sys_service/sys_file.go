// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type (
	IFile interface {
		InstallHook(state sys_enum.UploadEventState, hookFunc sys_model.FileHookFunc) int64
		UnInstallHook(savedHookId int64)
		CleanAllHook()
		Upload(ctx context.Context, in sys_model.FileUploadInput, userId int64) (*sys_model.FileUploadOutput, error)
		GetUploadFile(ctx context.Context, uploadId int64, userId int64, message ...string) (*sys_model.FileUploadOutput, error)
		SaveFile(ctx context.Context, storageAddr string, userId int64, info sys_model.FileUploadOutput) (*sys_entity.SysFile, error)
		UploadIDCard(ctx context.Context, in sys_model.OCRIDCardFileUploadInput, userId int64) (*sys_model.IDCardWithOCR, error)
		UploadBankCard(ctx context.Context, in sys_model.BankCardWithOCRInput, userId int64) (*sys_model.BankCardWithOCR, error)
		UploadBusinessLicense(ctx context.Context, in sys_model.OCRBusinessLicense, userId int64) (*sys_model.BusinessLicenseWithOCR, error)
		DownLoadFile(ctx context.Context, savePath string, url string) (string, error)
	}
)

var (
	localFile IFile
)

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
