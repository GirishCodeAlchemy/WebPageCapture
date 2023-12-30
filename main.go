package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type options struct {
	URL        string
	OutputType string
	OutputPath string
}

var opts options

func init() {
	flag.StringVar(&opts.OutputType, "type", "", "Specify 'pdf' to generate a PDF or 'image' to capture a screenshot")
	flag.StringVar(&opts.OutputPath, "filename", "", "(Optional) Output filename. Defaults to 'output.pdf' for PDF and 'output.jpg' for image.")
	flag.StringVar(&opts.URL, "url", "", "URL of the webpage to capture")
	flag.Usage = func() {
		printHelp()
	}
}

func launchBrowser() (context.Context, context.CancelFunc) {
	ctx, cancel := chromedp.NewContext(context.Background())
	return ctx, cancel
}

func generateScreenshot(ctx context.Context, url, outputPath string) error {
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

func generatePDF(ctx context.Context, url, outputPath string) error {
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

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
	); err != nil {
		return err
	}

	var pdfBuf []byte
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		pdfBuf, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
		return err
	})); err != nil {
		return err
	}

	err := os.WriteFile(outputPath, pdfBuf, 0644)
	if err != nil {
		return err
	}

	fmt.Println("PDF generated successfully.")
	return nil
}

func printHelp() {
	fmt.Println("Usage: web-capture-tool [flags] [URL] <pdf|image> [filename]")
	fmt.Println("\nCommand-line Arguments:")
	fmt.Println("  <URL>\t\t\tURL of the webpage to capture")
	fmt.Println("  <pdf|image>\t\tSpecify 'pdf' to generate a PDF or 'image' to capture a screenshot")
	fmt.Println("  [filename]\t\t(Optional) Output filename. Defaults to 'output.pdf' for PDF and 'output.jpg' for image.")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
	fmt.Println("\nExample:")
	fmt.Println("  web-capture-tool https://example.com pdf output.pdf")
	fmt.Println("  web-capture-tool --url=https://example.com --type=image --filename=output.jpg")
}

func validateFilename(filename string) error {
	// check if the filename has a valid extension.
	if filename == "" {
		return nil
	}

	ext := filepath.Ext(filename)
	if ext == "" {
		return fmt.Errorf("Filename must have an extension\n")
	}

	ext = ext[1:]
	validExtensions := map[string]bool{"pdf": true, "jpg": true, "png": true, "jpeg": true, "bmp": true}

	if !validExtensions[ext] {
		return fmt.Errorf("Wrong extension, Supported extensions: pdf, jpg, png\n")
	}

	return nil
}

func parseArgs() (options, error) {

	flag.Parse()

	if flag.NArg() > 0 && (flag.Arg(0) == "-h" || flag.Arg(0) == "--help") {
		printHelp()
		os.Exit(0)
	}

	args := flag.Args()
	if opts.URL == "" && len(args) > 0 {
		opts.URL = args[0]
	}
	if opts.OutputType == "" && len(args) > 1 {
		ext := filepath.Ext(args[1])
		if ext != "" {
			opts.OutputPath = args[1]
			if ext == ".pdf" {
				opts.OutputType = "pdf"
			} else {
				opts.OutputType = "image"
			}
		} else {
			opts.OutputType = args[1]
		}
	}
	if opts.OutputPath == "" && len(args) > 2 {
		opts.OutputPath = args[2]
	}

	if opts.URL == "" {
		fmt.Printf("\nMissing required flags or positional arguments. Use '-h' or '--help' for usage information.")
		printHelp()
		os.Exit(0)
	}

	if err := validateFilename(opts.OutputPath); err != nil {
		fmt.Printf("\nInvalid output filename: %s,  %v\n", opts.OutputPath, err)
		printHelp()
		os.Exit(0)
	}

	return opts, nil
}

func main() {

	opts, err := parseArgs()
	if err != nil {
		log.Fatal("Error:", err)
	}

	ctx, cancel := launchBrowser()
	defer cancel()

	switch opts.OutputType {
	case "image":
		if opts.OutputPath == "" {
			opts.OutputPath = "output.jpg"
		}
		err := generateScreenshot(ctx, opts.URL, opts.OutputPath)
		if err != nil {
			log.Fatal("Error capturing the screenshot:", err)
		}
	case "pdf", "":
		if opts.OutputPath == "" {
			opts.OutputPath = "output.pdf"
		}
		err := generatePDF(ctx, opts.URL, opts.OutputPath)
		if err != nil {
			log.Fatal("Error generating PDF:", err)
		}

	default:
		fmt.Println("Error: Invalid output type. Use '-h' or '--help' for usage information.")
		printHelp()

	}
}
