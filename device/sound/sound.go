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
	SampleRate int
	Src        io.Reader
}

type Sound struct {
	cfg    *Config
	ctx    *audio.Context
	player *audio.Player
}

func New(cfg *Config) (*Sound, error) {
	return &Sound{
		cfg: cfg,
		ctx: audio.NewContext(cfg.SampleRate),
	}, nil
}

// is the audio context ready?
func (s *Sound) IsReady() bool {
	return s.ctx.IsReady()
}

// start the sound player
func (s *Sound) Start() error {
	player, err := s.ctx.NewPlayer(s.cfg.Src)
	if err != nil {
		return err
	}
	s.player = player
	s.player.Play()
	return nil
}

//-----------------------------------------------------------------------------
