package utils

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}
