package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func main() {
	fmt.Println(time.Now().Format("15:04:05") + " : " + "start job..")
	s := gocron.NewScheduler(time.Local)
	s.Every(1).Day().At("11:59:59").Do(TaskShopee)
	s.StartBlocking()
}

func TaskShopee() {

	fmt.Println(time.Now().Format("15:04:05") + " : " + "start..")

	url, _ := launcher.New().Set("no-sandbox").Launch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	browser.MustSetCookies(&proto.NetworkCookie{
		Name:   "SPC_EC",
		Value:  "RzB3alVmd0dMQXhyejVJVPN9RHx/aifuJFDVEqTv3TKU4HhrwagNqRRkKUi7ZnW4aCpd3X7sGJ1FhJ+81PTfhtDVtnH8GQbl1EFSqclGpByzpZnmR6/1NdwfbICMlwmVlP/lOhOiJoWhN/f0NyhAmaE7M82aj65/l759JoCLNlo=",
		Domain: ".shopee.co.id",
	})

	page := browser.MustPage("https://shopee.co.id/BOLD%C3%AB-Set-Wajan-Super-Pan-Beige-i.72109688.1207801918")
	page.MustWaitLoad()

	fmt.Println(time.Now().Format("15:04:05") + " : " + "page loaded..")

	// variant := page.MustElementX("//button[contains(text(),'earphone putih')]")
	// variant.MustClick()

	wait := page.MustWaitRequestIdle()
	buy := page.MustElement("#main > div > div._193wCc > div.page-product.page-product--mall > div > div.product-briefing.flex.card.zINA0e > div.flex.flex-auto._3-GQHh > div > div:nth-child(5) > div > div > button.btn.btn-solid-primary.btn--l._3Kiuzg")
	buy.MustClick()
	wait()

	fmt.Println(time.Now().Format("15:04:05") + " : " + "barang di beli..")

	wait = page.MustWaitRequestIdle()
	checkout := page.MustElementX("//*[@id=\"main\"]/div/div[2]/div[2]/div/div[3]/div[2]/div[7]/button[4]/span")
	checkout.MustClick()
	wait()
	page.MustWaitIdle()
	page.MustWaitLoad()

	fmt.Println(time.Now().Format("15:04:05") + " : " + "barang di checkout..")

	wait = page.MustWaitRequestIdle()
	pesanan := page.MustElementX("//*[@id=\"main\"]/div/div[3]/div[2]/div[4]/div[2]/div[9]/button")
	pesanan.MustClick()
	wait()
	page.MustWaitLoad()
	fmt.Println(time.Now().Format("15:04:05") + " : " + "barang di buatkan pesanan..")

	wait = page.MustWaitRequestIdle()
	bayar := page.MustElementX("//*[@id=\"pay-button\"]")
	bayar.MustClick()
	wait()
	page.MustWaitLoad()

	fmt.Println(time.Now().Format("15:04:05") + " : " + "barang di bayar..")

	page.MustElementX("//*[@id=\"pin-popup\"]/div[1]/div[3]/div[1]")
	page.Keyboard.MustPress('')
	page.Keyboard.MustPress('')
	page.Keyboard.MustPress('')
	page.Keyboard.MustPress('')
	page.Keyboard.MustPress('')
	page.Keyboard.MustPress('')

	wait = page.MustWaitRequestIdle()
	confirmPin := page.MustElementX("//*[@id=\"pin-popup\"]/div[1]/div[4]/div[2]")
	confirmPin.MustClick()
	wait()

	fmt.Println("sippp")
}
