package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"github.com/google/uuid"
)

func GenerateID() string {
	return uuid.New().String()
}

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPositiveInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		value, err := strconv.Atoi(input)
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Vui long nhap mot so nguyen duong.")
	}
}

func GetPositiveFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Vui long nhap mot so thuc duong.")
	}
}

func GetOptionalPositiveFloat(prompt string, oldValue float64) float64 {

	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}
	value, err := strconv.ParseFloat(input, 64)
	if err != nil && value < 0 {
		fmt.Println("Gia tri ko hop le, giu nguyen gia tri cu.")
		return oldValue
	}
	return value
}

func GetNotEmptyValue(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}
		fmt.Println("Vui long nhap mot gia tri khong rong.")
	}
}

func GetOptionalValue(prompt string, oldValue string) string {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}
	return input
}

func ClearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing screen:", err)
	}
}

type HasID interface {
	GetID() int
}

func IsIdUnique[T HasID](id int, list []T) bool {
	for _, item := range list {
		if item.GetID() == id {
			return false
		}
	}
	return true
}
