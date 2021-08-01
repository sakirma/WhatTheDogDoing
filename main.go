package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	bString := binding.NewString()

	err := bString.Set("Binded Shiiiit")
	if err != nil {
		return
	}

	w.SetContent(container.NewVBox(
		widget.NewLabelWithData(bString),
		widget.NewButton("Get Random Recipe", func() {
			err := bString.Set("Wow a recipe")
			if err != nil {
				return 
			}
		}),
	))

	w.ShowAndRun()
}
