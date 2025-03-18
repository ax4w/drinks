package main

import "strings"

func firstEmoji(s string) string {
	for k, v := range nameToEmoji {
		if strings.Contains(strings.ToLower(s), k) {
			return v
		}
	}
	return ""
}

func firstEmojiFromStrings(strs []string) string {
	for _, v := range strs {
		if emoji, ok := nameToEmoji[v]; ok {
			return emoji
		}
	}
	return ""
}

var (
	nameToEmoji = map[string]string{
		"lemon":           "🍋",
		"lime":            "🍋",
		"orange":          "🍊",
		"blood orange":    "🍊",
		"grapefruit":      "🍊",
		"pineapple":       "🍍",
		"coconut":         "🥥",
		"strawberry":      "🍓",
		"raspberry":       "🍓",
		"blueberry":       "🫐",
		"cherry":          "🍒",
		"peach":           "🍑",
		"mango":           "🥭",
		"watermelon":      "🍉",
		"grape":           "🍇",
		"grapes":          "🍇",
		"banana":          "🍌",
		"apple":           "🍎",
		"rum":             "🥃",
		"vodka":           "🥃",
		"wine":            "🍷",
		"gin":             "🍸",
		"tequila":         "🍾",
		"whiskey":         "🥃",
		"brandy":          "🍷",
		"cointreau":       "🍊",
		"vermouth":        "🍸",
		"cola":            "🫧",
		"tonic water":     "💧",
		"ice":             "🧊",
		"soda water":      "🫧",
		"sprite":          "🫧",
		"fanta":           "🫧",
		"grenadine":       "🌺",
		"cranberry juice": "🍒",
		"pineapple juice": "🍍",
		"orange juice":    "🍊",
		"coconut milk":    "🥥",
		"coffee":          "☕",
		"simple syrup":    "🍯",
		"honey":           "🍯",
		"maple syrup":     "🍁",
		"agave syrup":     "🌵",
		"mint":            "🌿",
		"basil":           "🌿",
		"ginger":          "🫚",
		"cinnamon":        "🥄",
		"nutmeg":          "🌰",
		"chili":           "🌶️",
		"salt":            "🧂",
		"pepper":          "⚫",
		"olive":           "🫒",
		"lemon peel":      "🍋",
		"sugar rim":       "🍬",
		"cocktail cherry": "🍒",
		"ice cubes":       "🧊",
		"cucumber":        "🥒",
		"chocolate":       "🍫",
		"whipped cream":   "🍦",
		"egg":             "🥚",
		"bitters":         "🌿",
		"water":           "💧",
		"lime leaf":       "🍃",
		"rose petals":     "🌹",
	}
)
