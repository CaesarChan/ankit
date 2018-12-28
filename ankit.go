package ankit

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/pkg/errors"
)

// Note consists of fields.
type Note interface {
	Err() error
	Fields() []string
}

// Deck is a group of cards, which are generated by notes.
type Deck interface {
	Notes() (<-chan Note, error)
}

// Export write Deck's notes to io.Writer.
func Export(w io.Writer, d Deck) error {
	cw := csv.NewWriter(w)

	notes, err := d.Notes()
	if err != nil {
		return errors.Wrap(err, "cannot get notes")
	}

	for note := range notes {
		if err := note.Err(); err != nil {
			log.Printf("note error: %v", err)
			continue
		}

		fields := note.Fields()
		if err := cw.Write(fields); err != nil {
			return errors.Wrap(err, "cannot write field to csv")
		}
	}
	cw.Flush()

	return nil
}
