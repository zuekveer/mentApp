package runner

import (
	"context"
	"userservice/pkg/repository"
	"userservice/pkg/server"
)

type Runner struct {
	main       server.Server
	dependents []runner.StartStopper
}

func New(srv server.Server, ss []runner.StartStopper) *Runner {
	return &Runner {
		main:       srv,
		dependents: repo,
	}
}

func (r *Runner) Run() {
	ctx := context.Background()

	if err := r.dependents.Start(); err != nil {
		return
	}

	if err := r.dependents.Shutdown(); err != nil {
		return
	}

	if err := r.main.Start(); err != nil {
		return
	}

	//madeNew
	type StartStopper interface {
		Start()
		Shutdown()
	}

	func (strt *StartStopper) Start() error { 
		return nil
	}

	func (stp *StartStopper) Shutdown() error { 
		return nil
	}
}

//1) в мейн.го создать объект раннера (вызывать функцию нью) <|
// вызвать у объекта метод ран <|
//2) создать интерфейс СтартСтоппер с двумя методами стар и шутдавн <|
// пеередать вместо второго параметра в функцию нью слайс интерфейсов СтартСтоппер <|
// структура Раннер поле депендентс должна иметь тип слайса СтартСтоппер интерфейсов <|
//3) в методе ран нужно в цикле запустить все сервисы из слайса депендентс <|