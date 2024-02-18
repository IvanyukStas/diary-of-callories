package logger

type Logger struct{
	Info(error)
	Fatal(error)
	Worning(error)
}

