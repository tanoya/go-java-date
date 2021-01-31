package date

import (
	"strings"
	"time"
)

// FormaterEngine 时间格式化方式
type FormaterEngine int

// 目前支持的格式化方式
const (
	Go FormaterEngine = iota
	Java
)

func (fe *FormaterEngine) String(formater string) string {
	if fe == nil {
		return ""
	}

	var adapter Adapter
	switch *fe {
	case Java:
		adapter = new(JavaAdapter)
	case Go:
		adapter = new(GoAdapter)
	}
	return adapter.String(formater)

}

// Adapter 适配器，主要是进行格式化方式的适配
type Adapter interface {
	String(string) string
}

// JavaAdapter Java 适配器
type JavaAdapter int

// yyyy-MM-dd HH:mm:ss
// 暂时支持这几个字符
// Kv 保证有序
var kk = []string{"yyyy", "yy", "MM", "M", "dd", "d", "HH", "H", "mm", "m", "SSS", "ss", "s"}
var vv = []string{"2006", "06", "01", "1", "02", "2", "15", "3", "04", "4", ".000", "05", "5"}

// String 负责将Java时间方式格式化为Go的方式
func (ja *JavaAdapter) String(formater string) string {
	if formater == "" {
		return ""
	}

	for i, val := range kk {
		if strings.Contains(formater, val) {
			formater = strings.Replace(formater, val, vv[i], 1)
		}
	}
	return formater
}

// GoAdapter Go适配器
type GoAdapter int

// String go 自己本身支持的不需要实现
func (ga *GoAdapter) String(formater string) string {
	return formater
}

// Date 使用Java版本的format格式化方式
type Date struct {
	DT     *time.Time
	Engine FormaterEngine
}

// New 创建
func New() *Date {
	t := time.Now()
	d := Date{
		DT:     &t,
		Engine: Go, // 默认创建Go格式话的方式
	}
	return &d
}

// ParseJava 将Java字符串的日期转换为 Date 类型
func ParseJava(date, formater string) *Date {
	ret := Date{Engine: Java}
	formater = ret.Engine.String(formater)
	t, e := time.ParseInLocation(formater, date, time.Local)
	if e != nil {
		return nil
	}
	ret.DT = &t
	return &ret
}

// ParseJavaLocation 支持时区
// location 参考 zoneinfo_abbrs_windows.go 文件注释部分
// example Asia/Shanghai
func ParseJavaLocation(date, formater string, loc *time.Location) *Date {
	ret := Date{Engine: Java}
	formater = ret.Java().String(formater)
	t, e := time.ParseInLocation(formater, date, loc)
	if e != nil {
		return nil
	}
	ret.DT = &t
	return &ret
}

// Parse 将Go字符串转换为 Date 类型
func Parse(date, formater string) *Date {
	ret := Date{Engine: Go}
	t, e := time.Parse(formater, date)
	if e != nil {
		return nil
	}
	ret.DT = &t
	return &ret
}

// ParseLocation 支持时区
// location 参考 zoneinfo_abbrs_windows.go 文件注释部分
// example Asia/Shanghai
func ParseLocation(date, formater string, loc *time.Location) *Date {
	ret := Date{Engine: Go}
	t, e := time.ParseInLocation(formater, date, loc)
	if e != nil {
		return nil
	}
	ret.DT = &t
	return &ret

}

// Java 需要解析 Java 版本的格式
func (d *Date) Java() *Date {
	if d == nil {
		return d
	}
	d.Engine = Java
	return d
}

// String 将日期格式化输出
func (d *Date) String(formater string) string {
	if d == nil {
		return ""
	}
	s := d.Engine.String(formater)
	return d.DT.Format(s)
}

// ------------------ 操作部分 --------------------

// Add 时间增加
// du 操作的时间单元
// offset 时间偏移量 必须为正数
func (d *Date) Add(du time.Duration, offset int64) {
	if d == nil || offset < 0 {
		return
	}
	tmp := d.DT.Add(time.Duration(offset) * du)
	d.DT = &tmp
}

// Minus 时间递减
// du 操作的时间单元
// offset 时间偏移量 必须为正数
func (d *Date) Minus(du time.Duration, offset int64) {
	if d == nil || offset < 0 {
		return
	}
	tmp := d.DT.Add(-time.Duration(offset) * du)
	d.DT = &tmp
}
