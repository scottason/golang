// scraper.go
package main

import (
	// improt colly
	"encoding/csv"
	"encoding/json"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// iterating over the list of industry card
type Industry struct {
	Url, Image, Name string
}

func main() {
	// initialize struct slice
	var industries []Industry

	c := colly.NewCollector()

	//set a valid user_agent header

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	//connect to target site
	c.Visit("https://brightdata.com/")

	// HTML elements

	c.OnHTML(".elementor-element-6b05593c .section_cases__item", func(e *colly.HTMLElement) {

		url := e.Attr("href")

		image := e.ChildAttr(".elementor-image-box-img img", "data-lazy-src")

		name := e.ChildText(".elementor-image-box-content.elementor-image-box-title")

		// filter out unwanted data

		if url != "" || image != "" || name != "" {

			// initialize a new Industry instance

			industry := Industry{

				Url: url,

				Image: image,

				Name: name,
			}

			// add the industry instance to the list

			// of scraped industries

			industries = append(industries, industry)
		}
	})

	// --- export to CSV ---

	// open the output CSV file
	csvFile, csvErr := os.Create("industries.csv")
	// if the file creation fails
	if csvErr != nil {
		log.Fatalln("Failed to create the output CSV file", csvErr)
	}
	// release the resource allocated to handle
	// the file before ending the execution
	defer csvFile.Close()

	// create a CSV file writer
	writer := csv.NewWriter(csvFile)
	// release the resources associated with the
	// file writer before ending the execution
	defer writer.Flush()

	// add the header row to the CSV
	headers := []string{
		"url",
		"image",
		"name",
	}
	writer.Write(headers)

	// store each Industry product in the
	// output CSV file
	for _, industry := range industries {
		// convert the Industry instance to
		// a slice of strings
		record := []string{
			industry.Url,
			industry.Image,
			industry.Name,
		}
		// add a new CSV record
		writer.Write(record)
	}
	// open the output json FILE
	jsonFile, jsonErr := os.Create("industries.json")

	if jsonErr != nil {
		log.Fatalln("Failed to create the output json file", jsonErr)

	}
	defer jsonFile.Close()

	//convert industries to an indented json stringc
	jsonString, _ := json.MarshalIndent(industries, " ", " ")

	// write the JSON string to file

	jsonFile.Write(jsonString)

	// scraping logic

}
