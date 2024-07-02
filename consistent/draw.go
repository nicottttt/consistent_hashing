package consistent

import (
	"math"

	"github.com/fogleman/gg"
)

func DrawRing(c *Consistent, key string) {
	ring := c.GetRing()
	const S = 600 // Size of the image
	dc := gg.NewContext(S, S)

	// Draw the outer circle representing the hash ring
	dc.DrawCircle(S/2, S/2, float64(S/3))
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)
	dc.Stroke()

	// Draw hash points and server names
	for hash, server := range ring {
		angle := 2 * math.Pi * float64(hash) / 1024
		x := S/2 + float64(S/3)*math.Cos(angle)
		y := S/2 + float64(S/3)*math.Sin(angle)

		switch server {
		case "Server1":
			// Draw a small circle for the hash point
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(135, 206, 235) // Sky Blue: (135, 206, 235)
			dc.Fill()

			// // Draw the server name
			// dc.SetRGB(1, 0, 0) // Blue color for the text
			// dc.DrawStringAnchored(fmt.Sprintf("%s: %d", sh.server, sh.hash), x, y-10, 0.5, 0.5)
		case "Server2":
			// Draw a small circle for the hash point
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(255, 127, 80) // Coral Pink: (255, 127, 80)
			dc.Fill()
		case "Server3":
			// Draw a small circle for the hash point
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(50, 205, 50) // Lime Green: (50, 205, 50)
			dc.Fill()
		case "Server4":
			// Draw a small circle for the hash point
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(220, 20, 60) // Crimson Red: (220, 20, 60)
			dc.Fill()
		default:
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(255, 255, 255) // Crimson Red: (220, 20, 60)
			dc.Fill()
		}

	}

	//Key drawing
	angle := 2 * math.Pi * float64(c.GetHasher().Xxhash1024([]byte(key))) / 1024
	x := S/2 + float64(S/3)*math.Cos(angle)
	y := S/2 + float64(S/3)*math.Sin(angle)

	dc.DrawCircle(x, y, 5)
	dc.SetRGB(1, 0, 0)
	dc.Fill()

	// Draw the name
	dc.SetRGB(1, 0, 0)
	dc.DrawStringAnchored(key, x, y-10, 0.5, 0.5)

	// Save the image
	dc.SavePNG("hash_ring_diagram.png")
}
