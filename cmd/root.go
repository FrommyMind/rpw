/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var chars = []string{"a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z"}
var upperChars = []string{"A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z"}
var numers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var simpleSymbols = []string{"-"}
var symbols = []string{"@", "#", "$", "%", "^"}
var Length = 12
var Level = 0
var allChars = append(chars, numers...)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rpw",
	Short: "rpw (Random Password) is a password generator",
	Long:  `rpw (Random Password) is used for randomly generator password`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Args: func(cmd *cobra.Command, args []string) error {
		if Level < 0 || Level > 3 {
			return errors.New("Wrong level range")
		}
		if Length < 0 {
			return errors.New("Wrong length range")
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {

		switch Level {
		case 1:
			allChars = append(allChars, upperChars...)
		case 2:
			allChars = append(append(allChars, upperChars...), simpleSymbols...)
		case 3:
			allChars = append(append(append(allChars, upperChars...), simpleSymbols...), symbols...)
		}

		// for i := 0; i < len(allChars); i++ {
		// 	fmt.Print(string(allChars[i]))
		// }
		if len(args) == 0 {
			var random int
			println("Your password is:")
			for i := 0; i < Length; i++ {
				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)
				random = r1.Intn(len(allChars))

				print((allChars[random]))

			}
			println()
			return nil
		} else {
			return errors.New("rpw without sub command, please use flag")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rpw.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().IntVarP(&Length, "length", "l", 12, "The length of the password")
	rootCmd.PersistentFlags().IntVarP(&Level, "level", "e", 0, `The level of complexity
	0 : Only use lower case alphabet [a-z] and numers [0-9]
	1 : Use upper case alphabet [A-z],  plus level 0
	2 : Use symbols [-], plus level 1
	3 : Use symbols [@#$%+?], plus level 2`)
}
