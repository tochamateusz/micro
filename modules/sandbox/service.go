package sandbox

import "go.uber.org/zap"

type Sandbox struct {
	Test string
}

func New(log *zap.SugaredLogger) *Sandbox {
	log.Infow("Creating Sandbox")
	return &Sandbox{}
}
