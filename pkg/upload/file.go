package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"myblog/global"
	"myblog/pkg/util"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimPrefix(name,ext)
	fileName = util.EncodeMd5(fileName)
	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(dst string) bool {
	_,err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckContainExt(t FileType,name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeImage:
		for _,allowExt := range global.AppSetting.UploadImageAllowExts{
			if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(t FileType,f multipart.File) bool {
	content,_ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
		case TypeImage:
			if size >= global.AppSetting.UploadImageMaxSize {
				return true
			}
	}
	return false
}

func CheckPermission(dst string) bool{
	_,err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string,perm os.FileMode) error {
	err := os.MkdirAll(dst,perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader,dst string) error {
	src,err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out,err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_,err = io.Copy(out,src)
	return err

}