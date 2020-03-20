/*
 * Author: yangzhao
 * CreatedAt: 2020/3/20
 * Description: error
 */

package error

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//打印错误消息并终止程序运行
func FatalfIfErr(err error, msg string) {
	f, _ := exec.LookPath(os.Args[0])
	fileWithPath, _ := filepath.Abs(f)

	if err != nil {
		log.Fatalf("%s[error]%s %s[%s]%s %s%s%s：\n%s%v%s",
			redBg,
			reset,
			cyanBg,
			fileWithPath,
			reset,
			redBg,
			msg,
			reset,
			red,
			err,
			reset)
	}
}
