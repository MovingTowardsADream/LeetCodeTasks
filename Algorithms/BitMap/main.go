package main

import "fmt"

type BitMap struct {
	bits uint64
}

// Set устанавливает бит для данного значения
func (b *BitMap) Set(value int) {
	if value < 0 || value >= 64 {
		fmt.Println("Value out of range")
		return
	}
	b.bits |= 1 << value
}

// Clear сбрасывает бит для данного значения
func (b *BitMap) Clear(value int) {
	if value < 0 || value >= 64 {
		fmt.Println("Value out of range")
		return
	}
	b.bits &^= 1 << value
}

// Test проверяет, установлен ли бит для данного значения
func (b *BitMap) Test(value int) bool {
	if value < 0 || value >= 64 {
		fmt.Println("Value out of range")
		return false
	}
	return b.bits&(1<<value) != 0
}

func main() {
	bitmap := &BitMap{}

	// Установка битов
	bitmap.Set(1)
	bitmap.Set(3)
	bitmap.Set(5)

	// Проверка битов
	fmt.Println("Bit 1 set:", bitmap.Test(1)) // true
	fmt.Println("Bit 3 set:", bitmap.Test(3)) // true
	fmt.Println("Bit 5 set:", bitmap.Test(5)) // true
	fmt.Println("Bit 2 set:", bitmap.Test(2)) // false

	// Сброс бита
	bitmap.Clear(3)
	fmt.Println("Bit 3 set:", bitmap.Test(3)) // false
}
