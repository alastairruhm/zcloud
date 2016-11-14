package cmd

import (
	"bytes"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	RootCmd.SetArgs([]string{"version"})
	RootCmd.SetOutput(buf)

	err := RootCmd.Execute()

	if err != nil {
		t.Error(err)
	}

	actual := buf.String()
	expected := "zcloud client tool version " + VERSION + "\n"

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}

}
