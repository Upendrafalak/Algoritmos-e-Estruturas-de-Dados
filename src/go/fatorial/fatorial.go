package fatorial

func Fatorial(value int) int {
	if value == 1 {
		return 1
	}
	return value * Fatorial(value-1)
}
