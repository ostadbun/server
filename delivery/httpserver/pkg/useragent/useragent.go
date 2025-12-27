package useragent

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type DeviceInfo struct {
	// mobile | desktop
	Type string `json:"type"`

	// web | terminal
	Client string `json:"client"`

	// android | ios | windows | mac | linux
	OS string `json:"os"`
}

func ReadDeviceInfo(c *fiber.Ctx) DeviceInfo {

	ua := strings.ToLower(c.Get("User-Agent"))
	platform := strings.ToLower(strings.Trim(
		c.Get("Sec-CH-UA-Platform"),
		`"`,
	))
	mobile := c.Get("Sec-CH-UA-Mobile")

	// client
	client := "terminal"
	if strings.Contains(ua, "mozilla") {
		client = "web"
	}

	// os
	os := "unknown"
	switch {
	case platform == "android" || strings.Contains(ua, "android"):
		os = "android"
	case platform == "ios" || strings.Contains(ua, "iphone"):
		os = "ios"
	case platform == "windows" || strings.Contains(ua, "windows"):
		os = "windows"
	case platform == "macos" || strings.Contains(ua, "mac"):
		os = "mac"
	case platform == "linux" || strings.Contains(ua, "linux"):
		os = "linux"
	}

	// device type
	device := "desktop"
	if mobile == "?1" || strings.Contains(ua, "mobile") {
		device = "mobile"
	}

	return DeviceInfo{
		Type:   device,
		Client: client,
		OS:     os,
	}
}
