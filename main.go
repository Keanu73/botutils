package botutils

type cron struct{}
type db struct{}
type env struct{}

// Exports modules
var (
	Cron     = &cron{}
	Database = &db{}
	Env      = &env{}
)
