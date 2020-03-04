package moor

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	. "github.com/franela/goblin"
	"github.com/kennygrant/sanitize"
	log "github.com/sirupsen/logrus"
)

func TestHttpClient(t *testing.T) {
	g := Goblin(t)

	g.Describe("GET", func() {
		g.It("downloads and returns body for URL", func() {
			url := "https://blog.marahin.pl/?format=json"
			r, err := recorder.New(fmt.Sprintf("fixtures/%s", sanitize.BaseName(url)))
			if err != nil {
				log.Fatal(err)
			}

			defer r.Stop()
			result := Get(url, &http.Client{Transport: r})

			g.Assert(len(result) > 0).Eql(true)
		})

		g.It("removes N number of characters from the beginning of string", func() {
			url := "https://blog.marahin.pl/?format=json"
			r, err := recorder.New(fmt.Sprintf("fixtures/%s", sanitize.BaseName(url)))
			if err != nil {
				log.Fatal(err)
			}

			defer r.Stop()
			result := Get(url, &http.Client{Transport: r})

			g.Assert(result[0:BlockerCharactersAmount()] == "])}while(1);</x>").Equal(false)
		})
	})

	g.Describe("BlockerCharactersAmount", func() {
		g.It("returns default constant BLOCKER_CHARACTERS_AMOUNT if environment variable was not present", func() {
			g.Assert(os.Getenv("MOOR_BLOCKER_CHARACTERS_AMOUNT")).Equal("")
			g.Assert(BlockerCharactersAmount()).Equal(BLOCKER_CHARACTERS_AMOUNT)
		})

		g.It("returns custom amount if environment variable was present", func() {
			val := rand.Int()
			_ = os.Setenv("MOOR_BLOCKER_CHARACTERS_AMOUNT", fmt.Sprintf("%d", val))
			g.Assert(os.Getenv("MOOR_BLOCKER_CHARACTERS_AMOUNT") == "").Equal(false)
			g.Assert(BlockerCharactersAmount()).Equal(val)
		})
	})

	g.Describe("SanitizeUrl", func() {
		g.It("adds http:// in front of passed url if its not present", func() {
			g.Assert(SanitizeUrl("example.org")).Equal("http://example.org")
		})

		g.It("leaves http:// alone if its present already", func() {
			g.Assert(SanitizeUrl("http://example.org")).Equal("http://example.org")
		})

		g.It("leaves https:// alone if its present already", func() {
			g.Assert(SanitizeUrl("https://example.org")).Equal("https://example.org")
		})
	})
}
