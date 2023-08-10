package Utils

func CalculateCurrentLevelFromXP(xp int) int {
	if xp < 100 {
		return 0
	} else if xp >= 101 && xp < 224 {
		return 1
	} else if xp >= 225 && xp < 374 {
		return 2
	} else {
		levelCalc := xp - 375
		levelCalc = (levelCalc / 200) + 2
		return levelCalc
	}
}
