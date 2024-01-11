package server

type Server struct {
	config *Config
}

type Config struct {
	Port int8 `env:"SERVER_PORT"`
}

func (s *Server) Start() error {
	return nil
}

func (s *Server) Shutdown() error {
	return nil
}

func New(cfg *Config) *Server {
	return &Server{
		config: cfg,
	}
}

//описать методы старт и шатдавн для класса сервер. - <3
//В них ничего не возвращать (ретурн нил). Тоже самое для репозитория. - <3
//Их необходимо вызвать из мейн.го. В функции нью сервера и репозитория - <3
//добавить валидирование передаваемых конфигов. К примеру пароль
//от базы данных не может быть пустысм, а порт состоит из 2 цифр.
//Если ошибка валидации - выкидываю из функции нью
//вторым возвращаемым параметром ошибку (фмт.еррорф)