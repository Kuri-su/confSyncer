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
package confsyncer

import (
	"errors"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/Kuri-su/confSyncer/pkg/unit"
)

func ConfigPush(cmd *cobra.Command, args []string) {
	f := func() error {
		// check
		if !unit.IsDir(TmpDirPath) {
			return errors.New("not found dir")
		}

		// TODO package this code section
		// get config
		maps, err := GetFilesMaps()

		for _, pathStruct := range maps {
			copySrc := pathStruct.Dist
			copyDist := TmpDirPath + pathStruct.Src
			log.Println(copySrc, copyDist)
			err := unit.Copy(copySrc, copyDist)
			if err != nil {
				return err
			}
		}

		err = unit.GitCommitAndPush(TmpDirPath)
		if err != nil {
			color.Red(err.Error())
		}

		color.Green("Configs push finish!")
		return nil
	}

	err := f()
	if err != nil {
		color.Red(err.Error())
		return
	}
}
