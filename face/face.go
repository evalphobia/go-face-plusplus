package face

import (
	"github.com/evalphobia/go-face-plusplus/client"
	"github.com/evalphobia/go-face-plusplus/config"
	"github.com/evalphobia/go-face-plusplus/log"
)

// FaceService is service struct for Face API.
type FaceService struct {
	client *client.Client
	logger log.Logger
}

// New creates FaceService from Config data.
func New(conf config.Config) (*FaceService, error) {
	cli, err := conf.Client()
	if err != nil {
		return nil, err
	}

	return &FaceService{
		client: cli,
		logger: log.DefaultLogger,
	}, nil
}

// SetLogger sets internal API logger.
func (s *FaceService) SetLogger(logger log.Logger) {
	s.logger = logger
}
