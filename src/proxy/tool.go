package proxy

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
)

/**
生成MD5字符传
*/
func MakeMD5(sourceStr string) string {
	h := md5.New()
	h.Write([]byte(sourceStr))
	return hex.EncodeToString(h.Sum(nil))
}

/**
生成日志工具
@param basePath目录地址
*/
func MakeLogger(basePath string, fileName string) (*log.Logger, *os.File, error) {
	//判断目录是否存在，否则就创建
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		//目录不存在，则创建
		mkdirErr := os.MkdirAll(basePath, 0774)
		if nil != mkdirErr {
			return nil, nil, errors.New(fmt.Sprintf("生成失败,%s路径创建失败", basePath))
		}
	}
	exist := true
	filePath := fmt.Sprintf("%s%s.log", basePath, fileName) //生成文件名
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false //判断文件是否存在
	}
	var f *os.File
	var fe error
	if exist {
		//如果文件存在则更新内容
		f, fe = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0774)
	} else {
		//创建文件，写入内容
		f, fe = os.Create(filePath)
	}
	if nil != fe {
		return nil, nil, errors.New(fmt.Sprintf("生成失败，原因是文件写入失败,%v", fe))
	} else {
		logger := log.New(f, "", log.Llongfile)
		logger.SetFlags(log.LstdFlags) // 设置写入文件的log日志的格式
		return logger, f, nil
	}
}

/**
生成日志工具
@param basePath目录地址
*/
func SaveFileToLocal(basePath string, fileName string, data []byte) (string, error) {
	//判断目录是否存在，否则就创建
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		//目录不存在，则创建
		mkdirErr := os.MkdirAll(basePath, 0774)
		if nil != mkdirErr {
			return "", errors.New(fmt.Sprintf("生成失败,%s路径创建失败", basePath))
		}
	}
	exist := true
	filePath := fmt.Sprintf("%s%s", basePath, fileName) //生成文件名
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false //判断文件是否存在
	}
	var f *os.File
	var fe error
	if exist {
		//如果文件存在则更新内容
		f, fe = os.OpenFile(filePath, os.O_WRONLY, 0774)
	} else {
		//创建文件，写入内容
		f, fe = os.Create(filePath)
	}
	if nil != fe {
		return "", errors.New(fmt.Sprintf("生成失败，原因是文件写入失败,%v", fe))
	} else {
		_, err := f.Write(data)
		if nil == err {
			return filePath, nil
		} else {
			return "", errors.New(fmt.Sprintf("生成失败，原因是文件写入失败,%v", fe))
		}
	}
}

/**
读取本地文件
*/
func ReadLocalFile(basePath string, fileName string) (*os.File, error) {
	exist := true
	var f *os.File
	var fe error
	filePath := fmt.Sprintf("%s%s", basePath, fileName) //文件完整路径
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false //判断文件是否存在
		fe = errors.New(fmt.Sprintf("文件读取失败,%v", err))
	}

	if exist == true {
		//如果文件存在则更新内容
		f, fe = os.OpenFile(filePath, os.O_RDONLY, 0774)
		if nil != fe {
			fe = errors.New(fmt.Sprintf("文件读取失败,%v", fe))
		}
	}
	return f, fe
}

/**
进行base64编码
*/
func EncodeBase64(input []byte) string {
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

/**
解码base64
*/
func DecodeBase64(encodeString string) []byte {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		fmt.Println("DecodeBase64失败，原因:", err.Error())
	}
	return decodeBytes
}
