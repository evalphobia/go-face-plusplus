package beautify

import (
	"github.com/evalphobia/go-face-plusplus/client"
	"github.com/evalphobia/go-face-plusplus/config"
	"github.com/evalphobia/go-face-plusplus/log"
)

// BeautifyService is service struct for Beautify API.
type BeautifyService struct {
	client *client.Client
	logger log.Logger
}

// New creates BeautifyService from Config data.
func New(conf config.Config) (*BeautifyService, error) {
	cli, err := conf.Client()
	if err != nil {
		return nil, err
	}

	return &BeautifyService{
		client: cli,
		logger: log.DefaultLogger,
	}, nil
}

// SetLogger sets internal API logger.
func (s *BeautifyService) SetLogger(logger log.Logger) {
	s.logger = logger
}
