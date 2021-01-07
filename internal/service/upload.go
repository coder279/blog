package service

import (
	"errors"
	"mime/multipart"
	"myblog/global"
	"myblog/pkg/upload"
	"os"
)

type FileInfo struct {
	Name string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType,file multipart.File,fileheader *multipart.FileHeader) (*FileInfo,error){
	fileName := upload.GetFileName(fileheader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" +fileName
	if !upload.CheckContainExt(fileType,fileName) {
		return nil,errors.New("file suffix is not supported.")
	}
	if upload.CheckSavePath(uploadSavePath){
		if err := upload.CreateSavePath(uploadSavePath,os.ModePerm);err != nil {
			return nil,errors.New("failed to create save directory")
		}
	}
	if upload.CheckMaxSize(fileType,file){
		return nil,errors.New("exceeded maximum file limit")
	}
	if upload.CheckPermission(uploadSavePath){
		return nil,errors.New("insufficient file permissions")
	}
	if err := upload.SaveFile(fileheader,dst);err!=nil {
		return nil,err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name:fileName,AccessUrl:accessUrl},nil
}
