package main

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/bwmarrin/discordgo"
)

type Configuration struct {
    Token string
}

func main() {
    file, _ := os.Open("config.json")
    defer file.Close()

    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    }

    // Create a new Discord session using the provided bot token.
    dg, err := discordgo.New("Bot " + configuration.Token)
    if err != nil {
        fmt.Println("Error creating Discord session: ", err)
        return
    }

    // Register messageCreate as a callback for the messageCreate events.
    dg.AddHandler(messageCreate)

    // Open the websocket and begin listening.
    err = dg.Open()
    if err != nil {
        fmt.Println("Error opening Discord session: ", err)
    }

    // Wait until the bot is closed.
    fmt.Println("Moderation bot is now running. Press CTRL-C to exit.")
    <-make(chan struct{})
    return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Ignore all messages created by the bot itself
    // This isn't required in this specific example but it's a good practice.
    if m.Author.ID == s.State.User.ID {
        return
    }

    // If the message is "!Kick"
    if m.Content == "!Kick" {
        // Get the user object of the mentioned user
        user, err := s.User(m.Mentions[0].ID)
        if err != nil {
            fmt.Println("Error getting user: ", err)
            return
        }
        // Kick the user
        s.GuildMemberDeleteWithReason(m.GuildID, user.ID, "Kicked by moderator")
    }
}
