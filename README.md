# 1 GO语言基本使用
## 1.1 go环境配置
### 1.1.1 golang安装
建议开发环境为linux。
访问https://golang.google.cn/dl/，go语言版本需要大于1.15，建议使用最新。  
使用wget命令下载
wget https://golang.google.cn/dl/go1.16.4.linux-amd64.tar.gz
从示例命令中的go1.16.4.linux-amd64.tar.gz可以得出:
下载的go版本为go1.16.4 操作系统为linux，体系结构(大致等同于CPU运行的指令集)为amd64，如果intel和amd CPU体系结构一般均为amd64。
使用tar -C /usr/local -xzf进行解压
至此golang官方提供的开发工具已经完成安装
### 1.1.2 golang环境配置(GOPATH与GOROOT)
其中/usr/local为GOROOT前缀
默认GOPATH为~/go
上述~是指，操作系统为当前登录的用户创建的路径，其
GOPATH为工作时使用的地方，可以根据实际情况进行修改
GOROOT一般在安装完成之后不进行修改

GOROOT是go语言编译器、标准库、一些官方提供工具时以及编译使用的汇编文件存放的地方。一般不直接使用。
GOPATH路径下会有3个文件夹，分别为src、pkg、bin。
src为源代码，一般进行开发时编写代码的位置
pkg/mod为go下载程序编译时使用的用其他开发人员开发的模块文件(非标准库)存放的地方
bin为运行go get -u 命令时，会下载一些代码并编译成二进制可执行文件。该路径为这些二进制可执行文件存在的地方。
例如运行go get -u github.com/hhatto/gocloc/cmd/gocloc 后
可以在\$GOPATH/bin下发现新增一个名叫gocloc二进制可执行文件，这是一个挺好用的代码行数统计工具(不要现在就运行,环境还没配完)。

由于linux处于安全考虑 在运行当前工作路径(即运行pwd命令时，显示的路径就是当前工作路径)下文件时需要加上./
在bash(linux为每一个命令行、终端均会启动一个，响应你输入的进程名)中输入非当前路径下的命令时：
bash会查询环境变量PATH，寻找PATH下是否有二进制可执行文件名称与所执行的命令相同，如果有就执行该二进制文件，并将命令后输入的参数，作为启动参数传给该进程。
查看环境变量 echo \$PATH  (有\$符号的一般是环境变量的意思)
PATH一般为，通过符号:分开的若干文件夹路径

因此建议将\$GOROOT放入PATH环境变量
其为 export PATH=\$PATH:/usr/local/go/bin   
(/usr/local/go应该与$GOROOT相同)
该命令是 声明一个环境变量 PATH 其值为 (读取当前环境变量PATH的值,在该值之后加上:/usr/local/go/bin)
这样在终端中直接输入命令go时，终端才能找到所期望执行的可执行文件。
当然可以根据个人喜好决定是否将\$GOPATH/bin加入\$PATH
### 1.1.3 golang环境配置(GOPROXY与GOMOD)
golang官方提供了包代理网站，但是由于该网站部署在谷歌服务器上。
由于众所周知的原因导致，大陆地区基本无法使用。
但是默认下载所使用包时，会经由该代理，导致无法使用。
因此需要手动设置go proxy
输入go env，命令行显示go语言编译时的环境变量。
如果提示找不到go，那么请根据1.1.2检查\$PATH中是否含有/usr/local/go/bin。(查看环境变量命令 echo \$环境变量,例如:echo $PATH)
事实上输入go env应当等价于输入/usr/local/go/bin/go env

