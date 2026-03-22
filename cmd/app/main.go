package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/y-maeda1116/discord-trans-helper/internal/config"
	"github.com/y-maeda1116/discord-trans-helper/internal/translator"
	"github.com/bwmarrin/discordgo"
)

func main() {
	cfg := config.Load()

	// Initialize Discord session
	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatalf("Failed to create Discord session: %v", err)
	}
	defer dg.Close()

	// Register handlers
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "translate" {
			translator.HandleTranslate(s, i, cfg.DeepLAuthKey)
		}
	})
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %s#%s", s.State.User.Username, s.State.User.Discriminator)
	})

	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	// Open Discord connection
	if err := dg.Open(); err != nil {
		log.Fatalf("Failed to open Discord connection: %v", err)
	}

	// Register slash commands
	_, err = dg.ApplicationCommandCreate(dg.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "translate",
		Description: "翻訳コマンド（メッセージコンテキストメニューから使用）",
		Type:        discordgo.MessageApplicationCommand,
	})
	if err != nil {
		log.Printf("Failed to create command: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log.Println("Discord bot started. Press Ctrl+C to stop.")

	<-ctx.Done()
	log.Println("Shutting down gracefully...")
}
