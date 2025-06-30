package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetIntegerInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	receivedInteger, okError := strconv.Atoi(input)
	return receivedInteger, okError

}

func GetStringInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	return input, nil

}
