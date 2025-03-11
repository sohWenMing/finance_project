package server_tests

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	errorUtils "github.com/sohWenMing/finance_project/internal/utils/error_utils"
	executils "github.com/sohWenMing/finance_project/internal/utils/exec"
	httpInternal "github.com/sohWenMing/finance_project/internal/utils/http_internal"
)

var client *http.Client = httpInternal.DefaultClient

func TestMain(m *testing.M) {
	startDBCommand := executils.GenExecCommand("../..", "make", "start-postgres-image")
	stopDBCommand := executils.GenExecCommand("../..", "make", "stop-postgres-image")

	startServerCmd := executils.GenExecCommand("../..", "make", "run-server-image-background")
	stopServerCmd := executils.GenExecCommand("../..", "make", "stop-server-image")
	if err := startDBCommand.Run(); err != nil {
		os.Exit(1)
	}
	if err := startServerCmd.Run(); err != nil {
		os.Exit(1)
	}
	waitServerErr := waitForServer()
	if waitServerErr != nil {
		log.Fatal(waitServerErr)
	}
	code := m.Run()
	if err := stopServerCmd.Run(); err != nil {
		os.Exit(1)
	}
	if err := stopDBCommand.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(code)
}

func TestPing(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/ping", nil)
	errorUtils.AssertNoError(t, "TestPing", err)

	res, err := client.Do(req)
	errorUtils.AssertNoError(t, "TestPing", err)

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	errorUtils.AssertNoError(t, "TestPing", err)
	errorUtils.AssertVals(t, "TestPing", string(resBody), "OK")

}

func waitForServer() (err error) {
	for i := 0; i < 10; i++ {
		res, err := http.Get("http://localhost:8080/ping")
		if err == nil && res.StatusCode == http.StatusOK {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("server did not become ready in time")

}
