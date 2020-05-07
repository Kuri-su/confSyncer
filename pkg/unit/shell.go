/*
 * 	ConfSyncer - a little sync config files tool in the Linux.
 *     Copyright (C) 2020  Amatist_kurisu<misaki.zhcy@gmail.com>
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package unit

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunCommandInShell(command string) {
	commandWords := strings.Split(command, " ")
	cmd := exec.Command(commandWords[0], commandWords[1:]...)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdinPipe, _ := cmd.StdinPipe()

	var stdoutBuf, stderrBuf bytes.Buffer
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		_, _ = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, _ = io.Copy(stderr, stderrIn)
	}()
	go func() {
		_, _ = io.Copy(stdinPipe, os.Stdin)
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalln(err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

func RunCommandInShellGetOutput(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func MakeDirWithFilePath(path string) error {
	return os.MkdirAll(filepath.Dir(path), 0744)
}

func CreateFile(dir string) error {
	newFile, err := os.Create(dir)
	if err != nil {
		return err
	}
	newFile.Close()
	return nil
}

func WriteFile(dir string, context []byte) error {
	return ioutil.WriteFile(dir, context, 0744)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func Copy(src, dist string) error {
	// open source file
	originalFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer originalFile.Close()
	// create new file
	newFile, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer newFile.Close()
	// copy content
	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		return err
	}
	// flush file content to disk
	err = newFile.Sync()
	if err != nil {
		return err
	}
	return nil
}

func Move(src, dist string) error {
	return os.Rename(src, dist)
}
