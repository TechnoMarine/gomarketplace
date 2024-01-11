package common

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GetRandomName() string {
	names := []string{
		"roar",
		"studentw",
		"attend",
		"malboro",
		"conference",
		"language",
		"junior",
		"army",
		"analytical",
		"classic",
		"distracted",
		"attentive",
		"shade",
		"dearest",
		"wallow",
		"wanting",
		"alienated",
		"spoke",
	}

	return names[rand.Intn(len(names))]
}

func GetRandomRole() string {
	roles := []string{
		"Admin", "Moder", "Guest", "Customer",
	}

	return roles[rand.Intn(len(roles))]
}

func GetRandomArticleName() string {
	articleNames := []string{
		"Serendipity",
		"Luminescence",
		"Resilience",
		"Zenith",
		"Spectrum",
		"Equilibrium",
		"Pinnacle",
		"Ethereal",
		"Catalyst",
		"Utopia",
		"Harmony",
		"Echo",
		"Renaissance",
		"Quasar",
		"Ephemeral",
		"Nebula",
		"Cascade",
		"Vortex",
		"Panorama",
		"Jubilee",
	}

	return articleNames[rand.Intn(len(articleNames))]
}

func GetTextMock() string {
	return "Lorem ipsum is dolor set amet..."
}

func GetRandomStatuses() string {
	articleStatuses := []string{
		"Published",
		"Draft",
		"Under Review",
		"Scheduled",
		"Rejected",
	}

	return articleStatuses[rand.Intn(len(articleStatuses))]
}
