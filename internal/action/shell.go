package action

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/pkg/tpl"
	"io"
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

func (sh *shell) Dispatch(ctx *contract.Context, res http.ResponseWriter) (bool, error) {
	whName := tpl.MustRender(`{{ .Webhook.Name }}`, ctx)
	log.Infof("[%s] running action shell: %s", whName, sh.cmd)

	args := tpl.MustRenderAll(sh.args, ctx)
	cmd := exec.Command(sh.cmd, args...)
	setEnv(ctx, sh, cmd)

	body := []byte(tpl.MustRender(`{{ .Request.Body }}`, ctx))
	cmd.Stdin = bytes.NewBuffer(body)

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		msg := fmt.Sprintf("[%s] action failed: %s", whName, err)
		if sh.response {
			_, _ = res.Write([]byte(msg))
		}
		log.Error(msg)

		return false, err
	}

	wg := ioDispatcher(ctx, sh, res, stdoutPipe, stderrPipe)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		msg := fmt.Sprintf("[%s] action failed: %s", whName, err)
		if sh.response {
			_, _ = res.Write([]byte(fmt.Sprintf("%s\n", msg)))
		}
		log.Error(msg)

		return false, err
	}

	msg := fmt.Sprintf("[%s] action completed", whName)
	if sh.response {
		_, _ = res.Write([]byte(fmt.Sprintf("%s\n", msg)))
	}
	log.Info(msg)

	return true, nil
}

func ioDispatcher(ctx *contract.Context, sh *shell, res http.ResponseWriter, stdoutPipe io.ReadCloser, stderrPipe io.ReadCloser) *sync.WaitGroup {
	whName := tpl.MustRender(`{{ .Webhook.Name }}`, ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			msg := scanner.Text()
			if sh.response {
				_, _ = res.Write([]byte(fmt.Sprintf("%s\n", msg)))
				res.(http.Flusher).Flush()
			}
			log.Info(fmt.Sprintf("[%s] %s", whName, msg))
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			msg := scanner.Text()
			if sh.response {
				_, _ = res.Write([]byte(fmt.Sprintf("%s\n", msg)))
				res.(http.Flusher).Flush()
			}
			log.Warning(fmt.Sprintf("[%s] %s", whName, msg))
		}

		wg.Done()
	}()

	return &wg
}

func setEnv(ctx *contract.Context, sh *shell, cmd *exec.Cmd) {
	env := append(os.Environ(),
		tpl.MustRender(`WEBHUG_WEBHOOK={{ .Webhook.Name }}`, ctx),
		tpl.MustRender(`WEBHUG_REQUEST_METHOD={{ .Request.Method }}`, ctx),
		tpl.MustRender(`WEBHUG_REQUEST_REMOTE_ADDR={{ .Request.RemoteAddr }}`, ctx),
	)

	for _, envVar := range sh.env {
		env = append(env, tpl.MustRender(envVar, ctx))
	}

	cmd.Env = env
}
