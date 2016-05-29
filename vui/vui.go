package vui

type Ui struct {
	Screen [][]string
}

func (u *Ui) Output(msg string) {
	u.Info(msg)
}

func (u *Ui) Info(msg string) {
	line := []string{"info", msg}
	u.Screen = append(u.Screen, line)
}

func (u *Ui) Warn(msg string) {
	line := []string{"warn", msg}
	u.Screen = append(u.Screen, line)
}

func (u *Ui) Error(msg string) {
	line := []string{"error", msg}
	u.Screen = append(u.Screen, line)
}

func (u *Ui) Flush() {
	u.Screen = nil
}
