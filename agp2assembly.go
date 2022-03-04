package asmkit

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/op/go-logging"
)


// Agp2assembler
type Agp2assembler struct {
	Agpfile      string
	Assemblyfile string
}

// AGPLine github.com/tanghaibao/allhic
type AGPLine struct {
	object        string
	objectBeg     int
	objectEnd     int
	partNumber    int
	componentType byte
	isGap         bool
	strand        byte
	// As a gap
	gapLength       int
	gapType         string
	linkage         string
	linkageEvidence string
	// As a sequence chunk
	componentID  string
	componentBeg int
	componentEnd int
}

// AGP is a collection of AGPLines
type AGP struct {
	lines []AGPLine
}

// Add adds an AGPLine to the collection
// from github.com/tanghaibao/allhic
func (r *AGP) Add(row string) {
	words := strings.Fields(row)
	var line AGPLine
	line.object = words[0]
	line.objectBeg, _ = strconv.Atoi(words[1])
	line.objectEnd, _ = strconv.Atoi(words[2])
	line.partNumber, _ = strconv.Atoi(words[3])
	line.componentType = words[4][0]
	line.isGap = line.componentType == 'N' || line.componentType == 'U'
	if line.isGap {
		line.gapLength, _ = strconv.Atoi(words[5])
		line.gapType = words[6]
		line.linkage = words[7]
		line.linkageEvidence = words[8]
	} else {
		line.componentID = words[5]
		line.componentBeg, _ = strconv.Atoi(words[6])
		line.componentEnd, _ = strconv.Atoi(words[7])
		line.strand = words[8][0]
	}
	r.lines = append(r.lines, line)
}

type assemblyGroup struct {
	lines []string
}

func (r *Agp2assembler) AGP2Assembly() error {
	// input agp
	log.Noticef("Loading agpfile `%s`.", r.Agpfile)
	agp := new(AGP)
	fh, err := os.Open(r.Agpfile)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		agp.Add(scanner.Text())
	}
	// output file
	o, err := os.Create(r.Assemblyfile)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(o)

	var build strings.Builder
	var seq string
	var data []string
	idx := 1
	prevObject := ""
	gapLen := 100
	for _, line := range agp.lines {
		if line.object != prevObject {
			if prevObject != "" {
				data = append(data, build.String())
			}

			prevObject = line.object
			build.Reset()
		}
		if line.isGap {
			if gapLen != line.gapLength {
				gapLen = line.gapLength
			}
			continue
		} else {
			_, err = fmt.Fprintf(w,
				">%s %d %d\n",
				line.componentID,
				idx,
				line.componentEnd)
			if line.strand == '+' {
				seq = fmt.Sprintf(" %d", idx)

			} else {
				seq = fmt.Sprintf(" -%d", idx)
			}
			build.WriteString(seq)
			idx += 1
		}
	}
	// gap line
	// _, err = fmt.Fprintf(w, ">hic_gap_%d %d %d\n", idx, idx, gapLen)
	// gap := fmt.Sprintf(" %d ", idx)
	// last one
	data = append(data, build.String())
	build.Reset()

	for _, v := range data {
		v = strings.Replace(v, " ", "", 1)
		_, err = fmt.Fprintf(w, "%s\n", v) // strings.Replace(v, " ", gap, -1))
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	log.Notice("Done.")
	return nil
}

func (r *Agp2assembler) Run() error {
	err := r.AGP2Assembly()
	if err != nil {
		return err
	}
}
