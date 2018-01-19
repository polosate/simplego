package _7_lesson

func swap(a, b *int) {
	*a = *a - *b
	*b = *a + *b
	*a = *b - *a
}

func swap_bit(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
