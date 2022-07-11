package main

import (
	"betterdsc/ent"
	"betterdsc/ent/serveremote"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var (
	addEmoteCommand = &discordgo.ApplicationCommand{
		Name:        "addemote",
		Description: "Add an emote to the server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "emote-id",
				Description: "The emote to add",
				Required:    true,
			},
		},
	}
)

type Emote struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	ImageType string `json:"imageType"`
	Order     int
}

func getDefaultEmotes() map[string]Emote {
	client := &http.Client{}

	emotesChannel := make(chan []Emote, 26)
	wg := new(sync.WaitGroup)

	emotesByCode := make(map[string]Emote)
	wg.Add(1)
	go func() {
		getGlobalEmotes(client, emotesChannel)
		wg.Done()
	}()

	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			getTrendingEmotes(client, i*100, emotesChannel)
			wg.Done()
		}()
	}

	wg.Wait()
	close(emotesChannel)

	for result := range emotesChannel {
		for _, emote := range result {
			existingEmote, existsEmote := emotesByCode[emote.Code]
			if !existsEmote || existingEmote.Order < emote.Order {
				emotesByCode[emote.Code] = emote
			}
		}
	}

	return emotesByCode
}

func main() {
	godotenv.Load()
	token := os.Getenv("DISCORD_BOT_TOKEN")

	fmt.Println("Getting emote data")
	emoteByCode := getDefaultEmotes()

	client, err := ent.Open("sqlite3", "file:betterdsc.db?_fk=1")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	fmt.Println("Running migrations")
	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println("Starting bot")
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Logged in as: %v#%v\n", s.State.User.Username, s.State.User.Discriminator)
	})

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) { onMessageCreate(client, emoteByCode, s, m) })

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "addemote" {
			onAddEmote(client, s, i)
		}
	})

	// Setup connection
	err = discord.Open()
	if err != nil {
		panic(err)
	}
	defer discord.Close()
	defer fmt.Println("Closing bot")

	cmd, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", addEmoteCommand)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = discord.ApplicationCommandDelete(discord.State.User.ID, "", cmd.ID)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func getGlobalEmotes(client *http.Client, c chan []Emote) {
	url := "https://api.betterttv.net/3/cached/emotes/global"

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result := []Emote{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	for _, emote := range result {
		emote.Order = -1
	}

	c <- result
}

type TrendingEmote struct {
	Emote Emote `json:"emote"`
}

func getTrendingEmotes(client *http.Client, offset int, c chan []Emote) {
	url := "https://api.betterttv.net/3/emotes/shared/trending?limit=100&offset=" + fmt.Sprint(offset)

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result := []TrendingEmote{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	emotes := []Emote{}
	for _, trendingEmote := range result {
		trendingEmote.Emote.Order = offset
		emotes = append(emotes, trendingEmote.Emote)
	}

	c <- emotes
}

func onMessageCreate(client *ent.Client, emotesByCode map[string]Emote, s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	emote, _ := client.ServerEmote.Query().Where(
		serveremote.Code(m.Content),
		serveremote.ServerID(m.GuildID),
	).Only(context.Background())
	if emote != nil {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, m.Author.Username+":")
		s.ChannelMessageSend(m.ChannelID, "https://cdn.betterttv.net/emote/"+emote.EmoteID+"/2x."+emote.ImageType)
		return
	}

	if emote, ok := emotesByCode[m.Content]; ok {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, m.Author.Username+":")
		s.ChannelMessageSend(m.ChannelID, "https://cdn.betterttv.net/emote/"+emote.ID+"/2x."+emote.ImageType)
	}
}

func onAddEmote(client *ent.Client, s *discordgo.Session, i *discordgo.InteractionCreate) {
	emoteID := i.ApplicationCommandData().Options[0].Value.(string)

	result, err := http.Get("https://api.betterttv.net/3/emotes/" + emoteID)
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()

	if result.StatusCode != 200 {
		return
	}

	var emote Emote
	err = json.NewDecoder(result.Body).Decode(&emote)
	if err != nil {
		panic(err)
	}

	rowsChanged, err := client.ServerEmote.Delete().Where(
		serveremote.Code(emote.Code),
		serveremote.ServerID(i.GuildID),
	).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted " + fmt.Sprintf("%v", rowsChanged) + "emotes")

	serverEmote, err := client.ServerEmote.Create().SetEmoteID(emoteID).SetServerID(i.GuildID).SetCode(emote.Code).SetImageType(emote.ImageType).Save(context.Background())
	if err != nil {
		panic(err)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Added emote: " + emote.Code,
		},
	})

	fmt.Printf("Added emote %v to server %v\n", serverEmote.EmoteID, i.GuildID)
}
