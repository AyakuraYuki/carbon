package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, Parse(""))
		assert.Empty(t, Parse("").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.True(t, Parse("0").HasError())
		assert.True(t, Parse("xxx").HasError())

		assert.Empty(t, Parse("0").ToString())
		assert.Empty(t, Parse("xxx").ToString())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05", "").HasError())
		assert.True(t, Parse("2020-08-05", "0").HasError())
		assert.True(t, Parse("2020-08-05", "xxx").HasError())

		assert.Empty(t, Parse("2020-08-05", "").ToString())
		assert.Empty(t, Parse("2020-08-05", "0").ToString())
		assert.Empty(t, Parse("2020-08-05", "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, Now().Timestamp(), Parse("now").Timestamp())
		assert.Equal(t, Yesterday().Timestamp(), Parse("yesterday").Timestamp())
		assert.Equal(t, Tomorrow().Timestamp(), Parse("tomorrow").Timestamp())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-5").ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-05").ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-8-5 1:2:3").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:2:03").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:02:03").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 01:02:03").ToString())
		assert.Equal(t, Parse("2022-03-08T03:01:14-07:00").ToString(), Parse("2022-03-08T10:01:14Z").ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, Now().Timestamp(), Parse("now", PRC).Timestamp())
		assert.Equal(t, Yesterday().Timestamp(), Parse("yesterday", PRC).Timestamp())
		assert.Equal(t, Tomorrow().Timestamp(), Parse("tomorrow", PRC).Timestamp())
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", Parse("2020-8-5", PRC).ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", Parse("2020-8-05", PRC).ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", Parse("2020-08-05", PRC).ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-8-5 1:2:3", PRC).ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 1:2:03", PRC).ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 1:02:03", PRC).ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 01:02:03", PRC).ToString())
		assert.Equal(t, Parse("2022-03-08T03:01:14-07:00", PRC).ToString(), Parse("2022-03-08T10:01:14Z", PRC).ToString())
	})

	// https://github.com/dromara/carbon/issues/202
	t.Run("issue202", func(t *testing.T) {
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-08T09:02:48").ToString())
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-1-8T09:02:48").ToString())
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-08T9:2:48").ToString())
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-8T9:2:48").ToString())
	})

	// https://github.com/dromara/carbon/issues/232
	t.Run("issue232", func(t *testing.T) {
		assert.Equal(t, "0000-01-01 00:00:00 +0000 UTC", Parse("0000-01-01 00:00:00").ToString())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Parse("0001-01-01 00:00:00").ToString())
		assert.Equal(t, "", Parse("0001-00-00 00:00:00").ToString())
	})
}

