package consistent

import (
	"math"

	"github.com/fogleman/gg"
	"github.com/google/uuid"
)

func DrawRing(c *Consistent, keys ...string) {
	ring := c.GetRing()
	const S = 600 // Size of the image
	dc := gg.NewContext(S, S)

	// Draw the outer circle representing the hash ring
	dc.DrawCircle(S/2, S/2, float64(S/3))
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)
	dc.Stroke()

	server1, _ := uuid.Parse("0fef7c29-d38b-43bf-b383-eaeefcc21bc5")
	server2, _ := uuid.Parse("5142f8f6-e676-4859-b1c5-03b41912747d")
	server3, _ := uuid.Parse("230a3a8b-f9b7-42f2-a99c-25cb6038dea2")
	server4, _ := uuid.Parse("ae9fb244-e985-4254-bc09-54a3aab47060")

	// Draw hash points and server names
	for hash, server := range ring {
		angle := 2 * math.Pi * float64(hash) / 1024
		x := S/2 + float64(S/3)*math.Cos(angle)
		y := S/2 + float64(S/3)*math.Sin(angle)

		switch server {
		case server1.String():
			// Draw a small circle for the hash point
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(135, 206, 235) // Sky Blue: (135, 206, 235)
			dc.Fill()

			// // Draw the server name
			// dc.SetRGB(1, 0, 0) // Blue color for the text
			// dc.DrawStringAnchored(fmt.Sprintf("%s: %d", sh.server, sh.hash), x, y-10, 0.5, 0.5)
		case server2.String():
			// Draw a small circle for the hash point
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(255, 127, 80) // Coral Pink: (255, 127, 80)
			dc.Fill()
		case server3.String():
			// Draw a small circle for the hash point
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(50, 205, 50) // Lime Green: (50, 205, 50)
			dc.Fill()
		case server4.String():
			// Draw a small circle for the hash point
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(220, 20, 60) // Crimson Red: (220, 20, 60)
			dc.Fill()
		default:
			dc.DrawCircle(x, y, 5)
			dc.SetRGB(255, 255, 255) // black
			dc.Fill()
		}

	}

	// Draw the keys
	for _, key := range keys {
		angle := 2 * math.Pi * float64(c.GetHasher().hash_to_used(key)) / 1024
		x := S/2 + float64(S/3)*math.Cos(angle)
		y := S/2 + float64(S/3)*math.Sin(angle)

		dc.DrawCircle(x, y, 5)
		dc.SetRGB(1, 0, 0)
		dc.Fill()

		// Draw the name
		dc.SetRGB(1, 0, 0)
		dc.DrawStringAnchored(key, x, y-10, 0.5, 0.5)
	}

	// Save the image
	dc.SavePNG("hash_ring_diagram.png")
}
