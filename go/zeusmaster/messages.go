package zeusmaster

import (
	"errors"
	"strconv"
	"strings"
)

func ParsePidMessage(msg string) (int, string, error) {
	parts := strings.SplitN(msg, ":", 3)
	if parts[0] != "P" {
		return -1, "", errors.New("Wrong message type! Expected PidMessage, got: " + msg)
	}

	identifier := parts[2]
	pid, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, "", err
	}

	return pid, identifier, nil
}

func ParseFeatureMessage(msg string) (string, error) {
	parts := strings.SplitN(msg, ":", 2)
	if parts[0] != "F" {
		return "", errors.New("Wrong message type! Expected FeatureMessage, got: " + msg)
	}
	return strings.TrimSpace(parts[1]), nil
}

func ParseActionResponseMessage(msg string) (string, error) {
	parts := strings.SplitN(msg, ":", 2)
	if parts[0] != "R" {
		return "", errors.New("Wrong message type! Expected ActionResponseMessage, got: " + msg)
	}
	return parts[1], nil
}

func CreateSpawnSlaveMessage(identifier string) string {
	return "S:" + identifier
}

func CreateSpawnCommandMessage(identifier string) string {
	return "C:" + identifier
}

func ParseClientCommandRequestMessage(msg string) (string, string, error) {
	parts := strings.SplitN(msg, ":", 3)
	if parts[0] != "Q" {
		return "", "", errors.New("Wrong message type! Expected ClientCommandRequestMessage, got: " + msg)
	}

	command := parts[1]
	arguments := parts[2]

	return command, arguments, nil
}
