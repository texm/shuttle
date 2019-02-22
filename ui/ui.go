package ui

import (
	"log"
	"fmt"
	"time"
	"net/url"

	//"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/marcusolsson/tui-go"
	"github.com/texm/shuttle/bridge"
)

type StyledBox struct {
	Style string
	*tui.Box
}

func (s *StyledBox) Draw(p *tui.Painter) {
	p.WithStyle(s.Style, func(p *tui.Painter) {
		s.Box.Draw(p)
	})
}

type post struct {
	username string
	message  string
	time     string
}

var posts = []post{}

var logo = `
   _____ __  ____  ________________    ______
  / ___// / / / / / /_  __/_  __/ /   / ____/
  \__ \/ /_/ / / / / / /   / / / /   / __/   
 ___/ / __  / /_/ / / /   / / / /___/ /___   
/____/_/ /_/\____/ /_/   /_/ /_____/_____/										
`


func Main(brg *bridge.Bridge) {
	if (brg.IsLoggedIn) {
		ChatUI(brg)
	} else {
		LoginUI(brg)
	}
}

func LoginUI(brg *bridge.Bridge) {

	authToken := tui.NewEntry()

	userId := tui.NewEntry()
	userId.SetFocused(true)

	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("User Id"), tui.NewLabel("Auth Token"))
	form.AppendRow(userId, authToken)

	loginButton := tui.NewButton("Start Chatting")
	googleButton := tui.NewButton("Login With Google")

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, loginButton),
		tui.NewPadder(2, 0, googleButton),
	)

	mainMenu := tui.NewVBox(
		tui.NewPadder(10, 1, tui.NewLabel(logo)),
		tui.NewPadder(12, 0, tui.NewLabel("Welcome to Shuttle, Rocket.Chat CLI")),
		tui.NewPadder(1, 1, form),
		buttons,
	)
	
	mainMenu.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		mainMenu,
		tui.NewSpacer(),
	)

	root := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	loginButton.OnActivated(func(b *tui.Button) {
		err := brg.SetCredentials(userId.Text(), authToken.Text())
		if (err == nil) {
			ui.Quit()
		}
	})

	googleButton.OnActivated(func(b *tui.Button) {
		_ = brg.LoginWithGoogle()
		ui.Quit()
	})

	tui.DefaultFocusChain.Set(userId, authToken, loginButton, googleButton)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func ChatUI(brg *bridge.Bridge) {
	// SET UP SIDEBAR
	channelsResponse, err := brg.GetJoinedChannels(url.Values{})

	sidebar := tui.NewVBox()
	sidebar.Append(tui.NewHBox(tui.NewLabel("CHANNELS")))

	if err == nil {
		for _, c := range channelsResponse.Channels {
			sidebar.Append(tui.NewHBox(tui.NewLabel(c.Name)))
		}
	} else {
		fmt.Println(err)
	}

	sidebarScroll := tui.NewScrollArea(sidebar)
	sidebarBox := tui.NewVBox(sidebarScroll)
	sidebarBox.SetBorder(true)

	history := tui.NewVBox()

	for _, m := range posts {
		history.Append(tui.NewHBox(
			tui.NewLabel(m.time),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.username))),
			tui.NewLabel(m.message),
			tui.NewSpacer(),
		))
	}

	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		history.Append(tui.NewHBox(
			tui.NewLabel(time.Now().Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", "Putin"))),
			tui.NewLabel(e.Text()),
			tui.NewSpacer(),
		))
		input.SetText("")
	})

	root := tui.NewHBox(sidebarBox, chat)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
