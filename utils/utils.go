package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func MkdirAll(paths ...string) (err error) {
	if len(paths) == 0 {
		return
	}
	p := path.Join(paths...)
	if _, err = os.Stat(p); err != nil {
		err = os.MkdirAll(p, os.ModePerm)
	}
	return
}

//saveBaseDir := "./game-source"
//gameName := "MarshmallowRush"
//downloadDir := path.Join("/Users/xiaoliangcheng/Downloads/www.datinginfo.top/game", gameName)
//p := path.Join(saveBaseDir, gameName)
// %20 特殊符号转回空格
//fileRename(p)
// 复制已下载的游戏文件到指定目录
//copyDownloadGame(downloadDir, p)
func copyDownloadGame(downloadDir, p string) {
	command := exec.Command("cp", "-r", downloadDir, p)
	rs, err := command.CombinedOutput()

	// 下载的文件迁移
	fmt.Println(p, err, string(rs))
}

func FileRename(baseDir string) {
	_ = filepath.Walk(baseDir, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.Contains(p, "%20") {
			doRename(p)
		}
		return nil
	})
}

func doRename(f string) {
	// 重命名文件
	nf := strings.ReplaceAll(f, "%20", " ")
	err := os.Rename(f, nf)
	if err != nil {
		fmt.Println(f, err)
	} else {
		fmt.Println(f, nf)
	}
}

// 按错误日志下载的资源文件
func downloadFiles() {

}

func MD5(params string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(params))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
