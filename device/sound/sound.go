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