func TestParseByFormat(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, ParseByFormat("", DateFormat))
		assert.Empty(t, ParseByFormat("", DateFormat).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.True(t, ParseByFormat("0", DateFormat).HasError())
		assert.True(t, ParseByFormat("xxx", DateFormat, PRC).HasError())

		assert.Empty(t, ParseByFormat("0", DateFormat).ToString())
		assert.Empty(t, ParseByFormat("xxx", DateFormat, PRC).ToString())
	})

	t.Run("empty format", func(t *testing.T) {
		assert.True(t, ParseByFormat("2020-08-05", "").HasError())
		assert.Empty(t, ParseByFormat("2020-08-05", "").ToString())

	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, ParseByFormat("2020-08-05", DateFormat, "").HasError())
		assert.True(t, ParseByFormat("2020-08-05", DateFormat, "xxx").HasError())

		assert.Empty(t, ParseByFormat("2020-08-05", DateFormat, "").ToString())
		assert.Empty(t, ParseByFormat("2020-08-05", DateFormat, "xxx").ToString())
	})

	t.Run("invalid timestamp", func(t *testing.T) {
		assert.True(t, ParseByFormat("2020-08-05", TimestampFormat).HasError())
		assert.True(t, ParseByFormat("2020-08-05", TimestampMilliFormat).HasError())
		assert.True(t, ParseByFormat("2020-08-05", TimestampMicroFormat).HasError())
		assert.True(t, ParseByFormat("2020-08-05", TimestampNanoFormat).HasError())

		assert.Empty(t, ParseByFormat("2020-08-05", TimestampFormat).ToString())
		assert.Empty(t, ParseByFormat("2020-08-05", TimestampMilliFormat).ToString())
		assert.Empty(t, ParseByFormat("2020-08-05", TimestampMicroFormat).ToString())
		assert.Empty(t, ParseByFormat("2020-08-05", TimestampNanoFormat).ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", ParseByFormat("2020-08-05", DateFormat).ToString())
		assert.Equal(t, "0000-01-01 13:14:15 +0000 UTC", ParseByFormat("13:14:15", TimeFormat).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020-08-05 13:14:15", DateTimeFormat).ToString())

		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", ParseByFormat("2020-08-05", DateFormat, PRC).ToString())
		assert.Equal(t, "0000-01-01 13:14:15 +0805 LMT", ParseByFormat("13:14:15", TimeFormat, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("2020-08-05 13:14:15", DateTimeFormat, PRC).ToString())

		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s", PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s", PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒", PRC).ToString())
	})

	// https://github.com/dromara/carbon/issues/206
	t.Run("issue206", func(t *testing.T) {
		assert.Equal(t, "2023-11-11 04:34:00 +0000 UTC", ParseByFormat("1699677240", TimestampFormat).ToString())
		assert.Equal(t, "2023-11-11 04:34:00.666 +0000 UTC", ParseByFormat("1699677240666", TimestampMilliFormat).ToString())
		assert.Equal(t, "2023-11-11 04:34:00.666666 +0000 UTC", ParseByFormat("1699677240666666", TimestampMicroFormat).ToString())
		assert.Equal(t, "2023-11-11 04:34:00.666666666 +0000 UTC", ParseByFormat("1699677240666666666", TimestampNanoFormat).ToString())
	})
}

func TestParseByLayout(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, ParseByLayout("", DateLayout))
		assert.Empty(t, ParseByLayout("", DateLayout).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.True(t, ParseByLayout("0", DateFormat).HasError())
		assert.True(t, ParseByLayout("xxx", DateFormat, PRC).HasError())

		assert.Empty(t, ParseByLayout("0", DateFormat).ToString())
		assert.Empty(t, ParseByLayout("xxx", DateFormat, PRC).ToString())
	})

	t.Run("empty layout", func(t *testing.T) {
		assert.True(t, ParseByLayout("2020-08-05", "").HasError())
		assert.Empty(t, ParseByLayout("2020-08-05", "").ToString())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, ParseByLayout("2020-08-05", DateLayout, "").HasError())
		assert.True(t, ParseByLayout("2020-08-05", DateLayout, "xxx").HasError())

		assert.Empty(t, ParseByLayout("2020-08-05", DateLayout, "").ToString())
		assert.Empty(t, ParseByLayout("2020-08-05", DateLayout, "xxx").ToString())
	})

	t.Run("invalid timestamp", func(t *testing.T) {
		assert.True(t, ParseByLayout("2020-08-05", TimestampLayout).HasError())
		assert.True(t, ParseByLayout("2020-08-05", TimestampMilliLayout).HasError())
		assert.True(t, ParseByLayout("2020-08-05", TimestampMicroLayout).HasError())
		assert.True(t, ParseByLayout("2020-08-05", TimestampNanoLayout).HasError())

		assert.Empty(t, ParseByLayout("2020-08-05", TimestampLayout).ToString())
		assert.Empty(t, ParseByLayout("2020-08-05", TimestampMilliLayout).ToString())
		assert.Empty(t, ParseByLayout("2020-08-05", TimestampMicroLayout).ToString())
		assert.Empty(t, ParseByLayout("2020-08-05", TimestampNanoLayout).ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", ParseByLayout("2020-08-05", DateLayout).ToString())
		assert.Equal(t, "0000-01-01 13:14:15 +0000 UTC", ParseByLayout("13:14:15", TimeLayout).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout).ToString())

		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", ParseByLayout("2020-08-05", DateLayout, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout, PRC).ToString())

		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s 2006-01-02 15:04:05", PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05", PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒", PRC).ToString())
	})
}

