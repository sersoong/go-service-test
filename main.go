package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kardianos/service"
	"github.com/sersoong/go-service-test/web"
)

var serviceConfig = &service.Config{
	Name:        "WebService",
	DisplayName: "WebService",
	Description: "A small gin web service",
}

var logger service.Logger

func main() {

	// 构建服务对象
	prog := &Program{}
	s, err := service.New(prog, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 用于记录系统日志
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}

	cmd := os.Args[1]

	if cmd == "install" {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("安装成功")
	}
	if cmd == "uninstall" {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
	}

	// install, uninstall, start, stop 的另一种实现方式
	// err = service.Control(s, os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	logger.Info("开始服务")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	logger.Info("停止服务")
	return nil
}

func (p *Program) run() {
	// 此处编写具体的服务代码
	logger.Info("启动web服务")
	web.StartWeb()
}
