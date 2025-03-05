package server_tests

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	errorUtils "github.com/sohWenMing/finance_project/internal/utils/error_utils"
	executils "github.com/sohWenMing/finance_project/internal/utils/exec"
	httpInternal "github.com/sohWenMing/finance_project/internal/utils/http_internal"
)

var client *http.Client = httpInternal.DefaultClient

func TestMain(m *testing.M) {
	startServerCmd := executils.GenExecCommand("../..", "make", "run-server-image-background")
	stopServerCmd := executils.GenExecCommand("../..", "make", "stop-server-image")
	if err := startServerCmd.Run(); err != nil {
		os.Exit(1)
	}
	code := m.Run()
	if err := stopServerCmd.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(code)
}

func TestPing(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/ping", nil)
	if err != nil {
		fmt.Println("test failed at line 32")
	}
	res, err := client.Do(req)
	if !errorUtils.AssertNoError(t, "TestPing", err) {
		fmt.Println("test failed at line 37")
	}
	if !errorUtils.AssertVals(t, "TestPing", res.StatusCode, http.StatusOK) {
		fmt.Println("test failed at line 40")
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if !errorUtils.AssertNoError(t, "TestPing", err) {
		fmt.Println("test failed at line 46")
	}
	if !errorUtils.AssertVals(t, "TestPing", string(resBody), "OK") {
		fmt.Println("test failed at line 50")
	}
}
