package upload

import (
	"github.com/WayeeeX/go-gin-example/pkg/file"
	"github.com/WayeeeX/go-gin-example/pkg/logging"
	"github.com/WayeeeX/go-gin-example/pkg/setting"
	"log"
	"mime/multipart"
	"strings"
)

func GetLyricFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetLyricPath() + name
}

// GetLyricPath get save path
func GetLyricPath() string {
	return setting.AppSetting.LyricSavePath
}

// GetLyricFullPath get full save path
func GetLyricFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetLyricPath()
}
func CheckLyricExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.LyricAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckLyricSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.LyricMaxSize
}
