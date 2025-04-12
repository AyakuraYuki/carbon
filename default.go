package carbon

var (
	// DefaultLayout default layout
	// 默认布局模板
	DefaultLayout = DateTimeLayout

	// DefaultTimezone default timezone
	// 默认时区
	DefaultTimezone = UTC

	// DefaultWeekStartsAt default start date of the week
	// 默认一周开始日期
	DefaultWeekStartsAt = Monday

	// DefaultWeekendDays Default weekend days
	// 默认周末日期
	DefaultWeekendDays = []string{
		Saturday, Sunday,
	}

	// DefaultLocale default language locale
	// 默认语言区域
	DefaultLocale = "en"
)

// Default defines a Default struct.
// 定义 Default 结构体
type Default struct {
	Layout       string
	Timezone     string
	WeekStartsAt string
	WeekendDays  []string
	Locale       string
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
	if d.WeekStartsAt != "" {
		DefaultWeekStartsAt = d.WeekStartsAt
	}
	if len(d.WeekendDays) > 0 {
		DefaultWeekendDays = d.WeekendDays
	}
	if d.Locale != "" {
		DefaultLocale = d.Locale
	}
}

// ResetDefault resets default.
// 重置全局默认值
func ResetDefault() {
	DefaultLayout = DateTimeLayout
	DefaultTimezone = UTC
	DefaultWeekStartsAt = Sunday
	DefaultWeekendDays = []string{
		Saturday, Sunday,
	}
	DefaultLocale = "en"
}
