package storefront

func (s *Service) Broker(done <-chan interface{}) {
	for {
		select {
		case <-done:
			return
		case data := <-s.Notifier:
			for _, channel := range s.Clients {
				channel <- data
			}
		}
	}
}
