package clog

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

func TestSetLogLevel(t *testing.T) {
	SetLogLevel(TraceLevel)
	require.Equal(t, TraceLevel, Logger.Level)
}

func TestTrace(t *testing.T) {
	SetLogLevel(TraceLevel)
	buffer := &bytes.Buffer{}
	Logger.Out = buffer
	for n := 0; n < 5; n++ {
		Trace("Test", "test", map[string]string{
			"n": strconv.Itoa(n),
		})
	}
	output := buffer.String()
	Logger.Out = os.Stdout
	require.Contains(t, output, "\"level\":\"trace\",\"msg\":\"Test\",\"n\":\"1\"")
}

func TestDebug(t *testing.T) {
	SetLogLevel(DebugLevel)
	buffer := &bytes.Buffer{}
	Logger.Out = buffer
	for n := 0; n < 5; n++ {
		Debug("Test", "test", map[string]string{
			"n": strconv.Itoa(n),
		})
	}
	output := buffer.String()
	Logger.Out = os.Stdout
	require.Contains(t, output, "\"level\":\"debug\",\"msg\":\"Test\",\"n\":\"1\"")
}

func TestInfo(t *testing.T) {
	SetLogLevel(InfoLevel)
	buffer := &bytes.Buffer{}
	Logger.Out = buffer
	for n := 0; n < 5; n++ {
		Info("Test", "test", map[string]string{
			"n": strconv.Itoa(n),
		})
	}
	output := buffer.String()
	Logger.Out = os.Stdout
	require.Contains(t, output, "\"level\":\"info\",\"msg\":\"Test\",\"n\":\"1\"")
}

func TestWarn(t *testing.T) {
	SetLogLevel(WarnLevel)
	buffer := &bytes.Buffer{}
	Logger.Out = buffer
	for n := 0; n < 5; n++ {
		Info("Test", "test", map[string]string{
			"n": strconv.Itoa(n),
		})
	}
	output := buffer.String()
	Logger.Out = os.Stdout
	require.Contains(t, output, "\"level\":\"warn\",\"msg\":\"Test\",\"n\":\"1\"")
}

func TestError(t *testing.T) {
	SetLogLevel(ErrorLevel)
	buffer := &bytes.Buffer{}
	Logger.Out = buffer
	for n := 0; n < 5; n++ {
		Error("Test", "test", errors.New("test"), map[string]string{
			"n": strconv.Itoa(n),
		})
	}
	output := buffer.String()
	Logger.Out = os.Stdout
	require.Contains(t, output, "\"level\":\"error\",\"msg\":\"Test - ERROR: test\",\"n\":\"1\"")
}

func TestPanic(t *testing.T) {
	require.Panics(t, func() {
		SetLogLevel(ErrorLevel)
		buffer := &bytes.Buffer{}
		Logger.Out = buffer
		for n := 0; n < 5; n++ {
			Panic("Test", "test", errors.New("test"), map[string]string{
				"n": strconv.Itoa(n),
			})
		}
		output := buffer.String()
		Logger.Out = os.Stdout
		require.Contains(t, output, "\"level\":\"error\",\"msg\":\"Test - PANIC: test\",\"n\":\"1\"")
	})
}
