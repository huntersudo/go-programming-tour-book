package cmd
// todo https://golang2.eddycjy.com/posts/ch1/02-word2word/
import (
	"log"
	"strings"

	"github.com/go-programming-tour-book/tour/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下线线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

var str string
var mode int8
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

// Use：子命令的命令标识。
// Short：简短说明，在 help 输出的帮助信息中展示。
// Long：完整说明，在 help 输出的帮助信息中展示。
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

// ```
//[going@dev tour]$ go run main.go help word
//该子命令支持各种单词格式转换，模式如下：
//1：全部转大写
//2：全部转小写
//3：下划线转大写驼峰
//4：下划线转小写驼峰
//5：驼峰转下划线
//
//Usage:
//   word [flags]
//
//Flags:
//  -h, --help         help for word
//  -m, --mode int8    请输入单词转换的模式
//  -s, --str string   请输入单词内容
//
//```

//[going@dev tour]$ go run main.go word -s=eddycjy -m=1
//2021/07/12 10:35:55 输出结果: EDDYCJY
//[going@dev tour]$ go run main.go word -s=EDDYCJY -m=2
//2021/07/12 10:36:04 输出结果: eddycjy
//[going@dev tour]$ go run main.go word -s=EDDYCJY -m=3
//2021/07/12 10:36:09 输出结果: EDDYCJY
//[going@dev tour]$ go run main.go word -s=EDDYCJY -m=4
//2021/07/12 10:36:12 输出结果: eDDYCJY
//[going@dev tour]$ go run main.go word -s=EDDYCJY -m=5
//2021/07/12 10:36:18 输出结果: e_d_d_y_c_j_y
//[going@dev tour]$ go run main.go word -s=EddyCjy -m=5
//2021/07/12 10:36:34 输出结果: eddy_cjy