package services

import goaway "github.com/TwiN/go-away"

func CustomProfanityDetector() goaway.ProfanityDetector {
	profanities := []string{
		"cunt",
		"fag",
		"faggot",
		"fuck",
		"fucker",
		"nigga",
		"retard",
		"retarded",
		"shit",
	}

	profanityDetector := goaway.NewProfanityDetector().WithCustomDictionary(profanities, goaway.DefaultFalsePositives, goaway.DefaultFalseNegatives)

	return *profanityDetector
}
