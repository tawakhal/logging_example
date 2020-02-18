package logging

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestLogging(t *testing.T) {
	t.Fail()

	Setup()

	log.Info().Msg("This info message")
	log.Debug().Msg("This Debug message")
}
