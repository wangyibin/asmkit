package asmkit

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "asmkit",
	Short: "Toolkit for genome assembly",
	Long: `
                               __   .__  __   
	_____    ______ _____ |  | _|__|/  |_ 
	\__  \  /  ___//     \|  |/ /  \   __\
	 / __ \_\___ \|  Y Y  \    <|  ||  |  
	(____  /____  >__|_|  /__|_ \__||__|  
		 \/     \/      \/     \/       

 `,
	Version: Version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	bam2linkCmd := &cobra.Command{
		Use:   "bam2links <input.bam> <output.links>",
		Short: "Extract mnd links from bam",
		Long: `
bam2link function:
Given a bamfile, to extract links and store as mnd links file for juicebox assembly tools.
`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			bamfile := args[0]
			linkfile := args[1]
			p := Bam2linker{Bamfile: bamfile, Linkfile: linkfile}
			err := p.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	agp2assemblyCmd := &cobra.Command{
		Use:   "agp2assembly <input.agp> <output.assembly>",
		Short: "Convert agp file into 3d-dna assembly.",
		Long: `
agp2assembly function:
Convert agp file into 3d-dna assembly for juicebox assembly tool.		
`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			agpfile := args[0]
			assemblyfile := args[1]
			p := Agp2assembler{Agpfile: agpfile, Assemblyfile: assemblyfile}
			err := p.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	var mapq int
	bamStatCmd := &cobra.Command{
		Use:   "bamStat <input.bam>",
		Short: "Stat the number of different types of alignments.",
		Long: `
bamStat function:
Stat the number of different types of alignments, including unique || 
multiple || low quality alignments.
`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			bamfile := args[0]
			p := BamStater{Bamfile: bamfile, MapQ: mapq}
			err := p.Run()
			if err != nil {
				log.Fatal(err)
			}

		},
	}
	bamStatCmd.Flags().IntVarP(&mapq, "mapq", "", 10, "Minimum map quality")

	rootCmd.AddCommand(agp2assemblyCmd, bam2linkCmd, bamStatCmd)
}
