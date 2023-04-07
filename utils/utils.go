package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
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

//AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

//pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}
