package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	conf "online-ordering-system/config"
	ctl "online-ordering-system/controller"
	"online-ordering-system/logger"
	"online-ordering-system/model"
	rt "online-ordering-system/router"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// @title Online Ordering System API
// @version 1.0
// @description This is a online ordering  server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {

	//model 모듈 선언
	if mod, err := model.NewModel(); err != nil {
		//~생략
	} else if controller, err := ctl.NewCTL(mod); err != nil { //controller 모듈 설정
		//~생략
	} else if rt, err := rt.NewRouter(controller); err != nil { //router 모듈 설정
		//~생략
	} else {
		var configFlag = flag.String("config", "./config/.config.toml", "toml file to use for configuration")
		flag.Parse()
		cf := conf.GetConfig(*configFlag)

		// 로그 초기화
		if err := logger.InitLogger(cf); err != nil {
			fmt.Printf("init logger failed, err:%v\n", err)
			return
		}

		logger.Debug("ready server....")

		//http 서버 설정 변수
		mapi := &http.Server{
			Addr:           cf.Server.Port,
			Handler:        rt.Idx(),
			ReadTimeout:    0, //  5 * time.Second, 이전 값 현재 값은 테스트를 위해 설정함
			WriteTimeout:   0, // 10 * time.Second, 이전 값 현재 값은 테스트를 위해 설정함
			MaxHeaderBytes: 1 << 20,
		}

		//고루틴 서버 동작
		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		quit := make(chan os.Signal) //chan 선언
		// 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit //메세지 등록

		// 해당 context 타임아웃 설정, 5초후 server stop
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
		case <-ctx.Done():
			logger.Info("timeout of 5 seconds.")
		}
		logger.Info("Server exiting")
		//우아한 종료
	}

	if err := g.Wait(); err != nil {
		logger.Error(err)
	}

}
