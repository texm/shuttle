package ui

import (
	"fmt"
	"log"
	"net/url"
	"time"

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

var quited = false

var logo = `
   _____ __  ____  ________________    ______
  / ___// / / / / / /_  __/_  __/ /   / ____/
  \__ \/ /_/ / / / / / /   / / / /   / __/
 ___/ / __  / /_/ / / /   / / / /___/ /___
/____/_/ /_/\____/ /_/   /_/ /_____/_____/
`

func Main(brg *bridge.Bridge) {
	for quited == false {
		if brg.IsLoggedIn {
			ChatUI(brg)
		} else {
			LoginUI(brg)
		}
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
	ui.ClearKeybindings()
	if err != nil {
		log.Fatal(err)
	}

	loginButton.OnActivated(func(b *tui.Button) {
		err := brg.SetCredentials(userId.Text(), authToken.Text())
		if err == nil {
			ui.Quit()
		}
	})

	googleButton.OnActivated(func(b *tui.Button) {
		_ = brg.LoginWithGoogle()
		ui.Quit()
	})

	tui.DefaultFocusChain.Set(userId, authToken, loginButton, googleButton)

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func ChatUI(brg *bridge.Bridge) {

	state := brg.GetInterfaceState()

	messages, _ := brg.GetPastMessages(&state.CurChannel, 100)
	for i := len(messages) - 1; i >= 0; i-- {
		post := &post{}
		post.username = messages[i].User.UserName
		post.message = messages[i].Msg
		post.time = messages[i].Timestamp.Format("15:04")

		posts = append(posts, *post)
	}

	t := tui.NewTheme()
	normal := tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack}
	t.SetStyle("normal", normal)
	inverse := tui.Style{Bg: tui.ColorBlack, Fg: tui.ColorWhite}
	t.SetStyle("inverse", inverse)

	channelsResponse, err := brg.GetJoinedChannels(url.Values{})

	sidebar := tui.NewVBox()
	sidebar.Append(tui.NewHBox(tui.NewLabel("CHANNELS")))
	chs := tui.NewList()
	sidebar.Append(chs)

	if err == nil {
		for i, c := range channelsResponse.Channels {
			isSelected := false
			if state.CurChannel.Name == c.Name {
				isSelected = true
			}
			chs.AddItems(c.Name)
			if isSelected {
				chs.SetSelected(i)
			}
			// sidebar.Append(tui.NewHBox(tui.NewLabel(prefix + c.Name)))
		}
	} else {
		fmt.Println(err)
	}
	ts := tui.NewTheme()
	ts.SetStyle("list.item", tui.Style{Bg: tui.ColorCyan, Fg: tui.ColorMagenta})
	ts.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorRed, Fg: tui.ColorWhite})
	ts.SetStyle("normal", normal)

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

			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", brg.User.UserName))),
			tui.NewLabel(e.Text()),
			tui.NewSpacer(),
		))
		brg.SendMessage(e.Text(), &state.CurChannel)
		input.SetText("")
	})

	root := tui.NewHBox(sidebarBox, chat)
	ui, err := tui.New(root)

	go func() {
		for {
			RefreshChat(historyScroll, brg, chs.SelectedItem())
			ui.Repaint()
			time.Sleep(1 * time.Second)
		}
	}()

	ui.ClearKeybindings()
	ui.SetTheme(ts)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Ctrl+C", func() {
		ui.Quit()
		quited = true
	})

	ui.SetKeybinding("Esc", func() {
		input.SetFocused(false)
		chs.SetFocused(true)
		brg.SetPaneState(bridge.CHANNEL_PANE)
	})

	ui.SetKeybinding("Ctrl+[", func() {
		input.SetFocused(false)
		chs.SetFocused(true)
		brg.SetPaneState(bridge.CHANNEL_PANE)
	})

	ui.SetKeybinding("Ctrl+K", func() {
		chs.SetFocused(true)
		brg.SetPaneState(bridge.CHANNEL_PANE)
	})

	ui.SetKeybinding("Ctrl+N", func() {
		ChangeChSelected(chs, +1)
		input.SetFocused(false)
		chs.SetFocused(true)
		brg.SetPaneState(bridge.CHANNEL_PANE)
	})

	ui.SetKeybinding("Ctrl+P", func() {
		ChangeChSelected(chs, -1)
		input.SetFocused(false)
		chs.SetFocused(true)
		brg.SetPaneState(bridge.CHANNEL_PANE)
	})

	ui.SetKeybinding("I", func() {
		if brg.GetPaneState() == bridge.INPUT_PANE {
			return
		}
		chs.SetFocused(false)
		input.SetFocused(true)
	})

	input.OnChanged(func(e *tui.Entry) {
		if brg.GetPaneState() != bridge.INPUT_PANE {
			input.SetText("")
			brg.SetPaneState(bridge.INPUT_PANE)
		}
	})
	brg.SetPaneState(bridge.MESSAGE_PANE)

	ui.SetKeybinding("Ctrl+U", func() {
		chs.SetFocused(false)
		input.SetFocused(false)

		// historyScroll.SetFocused(true)
		historyScroll.SetAutoscrollToBottom(false)
		historyScroll.Scroll(0, -2)
		// historyScroll.ScrollToBottom()
		ui.Repaint()
	})

	ui.SetKeybinding("Ctrl+D", func() {
		chs.SetFocused(false)
		input.SetFocused(false)
		historyScroll.SetFocused(true)
		historyScroll.Scroll(0, 2)
		historyScroll.SetAutoscrollToBottom(false)
	})

	ui.SetKeybinding("F", func() {
		if brg.GetPaneState() == bridge.INPUT_PANE {
			return
		}
		chs.SetFocused(false)
		input.SetFocused(false)
		historyScroll.SetFocused(true)
		historyScroll.ScrollToBottom()
		historyScroll.SetAutoscrollToBottom(true)
	})

	ui.SetKeybinding("G", func() {
		if brg.GetPaneState() == bridge.INPUT_PANE {
			return
		}
		chs.SetFocused(false)
		input.SetFocused(false)
		historyScroll.SetFocused(true)
		historyScroll.SetAutoscrollToBottom(false)
		historyScroll.ScrollToTop()
	})

	ui.SetKeybinding("Enter", func() {
		state := brg.GetPaneState()
		switch state {
		case bridge.CHANNEL_PANE:
			RefreshChat(historyScroll, brg, chs.SelectedItem())
			historyScroll.ScrollToBottom()
			historyScroll.SetAutoscrollToBottom(true)
		case bridge.MESSAGE_PANE:
		}
	})

	chs.OnItemActivated(func(chs *tui.List) {
		ui.Repaint()
		ChangeChSelected(chs, 1)
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}

	tui.DefaultFocusChain.Set(input, chs)

	for {
	}
}

