package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	conf "online-ordering-system/config"
	ctl "online-ordering-system/controller"
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

// @BasePath /v1
// swagger API 선언
// func setupSwagger(r *gin.Engine) {
// 	r.GET("/", func(c *gin.Context) {
// 		c.Redirect(http.StatusFound, "/swagger/index.html")
// 	})

// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// }

func main() {

	//model 모듈 선언
	if mod, err := model.NewModel(); err != nil {
		//~생략
	} else if controller, err := ctl.NewCTL(mod); err != nil { //controller 모듈 설정
		//~생략
	} else if rt, err := rt.NewRouter(controller); err != nil { //router 모듈 설정
		//~생략
	} else {
		config := conf.GetConfig("./config/.config.toml")
		fmt.Println("config.Server.Port", config.Server.Port)
		fmt.Println("config.Server.Mode", config.Server.Mode)
		fmt.Println("config.DB[account][pass]", config.DB["account"]["pass"])
		fmt.Println("work", config.Work)
		fmt.Println("work", config.Work[0].Desc)
		// mapi := &http.Server{
		// 	Addr:           ":8080",
		// 	Handler:        rt.Idx(),
		// 	ReadTimeout:    5 * time.Second,
		// 	WriteTimeout:   10 * time.Second,
		// 	MaxHeaderBytes: 1 << 20,
		// }

		//http 서버 설정 변수
		mapi := &http.Server{
			Addr:           config.Server.Port,
			Handler:        rt.Idx(),
			ReadTimeout:    0,
			WriteTimeout:   0,
			MaxHeaderBytes: 1 << 20,
		}

		//고루틴 서버 동작
		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		var configFlag = flag.String("config", "./conf/config.toml", "toml file to use for configuration")
		flag.Parse()
		cf := conf.NewConfig(*configFlag)

		// 로그 초기화
		if err := logger.InitLogger(cf); err != nil {
			fmt.Printf("init logger failed, err:%v\n", err)
			return
		}

		// flag.Parse()
		// fmt.Println(flag.Args())

		// var port string
		// flag.StringVar(&port, "port", "7070", "port to listen on")
		// // port := flag.Int("port", 8080, "포트번호")
		// var conf string
		// flag.StringVar(&conf, "config", "./conf.toml", "config file to use")

		// pMod := flag.String("mode", "debug", "service mode")
		// flag.Parse()
		// fmt.Println(port)
		// fmt.Println(conf)
		// fmt.Println(*pMod)

		// port := flag.Int("port", 8080, "포트번호")
		// conf2 := flag.String("config", "./", "config")
		// fmt.Println(*port, *conf2)
		// flag.Parse()
		// fmt.Println("port", *port)

		// middleware 설정
		// setupSwagger(rt.Idx())

		stopSig := make(chan os.Signal) //chan 선언
		// 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
		signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM)
		<-stopSig //메세지 등록

		// 해당 context 타임아웃 설정, 5초후 server stop
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
		case <-ctx.Done():
			fmt.Println("timeout 5 seconds.")
		}
		fmt.Println("Server stop")
		//우아한 종료
	}

	g.Wait()
	//~생략

}
