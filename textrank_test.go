package textrank

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRegularUsage(t *testing.T) {
	rawText := "Over the past fortnight we asked you to nominate your top extensions for the GNOME desktop. And you did just that. Having now sifted through the hundreds of entries, we’re ready to reveal your favourite GNOME Shell extensions. GNOME 3 (which is more commonly used with the GNOME Shell) has an extension framework that lets developers (and users) extend, build on, and shape how the desktop looks, acts and functions. Dash to Dock takes the GNOME Dash — this is the ‘favourites bar’ that appears on the left-hand side of the screen in the Activities overlay — and transforms it into a desktop dock. And just like Plank, Docky or AWN you can add app launchers, rearrange them, and use them to minimise, restore and switch between app windows. Dash to Dock has many of the common “Dock” features you’d expect, including autohide and intellihide, a fixed-width mode, adjustable icon size, and custom themes. My biggest pet peeve with GNOME Shell is its legacy app tray that hides in the bottom left of the screen. All extraneous non-system applets, indicators and tray icons hide down here. This makes it a little harder to use applications that rely on a system tray presence, like Skype, Franz, Telegram, and Dropbox. TopIcons Plus is the quick way to put GNOME system tray icons back where they belong: on show and in reach. The extension moves legacy tray icons from the bottom left of Gnome Shell to the right-hand side of the top panel. A well-stocked settings panel lets you adjust icon opacity, color, padding, size and tray position. Dive into the settings to adjust the sizing, styling and positioning of icons. Like the popular daily stimulant of choice, the Caffeine GNOME extension keeps your computer awake. It couldn’t be simpler to use: just click the empty mug icon. An empty cup means you’re using normal auto suspend rules – e.g., a screensaver – while a freshly brewed cup of coffee means auto suspend and screensaver are turned off. The Caffeine GNOME extension supports GNOME Shell 3.4 or later. Familiar with applications like Guake and Tilda? If so, you’ll instantly see the appeal of the (superbly named) Drop Down Terminal GNOME extension. When installed just tap the key above the tab key (though it can be changed to almost any key you wish) to get instant access to the command line. Want to speed up using workspaces? This simple tool lets you do just that. Once installed you can quickly switch between workspaces by scrolling over the top panel - no need to enter the Activities Overlay!"
	rule := CreateDefaultRule()
	language := CreateDefaultLanguage()
	algorithm := CreateDefaultAlgorithm()

	Append(rawText, language, rule, 1)
	Ranking(1, algorithm)

	FindPhrases(1)
	FindSingleWords(1)
	FindSentencesByRelationWeight(1,6)
	FindSentencesByPhraseChain(1, []string{
		"gnome",
		"shell",
		"favourite",
	})
	FindSentencesFrom(1, 2, 2)

	assert.Equal(t, 1, 1)
}

func TestExtensiveAsyncUsage(t *testing.T) {
	exitTest := false
	stream := make(chan string)

	go func() {
		i := 1

		for {
			rawText := <-stream
			rule := CreateDefaultRule()
			language := CreateDefaultLanguage()
			algorithm := CreateDefaultAlgorithm()

			Append(rawText, language, rule, 2)
			Ranking(2, algorithm)
			//FindPhrases(2)

			if i == 5 {
				exitTest = true
			}

			i++
		}
	}()

	go func() {
		rawTexts := []string{
			"Over the past fortnight we asked you to nominate your top extensions for the GNOME desktop.",
			"And you did just that. Having now sifted through the hundreds of entries, we’re ready to reveal your favourite GNOME Shell extensions. GNOME 3 (which is more commonly used with the GNOME Shell) has an extension framework that lets developers (and users) extend, build on, and shape how the desktop looks, acts and functions.",
			"Dash to Dock takes the GNOME Dash — this is the ‘favourites bar’ that appears on the left-hand side of the screen in the Activities overlay — and transforms it into a desktop dock. And just like Plank, Docky or AWN you can add app launchers, rearrange them, and use them to minimise, restore and switch between app windows. Dash to Dock has many of the common “Dock” features you’d expect, including autohide and intellihide, a fixed-width mode, adjustable icon size, and custom themes.",
			"My biggest pet peeve with GNOME Shell is its legacy app tray that hides in the bottom left of the screen. All extraneous non-system applets, indicators and tray icons hide down here. This makes it a little harder to use applications that rely on a system tray presence, like Skype, Franz, Telegram, and Dropbox. TopIcons Plus is the quick way to put GNOME system tray icons back where they belong: on show and in reach. The extension moves legacy tray icons from the bottom left of Gnome Shell to the right-hand side of the top panel. A well-stocked settings panel lets you adjust icon opacity, color, padding, size and tray position. Dive into the settings to adjust the sizing, styling and positioning of icons. Like the popular daily stimulant of choice, the Caffeine GNOME extension keeps your computer awake.",
			"It couldn’t be simpler to use: just click the empty mug icon. An empty cup means you’re using normal auto suspend rules – e.g., a screensaver – while a freshly brewed cup of coffee means auto suspend and screensaver are turned off. The Caffeine GNOME extension supports GNOME Shell 3.4 or later. Familiar with applications like Guake and Tilda? If so, you’ll instantly see the appeal of the (superbly named) Drop Down Terminal GNOME extension. When installed just tap the key above the tab key (though it can be changed to almost any key you wish) to get instant access to the command line. Want to speed up using workspaces? This simple tool lets you do just that. Once installed you can quickly switch between workspaces by scrolling over the top panel - no need to enter the Activities Overlay!",
		}

		for _, rawText := range rawTexts {
			stream <- rawText
		}
	}()

	for !exitTest {
		time.Sleep(time.Second * 1)
	}
}
