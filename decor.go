package ansi

type Decor = string

const (
	Red    Decor = "\033[0;31m"
	Green  Decor = "\033[0;32m"
	Yellow Decor = "\033[1;33m"
	Blue   Decor = "\033[1;34m"
	Purple Decor = "\033[1;35m"
	Cyan   Decor = "\033[1;36m"

	Bold      Decor = "\033[1m"
	Underline Decor = "\033[4m"
	Reversed  Decor = "\033[7m"

	Reset Decor = "\033[0m"
)
