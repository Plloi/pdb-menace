package menace

import (
	"github.com/Plloi/pdb-cmdr/pkg/router"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Setup Registers package commands to a command router
func Setup(r *router.CommandRouter) {
	log.Info("Setting up Menace 'Storat Command' command")
	r.RegisterCommand("start", "Start a new Game of TicTacToe", NewGame)
}

func NewGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Not Implemented")
}
