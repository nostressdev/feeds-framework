package feeds

import (
	"github.com/nostressdev/feeds-framework/internal/storage"
	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/signer"
)

type FeedsService struct {
	*Config
}

type Config struct {
	*signer.Signer
	Storage storage.FeedsStorage
}

func New(config *Config) proto.FeedsServer {
	return &FeedsService{
		Config: config,
	}
}
