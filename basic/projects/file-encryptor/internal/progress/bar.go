package progress

import (
	"fmt"
	"strings"
	"time"
)

// Bar represents a progress bar
type Bar struct {
	total       int64
	current     int64
	width       int
	startTime   time.Time
	lastUpdate  time.Time
	description string
}

// NewBar creates a new progress bar
func NewBar(total int64, description string) *Bar {
	return &Bar{
		total:       total,
		current:     0,
		width:       50,
		startTime:   time.Now(),
		lastUpdate:  time.Now(),
		description: description,
	}
}

// Update updates the progress bar
func (b *Bar) Update(current int64) {
	b.current = current

	// Throttle updates to avoid excessive printing
	if time.Since(b.lastUpdate) < 100*time.Millisecond && current < b.total {
		return
	}
	b.lastUpdate = time.Now()

	b.render()
}

// Finish completes the progress bar
func (b *Bar) Finish() {
	b.current = b.total
	b.render()
	fmt.Println()
}

// render draws the progress bar
func (b *Bar) render() {
	percentage := float64(b.current) / float64(b.total) * 100
	filled := int(float64(b.width) * float64(b.current) / float64(b.total))

	bar := strings.Repeat("█", filled) + strings.Repeat("░", b.width-filled)

	// Calculate speed
	elapsed := time.Since(b.startTime).Seconds()
	speed := float64(b.current) / elapsed / 1024 / 1024 // MB/s

	// Format sizes
	currentMB := float64(b.current) / 1024 / 1024
	totalMB := float64(b.total) / 1024 / 1024

	// Print progress bar
	fmt.Printf("\r[%s] %.1f%% | %.2f MB/%.2f MB | %.2f MB/s",
		bar, percentage, currentMB, totalMB, speed)
}

// FormatBytes formats bytes to human-readable format
func FormatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// FormatDuration formats duration to human-readable format
func FormatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%.0fms", d.Seconds()*1000)
	}
	if d < time.Minute {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	return fmt.Sprintf("%.1fm", d.Minutes())
}
