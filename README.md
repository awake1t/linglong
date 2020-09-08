# linglong

  一款甲方资产巡航扫描系统。系统定位是发现资产，进行端口爆破。帮助企业更快发现弱口令问题。主要功能包括: `资产探测`、`端口爆破`、`定时任务`、`管理后台识别`、`报表展示`。

### 简介	

​	企业资产的全面收集才是进行下一步的黑盒漏洞探测前提。linglong使用masscan无限循环进行资产探测，配合nmap进行端口指纹识别。保证企业更快发现企业暴露资产。  

​	**关于管理后台识别**

​	不论甲方乙方。大家在渗透一个网站的时候，很多时候都想尽快找到后台地址。而linglong对自己的资产库进行Title识别。然后根据title关键字、url关键字、body关键字(比如url中包含login、body中包含username/password)进行简单区分后台。帮助我们渗透中尽快锁定后台。  

​	**关于资产更新**

​	masscan可以无限扫描，但是对于失效资产我们也不能一直保存。linglong通过动态设置资产扫描周期，对于N个扫描周期都没有扫描到的资产会进行删除。保存资产的时效性


![image](https://github.com/awake1t/linglong/tree/master/configs/img/index.png)
![image](https://github.com/awake1t/linglong/tree/master/configs/img/iplist.png)
![image](https://github.com/awake1t/linglong/tree/master/configs/img/tasklist.png)
![image](https://github.com/awake1t/linglong/tree/master/configs/img/addtask.png)
![image](https://github.com/awake1t/linglong/tree/master/configs/img/tasklog.png)
![image](https://github.com/awake1t/linglong/tree/master/configs/img/setting.png)


### 本地安装

先保证系统安装了 `Massscan`、`Nmap`、`Mysql`.

```
# 下载
git clone https://github.com/awake1t/linglong
cd linglong

# 导入数据库
mysql -uroot -p < linglong.sql

# 修改配置文件中数据库密码为你的数据库密码,文件第16行
vim ./configs/config.yaml 

# 运行
chmod +x linglonglinux && ./linglonglinux

访问http://127.0.0.1:18000 ，进入登陆界面。 账号:awake 密码:awakehhhh 。进入设置界面，配置你要扫描资产、端口。安装完成
```


### 部署到服务端

> 因为代码是前后端分离。上面的安装方式用自己电脑安装没问题。 如果部署到服务端需要修改前端请求后端的地址。使用如下命令把`YourServerDomain`。替换成你的服务器域名，比如github.com。就可以咯～

```
sed -i 's#http://127.0.0.1:18000#http://YourServerDomain#g' ./dist/js/app.48c176d1.js && sed -i 's#http://127.0.0.1:18000#http://YourServerDomain#g' ./dist/js/app.48c176d1.js.map
```



### 已完成

- [x] masscan+namp巡航扫描及时发现资产
- [x] 创建定时爆破任务(FTP/SSH/SMB/MSSQL/MYSQL/POSTGRESQL/MONGOD)
- [x] 管理后台识别
- [x] 结果下载
- [x] 报表展示


### 更多
> 表哥，我想要个你的Star (๑•̀ㅂ•́)و✧



### 致谢

https://github.com/ysrc/xunfeng



