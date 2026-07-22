//-----------------------------------------------------------------------------
/*

Ebiten Audio Consumer

*/
//-----------------------------------------------------------------------------

package sound

import (
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

//-----------------------------------------------------------------------------

type Config struct {
	Enable     bool
	SampleRate int
	Src        io.Reader
}

type Sound struct {
	cfg    Config
	ctx    *audio.Context
	player *audio.Player
}

func New(cfg Config) (*Sound, error) {
	if !cfg.Enable {
		return &Sound{
			cfg: cfg,
		}, nil
	}
	return &Sound{
		cfg: cfg,
		ctx: audio.NewContext(cfg.SampleRate),
	}, nil
}

// is the audio context ready?
func (s *Sound) IsReady() bool {
	if !s.cfg.Enable {
		return false
	}
	return s.ctx.IsReady()
}

// start the sound player
func (s *Sound) Start() error {
	if !s.cfg.Enable {
		return nil
	}
	player, err := s.ctx.NewPlayer(s.cfg.Src)
	if err != nil {
		return err
	}
	s.player = player
	s.player.Play()
	return nil
}

//-----------------------------------------------------------------------------