func TestParseWithLayouts(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, ParseWithLayouts("", []string{DateTimeLayout}))
		assert.Empty(t, ParseWithLayouts("", []string{DateTimeLayout}).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.True(t, ParseWithLayouts("0", []string{DateTimeLayout}).HasError())
		assert.True(t, ParseWithLayouts("xxx", []string{DateTimeLayout}).HasError())

		assert.Empty(t, ParseWithLayouts("0", []string{DateTimeLayout}).ToString())
		assert.Empty(t, ParseWithLayouts("xxx", []string{DateTimeLayout}).ToString())
	})

	t.Run("empty layouts", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseWithLayouts("2020-08-05 13:14:15", []string{}).ToString())
		assert.Equal(t, "2006-01-02 15:04:05", ParseWithLayouts("2020-08-05 13:14:15", []string{}).CurrentLayout())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, ParseWithLayouts("2020-08-05 13:14:15", []string{DateLayout}, "").HasError())
		assert.True(t, ParseWithLayouts("2020-08-05 13:14:15", []string{DateLayout}, "xxx").HasError())

		assert.Empty(t, ParseWithLayouts("2020-08-05 13:14:15", []string{DateLayout}, "").ToString())
		assert.Empty(t, ParseWithLayouts("2020-08-05 13:14:15", []string{DateLayout}, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		c := ParseWithLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"})
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c.ToString())
		assert.Equal(t, "2006|01|02 15|04|05", c.CurrentLayout())
	})

	t.Run("with timezone", func(t *testing.T) {
		c := ParseWithLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"}, PRC)
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c.ToString())
		assert.Equal(t, "2006|01|02 15|04|05", c.CurrentLayout())
	})
}

func TestParseWithFormats(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, ParseWithFormats("", []string{DateTimeFormat}))
		assert.Empty(t, ParseWithFormats("", []string{DateTimeFormat}).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.True(t, ParseWithFormats("0", []string{DateTimeLayout}).HasError())
		assert.True(t, ParseWithFormats("xxx", []string{DateTimeLayout}).HasError())

		assert.Empty(t, ParseWithFormats("0", []string{DateTimeLayout}).ToString())
		assert.Empty(t, ParseWithFormats("xxx", []string{DateTimeLayout}).ToString())
	})

	t.Run("empty layouts", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseWithFormats("2020-08-05 13:14:15", []string{}).ToString())
		assert.Equal(t, "2006-01-02 15:04:05", ParseWithFormats("2020-08-05 13:14:15", []string{}).CurrentLayout())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, ParseWithFormats("2020-08-05 13:14:15", []string{DateFormat}, "").HasError())
		assert.True(t, ParseWithFormats("2020-08-05 13:14:15", []string{DateFormat}, "xxx").HasError())

		assert.Empty(t, ParseWithFormats("2020-08-05 13:14:15", []string{DateFormat}, "").ToString())
		assert.Empty(t, ParseWithFormats("2020-08-05 13:14:15", []string{DateFormat}, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		c := ParseWithFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"})
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c.ToString())
		assert.Equal(t, "2006|01|02 15|04|05", c.CurrentLayout())
	})

	t.Run("with timezone", func(t *testing.T) {
		c := ParseWithFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"}, PRC)
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c.ToString())
		assert.Equal(t, "2006|01|02 15|04|05", c.CurrentLayout())
	})
}

