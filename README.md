# confsyncer

A little sync files tool in the **Linux**.

## What's this

`confsyncer` is a little tool about push and pull files in git repo,  基于此 你可以 很方便的将一些配置文件在 多台机器中进行同步.

## Features

- 手动 Push 监视中的 文件到指定的 git 仓库
- 自动/手动 从指定的 git 仓库 拉取文件到本地的指定位置

## Install
```shell
$ wget xxxxxx
```

#### or
```shell
# use docker-compose with "gen"
$ wget xxxxxx
$ cp xxxxx
$ gen xxxx
```

## How to Use

### In Host

// crontab 定期 confsyncer pull 即可

### In Container 

// confsyncerGen + dockerCompose  即可



## docker-compose.yaml Gen

// TODO docs

## LICENSE

GPL-3.0
