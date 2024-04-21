package main

import (
	"fmt"
	"math"
	"time"
	"strings"

	"github.com/rivo/tview"
)

const (
	initialDistanceKmBetweenEarthAndMoon = 384467.0 
	annualDriftCm                        = 3.8      
)

func main() {
	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("Initializing...").
		SetScrollable(true)

	go func() {
		startTime := time.Now()
		for {
			currentTime := time.Now()
			elapsedYears := currentTime.Sub(startTime).Hours() / 24 / 365
			currentDistance := initialDistanceKmBetweenEarthAndMoon + (annualDriftCm * elapsedYears / 100000)

			
			solarSystemAnimation := animateSolarSystem(elapsedYears)

			distanceText := fmt.Sprintf("[yellow]Current distance from Earth to Moon: [red]%.2f km", currentDistance)
			fullText := solarSystemAnimation + "\n" + distanceText
			app.QueueUpdateDraw(func() {
				textView.SetText(fullText)
			})
			time.Sleep(250 * time.Millisecond) 
		}
	}()

	if err := app.SetRoot(textView, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}


func animateSolarSystem(elapsedYears float64) string {
	
	earthOrbitPosition := int(math.Mod(elapsedYears*5, 40))
	moonOrbitPosition := int(math.Mod(elapsedYears*20, 10))

	// Added simple calculations for planets' distances from the sun, approximations in AU (Astronomical Unit)
	mercuryDistance := 0.39 // AU from the sun
	venusDistance := 0.72   // AU from the sun
	earthDistance := 1.00   // AU from the sun
	marsDistance := 1.52    // AU from the sun
	jupiterDistance := 5.20 // AU from the sun
	saturnDistance := 9.58  // AU from the sun
	uranusDistance := 19.22 // AU from the sun
	neptuneDistance := 30.05 // AU from the sun

	solarSystem := fmt.Sprintf(`[yellow]Simplified Solar System with Orbits and Planets and Distances from the Sun in AU:[white]
Mercury: %s `+generateOrbitWithPlanet(15, int(math.Mod(elapsedYears*30, 15)), "[cyan]☿[white]")+` (Distance: %.2f AU)
Venus:   %s `+generateOrbitWithPlanet(25, int(math.Mod(elapsedYears*24, 25)), "[magenta]♀[white]")+` (Distance: %.2f AU)
Earth:   %s `+generateOrbitWithPlanet(40, earthOrbitPosition, "[green]●[white]")+` (Distance: %.2f AU)
Moon:    `+generateOrbitWithPlanet(10, moonOrbitPosition, "[red]○[white]")+`
Mars:    %s `+generateOrbitWithPlanet(45, int(math.Mod(elapsedYears*13, 45)), "[red]♂[white]")+` (Distance: %.2f AU)
Jupiter: %s `+generateOrbitWithPlanet(55, int(math.Mod(elapsedYears*2, 55)), "[yellow]♃[white]")+` (Distance: %.2f AU)
Saturn:  %s `+generateOrbitWithPlanet(65, int(math.Mod(elapsedYears*0.5, 65)), "[orange]♄[white]")+` (Distance: %.2f AU)
Uranus:  %s `+generateOrbitWithPlanet(75, int(math.Mod(elapsedYears*0.2, 75)), "[blue]♅[white]")+` (Distance: %.2f AU)
Neptune: %s `+generateOrbitWithPlanet(85, int(math.Mod(elapsedYears*0.1, 85)), "[darkblue]♆[white]")+` (Distance: %.2f AU)`,
		mercuryDistance, venusDistance, earthDistance, marsDistance, jupiterDistance, saturnDistance, uranusDistance, neptuneDistance)
	return solarSystem
}


func generateOrbit(length int) string {
	return "-" + strings.Repeat("-", length) + "-"
}


func generateOrbitWithPlanet(length, position int, marker string) string {
    orbits := strings.Repeat("-", length)
    pre := ""
    post := ""
    if position < len(orbits) {
        pre = strings.Repeat("-", position)
        postLength := length - len(marker) - len(pre) 
        if postLength > 0 {
            post = strings.Repeat("-", postLength)
        }
        orbits = pre + marker + post
    } else {
        orbits = pre + marker + strings.Repeat("-", length - len(marker))
    }
    return "-" + orbits + "-"
}