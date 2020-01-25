package action

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/phramz/webhug/internal/contract"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

type shell struct {
	cmd      string
	args     []string
	env      []string
	response bool
}

func (sh *shell) HasResponse() bool {
	return sh.response
}

func (sh *shell) Dispatch(wh contract.Webhook, rq *http.Request, res http.ResponseWriter) {
	log.Infof("[%s] running action shell: %s", wh.GetName(), sh.cmd)

	jsonHdr, _ := json.Marshal(rq.Header)

	cmd := exec.Command(sh.cmd, sh.args...)
	setEnv(wh, rq, jsonHdr, sh, cmd)

	body, _ := ioutil.ReadAll(rq.Body)
	cmd.Stdin = bytes.NewBuffer(body)

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		log.Errorf("[%s] action failed: %s", wh.GetName(), err)
	}

	wg := ioDispatcher(stdoutPipe, wh, sh, res, stderrPipe)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		msg := fmt.Sprintf("[%s] action failed: %s", wh.GetName(), err)
		if sh.response {
			res.Write([]byte(msg))
		}
		log.Error(msg)
	}
}

func ioDispatcher(stdoutPipe io.ReadCloser, wh contract.Webhook, sh *shell, res http.ResponseWriter, stderrPipe io.ReadCloser) *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			msg := fmt.Sprintf("[%s] %s\n", wh.GetName(), scanner.Text())
			if sh.response {
				res.Write([]byte(msg))
				res.(http.Flusher).Flush()
			}
			log.Info(msg)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			msg := fmt.Sprintf("[%s] %s\n", wh.GetName(), scanner.Text())
			if sh.response {
				res.Write([]byte(msg))
				res.(http.Flusher).Flush()
			}
			log.Warning(msg)
		}

		wg.Done()
	}()

	return &wg
}

func setEnv(wh contract.Webhook, rq *http.Request, jsonHdr []byte, sh *shell, cmd *exec.Cmd) {
	env := append(os.Environ(),
		fmt.Sprintf("WEBHUG_WEBHOOK=%s", wh.GetName()),
		fmt.Sprintf("WEBHUG_REQUEST_METHOD=%s", rq.Method),
		fmt.Sprintf("WEBHUG_REQUEST_REMOTE_ADDR=%s", rq.RemoteAddr),
		fmt.Sprintf("WEBHUG_REQUEST_HEADER=%s", jsonHdr),
	)

	env = append(env, sh.env...)
	cmd.Env = env
}
