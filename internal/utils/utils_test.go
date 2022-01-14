package utils

import (
	"fmt"
	"testing"
)

func TestGenerateSerialNumber(t *testing.T) {
	id := GenerateSerialNumber()
	fmt.Println(id)
}
