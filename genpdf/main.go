package main

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
)

func main() {
	// create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctx,
		chromedp.WithRunnerOptions(
			runner.Flag("headless", ""),
			runner.Flag("no-sandbox", ""),
			runner.Flag("no-gpu", ""),
		),
	)
	if err != nil {
		panic(err)
	}

	// run task list
	var pdfBuf []byte
	err = c.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("http://localhost:3000"),
		chromedp.Sleep(2 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context, h cdp.Executor) error {
			printToPDF := &page.PrintToPDFParams{
				PrintBackground:   true,
				PreferCSSPageSize: true,
			}

			// take page screenshot
			buf, err := printToPDF.Do(ctx, h)
			if err != nil {
				return err
			}

			pdfBuf = buf
			return nil
		}),
	})
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("static/resume.pdf", pdfBuf, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
