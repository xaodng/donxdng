github是一个基于git的代码托管平台

1.配置Git
首先在本地创建ssh key；
ssh-keygen -t rsa -C "your_email@youremail.com"

后面的your_email@youremail.com改为你在github上注册的邮箱，之后会要求确认路径和输入密码，我们这使用默认的一路回车就行。成功的话会在~/下生成.ssh文件夹，进去，打开id_rsa.pub，复制里面的key。
回到github上，进入 Account Settings（账户配置），左边选择SSH Keys，Add SSH Key,title随便填，粘贴在你电脑上生成的key。


为了验证是否成功，在git bash下输入：
ssh -T git@github.com

接下来我们要做的就是把本地仓库传到github上去，在此之前还需要设置username和email，因为github每次commit都会记录他们。
$ git config --global user.name "your name"
$ git config --global user.email "your_email@youremail.com"

远程库与本地库之间的操作：
yourName和yourRepo表示你再github的用户名和刚才新建的仓库
本地库关联远程库，在本地仓库目录运行命令：
$ git remote add origin git@github.com:yourName/yourRepo.git

从远程克隆一份到本地可以通过git clone
Git支持HTTPS和SSH协议，SSH速度更快
$ git clone git@github.com:yourName/yourRepo.git

2.工作流
查看状态：
$ git status

提交、上传：
接下来在本地仓库里添加一些文件，比如README，

$ git add README
$ git commit -m "first commit"

删除
$ git rm README
$ git commit -m "first commit"

上传到github：
$ git push origin master

git push命令会将本地仓库推送到远程服务器。
git pull命令则相反。









