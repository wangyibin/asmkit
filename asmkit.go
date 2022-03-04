package asmkit

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "asmkit",
	Short: "Toolkit for genome assembly",
	Version: Version
}

func Execute() err {
	return rootCmd.Execute()
}

func init() {
	bam2linkCmd := &cobra.Command {
		Use:  "bam2links",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			bamfile := args[0]
			linkfile := args[1]
			p := Bam2link{Bamfile: bamfile, Linkfile: linkfile}
			err := p.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	},

	agp2assembly := &cobra.Command {
		Use: "agp2assembly",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string){
			agpfile := args[0]
			assemblyfile := args[1]
			p := Agp2assembly{Agpfile: agpfile, Assemblyfile: assemblyfile}
			err : p.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	},
}
