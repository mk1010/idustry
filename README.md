# 1 GO语言基本使用
## 1.1 go环境配置
### 1.1.1 golang安装
建议开发环境为linux。  
访问https://golang.google.cn/dl/  ，go语言版本需要大于1.15，建议使用最新。  
使用wget命令下载  
wget https://golang.google.cn/dl/go1.16.4.linux-amd64.tar.gz  
从示例命令中的go1.16.4.linux-amd64.tar.gz可以得出:  
下载的go版本为go1.16.4 操作系统为linux，体系结构(大致等同于CPU运行的指令集)为amd64，例如intel和amd CPU体系结构一般均为amd64。  
使用tar -C /usr/local -xzf进行解压  
至此golang官方提供的开发工具已经完成安装  
### 1.1.2 golang环境配置(GOPATH与GOROOT)
其中/usr/local为GOROOT前缀  
默认GOPATH为~/go  
上述~是指，操作系统为当前登录的用户创建的路径。  
GOPATH为工作时使用的地方，可以根据实际情况进行修改  
GOROOT一般在安装完成之后不进行修改  

GOROOT是go语言编译器、标准库、一些官方提供工具时以及编译使用的汇编文件存放的地方。一般不直接使用。  
GOPATH路径下会有3个文件夹，分别为src、pkg、bin。  
src为源代码，一般进行开发时编写代码的位置  
pkg/mod为go下载程序编译时使用的用其他开发人员开发的模块文件(非标准库)存放的地方  
bin为运行go get -u 命令时，会下载一些代码并编译成二进制可执行文件。该路径为这些二进制可执行文件存在的地方。  
例如运行go get -u github.com/hhatto/gocloc/cmd/gocloc 后  
可以在\$GOPATH/bin下发现新增一个名叫gocloc二进制可执行文件，这是一个挺好用的代码行数统计工具(不要现在就运行,环境还没配完)。  
  
