package image_generation

import (
	"Twipex_project/database"
	"Twipex_project/twitter"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type variation struct {
	rankname     string
	rp           string
	rpup         string
	kills        string
	killup       string
	damage       string
	damageup     string
	wins         string
	winup        string
	rpupsend     string
	killupsend   string
	damageupsend string
	winupsend    string
	date         string
}

type userMetainfo struct {
	id       string
	avatar   string
	rank     string
	platform string
	legend   string
}

func PostImage(now string) {
	log.Printf(now + "work start")
	senddate := getTime().Format("02")
	queue := database.GetPostUser(now)
	for i := range queue {
		userinfo := database.GetOne(queue[i].AccountId)
		if senddate != userinfo.Lastsenddate {
			userdata := getApexData(userinfo.Platform, userinfo.UserId)
			if userdata == nil {
				log.Printf("Failed to get data | platform:%v id:%v twitter:%v", userinfo.Platform, userinfo.UserId, userinfo.AccountName)
			}
			if userdata != nil {
				retrievedata := userdata[0]

				variation := makeImage(retrievedata, userinfo)
				variation.saveLog(userinfo)
				variation.updateUserData(userinfo)

				if queue[i].SendInterval == "week" {
					weekday := getTime().Weekday()
					if weekday == 1 {
						err := twitter.PostTweet(userinfo)
						if err == nil {
							log.Printf("send successfully @%v", userinfo.AccountName)
							variation.updateLastMadeTime(userinfo)
						}
					}
					if userinfo.Lastsenddate == "" {
						variation.updateLastMadeTime(userinfo)
					}
				} else {
					err := twitter.PostTweet(userinfo)
					if err == nil {
						log.Printf("send successfully @%v", userinfo.AccountName)
						variation.updateLastMadeTime(userinfo)
					}
				}
				time.Sleep(time.Second * 2)
			}
		}
	}
	log.Printf(now + "work end")
}

func makeImage(newdata apexLawData, beforedata database.UserData) variation {
	usermetainfo := userMetainfo{
		id:       beforedata.UserId,
		avatar:   newdata.Data.PlatformInfo.AvatarURL,
		rank:     newdata.Data.Segments[0].Stats.Rankscore.Rankmetadata.Rankname,
		platform: beforedata.Platform,
		legend:   beforedata.Legend,
	}
	if beforedata.Predator == "on" && newdata.Data.Segments[0].Stats.Rankscore.Value >= 10000 {
		usermetainfo.rank = "Apex Predator"
	}

	dc := gg.NewContext(1080, 607)
	dc.DrawRectangle(0, 0, 1080, 607)

	flagImage := openImage("background/" + beforedata.Legend + ".png")
	dc.DrawImage(flagImage, 0, 0)

	area := openImage("area.png")
	dc.DrawImage(area, 0, 0)

	makeQrcode(beforedata.AccountId)
	usermetainfo.drawUserMetainfo(dc)

	variation := calculateValues(newdata, beforedata)
	variation.drawRp(dc)
	variation.drawKills(dc)
	variation.drawDamage(dc)
	if beforedata.Winad != "" {
		variation.drawWins(dc)
	}

	drawDate(dc, beforedata)

	dc.SavePNG("data.png")

	return variation
}

func drawDate(dc *gg.Context, beforedata database.UserData) {
	dc.SetRGB(0.85, 0.85, 0.85)
	setSize(24, dc)
	today := getTime().Format("2006/01/02")
	yesterday := getTime().AddDate(0, 0, -1).Format("2006/01/02")

	if beforedata.SendInterval == "week" {
		if beforedata.LastMadeDate == yesterday || beforedata.LastMadeDate == "" {
			dc.DrawString(today, 920, 595)
		} else {
			dc.DrawString(beforedata.LastMadeDate+" ~ "+today, 800, 595)
		}
	} else {
		dc.DrawString(today, 920, 595)
	}
}

func makeQrcode(twitterid string) {
	qrCode, _ := qr.Encode("https://twipex.herokuapp.com//data/"+twitterid, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("image_generation/material/qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}

func (u userMetainfo) drawUserMetainfo(dc *gg.Context) {
	getAvatar(u.avatar)
	avatar := openImage("avatar.jpg")
	if avatar == nil {
		avatar = openImage("avatar_failed.jpg")
	}
	platform := openImage(u.platform + "48.png")
	rank := openImage("rank/" + u.rank + ".jpg")
	dc.DrawImage(resize.Resize(156, 156, rank, resize.Lanczos3), 274, 44)
	dc.DrawImage(resize.Resize(208, 208, avatar, resize.Lanczos3), 20, 20)
	dc.DrawImage(platform, 476, 60)

	setFont()
	setSize(20, dc)
	dc.SetRGB(0.85, 0.85, 0.85)
	if u.platform == "origin" {
		dc.DrawString("Origin", 524, 100)
	} else if u.platform == "psn" {
		dc.DrawString("PSN", 524, 100)
	} else if u.platform == "xbl" {
		dc.DrawString("XBOX", 524, 100)
	}
	setSize(60, dc)
	dc.DrawString(u.id, 476, 165)

	legendbunner := openImage("tiles/" + u.legend + ".png")
	dc.DrawImage(resize.Resize(250, 279, legendbunner, resize.Lanczos3), 240, 275)
	flagImage := openImage("qrcode.png")
	dc.SetRGB(0.85, 0.85, 0.85)
	setSize(38, dc)
	dc.DrawString(u.legend, 40, 320)
	dc.SetRGB(0.75, 0.75, 0.75)
	setSize(20, dc)
	dc.DrawString("Active Legend", 40, 280)
	dc.DrawString("↑ Check details", 52, 520)

	dc.DrawImage(resize.Resize(140, 140, flagImage, resize.Lanczos3), 54, 350)
}

func (v variation) drawRp(dc *gg.Context) {
	dc.SetRGB(0.85, 0.85, 0.85)
	setSize(50, dc)

	dc.DrawString(v.rp, 360, 280)

	dc.SetRGB(0.75, 0.75, 0.75)
	dc.DrawString("RP", 280, 280)
	setSize(20, dc)
	dc.DrawString("To the next rank", 700, 280)
	setSize(30, dc)
	nextrank := strconv.Itoa(calculateNextRank(v))
	if nextrank == "0" {
		nextrank = "-"
	}
	dc.DrawString(nextrank, 875, 280)

	if v.rpup != "0" && v.rpupsend != v.rp {
		setSize(40, dc)
		val, _ := strconv.Atoi(v.rpup)
		if val > 0 {
			dc.SetRGB(0.8, 0.3, 0.3)
			dc.DrawString(v.rpupsend, 545, 280)
			flagImage := openImage("up32.png")
			dc.DrawImage(flagImage, 510, 250)
		} else {
			dc.SetRGB(0.25, 0.67, 0.83)
			dc.DrawString(v.rpupsend, 545, 280)
			flagImage := openImage("down32.png")
			dc.DrawImage(flagImage, 510, 250)
		}

	}
}

func (v variation) drawKills(dc *gg.Context) {
	dc.SetRGB(0.75, 0.75, 0.75)
	setSize(40, dc)
	dc.DrawString("Kills", 490, 380)
	dc.SetRGB(0.85, 0.85, 0.85)
	dc.DrawString(v.kills, 660, 380)

	if v.killup != "0" && v.killup != v.kills {
		setSize(40, dc)
		dc.SetRGB(0.8, 0.3, 0.3)
		dc.DrawString(v.killup, 885, 380)
		flagImage := openImage("up32.png")
		dc.DrawImage(flagImage, 850, 350)
	}
}

func (v variation) drawDamage(dc *gg.Context) {
	dc.SetRGB(0.75, 0.75, 0.75)
	setSize(40, dc)
	dc.DrawString("Damage", 490, 440)
	dc.SetRGB(0.85, 0.85, 0.85)
	dc.DrawString(v.damage, 660, 440)

	if v.damageup != "0" && v.damageup != v.damage {
		setSize(40, dc)
		dc.SetRGB(0.8, 0.3, 0.3)
		dc.DrawString(v.damageup, 885, 440)
		flagImage := openImage("up32.png")
		dc.DrawImage(flagImage, 850, 410)
	}
}

func (v variation) drawWins(dc *gg.Context) {
	dc.SetRGB(0.75, 0.75, 0.75)
	setSize(40, dc)
	dc.DrawString("Wins", 490, 500)
	dc.SetRGB(0.85, 0.85, 0.85)
	dc.DrawString(v.wins, 660, 500)

	if v.winup != "0" && v.winup != v.wins {
		setSize(40, dc)
		dc.SetRGB(0.8, 0.3, 0.3)
		dc.DrawString(v.winup, 885, 500)
		flagImage := openImage("up32.png")
		dc.DrawImage(flagImage, 850, 470)
	}
}

func (v variation) saveLog(u database.UserData) {
	t := getTime()
	//For the first time, do not saveLog the log
	if u.Rp != 0 {
		database.LogInsert(u.AccountId, v.rp, v.rpup, v.killup, v.damageup, v.winup, t)
	}
}

func (v variation) updateUserData(u database.UserData) {
	t := getTime().Format("02")
	database.UpdateUserData(u.AccountId, u.Legend, t, v.rankname, v.rp, v.kills, v.damage, v.wins)
}

func (v variation) updateLastMadeTime(u database.UserData) {
	t := getTime().Format("2006/01/02")
	database.UpdateLastMade(u.AccountId, v.rp, v.kills, v.damage, v.wins, t)

}
