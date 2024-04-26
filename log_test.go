package arlogf

import (
	"testing"
)

func Test_CreateLogger(T *testing.T) {
	logger := NewLogger(false)

	logger.Builder("test1").Print("Testing")
	logger.Builder("test2").Log().Print("Testing")
	logger.Builder("test2-f").Log().Printf("%s", "Testing")
	logger.Builder("test3").Warn().Print("Testing")
	logger.Builder("test4").Error("").Print("Testing")
	logger.Builder("test5").Fatal("").Print("Testing")
}
