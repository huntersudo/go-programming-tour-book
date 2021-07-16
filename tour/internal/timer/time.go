package timer

import "time"

// GetNowTime 而在 Go 语言中使用 Location 来表示地区相关的时区，一个 Location 可能表示多个时区。
// 在标准库 time 上，提供了 Location 的两个实例：Local 和 UTC。Local 代表当前系统本地时区；
//UTC 代表通用协调时间，也就是零时区，在默认值上，标准库 time 使用的是 UTC 时区。
// 时区信息既浩繁又多变，Unix 系统以标准格式存于文件中，这些文件位于 /usr/share/zoneinfo，
//而本地时区可以通过 /etc/localtime 获取，这是一个符号链接，指向 /usr/share/zoneinfo 中某一个时区。
//比如我本地电脑指向的是：/var/db/timezone/zoneinfo/Asia/Shanghai。

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// GetCalculateTime 时间推算
// ParseDuration 方法用于在字符串中解析出 duration（持续时间），其支持的有效单位有"ns”, “us” (or “µ s”), “ms”, “s”, “m”, “h”，
// 例如：“300ms”, “-1.5h” or “2h45m”。而在 Add 方法中，我们可以将其返回的 duration 传入，就可以得到当前 timer 时间加上 duration 后所得到的最终时间
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), nil
}
// const (
//	Nanosecond  Duration = 1
//	Microsecond          = 1000 * Nanosecond
//	Millisecond          = 1000 * Microsecond
//	Second               = 1000 * Millisecond
//	Minute               = 60 * Second
//	Hour                 = 60 * Minute
//)
//
//...
//timer.GetNowTime().Add(time.Second * 60)