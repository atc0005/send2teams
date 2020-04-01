package main

import (
	"fmt"
	"log"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify"
	"github.com/atc0005/send2teams/config"
)

// testCase2 generates a message card with an empty section to confirm whether
// the go-teams-notify package properly drops the empty section JSON array.
func testCase2(cfg *config.Config) goteamsnotify.MessageCard {

	// setup message card
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = cfg.MessageTitle
	msgCard.Text = "Test Case 2 (top-level text content)"
	msgCard.ThemeColor = cfg.ThemeColor

	// TODO: This results in an empty JSON sections array. This is what we're
	// effectively asking for here. Is this an error condition that the library
	// should be handling?
	//
	//testEmptySection := goteamsnotify.NewMessageCardSection()
	testEmptySection := goteamsnotify.MessageCardSection{}
	fmt.Println("Calling AddSection from Test Case 2 with empty section")
	msgCard.AddSection(&testEmptySection)
	msgCard.AddSection(nil)
	msgCard.AddSection(nil)
	msgCard.AddSection(nil)
	msgCard.AddSection(nil)

	structDetails, err := goteamsnotify.FormatAsCodeBlock(fmt.Sprintf("This message card's fields: %+v", msgCard))
	if err == nil {
		msgCard.Text = structDetails
	}

	return msgCard
}

