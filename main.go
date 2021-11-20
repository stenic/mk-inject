package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"log"

	"github.com/NoUseFreak/go-vembed"
	"github.com/spf13/cobra"
)

var label string
var inplace bool

func init() {
	rootCmd.PersistentFlags().StringVarP(&label, "label", "l", "", "Replace a specific tag")
	rootCmd.PersistentFlags().BoolVarP(&inplace, "inplace", "i", false, "Edit in place")
	rootCmd.Version = fmt.Sprintf(
		"%s, build %s",
		vembed.Version.GetGitSummary(),
		vembed.Version.GetGitCommit(),
	)
}

var tagRegexTpl = `<!--\s+mk-inject:%s:%s[^-]+-->`
var optionRegex = `(\w+)="([^"]+)"`

func main() {
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "mk-inject [--label labelName] file",
	Short: "mk-inject",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]
		injectString, err := readStdin()
		if err != nil {
			log.Fatalf("Failed capturing input - %s", err.Error())
		}

		b, err := ioutil.ReadFile(filePath) // just pass the file name
		if err != nil {
			log.Fatalf("Could not open file - %s", err.Error())
		}

		original := string(b)

		startReg := regexp.MustCompile(fmt.Sprintf(tagRegexTpl, "start", label))
		endReg := regexp.MustCompile(fmt.Sprintf(tagRegexTpl, "end", label))
		optReg := regexp.MustCompile(optionRegex)

		startIdx := startReg.FindAllStringIndex(original, -1)
		endIdx := endReg.FindAllStringIndex(original, -1)

		out := ""
		start := 0
		for i := 0; i < len(endIdx); i++ {
			startTag := original[startIdx[i][0]:startIdx[i][1]]
			prefix := []string{""}
			suffix := []string{""}
			modes := optReg.FindAllStringSubmatch(startTag, -1)
			for _, mode := range modes {
				switch mode[1] {
				case "prefix":
					prefix = append(prefix, mode[2])
				case "suffix":
					suffix = append([]string{mode[2]}, suffix...)
				default:
					log.Fatalf("Unknown modifier %s", mode[1])
				}
			}

			out += fmt.Sprintf(
				"%s%s\n%s\n%s%s",
				original[start:startIdx[i][1]],
				strings.Join(prefix, "\n"),
				injectString,
				strings.Join(suffix, "\n"),
				original[endIdx[i][0]:endIdx[i][1]],
			)

			// reset new startPoint
			start = endIdx[i][1]
		}
		// add remainder
		out += original[start:]

		if inplace {
			info, err := os.Stat(filePath)
			if err != nil {
				log.Fatalf("Failed to access file - %s", err)
			}
			ioutil.WriteFile(filePath, []byte(out), info.Mode().Perm())
			fmt.Println("Inplace")
		} else {
			fmt.Print(out)
		}

		return nil
	},
}

func readStdin() (string, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return "", nil
	}
	var stdin []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = append(stdin, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strings.Join(stdin, "\n"), nil
}
