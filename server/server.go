package server

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/jander/golog/logger"
	"github.com/kardianos/service"
)

type Service struct {
	name        string
	displayName string
	desc        string
}

func On(n, dn, d string) *Service {
	return &Service{
		name:        n,
		displayName: dn,
		desc:        d,
	}

}

func (sv *Service) Start(callBack func()) {
	if len(sv.name) == 0 {
		color.Error.Render("Service init faild,must first call On.\n")
		return
	}
	//name, displayName, desc := config.GetServiceConfig()
	p := &program{callBack}
	sc := &service.Config{
		Name:        sv.name,
		DisplayName: sv.displayName,
		Description: sv.desc,
	}
	s, err := service.New(p, sc)
	//var s, err = service.NewService(name, displayName, desc)
	if err != nil {
		color.Error.Render("%s unable to start: %s", sv.displayName, err)
		return
	}

	color.Info.Render("Service \"%s\" do.\n", sv.displayName)

	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			{
				err = s.Install()
				if err != nil {
					color.Error.Render("Failed to install: %s\n", err)
					return
				}
				color.Info.Render("Service \"%s\" installed.\n", sv.displayName)
			}
		case "remove":
			{
				err = s.Uninstall()
				if err != nil {
					color.Error.Render("Failed to remove: %s\n", err)
					return
				}
				color.Info.Render("Service \"%s\" removed.\n", sv.displayName)
			}
		case "run":
			{
				err = s.Run()
				if err != nil {
					color.Error.Render("Failed to run: %s\n", err)
					return
				}
				color.Info.Render("Service \"%s\" run.\n", sv.displayName)
			}
		case "start":
			{
				err = s.Start()
				if err != nil {
					color.Error.Render("Failed to start: %s\n", err)
					return
				}
				color.Info.Render("starting check service:", sv.displayName)

				ticker := time.NewTicker(1 * time.Second)
				<-ticker.C

				var s ServiceTools
				st, err := s.IsStart(sv.name)
				if err != nil {
					color.Error.Render(err)
					return
				} else {
					if st == Stopped || st == StopPending {
						color.Error.Render("Service \"%s\" is Stopped.\n", sv.displayName)
						color.Error.Render("can't to start service.")
						return
					}
				}
				color.Info.Render("Service \"%s\" started.\n", sv.displayName)
			}
		case "stop":
			{
				err = s.Stop()
				if err != nil {
					color.Error.Render("Failed to stop: %s\n", err)
					return
				}
				var s ServiceTools
				st, err := s.IsStart(sv.name)
				if err != nil {
					fmt.Println(err)
					return
				} else {
					if st == Running || st == StartPending {
						color.Error.Render("Service \"%s\" is Started.\n", sv.displayName)
						color.Error.Render("can't to stop service.")
						return
					}
				}
				color.Info.Render("Service \"%s\" stopped.\n", sv.displayName)
			}
		}
		return
	}

	color.Note.Render("Failed to read args\n")

	if err = s.Run(); err != nil {
		logger.Error(err)
	}
}

type program struct {
	callBack func()
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *program) run() {
	p.callBack()
}
func (p *program) Stop(s service.Service) error {
	return nil
}

type IServiceTools interface {
	IsStart(name string) (status int, err error)
}
