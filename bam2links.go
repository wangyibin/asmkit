package asmkit

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/biogo/hts/bam"
)

// Bam2linker
type Bam2linker struct {
	Bamfile  string
	Linkfile string
}

// LinkLine stores link information
type LinkLine struct {
	strand1 int
	r1      string
	pos1    int
	order1  int
	strand2 int
	r2      string
	pos2    int
	order2  int
	suffix  string
}

func (r *Bam2linker) ExtractLinks() error {
	f, err := os.Create(r.Linkfile)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)

	var l LinkLine
	l.suffix = "1 - - 1 - - -"
	l.strand1, l.strand2 = 0, 0

	log.Notice("Start processing bam ...")
	fh, err := os.Open(r.Bamfile)
	if err != nil {
		return err
	}
	br, err := bam.NewReader(fh, 0)
	if err != nil {
		return err
	}
	for {
		rec, err := br.Read()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		// Filtering: Multiple | Unmapped | Secondary | Duplicate | Supplementary
		if rec.MapQ == 0 || rec.Flags&3844 != 0 {
			continue
		}
		l.r1 = rec.Ref.Name()
		l.r2 = rec.MateRef.Name()
		l.pos1 = rec.Pos
		l.pos2 = rec.MatePos
		// Get strand of read
		if rec.Flags&16 == 0 {
			l.strand1 = 0
		} else {
			l.strand1 = 16
		}
		// Get strand of mate
		if rec.Flags&32 == 0 {
			l.strand2 = 0
		} else {
			l.strand2 = 16
		}
		// Get order by ref id
		if rec.Ref.ID() <= rec.MateRef.ID() {
			l.order1 = 0
			l.order2 = 1
		} else {
			l.order1 = 1
			l.order2 = 0
		}
		_, err = fmt.Fprintf(w,
			"%d %s %d %d %d %s %d %d %s\n",
			l.strand1,
			l.r1,
			l.pos1,
			l.order1,
			l.strand2,
			l.r2,
			l.pos2,
			l.order2,
			l.suffix,
		)

	}
	err = w.Flush()
	if err != nil {
		return err
	}
	log.Notice("Done.")
	return nil
}

func (r *Bam2linker) Run() error {
	err := r.ExtractLinks()
	if err != nil {
		return err
	}
	return nil
}
