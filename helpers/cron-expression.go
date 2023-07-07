package helpers

const (
	CronEveryMinute     = "* * * * *"
	CronEvery5Minute    = "*/5 * * * *"
	CronEvery10Minute   = "*/10 * * * *"
	CronEvery15Minute   = "*/15 * * * *"
	CronEvery30Minute   = "*/30 * * * *"
	CronEveryHour       = "0 * * * *"
	CronEveryDay        = "0 0 * * *"
	CronEveryWeek       = "0 0 * * 0"
	CronEveryMonth      = "0 0 1 * *"
	CronMidnight        = "0 0 * * *"
	CronTwiceDaily      = "0 0,12 * * *"
	CronWeekdayMornings = "0 9 * * 1-5"
	CronWeekendEvenings = "0 18 * * 6,0"
)
