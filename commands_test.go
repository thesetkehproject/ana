package main

import (
	"strings"
	"testing"
)

func TestHelpCmd(t *testing.T) {
	helpMsg := HelpCmd("!")
	if !strings.Contains(helpMsg, "Commands: !help, !ddg/search") {
		t.Errorf("HelpCmd(!), got %v, want 'Commands: !help, !ddg/search'\n", helpMsg)
	}
}

func TestWikiCmd(t *testing.T) {
	config := &Config{}
	config.WikiLink = "WIKIURL"
	wikiUrl := WikiCmd(config)
	if !strings.Contains(wikiUrl, "WIKIURL") {
		t.Errorf("WikiCmd(%v), got %v, want 'WikiUrl'\n", config, wikiUrl)
	}
}

func TestHomePageCmd(t *testing.T) {
	config := &Config{}
	config.Homepage = "HOMEURL"
	homeUrl := HomePageCmd(config)
	if !strings.Contains(homeUrl, "HOMEURL") {
		t.Errorf("HomePageCmd(%v), got %v, want 'HOMEURL'\n", config, homeUrl)
	}
}

func TestForumCmd(t *testing.T) {
	config := &Config{}
	config.Forums = "FORUMURL"
	forumUrl := ForumCmd(config)
	if !strings.Contains(forumUrl, "FORUMURL") {
		t.Errorf("ForumCmd(%v), got %v, want 'FORUMURL'\n", config, forumUrl)
	}
}
