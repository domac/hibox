package restart

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const (
	ENV_NAME     = "RST_LISTENER"
	QUIT_TIMEOUT = 5 * time.Second
)

type innerListener struct {
	Addr     string `json:"addr"`
	FD       int    `json:"fd"`
	Filename string `json:"filename"`
}

func importListener(addr string) (net.Listener, error) {
	listenerEnv := os.Getenv(ENV_NAME)
	if listenerEnv == "" {
		return nil, fmt.Errorf("unable to find %s environment variable", listenerEnv)
	}
	fmt.Printf("import environment variable %s\n", listenerEnv)

	var l innerListener
	err := json.Unmarshal([]byte(listenerEnv), &l)
	if err != nil {
		return nil, err
	}

	if l.Addr != addr {
		return nil, fmt.Errorf("unable to find listener for %v = %v\n", l.Addr, addr)
	}

	listenerFile := os.NewFile(uintptr(l.FD), l.Filename)
	if listenerFile == nil {
		return nil, fmt.Errorf("unable to create listener file")
	}
	defer listenerFile.Close()

	ln, err := net.FileListener(listenerFile)
	if err != nil {
		return nil, err
	}
	return ln, nil
}

func createListener(addr string) (net.Listener, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return ln, nil
}

func ImportOrCreateListener(addr string) (net.Listener, error) {
	ln, err := importListener(addr)
	if err == nil {
		fmt.Printf("Imported listener file descriptor for %v\n", addr)
		return ln, err
	}

	println(">>>>", err.Error())

	ln, err = createListener(addr)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Create listener file descriptor for %v \n", addr)
	return ln, nil
}

type Server struct {
	srv  *http.Server
	ln   net.Listener
	addr string
}

func StartServer(addr string, ln net.Listener) (*Server, error) {
	if ln == nil {
		return nil, errors.New("listener is null")
	}
	server := &Server{
		srv:  &http.Server{Addr: addr},
		ln:   ln,
		addr: addr,
	}
	go server.srv.Serve(server.ln)
	return server, nil
}

func (s *Server) shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) getListenerFile() (*os.File, error) {
	switch t := s.ln.(type) {
	case *net.TCPListener:
		return t.File()
	case *net.UnixListener:
		return t.File()
	}
	return nil, fmt.Errorf("unsupported Listener:%T", s.ln)
}

//fork 子进程
func (s *Server) forkChild() (*os.Process, error) {
	lnFile, err := s.getListenerFile()
	if err != nil {
		return nil, err
	}
	defer lnFile.Close()

	l := innerListener{
		Addr:     s.addr,
		FD:       3,
		Filename: lnFile.Name(),
	}

	listenerEnv, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	files := []*os.File{
		os.Stdin,
		os.Stdout,
		os.Stderr,
		lnFile,
	}

	execName, err := os.Executable()
	if err != nil {
		return nil, err
	}

	environment := append(os.Environ(), ENV_NAME+"="+string(listenerEnv))

	p, err := os.StartProcess(execName, []string{execName}, &os.ProcAttr{
		Dir:   filepath.Dir(execName),
		Env:   environment,
		Files: files,
		Sys:   &syscall.SysProcAttr{},
	})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func WaitForSignals(server *Server) error {
	signalCh := make(chan os.Signal, 1024)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-signalCh:
			fmt.Printf("%v signal received\n", s)
			switch s {
			case syscall.SIGHUP: //平滑升级

				p, err := server.forkChild()
				if err != nil {
					fmt.Printf("unable to fork child: %v\n", err)
					continue
				}

				if p == nil {
					fmt.Printf("unable to fork child which process is null\n")
					continue
				}
				fmt.Println("hup now !")
				ctx, cancel := context.WithTimeout(context.Background(), QUIT_TIMEOUT)
				defer cancel()
				return server.shutdown(ctx)

			case syscall.SIGUSR2: //fork
				p, err := server.forkChild()
				if err != nil {
					fmt.Printf("unable to fork child: %v\n", err)
					continue
				}

				if p == nil {
					fmt.Printf("unable to fork child which process is null\n")
					continue
				}
				fmt.Printf("Fork child - pid:%v\n", p.Pid)
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM: //强制退出
				fmt.Println("quit now !")
				ctx, cancel := context.WithTimeout(context.Background(), QUIT_TIMEOUT)
				defer cancel()
				return server.shutdown(ctx)
			}
		}
	}
}