由于linux出于安全考虑 在运行当前工作路径(即运行pwd命令时，显示的路径就是当前工作路径)下的可执行文件时需要加上./  
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
一个项目仓库中存在若干分支，各分支由若干结点组成，分支之间具有相互继承关系，结点之间存在父子关系。  
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
如果按照之前教程，这个路径可用在终端上输入 cd \$GOPATH/src/idustry进行跳转，得到的就是本项目的根目录。  
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
go build会根据该项目的main包，生成一个二进制可执行文件 ldflags为编译时参数，影响编译。具体可以查阅文档
https://golang.org/cmd/go/  
(当然也可以不写 直接go build)go build时工作路径(pwd输出的那个)下应当是个main包  
# 2 启动配置及过程
## 2.1 init函数
init函数见  
https://books.studygolang.com/go42/content/42_07_package.html#76-%E5%8C%85%E7%9A%84%E5%88%9D%E5%A7%8B%E5%8C%96  
该书本章节简要介绍了，go语言中的包。  
由于本书编写时间较早，因此建议给出的包管理工具并不是go mod。  
init函数具有先后次序，对于某个包，先进行包级变量的初始化，紧接着进行init函数调用。  
而各包的初始化也具有先后顺序。  
其初始化顺序是，从main包开始，根据main包依赖的其他包，自上而下进行初始化后再初始化main包。  
如果初始化时该包依赖了其他包，也是先对依赖包自上而下进行初始化。  
每个包只会被初始化一次。  
init函数使用规则，见参考用书。  
在本项目中main包通过将依赖包重命名为 _  
显示声明需要进行包init初始化，而不进行调用  
main包依赖，本项目的bash包  
bash包的目的就是替代启动脚本  
之所以不使用启动脚本的原因是，考虑到项目实际运行环境的复杂多变，例如windows，非容器。  
以及实际使用人员的计算机知识匮乏，就不要再来个启动脚本折腾了。  
在所有包的init函数执行完毕后，开始执行main包的main函数。
## 2.2 配置文件介绍
### 2.2.1 日志配置文件
在本项目bash/conf.go文件中包含了日志输出的配置文件。  
日志的配置文件是采用uber开源的zap所依赖的日志配置文件。  
dubbo-go对zap进行了二次开发及封装，因此本项目统一日志包为  
github.com/apache/dubbo-go/common/logger  
日志文件完成了保存过去7天日志文件(没有容器，只能这样了)，按每天每小时进行日志划分。
日志配置文件各项说明可以查看zap源码  
go.uber.org/zap下存在文件名为config.go的源码文件，其中包含结构体Config。参考源代码注释即可。
### 2.2.2 dubbo配置文件
本项目dubbo配置文件在./conf/server.yml中  
也可以参考 github.com/apache/dubbo-go/config 下provider_config.go中的源代码ProviderConfig struct进行阅读。
1. application:
应用名与版本号作为键值，其不允许跨版本调用。  
如果修改后的服务及接口对以前保持兼容，那么不需要修改版本号。  
其余配置项不是很重要(吧)。  
2. registries
注册中心名称zk，使用的协议zookeeper,其余的参阅dubbo-go源代码
查看方式如下:
"github.com/apache/dubbo-go/registry/zookeeper"
将上述包中的zookper替换成其他可用注册协议例如etcdv3
在该文件夹下的registry.go文件中的init函数会声明，其实例对应配置文件中的键值。  
同时调用之前也应该import这个包进来，保证init被调用。本项目在main包的main.go中import了zookpeer。  
如果需要切换成etcdv3等其他协议需要同时修改配置文件以及main包import的依赖包为
"github.com/apache/dubbo-go/registry/etcdv3"  
还有一项比较重要的是zone
声明注册中心的机房信息，根据这个防止跨机房调用。
当然 也可以输入账号密码 这里默认为空。  
3. services
作为服务提供方所提供的服务，服务配置文件采取map[string]ServiceConfig形式提供  
服务名为驼峰命名法，由于dubbo一开始为java语言开发，因此为了保证兼容性，首字母需要小写。  
需要声明：该服务注册的注册中心、底层协议、接口名(对应java的包名，对于dubbo-go而言没什么用)、负载均衡方法、集群容错策略。
声明该服务为具有的方法:
具体可选配置见源码。  
本项目进行了接口名、接口负载均衡方法和重试次数设置。
注意这里的负载均衡方法与上面声明的调用时机不一样。上面是当服务进行多注册中心注册时，选取注册中心的方法。接口内是调用接口时选择的负载均衡方法。  
4. protocols
声明底层协议的名称 ip与port。  
也可以不指定默认分配。  
5. protocol_conf
字面意思，一般为getty配置相同。不一定会用到
### 2.2.3 基础设施配置文件
dbs声明连接的数据库，为数组类型  
db为数据库database为数据库名称，setting为数据库设置，该设置会在数据库连接时传入。  
配置文件支持读写分离。
write为写数据库配置，包含用户名、密码、ip+port及consul(未启用)  
read为读数据库配置，为数组类型。  
选取时，默认采取随机选取读数据库进行读取。 
注意默认负载均衡算法将一定程度上集群性能，因此建议开启sticky(粘性)链接， 因为按照设计不会给服务端配置反向代理。  
并且需要维护链接状态。
负载均衡算法也可以考虑least active。  
本项目配置文件为单机。  
env声明gin框架的运行环境，目前没什么用。  
rmq_naming_service、redis_cluster_name与redis_hosts是字面意思。  
# 3 NC-Link协议
适配器项目地址  
https://github.com/mk1010/industry_adaptor  
## 3.1 NC-Link服务介绍
服务由五个接口组成，接口实现在 ./service 下。
nclink_base.go中的NcLinkServiceProvider结构体，是服务提供的基础结构体。
其持有的方法，即为RPC所调用的方法。接口定义在protobuf生成的go文件中声明。  
其Reference方法返回的字符串，需要与配置文件中服务名相同。  
dubbo通过该字符串进行配置文件查询。  
由于proto3默认，也仅支持optional语义。  
因此下文描述字段均为optional语义。  
1. NclinkAuth 
未被启用  
2. NCLinkGetMeta 
适配器向代理请求数据时使用   
查询时请求包含5个slice分别为5种元数据ID  
该接口从数据库中查询，并按指定结构对数据进行组织。  
值得一提的是，如果某ID未查询到对应数据，采取的策略是不对baseResp中的error进行处理。
这样的调用不会返回error。  
3. NCLinkSendBasicData
未被启用 
其设计是进行byte型数据(例如图片)发送的通用接口，根据message kind进行结构解析。  
由于byte型数据结构及组织形式还未确定，因此未被启用。
建议在proto文件中定义枚举值
4. NCLinkSendData
现有逻仅仅对数据库结构和请求结构体进行了转换，转换后写入数据库。  
5. NCLinkSubscribe
订阅接口，为双向流式接口。
在适配器和代理之间进行发布订阅通信模型实现。
其实现为循环订阅，即当网络发生异常时，进行不停循环订阅。周期为3秒进行一次订阅操作。  
适配器发送的第一条消息必须为NCLinkSub消息  
其实现在industry_adaptor/task/topic/common.go中  
随后通过switch选择处理函数  
这里默认没有找到处理函数的主题订阅进行拒绝。  
默认行为可以根据实际情况进行修改。例如:如果rmq对应的topic存在,  调用一个公共的数据转发函数。  
在实现中 适配器的所有订阅，最终都会转换成为对rmq的订阅。通过为rmq中各topic维护的sync.Map，向负载处理适配器链接的协程写入数据。以实现协程中数据分发。  
## 3.2 NCLink元数据
元数据结构见  
industry_adaptor/nclink/NC-Link.proto
1. NCLinkAdaptor
适配器对象，其由适配器ID、名称、适配器类型、摘要注释、运行时监控的NCLink设备ID以及由用户自定义的设备配置组成。  
由于设备配置文件读取后决定NCLink设备启动方式、监控方式以及元数据的交互方式。适配器访问NCLink设备元数据时，应该通过进程唯一的NCLink设备哈希表进行查询。  
如果哈希表中没有对应键值对，那么应当向代理发起查询请求，并将请求结果写入哈希表中。
2. NCLinkDevice
设备对象，其由设备ID、名称、设备类型、摘要注释、设备组、设备组件ID、设备组件配置、NCLink数据采集所使用的数据项组成。  
与适配器结构类似，读取组件配置后决定NCLink组件启动方式及消息整理方法。其持有的NCLinkDataItem应当为指针持有，每个DataItem应当全局唯一。  
其标识了该设备会产生数据的具体类型与结构。同样的组件元数据也由一个全局唯一的哈希表进行管理。
3. NCLinkComponent
组件对象，其由组件ID、名称、组件类型、摘要注释、NCLinkDataInfo对象组成。组件为NCLink协议中划分的最小独立存在结点。  
其记录了该组件会采样得到的数据类型(NCLinkDataItem)以及采样周期(NCLinkSampleInfo)。所有数据均属于某一组件，组件隶属于某一设备，设备隶属于某一适配器采集。值得一提的是，这种关系并不是强关联，而是允许动态迁移。  
**数据类型与采样周期会被多组件共享，因此禁止修改。**
## 3.3 NCLink对象实例
由于go禁止循环依赖，但是为了给各层级结构提供互操作的可能性。其结构见
industry_adaptor/task/common 
industry_adaptor/task/adaptor  
industry_adaptor/task/device  
industry_adaptor/task/component
除其中common包外，各包均有common.go文件，其中声明工厂函数，根据type进行实例化。  
实例化存在common包中，common包中定义了API以实现互操作。  
common包中还包含了sync.Map，其具体键值类型见注释。  
因此可以通过ID获取实例(如果实例存在并有效)。