// testCase1 generates a message card with a number of useful sections and
// content. This includes activity fields, gallery images, code snippet
// section, code block section and branding trailer section. This test case is
// close to, but not quite what the bounce application might generate from its
// input.
func testCase1(cfg *config.Config) goteamsnotify.MessageCard {

	// setup message card
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = cfg.MessageTitle
	msgCard.Text = "Test Case 1 (top-level text content)"
	msgCard.ThemeColor = cfg.ThemeColor

	mainMsgSection := goteamsnotify.NewMessageCardSection()

	// This represents what the user would provide via CLI flag:
	mainMsgSection.Text = cfg.MessageText + " (section text)"

	//log.Printf("msgCard before adding mainMsgSection: %+v", msgCard)
	msgCard.AddSection(mainMsgSection)
	//log.Printf("msgCard after adding mainMsgSection: %+v", msgCard)

	/*

		Code Snippet Sample Section

	*/

	codeSnippetSampleSection := goteamsnotify.NewMessageCardSection()
	codeSnippetSampleSection.StartGroup = true

	codeSnippetSampleSection.Title = "Code Snippet Sample Section"

	// This represents something programatically generated:
	unformattedTextSample := "GET request received on /api/v1/echo/json endpoint"
	formattedTextSample, err := goteamsnotify.FormatAsCodeSnippet(unformattedTextSample)
	if err != nil {

		log.Printf("error formatting text as code snippet: %#v", err)
		log.Printf("Current state of section: %+v", codeSnippetSampleSection)

		log.Println("Using unformattedTextSample")
		codeSnippetSampleSection.Text = unformattedTextSample
	} else {
		log.Println("Using formattedTextSample")
		codeSnippetSampleSection.Text = formattedTextSample
		msgCard.AddSection(codeSnippetSampleSection)
	}

	/*

		Code Block Sample Section

	*/

	codeBlockSampleSection := goteamsnotify.NewMessageCardSection()
	codeBlockSampleSection.Title = "Code Block Sample Section"

	// This represents something programatically generated:
	sampleJSONInput := `{"result":{"sourcetype":"mongod","count":"8"},"sid":"scheduler_admin_search_W2_at_14232356_132","results_link":"http://web.example.local:8000/app/search/@go?sid=scheduler_admin_search_W2_at_14232356_132","search_name":null,"owner":"admin","app":"search"}`
	formattedTextSample, err = goteamsnotify.FormatAsCodeBlock(sampleJSONInput)
	if err != nil {

		log.Printf("error formatting text as code snippet: %#v", err)
		log.Printf("Current state of section: %+v", codeBlockSampleSection)

		log.Println("Using unformattedTextSample")
		codeBlockSampleSection.Text = unformattedTextSample
	} else {
		log.Println("Using formattedTextSample")
		codeBlockSampleSection.Text = formattedTextSample
	}

	msgCard.AddSection(codeBlockSampleSection)

	/*
		Activity section
	*/

	activitySection := goteamsnotify.NewMessageCardSection()
	activitySection.Title = "Title: Testing activity properties"
	activitySection.ActivityText = "ActivityText: Adam did something today."
	activitySection.ActivitySubtitle = "ActivitySubtitle: Hopefully it was useful"
	activitySection.ActivityImage = "https://avatars2.githubusercontent.com/u/36716992"
	msgCard.AddSection(activitySection)

	/*
		Hero Image section
	*/

	heroImageSection := goteamsnotify.NewMessageCardSection()
	heroImageSection.Title = "Testing hero image"
	heroImageSection.Text = "Unfortunately this property is not supported by Microsoft Teams."
	// if err := heroImageSection.AddHeroImage(
	// 	"https://live.staticflickr.com/3551/3388550814_0f4ac0d1a0.jpg",
	// 	"https://search.creativecommons.org/photos/78cdb549-3270-48be-9df3-84d53ab3d245",
	// ); err != nil {
	// 	fmt.Printf("failed to add hero image: %s", err)
	// 	os.Exit(1)
	// }
	heroImage := goteamsnotify.NewMessageCardSectionImage()
	heroImage.Image = "https://live.staticflickr.com/3551/3388550814_0f4ac0d1a0.jpg"
	heroImage.Title = "https://search.creativecommons.org/photos/78cdb549-3270-48be-9df3-84d53ab3d245"
	if err := heroImageSection.AddHeroImage(heroImage); err != nil {
		fmt.Printf("failed to add hero image: %s", err)
		os.Exit(1)
	}
	if err := heroImageSection.AddHeroImageStr(
		"https://live.staticflickr.com/3551/3388550814_0f4ac0d1a0.jpg",
		"https://search.creativecommons.org/photos/78cdb549-3270-48be-9df3-84d53ab3d245",
	); err != nil {
		fmt.Printf("failed to add hero image: %s", err)
		os.Exit(1)
	}
	msgCard.AddSection(heroImageSection)

	/*
		Image Gallery section
	*/

	galleryImageSection := goteamsnotify.NewMessageCardSection()
	bannerImg := goteamsnotify.NewMessageCardSectionImage()
	bannerImg.Image = "https://live.staticflickr.com/3551/3388550814_0f4ac0d1a0.jpg"
	bannerImg.Title = "https://search.creativecommons.org/photos/78cdb549-3270-48be-9df3-84d53ab3d245"
	if err := galleryImageSection.AddImage(bannerImg); err != nil {
		fmt.Printf("failed to add image: %s", err)
		os.Exit(1)
	}

	if err := galleryImageSection.AddImage(goteamsnotify.MessageCardSectionImage{
		Image: "https://farm3.staticflickr.com/2359/2149071817_0c0f7fd539.jpg",
		Title: "https://search.creativecommons.org/photos/4393a3f3-ea51-438c-89da-1e3fa468d80b",
	}); err != nil {
		fmt.Printf("failed to add hero image: %s", err)
		os.Exit(1)
	}
	galleryImageSection.Title = "Testing gallery images"
	msgCard.AddSection(galleryImageSection)

	/*
		Bad Data: Image Gallery section
	*/

	// badGalleryImageSection := goteamsnotify.NewMessageCardSection()
	// badBannerImg := goteamsnotify.NewMessageCardSectionImage()
	// badBannerImg.Image = ""
	// badBannerImg.Title = ""

	// // This doesn't check the return code
	// //badGalleryImageSection.AddImage(badBannerImg)

	// // Let's do that
	// if err := badGalleryImageSection.AddImage(badBannerImg); err != nil {
	// 	fmt.Printf("failed to add section image: %s", err)
	// 	os.Exit(1)
	// }
	// badGalleryImageSection.Title = "Testing empty fields for MessageCardSectionImage type"
	// msgCard.AddSection(badGalleryImageSection)

	/*
		Branding trailer section
	*/

	trailerSection := goteamsnotify.NewMessageCardSection()
	trailerSection.Text = config.MessageTrailer()
	trailerSection.StartGroup = true

	//log.Printf("msgCard before adding trailerSection: %+v", msgCard)
	msgCard.AddSection(trailerSection)
	//log.Printf("msgCard after adding trailerSection: %+v", msgCard)

	return msgCard

}
