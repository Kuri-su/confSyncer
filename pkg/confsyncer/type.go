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

type Config struct {
	GitRepo             string `yaml:"gitRepo"`
	GitPullTimeInternal int    `yaml:"gitPullTimeInternal"`
	Maps                []Path `yaml:"maps"`
}

type Path struct {
	Local       string `yaml:"local"`
	GitRepoPath string `yaml:"gitRepoPath"`
}
