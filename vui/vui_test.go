package vui

import (
	"testing"

	"github.com/wata727/spaceship/vui"
)

func TestVuiOutput(t *testing.T) {
	var ui = &vui.Ui{}

	expected_level := "info"
	expected_first_text := "testing..."
	expected_second_text := "finish!"

	ui.Output("testing...")
	ui.Output("finish!")

	if expected_level != ui.Screen[0][0] {
		t.Errorf("Log Level don't match.\ngot %v\nwant %v", ui.Screen[0][0], expected_level)
	}
	if expected_first_text != ui.Screen[0][1] {
		t.Errorf("Log Text don't match.\ngot %v\nwant %v", ui.Screen[0][1], expected_first_text)
	}
	if expected_level != ui.Screen[1][0] {
		t.Errorf("Add Log Level don't match.\ngot %v\nwant %v", ui.Screen[0][0], expected_level)
	}
	if expected_second_text != ui.Screen[1][1] {
		t.Errorf("Add Log Text don't match.\ngot %v\nwant %v", ui.Screen[0][1], expected_second_text)
	}
}

func TestVuiInfo(t *testing.T) {
	var ui = &vui.Ui{}

	expected_level := "info"
	expected_first_text := "testing..."
	expected_second_text := "finish!"

	ui.Info("testing...")
	ui.Info("finish!")

	if expected_level != ui.Screen[0][0] {
		t.Errorf("Log Level don't match.\ngot %v\nwant %v", ui.Screen[0][0], expected_level)
	}
	if expected_first_text != ui.Screen[0][1] {
		t.Errorf("Log Text don't match.\ngot %v\nwant %v", ui.Screen[0][1], expected_first_text)
	}
	if expected_level != ui.Screen[1][0] {
		t.Errorf("Add Log Level don't match.\ngot %v\nwant %v", ui.Screen[0][0], expected_level)
	}
	if expected_second_text != ui.Screen[1][1] {
		t.Errorf("Add Log Text don't match.\ngot %v\nwant %v", ui.Screen[0][1], expected_second_text)
	}
}

func TestVuiWarn(t *testing.T) {
	var ui = &vui.Ui{}

	expected_level := "warn"
	expected_first_text := "testing..."
	expected_second_text := "finish!"

	ui.Warn("testing...")
	ui.Warn("finish!")

	if expected_level != ui.Screen[0][0] {
		t.Errorf("Log Level don't match.\ngot %v\nwant %v", ui.Screen[0][0], expected_level)
	}
	if expected_first_text != ui.Screen[0][1] {
		t.Errorf("Log Text don't match.\ngot %v\nwant %v", ui.Screen[0][1], expected_first_text)
	}
	if expected_level != ui.Screen[1][0] {
		t.Errorf("Add Log Level don't match.\ngot %v\nwant %v", ui.Screen[0][0], expected_level)
	}
	if expected_second_text != ui.Screen[1][1] {
		t.Errorf("Add Log Text don't match.\ngot %v\nwant %v", ui.Screen[0][1], expected_second_text)
	}
}

func TestVuiError(t *testing.T) {
	var ui = &vui.Ui{}

	expected_level := "error"
	expected_first_text := "testing..."
	expected_second_text := "finish!"

	ui.Error("testing...")
	ui.Error("finish!")

	if expected_level != ui.Screen[0][0] {
		t.Errorf("Log Level don't match.\ngot %v\nwant %v", ui.Screen[0][0], expected_level)
	}
	if expected_first_text != ui.Screen[0][1] {
		t.Errorf("Log Text don't match.\ngot %v\nwant %v", ui.Screen[0][1], expected_first_text)
	}
	if expected_level != ui.Screen[1][0] {
		t.Errorf("Add Log Level don't match.\ngot %v\nwant %v", ui.Screen[0][0], expected_level)
	}
	if expected_second_text != ui.Screen[1][1] {
		t.Errorf("Add Log Text don't match.\ngot %v\nwant %v", ui.Screen[0][1], expected_second_text)
	}
}

func TestVuiFlush(t *testing.T) {
	var ui = &vui.Ui{}

	ui.Output("testing...")

	if len(ui.Screen) != 1 {
		t.Errorf("Vui Output not working.\ngot %v", len(ui.Screen))
	}

	ui.Flush()
	if len(ui.Screen) != 0 {
		t.Errorf("Vui Flush not working.\ngot %v", len(ui.Screen))
	}
}
