
## Git同时向多个仓库推送代码

配置文件参考：

```conf
[core]
	repositoryformatversion = 0
	filemode = false
	bare = false
	logallrefupdates = true
	symlinks = false
	ignorecase = true
[user]
	email = mrivandu@hotmail.com
	name = cxwn
[remote "origin"]
	url = https://github.com/cxwn/notes.git
	fetch = +refs/heads/*:refs/remotes/origin/*
	url = https://gitee.com/cxwn/notes.git
[branch "master"]
	remote = origin
	merge = refs/heads/master
```

命令参考：

```bash
git config user.email "mrivandu@hotmail.com"        
git config user.name "cxwn"
git remote add origin https://github.com/cxwn/notes.git
git remote set-url --add origin https://gitee.com/cxwn/notes.git
```
注意：首先需要在GitHub和Gitee建立相关的仓库。本例中，两个仓库的用户名和密码、仓库、路径全部一致，不一致的情况尚未测试。
