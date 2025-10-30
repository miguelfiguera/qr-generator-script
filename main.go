package main

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	// Create application
	myApp := app.New()
	myWindow := myApp.NewWindow("QR Code Generator")
	myWindow.Resize(fyne.NewSize(600, 500))

	// Create UI elements
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter URL (e.g., https://www.example.com)")

	// QR code image display
	var qrImage *canvas.Image
	qrContainer := container.NewCenter()

	statusLabel := widget.NewLabel("Enter a URL and click 'Generate QR Code'")
	statusLabel.Alignment = fyne.TextAlignCenter

	// Save button (initially disabled)
	var currentQRPath string
	saveBtn := widget.NewButton("Save QR Code", func() {
		if currentQRPath == "" {
			dialog.ShowError(fmt.Errorf("No QR code to save"), myWindow)
			return
		}

		// Create file save dialog
		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()

			// Read the QR code file
			data, err := os.ReadFile(currentQRPath)
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}

			// Write to selected file
			_, err = writer.Write(data)
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}

			statusLabel.SetText(fmt.Sprintf("QR Code saved successfully to %s", writer.URI().Path()))
		}, myWindow)

		saveDialog.SetFileName("qrcode.png")
		saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png"}))
		saveDialog.Show()
	})
	saveBtn.Disable()

	// Generate button
	generateBtn := widget.NewButton("Generate QR Code", func() {
		url := urlEntry.Text
		if url == "" {
			statusLabel.SetText("Please enter a URL")
			dialog.ShowError(fmt.Errorf("URL cannot be empty"), myWindow)
			return
		}

		// Generate QR code
		statusLabel.SetText("Generating QR code...")

		// Create temp file for QR code
		tempDir := os.TempDir()
		currentQRPath = filepath.Join(tempDir, "qrcode_temp.png")

		err := qrcode.WriteFile(url, qrcode.Medium, 256, currentQRPath)
		if err != nil {
			statusLabel.SetText("Error generating QR code")
			dialog.ShowError(err, myWindow)
			return
		}

		// Load and display the QR code
		qrImage = canvas.NewImageFromFile(currentQRPath)
		qrImage.FillMode = canvas.ImageFillOriginal
		qrImage.SetMinSize(fyne.NewSize(256, 256))

		qrContainer.Objects = []fyne.CanvasObject{qrImage}
		qrContainer.Refresh()

		statusLabel.SetText("QR Code generated successfully!")
		saveBtn.Enable()
	})

	// Clear button
	clearBtn := widget.NewButton("Clear", func() {
		urlEntry.SetText("")
		qrContainer.Objects = []fyne.CanvasObject{}
		qrContainer.Refresh()
		statusLabel.SetText("Enter a URL and click 'Generate QR Code'")
		currentQRPath = ""
		saveBtn.Disable()
	})

	// Create layout
	inputContainer := container.NewVBox(
		widget.NewLabel("QR Code Generator"),
		widget.NewSeparator(),
		urlEntry,
		container.NewGridWithColumns(3, generateBtn, saveBtn, clearBtn),
		widget.NewSeparator(),
		statusLabel,
	)

	// Instructions
	instructions := widget.NewLabel("Enter a URL and generate a QR code.\nThe QR code will be displayed below.")
	instructions.Wrapping = fyne.TextWrapWord
	instructions.TextStyle = fyne.TextStyle{Italic: true}
	instructions.Alignment = fyne.TextAlignCenter

	// Add some styling
	instructionsCard := widget.NewCard("", "", instructions)

	// Main layout
	content := container.NewBorder(
		inputContainer,
		instructionsCard,
		nil,
		nil,
		qrContainer,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
