package unit

import (
	"fmt"

	"github.com/fatih/color"
)

func GitPull(dir string) error {
	output, err := RunCommandInShellGetOutput("cd " + dir + " && git pull origin master")
	color.HiBlue(output)
	return err
}

func GitPush(dir string) error {
	output, err := RunCommandInShellGetOutput(fmt.Sprintf(`cd %s && git commit -m "sync push config" ; git push origin master`, dir))
	fmt.Println(output)
	return err
}

func GitClone(gitRepoPath string, dir string) error {
	output, err := RunCommandInShellGetOutput("git clone " + gitRepoPath + " " + dir)
	fmt.Println(output)
	return err
}
