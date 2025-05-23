package carbon

var (
	// DefaultLayout default layout
	// 默认布局模板
	DefaultLayout = DateTimeLayout

	// DefaultTimezone default timezone
	// 默认时区
	DefaultTimezone = UTC

	// DefaultLocale default language locale
	// 默认语言区域
	DefaultLocale = "en"

	// DefaultWeekStartsAt default start date of the week
	// 默认一周开始日期
	DefaultWeekStartsAt = Monday

	// DefaultWeekendDays default weekend days of the week
	// 默认周末日期
	DefaultWeekendDays = []Weekday{
		Saturday, Sunday,
	}
)

// Default defines a Default struct.
// 定义 Default 结构体
type Default struct {
	Layout       string
	Timezone     string
	Locale       string
	WeekStartsAt Weekday
	WeekendDays  []Weekday
}

// SetDefault sets default.
// 设置全局默认值
func SetDefault(d Default) {
	if d.Layout != "" {
		DefaultLayout = d.Layout
	}
	if d.Timezone != "" {
		DefaultTimezone = d.Timezone
	}
	if d.Locale != "" {
		DefaultLocale = d.Locale
	}
	if d.WeekStartsAt.String() != "" {
		DefaultWeekStartsAt = d.WeekStartsAt
	}
	if len(d.WeekendDays) > 0 {
		DefaultWeekendDays = d.WeekendDays
	}
}

// ResetDefault resets default.
// 重置全局默认值
func ResetDefault() {
	DefaultLayout = DateTimeLayout
	DefaultTimezone = UTC
	DefaultLocale = "en"
	DefaultWeekStartsAt = Monday
	DefaultWeekendDays = []Weekday{
		Saturday, Sunday,
	}
}
