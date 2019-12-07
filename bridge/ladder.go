package bridge

import (
	"github.com/nokka/slashdiablo-launcher/ladder"
	"github.com/nokka/slashdiablo-launcher/log"
	"github.com/therecipe/qt/core"
)

// LadderBridge is the connection between QML and the ladder model.
type LadderBridge struct {
	core.QObject

	// Dependencies.
	ladderService ladder.Service
	logger        log.Logger

	// Properties.
	_ bool `property:"loading"`
	_ bool `property:"error"`

	// Models.
	LadderModel *core.QAbstractListModel `property:"characters"`

	// Slots.
	_ func(mode string) `slot:"getLadder"`
}

// Connect will connect the QML signals to functions in Go.
func (b *LadderBridge) Connect() {
	b.ConnectGetLadder(b.getLadder)
}

func (b *LadderBridge) getLadder(mode string) {
	go func() {
		// Tell the GUI that we're fetching data.
		b.SetLoading(true)

		// Set the ladder characters on the model.
		err := b.ladderService.SetLadderCharacters(mode)

		// Stop loading when we're done fetching ladder data.
		b.SetLoading(false)

		if err != nil {
			b.logger.Error(err)
			b.SetError(true)
			return
		}
	}()

	return
}

// NewLadder creates a new ladder bridge with all dependencies set up.
func NewLadder(ls ladder.Service, lm *ladder.TopLadderModel, logger log.Logger) *LadderBridge {
	l := NewLadderBridge(nil)

	// Setup dependencies.
	l.ladderService = ls
	l.logger = logger

	// Setup model.
	l.SetCharacters(lm)

	// Set initial state.
	l.SetLoading(false)
	l.SetError(false)

	return l
}
