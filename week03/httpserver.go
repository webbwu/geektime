package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func myHttpServer() error {
	http.HandleFunc("/", MyHandler)
	return http.ListenAndServe("127.0.0.0:8000", nil)
}

func signalNotify() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出", s)
			return errors.New("get signal")
		case syscall.SIGUSR1, syscall.SIGUSR2:
			fmt.Println("自定义信号", s)
		/* process */
		default:
			fmt.Println("other", s)
		}
	}
	return nil
}

func main() {
	group := new(errgroup.Group)
	group.Go(myHttpServer)
	group.Go(signalNotify)

	if err := group.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("successfully!")
	}

}