func TestFormat2Layout(t *testing.T) {
	c := Parse("2020-08-05 13:14:15.999999999")

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, c.Layout(AtomLayout), c.Format(AtomFormat))
		assert.Equal(t, c.Layout(ANSICLayout), c.Format(ANSICFormat))
		assert.Equal(t, c.Layout(CookieLayout), c.Format(CookieFormat))
		assert.Equal(t, c.Layout(KitchenLayout), c.Format(KitchenFormat))
		assert.Equal(t, c.Layout(RssLayout), c.Format(RssFormat))
		assert.Equal(t, c.Layout(RubyDateLayout), c.Format(RubyDateFormat))
		assert.Equal(t, c.Layout(UnixDateLayout), c.Format(UnixDateFormat))

		assert.Equal(t, c.Layout(RFC1036Layout), c.Format(RFC1036Format))
		assert.Equal(t, c.Layout(RFC1123Layout), c.Format(RFC1123Format))
		assert.Equal(t, c.Layout(RFC1123ZLayout), c.Format(RFC1123ZFormat))
		assert.Equal(t, c.Layout(RFC2822Layout), c.Format(RFC2822Format))
		assert.Equal(t, c.Layout(RFC3339Layout), c.Format(RFC3339Format))
		assert.Equal(t, c.Layout(RFC3339MilliLayout), c.Format(RFC3339MilliFormat))
		assert.Equal(t, c.Layout(RFC3339MicroLayout), c.Format(RFC3339MicroFormat))
		assert.Equal(t, c.Layout(RFC3339NanoLayout), c.Format(RFC3339NanoFormat))
		assert.Equal(t, c.Layout(RFC7231Layout), c.Format(RFC7231Format))
		assert.Equal(t, c.Layout(RFC822Layout), c.Format(RFC822Format))
		assert.Equal(t, c.Layout(RFC822ZLayout), c.Format(RFC822ZFormat))
		assert.Equal(t, c.Layout(RFC850Layout), c.Format(RFC850Format))

		assert.Equal(t, c.Layout(ISO8601Layout), c.Format(ISO8601Format))
		assert.Equal(t, c.Layout(ISO8601MilliLayout), c.Format(ISO8601MilliFormat))
		assert.Equal(t, c.Layout(ISO8601MicroLayout), c.Format(ISO8601MicroFormat))
		assert.Equal(t, c.Layout(ISO8601NanoLayout), c.Format(ISO8601NanoFormat))

		assert.Equal(t, c.Layout(ISO8601ZuluLayout), c.Format(ISO8601ZuluFormat))
		assert.Equal(t, c.Layout(ISO8601ZuluMilliLayout), c.Format(ISO8601ZuluMilliFormat))
		assert.Equal(t, c.Layout(ISO8601ZuluMicroLayout), c.Format(ISO8601ZuluMicroFormat))
		assert.Equal(t, c.Layout(ISO8601ZuluNanoLayout), c.Format(ISO8601ZuluNanoFormat))

		assert.Equal(t, c.Layout(FormattedDateLayout), c.Format(FormattedDateFormat))
		assert.Equal(t, c.Layout(FormattedDayDateLayout), c.Format(FormattedDayDateFormat))

		assert.Equal(t, c.Layout(DayDateTimeLayout), c.Format(DayDateTimeFormat))
		assert.Equal(t, c.Layout(DateTimeLayout), c.Format(DateTimeFormat))
		assert.Equal(t, c.Layout(DateTimeMilliLayout), c.Format(DateTimeMilliFormat))
		assert.Equal(t, c.Layout(DateTimeMicroLayout), c.Format(DateTimeMicroFormat))
		assert.Equal(t, c.Layout(DateTimeNanoLayout), c.Format(DateTimeNanoFormat))
		assert.Equal(t, c.Layout(ShortDateTimeLayout), c.Format(ShortDateTimeFormat))
		assert.Equal(t, c.Layout(ShortDateTimeMilliLayout), c.Format(ShortDateTimeMilliFormat))
		assert.Equal(t, c.Layout(ShortDateTimeMicroLayout), c.Format(ShortDateTimeMicroFormat))
		assert.Equal(t, c.Layout(ShortDateTimeNanoLayout), c.Format(ShortDateTimeNanoFormat))

		assert.Equal(t, c.Layout(DateLayout), c.Format(DateFormat))
		assert.Equal(t, c.Layout(DateMilliLayout), c.Format(DateMilliFormat))
		assert.Equal(t, c.Layout(DateMicroLayout), c.Format(DateMicroFormat))
		assert.Equal(t, c.Layout(DateNanoLayout), c.Format(DateNanoFormat))
		assert.Equal(t, c.Layout(ShortDateLayout), c.Format(ShortDateFormat))
		assert.Equal(t, c.Layout(ShortDateMilliLayout), c.Format(ShortDateMilliFormat))
		assert.Equal(t, c.Layout(ShortDateMicroLayout), c.Format(ShortDateMicroFormat))
		assert.Equal(t, c.Layout(ShortDateNanoLayout), c.Format(ShortDateNanoFormat))

		assert.Equal(t, c.Layout(TimeLayout), c.Format(TimeFormat))
		assert.Equal(t, c.Layout(TimeMilliLayout), c.Format(TimeMilliFormat))
		assert.Equal(t, c.Layout(TimeMicroLayout), c.Format(TimeMicroFormat))
		assert.Equal(t, c.Layout(TimeNanoLayout), c.Format(TimeNanoFormat))
		assert.Equal(t, c.Layout(ShortTimeLayout), c.Format(ShortTimeFormat))
		assert.Equal(t, c.Layout(ShortTimeMilliLayout), c.Format(ShortTimeMilliFormat))
		assert.Equal(t, c.Layout(ShortTimeMicroLayout), c.Format(ShortTimeMicroFormat))
		assert.Equal(t, c.Layout(ShortTimeNanoLayout), c.Format(ShortTimeNanoFormat))
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, c.Layout(AtomLayout, PRC), c.Format(AtomFormat, PRC))
		assert.Equal(t, c.Layout(ANSICLayout, PRC), c.Format(ANSICFormat, PRC))
		assert.Equal(t, c.Layout(CookieLayout, PRC), c.Format(CookieFormat, PRC))
		assert.Equal(t, c.Layout(KitchenLayout, PRC), c.Format(KitchenFormat, PRC))
		assert.Equal(t, c.Layout(RssLayout, PRC), c.Format(RssFormat, PRC))
		assert.Equal(t, c.Layout(RubyDateLayout, PRC), c.Format(RubyDateFormat, PRC))
		assert.Equal(t, c.Layout(UnixDateLayout, PRC), c.Format(UnixDateFormat, PRC))

		assert.Equal(t, c.Layout(RFC1036Layout, PRC), c.Format(RFC1036Format, PRC))
		assert.Equal(t, c.Layout(RFC1123Layout, PRC), c.Format(RFC1123Format, PRC))
		assert.Equal(t, c.Layout(RFC1123ZLayout, PRC), c.Format(RFC1123ZFormat, PRC))
		assert.Equal(t, c.Layout(RFC2822Layout, PRC), c.Format(RFC2822Format, PRC))
		assert.Equal(t, c.Layout(RFC3339Layout, PRC), c.Format(RFC3339Format, PRC))
		assert.Equal(t, c.Layout(RFC3339MilliLayout, PRC), c.Format(RFC3339MilliFormat, PRC))
		assert.Equal(t, c.Layout(RFC3339MicroLayout, PRC), c.Format(RFC3339MicroFormat, PRC))
		assert.Equal(t, c.Layout(RFC3339NanoLayout, PRC), c.Format(RFC3339NanoFormat, PRC))
		assert.Equal(t, c.Layout(RFC7231Layout, PRC), c.Format(RFC7231Format, PRC))
		assert.Equal(t, c.Layout(RFC822Layout, PRC), c.Format(RFC822Format, PRC))
		assert.Equal(t, c.Layout(RFC822ZLayout, PRC), c.Format(RFC822ZFormat, PRC))
		assert.Equal(t, c.Layout(RFC850Layout, PRC), c.Format(RFC850Format, PRC))

		assert.Equal(t, c.Layout(ISO8601Layout, PRC), c.Format(ISO8601Format, PRC))
		assert.Equal(t, c.Layout(ISO8601MilliLayout, PRC), c.Format(ISO8601MilliFormat, PRC))
		assert.Equal(t, c.Layout(ISO8601MicroLayout, PRC), c.Format(ISO8601MicroFormat, PRC))
		assert.Equal(t, c.Layout(ISO8601NanoLayout, PRC), c.Format(ISO8601NanoFormat, PRC))

		assert.Equal(t, c.Layout(ISO8601ZuluLayout, PRC), c.Format(ISO8601ZuluFormat, PRC))
		assert.Equal(t, c.Layout(ISO8601ZuluMilliLayout, PRC), c.Format(ISO8601ZuluMilliFormat, PRC))
		assert.Equal(t, c.Layout(ISO8601ZuluMicroLayout, PRC), c.Format(ISO8601ZuluMicroFormat, PRC))
		assert.Equal(t, c.Layout(ISO8601ZuluNanoLayout, PRC), c.Format(ISO8601ZuluNanoFormat, PRC))

		assert.Equal(t, c.Layout(FormattedDateLayout, PRC), c.Format(FormattedDateFormat, PRC))
		assert.Equal(t, c.Layout(FormattedDayDateLayout, PRC), c.Format(FormattedDayDateFormat, PRC))

		assert.Equal(t, c.Layout(DayDateTimeLayout, PRC), c.Format(DayDateTimeFormat, PRC))
		assert.Equal(t, c.Layout(DateTimeLayout, PRC), c.Format(DateTimeFormat, PRC))
		assert.Equal(t, c.Layout(DateTimeMilliLayout, PRC), c.Format(DateTimeMilliFormat, PRC))
		assert.Equal(t, c.Layout(DateTimeMicroLayout, PRC), c.Format(DateTimeMicroFormat, PRC))
		assert.Equal(t, c.Layout(DateTimeNanoLayout, PRC), c.Format(DateTimeNanoFormat, PRC))
		assert.Equal(t, c.Layout(ShortDateTimeLayout, PRC), c.Format(ShortDateTimeFormat, PRC))
		assert.Equal(t, c.Layout(ShortDateTimeMilliLayout, PRC), c.Format(ShortDateTimeMilliFormat, PRC))
		assert.Equal(t, c.Layout(ShortDateTimeMicroLayout, PRC), c.Format(ShortDateTimeMicroFormat, PRC))
		assert.Equal(t, c.Layout(ShortDateTimeNanoLayout, PRC), c.Format(ShortDateTimeNanoFormat, PRC))

		assert.Equal(t, c.Layout(DateLayout, PRC), c.Format(DateFormat, PRC))
		assert.Equal(t, c.Layout(DateMilliLayout, PRC), c.Format(DateMilliFormat, PRC))
		assert.Equal(t, c.Layout(DateMicroLayout, PRC), c.Format(DateMicroFormat, PRC))
		assert.Equal(t, c.Layout(DateNanoLayout, PRC), c.Format(DateNanoFormat, PRC))
		assert.Equal(t, c.Layout(ShortDateLayout, PRC), c.Format(ShortDateFormat, PRC))
		assert.Equal(t, c.Layout(ShortDateMilliLayout, PRC), c.Format(ShortDateMilliFormat, PRC))
		assert.Equal(t, c.Layout(ShortDateMicroLayout, PRC), c.Format(ShortDateMicroFormat, PRC))
		assert.Equal(t, c.Layout(ShortDateNanoLayout, PRC), c.Format(ShortDateNanoFormat, PRC))

		assert.Equal(t, c.Layout(TimeLayout, PRC), c.Format(TimeFormat, PRC))
		assert.Equal(t, c.Layout(TimeMilliLayout, PRC), c.Format(TimeMilliFormat, PRC))
		assert.Equal(t, c.Layout(TimeMicroLayout, PRC), c.Format(TimeMicroFormat, PRC))
		assert.Equal(t, c.Layout(TimeNanoLayout, PRC), c.Format(TimeNanoFormat, PRC))
		assert.Equal(t, c.Layout(ShortTimeLayout, PRC), c.Format(ShortTimeFormat, PRC))
		assert.Equal(t, c.Layout(ShortTimeMilliLayout, PRC), c.Format(ShortTimeMilliFormat, PRC))
		assert.Equal(t, c.Layout(ShortTimeMicroLayout, PRC), c.Format(ShortTimeMicroFormat, PRC))
		assert.Equal(t, c.Layout(ShortTimeNanoLayout, PRC), c.Format(ShortTimeNanoFormat, PRC))
	})
}