修改go env
输入命令go env -w GOPROXY=https://goproxy.cn,direct 
goproxy.cn是七牛云提供，可在中国大陆访问的golang代理网址
go mod是官方自go v1.11提供的包管理工具，直至今天已经占据了约90%的go语言包管理工具份额。本项目也使用了go mod进行包管理。
因此需要启用包管理工具
go env -w GO111MODULE=on
### 1.1.4 git的基本使用
本项目的git地址 https://github.com/mk1010/industry_adaptor.git
可以cd 至\$GOPATH/src下使用
git clone 某个项目的git地址
将远程仓库进行复制到本地，创建一个本地仓库。
git仓库分为远程仓库与本地仓库，在代码托管网站(例如github)上保存的为远程仓库。
在主机本地上保存的为本地仓库，本地仓库位置为该项目的根目录下有一个名称为.git的隐藏文件夹，该文件夹为git工具维护的本地仓库。请不要直接修改该文件夹。
一个项目仓库中存在若干分支，各分支由若干结点组成，分支直接具有相互继承关系，结点之间存在父子关系。
各结点实质为项目代码状态。
git commit命令将会在本地仓库修改之前的结点基础上 加上开发者进行的修改后，生成一个新的本地结点。
git push是将本地仓库的最新结点推送给远程仓库。
git pull是获取远程仓库中分支的最新结点。
结点冲突请自行bing/google git merge,在没有commit之前，git stash 也挺好用的(x)
## 1.2 go mod 
### 1.2.1 go module
1. module就是模块。
模块的具体使用可以参看go语言圣经的第二章第六节，包与模块就是同一个东西。
https://books.studygolang.com/gopl-zh/ch2/ch2-06.html
首先go语言中代码都会保存在文件中，文件隶属于某个模块，而模块与项目路径有关。以下使用./代指本项目的根目录。
如果安装之前教程，这个路径可用在终端上输入 cd \$GOPATH/src/idustry进行跳转，得到的就是本项目的根目录。
可以看到./task下有若干go语言源代码文件(.go文件)
由于这些文件在同一个路径下，因此文件第一行声明其所属的go模块应当一致，示例为task。
2. 开发规范
模块名应当与所在文件夹名字一致。
同一模块代码，在该模块不同文件中是没有区别的。但是一个文件最好是某个功能的全部实现与定义，并将该功能作为文件名，方便阅读。
模块全称为go.mod声明的项目模块名(见1.2.2)加上相对路径。
例如本项目的模块名为github.com/mk1010/idustry
那么其他模块需要引用task模块，应当import
github.com/mk1010/idustry/task
编写代码使用时即可通过task.sth(something) 进行使用
3. 模块不具有层级结构
   本项目./modules中下面有go代码文件与其他文件夹
   这些文件夹以及文件夹里面的文件与go代码文件不存在任何关系
   例如./modules中的rmq文件夹
   github.com/mk1010/idustry/modules/rmq
   使用为rmq.sth
   github.com/mk1010/idustry/modules
   使用为
   modules.sth
4. import 同名模块处理
例如本项目的main包的main.go中
包含了 github.com/mk1010/idustry/config 与
github.com/apache/dubbo-go/config
从路径上看这是2个不同的包，但有着同样的包名
因此 可以通过import 别名解决
dubboConfig "github.com/apache/dubbo-go/config"
5. 模块命名规范
   任何模块命名不允许含有非ASCII字符,例如中文
项目模块名如果需要发布应当为项目代码托管的git地址。(自己本地写的玩具代码就不用了)
6. 禁止循环依赖
   模块 A依赖模块B时，模块B不允许依赖模块A。当然实际上比这个例子更严格。
    示例为直接依赖，实际上要求间接依赖也不允许循环。
    即A依赖B、B依赖C、C依赖D等等。这中间B、C、D均不允许依赖模块A。
    如果两个模块需要相互调用，可行的办法是，把相互调用部分放在同一个模块A，模块B与C调用A即可。
### 1.2.2 go mod 使用
开发者需要为各个项目进行包管理
在各使用go mod进行包管理的项目的根路径中，会看到go.mod与go.sum文件。
如果没有，那么说明该项目刚刚创建或者不使用go mod进行管理
cd命令为切换当前工作路径，切换到该项目的根路径下
执行 go mod init
该命令会创建这2个文件

go.mod记录了当前项目依赖的非标准库的额外模块(包)
go.mod文件第一行声明了该项目的模块名
module github.com/mk1010/idustry

接下来声明go语言，语法版本。因此编译器需要大于等于该版本

require声明该项目依赖的第三方模块及版本号

go mod tidy命令会移除该项目不再依赖的模块，以及加入新依赖的模块
go.sum记录各模块的h1值，防止源代码被纂改。

go mod download会下载该项目所依赖的所有模块，及这些模块依赖的模块，依次递归。
下载到$GOPATH/pkg/mod中。

go build ldflags='-s -w'
go build会根据该项目的main包，生成一个二进制可执行文件 ldflags为编译时参数，影响编译。具体可以查阅文档。(当然也可以不写 直接go build)go build时工作路径(pwd输出的那个)下应当是个main包
