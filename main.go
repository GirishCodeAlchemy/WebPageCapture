package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func generateScreenshot(url, outputPath string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, fullScreenshot(url, 500, &buf)); err != nil {
		return err
	}

	if err := os.WriteFile(outputPath, buf, 0644); err != nil {
		return err
	}

	fmt.Println("Screenshot captured successfully:", outputPath)
	return nil
}

func fullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second), // Adjust this delay as needed for page loading
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.FullScreenshot(res, quality),
	}
}

func generatePDF(url, outputPath string) error {
	// Launch a headless Chrome browser using chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set up options for generating PDF
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("disable-software-rasterizer", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("hide-scrollbars", true),
		chromedp.Flag("window-size", "1920,1080"),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)
	defer allocCancel()

	// Launch a new browser context
	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	// Run tasks to navigate to the page and wait for body visibility
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
	); err != nil {
		return err
	}

	// Create a PDF using the DevTools protocol
	var pdfBuf []byte
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		pdfBuf, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
		return err
	})); err != nil {
		return err
	}

	// Save the generated PDF to a file
	err := os.WriteFile(outputPath, pdfBuf, 0644)
	if err != nil {
		return err
	}

	fmt.Println("PDF generated successfully.")
	return nil
}

func main() {
	// Check if at least two arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run yourprogram.go <URL> <pdf|image> [filename]")
		return
	}

	url := os.Args[1]
	outputType := os.Args[2]

	outputPath := ""

	if len(os.Args) > 3 {
		outputPath = os.Args[3]
	}

	switch outputType {
	case "image":
		if outputPath == "" {
			outputPath = "output.jpg"
		}
		err := generateScreenshot(url, outputPath)
		if err != nil {
			log.Fatal("Error capturing the screenshot:", err)
		}
	default:
		if outputPath == "" {
			outputPath = "output.pdf"
		}
		err := generatePDF(url, outputPath)
		if err != nil {
			log.Fatal("Error generating PDF:", err)
		}
	}
}