func ChangeChSelected(l *tui.List, delta int) {
	selected := l.Selected()
	newSelected := selected + delta
	if newSelected >= l.Length() {
		newSelected = l.Length() - 1
	}

	if newSelected < 0 {
		newSelected = 0
	}

	if selected == newSelected {
		return
	}

	l.SetSelected(newSelected)
}

func RefreshChat(historyScroll *tui.ScrollArea, brg *bridge.Bridge, channelName string) {

	messages, _ := brg.GetPastMessagesByName(channelName, 200)
	if len(messages) == 0 {
		return
	}

	posts = []post{}
	history := tui.NewVBox()

	for i := len(messages) - 1; i >= 0; i-- {
		post := &post{}
		post.username = messages[i].User.UserName
		post.message = messages[i].Msg
		post.time = messages[i].Timestamp.Format("15:04")

		posts = append(posts, *post)
	}

	for _, m := range posts {
		history.Append(tui.NewHBox(
			tui.NewLabel(m.time),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.username))),
			tui.NewLabel(m.message),
			tui.NewSpacer(),
		))
	}
	historyScroll.Widget = history
}

func KeepRefreshChat(historyScroll *tui.ScrollArea, brg *bridge.Bridge, channelName string) {
	for {
		RefreshChat(historyScroll, brg, channelName)
		time.Sleep(2 * time.Second)
	}
}
