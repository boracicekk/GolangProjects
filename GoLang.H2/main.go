package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/mbndr/figlet4go"
)

func logo(yamlFile string, yamlFile2 string) {
	// YAVUZLAR
	yavuzlarASCII := figlet4go.NewAsciiRender()
	yavuzlarOptions := figlet4go.NewRenderOptions()
	yavuzlarOptions.FontName = "larry3d"
	yavuzlarOptions.FontColor = []figlet4go.Color{figlet4go.ColorBlue}
	yavuzlarStr, _ := yavuzlarASCII.RenderOpts(yamlFile, yavuzlarOptions)

	fmt.Println(yavuzlarStr)

	// Web Scraper
	webScraperASCII := figlet4go.NewAsciiRender()
	webScraperOptions := figlet4go.NewRenderOptions()
	webScraperOptions.FontName = "block"
	webScraperOptions.FontColor = []figlet4go.Color{figlet4go.ColorRed}
	webScraperStr, _ := webScraperASCII.RenderOpts(yamlFile2, webScraperOptions)

	fmt.Println(webScraperStr)
}

// help function introduce the command's and meanings.
func help(learn string) {
	if learn == "--help" {
		fmt.Println("First Website Filter Commands:\n --date : the filter command hides the new's date time parameter.\n --desc : The filter command filters the description information of the news parameter.\n Also you can use both of filter command like ' --desc --date '.")
		fmt.Println("Second Website Filter Commands:\n --bookname : hides bookname parameter.\n--price : hides bookname parameter.")
		fmt.Println("Third Website Filter Commands:\n --productname: hides product's name parameter.\n--price : hides product's price parameter.")
	} else if learn == "--continue" {
		return
	}
}
func main() {
	logo("YAVUZLAR", "Web Scraper")
	var urlNumb int
	var filter1 string
	var learn string
	fmt.Println("Select the website you want to scrape:\n1)The Hacker News\n2)BkmKitap[Defter]\n3)N11[Bilgisayar]")
	fmt.Scan(&urlNumb)
	fmt.Println("You can type --help to learn Filtering Commands or type --continue:")
	fmt.Scan(&learn)
	help(learn)

	if urlNumb == 1 {
		url1 := "https://thehackernews.com"
		resp1, err1 := http.Get(url1)
		if err1 != nil {
			fmt.Println("HTTP GET ERROR:", err1)
			return
		}
		if resp1.StatusCode != 200 {
			fmt.Println("HTTP request failed. \n HTTP STATUS CODE:", resp1.StatusCode)
		} else {
			doc, _ := goquery.NewDocumentFromReader(resp1.Body)
			fmt.Println("Enter your filter/s:")
			fmt.Scan(&filter1)

			if filter1 == "--desc --date" || filter1 == "--date --desc" {
				doc.Find(".clear.home-right").Each(func(i int, selection *goquery.Selection) {
					title := selection.Find("h2").Text()
					fmt.Println("\n" + title)
				})
			} else if filter1 == "--desc" {
				doc.Find(".clear.home-right").Each(func(i int, selection *goquery.Selection) {
					title := selection.Find("h2").Text()
					date := selection.Find("span.h-datetime").Text()
					fmt.Println("\n" + title + " " + date)
				})
			} else if filter1 == "--date" {
				doc.Find(".clear.home-right").Each(func(i int, selection *goquery.Selection) {
					title := selection.Find("h2").Text()
					desc := selection.Find(".home-desc").Text()
					fmt.Println("\n" + title + " " + desc)
				})
			} else {
				fmt.Println("Enter a valid filter!")
			}

		}
	} else if urlNumb == 2 {
		url2 := "https://www.bkmkitap.com/butik-defterler-4374"
		resp2, err2 := http.Get(url2)
		if err2 != nil {
			fmt.Println("HTTP GET ERROR:", err2)
			return
		}
		if resp2.StatusCode != 200 {
			fmt.Println("HTTP request failed. \n HTTP STATUS CODE:", resp2.StatusCode)
		} else {
			doc, _ := goquery.NewDocumentFromReader(resp2.Body)
			fmt.Println("Enter your filter/s:")
			fmt.Scan(&filter1)
			if filter1 == "null" {
				doc.Find(".col.col-12.p-left").Each(func(i int, selection *goquery.Selection) {
					bookname := selection.Find(".fl.col-12.text-description.detailLink").Text()
					price := selection.Find(".col.col-12.currentPrice").Text()
					fmt.Println("\n" + bookname + "" + price)
				})
			} else if filter1 == "--bookname" {
				doc.Find(".col.col-12.p-left").Each(func(i int, selection *goquery.Selection) {
					price := selection.Find(".col.col-12.currentPrice").Text()
					fmt.Println("\n" + price)
				})
			} else if filter1 == "--price" {
				doc.Find(".col.col-12.p-left").Each(func(i int, selection *goquery.Selection) {
					bookname := selection.Find(".fl.col-12.text-description.detailLink").Text()
					fmt.Println("\n" + bookname)
				})
			} else {
				fmt.Println("Enter a valid filter!")
			}

		}
	} else if urlNumb == 3 {
		url3 := "https://www.n11.com/bilgisayar"
		resp3, err3 := http.Get(url3)
		if err3 != nil {
			fmt.Println("HTTP GET ERROR:", err3)
			return
		}
		if resp3.StatusCode != 200 {
			fmt.Println("HTTP request failed. \n HTTP STATUS CODE:", resp3.StatusCode)
		} else {
			doc, _ := goquery.NewDocumentFromReader(resp3.Body)
			fmt.Println("Enter your filter/s:")
			fmt.Scan(&filter1)
			if filter1 == "null" {
				doc.Find(".list-ul").Each(func(i int, selection *goquery.Selection) {
					productname := selection.Find(".productName").Text()
					price := selection.Find(".newPrice.cPoint.priceEventClick").Text()
					fmt.Println("\n" + productname + "" + price)
				})
			} else if filter1 == "--productname" {
				doc.Find(".list-ul").Each(func(i int, selection *goquery.Selection) {
					price := selection.Find(".newPrice.cPoint.priceEventClick").Text()
					fmt.Println("\n" + price)
				})
			} else if filter1 == "--price" {
				doc.Find(".list-ul").Each(func(i int, selection *goquery.Selection) {
					productname := selection.Find(".productName").Text()
					fmt.Println("\n" + productname)
				})
			} else {
				fmt.Println("Enter a valid filter!")
			}

		}
	} else {
		fmt.Println("Enter a valid value!")
	}

}
