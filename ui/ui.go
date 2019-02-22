package ui

import (
	"log"
	"fmt"

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

var posts = []post{
	{username: "john", message: "hi, what's up?", time: "14:41"},
	{username: "jane", message: "not much", time: "14:43"},
}

var logo = `
   _____ __  ____  ________________    ______
  / ___// / / / / / /_  __/_  __/ /   / ____/
  \__ \/ /_/ / / / / / /   / / / /   / __/   
 ___/ / __  / /_/ / / /   / / / /___/ /___   
/____/_/ /_/\____/ /_/   /_/ /_____/_____/										
`


func Main(brg *bridge.Bridge) {
	if (brg.IsLoggedIn) {
		fmt.Println("logged in!")
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
