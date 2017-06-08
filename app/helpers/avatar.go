package helpers

import (
	"math/rand"
	"time"
)

func getRandomAvatarTheme() string {
	var themes = []string{"frogideas", "sugarsweets", "heatwave", "daisygarden", "seascape", "summerwarmth", "berrypie", "bythepool", "duskfalling"}
	rand.Seed(time.Now().Unix())
	return themes[rand.Intn(len(themes))]
}

// GetAvatarURL func
func GetAvatarURL(email string) string {
	theme := getRandomAvatarTheme()
	return "/avatar/" + email + "?theme=" + theme + "&numcolors=4&fmt=svg&size=460"
}

// GetLogoURL func
func GetLogoURL(handle string) string {
	theme := getRandomAvatarTheme()
	return "/avatar/" + handle + "?theme=" + theme + "&numcolors=4&fmt=svg&size=460"
}
