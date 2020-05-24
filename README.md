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

1. init confsyncer's config

    ```shell
    # init confsyncer's config
    $ confsyncer init
    init confsyncer's configfile in ~/.confsyncer/config.yaml
    
    This is your config: 
    {
        "gitpulltimeinternal": 30,
        "maps": [
            {
                "src": "// TODO SourceFilePathOfGitRepo",
                "dist": "// TODO FilePathOfLocal"
            }
        ],
        "gitrepo": "git@gitlab.com:examples/examples.git"
    }
    
    you should modify it before use.
    ```

2. modify confsyncer 

3. create the git repo and push your config files to repo

   ```shell
   # show confsyncer's config
   $ confsyncer config
   xxxxxx
   
   # commit && push files to git repo
   $ confsyncer push
   /tmp/a->/tmp/b job success
   /tmp/c->/tmp/d job success
   /tmp/e->/tmp/f job failed
   ```

4. modify crontab file to pull config every 1m

    ```shell
    # 
    * * * * * bash -c "confsyncer pull"
    ```

5. 或者你也可以手动的拉取

    ```shell
    # pull config files
    $ confsyncer pull
    /tmp/a->/tmp/b job success
    /tmp/c->/tmp/d job success
    /tmp/e->/tmp/f job failed
    ```

### In Container 

1. init confsyncer's config 

   ```shell
   # init confsyncer's config
   $ confsyncer init
   init confsyncer's configfile in ~/.confsyncer/config.yaml
   
   This is your config: 
   {
       "gitpulltimeinternal": 30,
       "maps": [
           {
               "src": "// TODO SourceFilePathOfGitRepo",
               "dist": "// TODO FilePathOfLocal"
           }
       ],
       "gitrepo": "git@gitlab.com:examples/examples.git"
   }
   you should modify it before use.
   ```

2. modify confsyncer 

3. run confsyncer-gen with you confsyncer config

   ```shell
   # 
   $ confsyncer-gen -f ~/.confsyncer/config.yaml -o ./
# 然后使用 docker-compose 启动服务即可 , 重新生成后也是一样的操作和结果
   $ docker-compose up -d
   ```
   

## docker-compose.yaml Gen

// TODO docs

## LICENSE

GPL-3.0
