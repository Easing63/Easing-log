/*
Package logch
@Time: 2022/9/23 15:29
@Description: simple log
*/
package logch

import "github.com/sirupsen/logrus"

// MsgOut 首先日志得能够以文件的方式输出
func MsgOut(fileName string, filed...string) {
	// 设置日志级别
	logrus.SetLevel(logrus.InfoLevel)
	// 添加字段
	logrus.WithFields(logrus.Fields{

	})
}
