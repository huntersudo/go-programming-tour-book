package cmd
// https://golang2.eddycjy.com/posts/ch1/03-time2format/
import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-programming-tour-book/tour/internal/timer"

	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// 第一个格式：通过调用 Format 方法设定约定的 2006-01-02 15:04:05 格式来进行时间的标准格式化。
// 第二个格式：通过调用 Unix 方法返回 Unix 时间，就是我们通俗说的时间戳，其值为自 UTC 1970 年 1 月 1 日起经过的秒数。
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTime string
var duration string
// 在上述代码中，一共支持了三种常用时间格式的处理，分别是：时间戳、2006-01-02 以及 2006-01-02 15:04:05。
// 在时间格式处理上，我们调用了 strings.Contains 方法，对空格进行了包含判断，
//若存在则按既定的 2006-01-02 15:04:05 格式进行格式化，否则以 2006-01-02 格式进行处理，若出现异常错误，则直接按时间戳的方式进行转换处理。
var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	// time 子命令
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间，有效单位为时间戳或已格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}

// [going@dev tour]$ go run main.go time now
//2021/07/12 11:08:00 输出结果: 2021-07-12 11:08:00, 1626059280
//[going@dev tour]$ go run main.go time calc -c="2029-09-04 12:02:33" -d=5m
//2021/07/12 11:08:17 输出结果: 2029-09-04 12:07:33, 1883218053
//[going@dev tour]$ go run main.go time calc -c="2029-09-04 12:02:33" -d=-2h
//2021/07/12 11:08:30 输出结果: 2029-09-04 10:02:33, 1883210553