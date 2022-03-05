package asmkit

import (
	"fmt"
	"io"
	"os"

	"github.com/biogo/hts/bam"
)

type BamStater struct {
	Bamfile string
	MapQ    int
}

func (r *BamStater) BamStat() error {

	// convert int to byte
	Mapq := byte(r.MapQ)

	totalNums := 0
	uniqueNums := 0
	multiNums := 0
	lowqNums := 0

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

		// Calculating the number of unique
		// || multi || low quality alignments
		if rec.MapQ < Mapq {
			if rec.MapQ == 0 {
				multiNums++
			} else {
				lowqNums++
			}
		} else {
			uniqueNums++
		}
		totalNums++
	}
	fmt.Printf(
		"unique alignments: %d\nmulti alignments: %d\nlow quality alignments: %d\ntotal alignments: %d",
		uniqueNums,
		multiNums,
		lowqNums,
		totalNums,
	)

	return nil
}

func (r *BamStater) Run() error {
	err := r.BamStat()
	if err != nil {
		return err
	}
	return nil
}
