package storefront

type Service struct {
	Clients  []chan string
	Notifier chan string
}

func New() *Service {
	return &Service{
		Clients:  make([]chan string, 0),
		Notifier: make(chan string),
	}
}
