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

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Kuri-su/confSyncer/pkg/confsyncer/ctl"
	"github.com/Kuri-su/confSyncer/pkg/unit"
)

func ConfigPush(cmd *cobra.Command, args []string) {
	f := func() error {
		if !unit.IsDir(ctl.TmpDirPath) {
			return errors.New("not found dir")
		} else {
			err := unit.GitPush(ctl.TmpDirPath)
			if err != nil {
				color.Red(err.Error())
			}
		}

		maps := viper.GetStringMapString("maps")

		for src, dist := range maps {
			err := unit.Copy(dist, ctl.TmpDirPath+src)
			if err != nil {
				return err
			}
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
