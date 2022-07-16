package formater

func Formater(number string) string {
	for i := len(number) - 3; i > 0; i -= 3 {
		number = number[:i] + "," + number[i:]
	}
	return number
}
