package image_generation

import (
	"Twipex_project/database"
	"strconv"
)

func calculateValues(newdata apexLawData, beforedata database.UserData) variation {
	for i := 0; i < len(newdata.Data.Segments); i++ {
		if newdata.Data.Segments[i].Metadata.Name == beforedata.Legend {
			if beforedata.Legend != beforedata.BeforeLegend {
				va := variation{
					rankname:     newdata.Data.Segments[0].Stats.Rankscore.Rankmetadata.Rankname,
					rp:           strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value)),
					rpup:         strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.Rp)),
					kills:        strconv.Itoa(int(newdata.Data.Segments[i].Stats.Kills.Value)),
					killup:       "0",
					damage:       strconv.Itoa(int(newdata.Data.Segments[i].Stats.Damage.Value)),
					damageup:     "0",
					wins:         strconv.Itoa(int(newdata.Data.Segments[i].Stats.Wins.Value)),
					winup:        "0",
					rpupsend:     strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.LastMadeRp)),
					killupsend:   "0",
					damageupsend: "0",
					winupsend:    "0",
					date:         getTime().Format("2006/01/02"),
				}
				return va
			}

			va := variation{
				rankname:     newdata.Data.Segments[0].Stats.Rankscore.Rankmetadata.Rankname,
				rp:           strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value)),
				rpup:         strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.Rp)),
				kills:        strconv.Itoa(int(newdata.Data.Segments[i].Stats.Kills.Value)),
				killup:       strconv.Itoa(int(newdata.Data.Segments[i].Stats.Kills.Value) - (beforedata.Kills)),
				damage:       strconv.Itoa(int(newdata.Data.Segments[i].Stats.Damage.Value)),
				damageup:     strconv.Itoa(int(newdata.Data.Segments[i].Stats.Damage.Value) - (beforedata.Damage)),
				wins:         strconv.Itoa(int(newdata.Data.Segments[i].Stats.Wins.Value)),
				winup:        strconv.Itoa(int(newdata.Data.Segments[i].Stats.Wins.Value) - (beforedata.Wins)),
				rpupsend:     strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.LastMadeRp)),
				killupsend:   strconv.Itoa(int(newdata.Data.Segments[i].Stats.Kills.Value) - (beforedata.LastMadeKills)),
				damageupsend: strconv.Itoa(int(newdata.Data.Segments[i].Stats.Damage.Value) - (beforedata.LastMadeDamage)),
				winupsend:    strconv.Itoa(int(newdata.Data.Segments[i].Stats.Wins.Value) - (beforedata.LastMadeWins)),
				date:         getTime().Format("2006/01/02"),
			}
			return va
		}
	}

	//選択したレジェンドのデータがなかったときの処理
	va := variation{
		rankname: newdata.Data.Segments[0].Stats.Rankscore.Rankmetadata.Rankname,
		rp:       strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value)),
		rpup:     strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.Rp)),
		kills:    "0",
		killup:   "0",
		damage:   "0",
		damageup: "0",
		wins:     "0",
		winup:    "0",
	}
	return va
}

func calculateNextRank(r variation) int {
	val, _ := strconv.Atoi(r.rp)
	if val < 300 {
		return 300 - val
	}
	if val < 600 {
		return 600 - val
	}
	if val < 900 {
		return 900 - val
	}
	if val < 1200 {
		return 1200 - val
	}
	if val < 1600 {
		return 1600 - val
	}
	if val < 2000 {
		return 2000 - val
	}
	if val < 2400 {
		return 2400 - val
	}
	if val < 2800 {
		return 2800 - val
	}
	if val < 3300 {
		return 3300 - val
	}
	if val < 3800 {
		return 3800 - val
	}
	if val < 4300 {
		return 4300 - val
	}
	if val < 4800 {
		return 4800 - val
	}
	if val < 5400 {
		return 5400 - val
	}
	if val < 6000 {
		return 6000 - val
	}
	if val < 6600 {
		return 6600 - val
	}
	if val < 7200 {
		return 7200 - val
	}
	if val < 7900 {
		return 7900 - val
	}
	if val < 8600 {
		return 8600 - val
	}
	if val < 9300 {
		return 9300 - val
	}
	if val < 10000 {
		return 10000 - val
	}
	if val >= 10000 {
		return 0
	}
	return 0
}
