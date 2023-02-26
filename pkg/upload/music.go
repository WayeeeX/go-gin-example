package upload

import (
	"github.com/WayeeeX/go-gin-example/pkg/file"
	"github.com/WayeeeX/go-gin-example/pkg/logging"
	"github.com/WayeeeX/go-gin-example/pkg/setting"
	"log"
	"mime/multipart"
	"strings"
)

func GetMusicFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetMusicPath() + name
}

// GetMusicPath get save path
func GetMusicPath() string {
	return setting.AppSetting.MusicSavePath
}

// GetMusicFullPath get full save path
func GetMusicFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetMusicPath()
}
func CheckMusicExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.MusicAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckMusicSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.MusicMaxSize
}
