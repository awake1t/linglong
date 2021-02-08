# linglong

  一款甲方资产巡航扫描系统。系统定位是发现资产，进行端口爆破。帮助企业更快发现弱口令问题。主要功能包括: `资产探测`、`端口爆破`、`定时任务`、`管理后台识别`、`报表展示`。

## 简介	
> 企业资产的全面收集才是进行下一步的黑盒漏洞探测前提。linglong使用masscan无限循环进行资产探测，配合nmap进行端口指纹识别。保证企业更快发现企业暴露资产。

![image](https://github.com/awake1t/linglong/blob/master/img/index.png)
![image](https://github.com/awake1t/linglong/blob/master/img/iplist.png)
![image](https://github.com/awake1t/linglong/blob/master/img/tasklist.png)
![image](https://github.com/awake1t/linglong/blob/master/img/addtask.png)
![image](https://github.com/awake1t/linglong/blob/master/img/tasklog.png)
![image](https://github.com/awake1t/linglong/blob/master/img/setting.png)

### 功能清单

- [x] masscan+namp巡航扫描及时发现资产
- [x] 创建定时爆破任务(FTP/SSH/SMB/MSSQL/MYSQL/POSTGRESQL/MONGOD)
- [x] 管理后台识别
- [x] 结果下载
- [x] 报表展示
- [x] docker部署 - 方便体验、部署 [21-02-08] 
- [ ] CMS识别 - 结合威胁情报、如果某个CMS爆出漏洞，可以快速定位企业内部有多少资产
- [ ] poc扫描 - 还在考虑怎么加入。主要是考虑poc的长期更新。目前暂定调用xrayPoc,Xray牛逼！


**关于管理后台识别**

    不论甲方乙方。大家在渗透一个网站的时候，很多时候都想尽快找到后台地址。而linglong对自己的资产库进行Title识别。然后根据title关键字、url关键字、body关键字(比如url中包含login、body中包含username/password)进行简单区分后台。帮助我们渗透中尽快锁定后台。  

**关于资产更新**

    masscan可以无限扫描，但是对于失效资产我们也不能一直保存。linglong通过动态设置资产扫描周期，对于N个扫描周期都没有扫描到的资产会进行删除。保存资产的时效性


## 安装

### Docker安装
```
https://github.com/awake1t/linglong
cd linglong
docker-compose up -d
```

运行结束后访问 http://ip:18000
账号:linglong
密码:linglong5s



### 本地安装

先保证系统安装了 `Massscan`、`Nmap`、`Mysql`.


```
// 下载
git clone https://github.com/awake1t/linglong
cd linglong

// 导入数据库
mysql -uroot -p < mysql/init.sql

// 修改配置文件中数据库密码为你的数据库密码,文件第16行
vim ./configs/config.yaml 

// 运行
chmod +x linglong && ./linglong

访问http://127.0.0.1:18000 ，进入登陆界面。 账号:linglong 密码:linglong5s 。进入设置界面，配置你要扫描资产、端口。安装完成
```

> 因为代码是前后端分离。上面的安装方式部署在自己电脑上是没问题的。 但是如果部署到服务器上，你会发现。进入登陆界面。点击登陆没有任何反应。因为为了不让大家安装vue，vue的代码是写死的。
需要修改前端请求后端的地址。使用如下命令把`http://10.10.10.10:80`。替换成你的服务器协议+ip+port。就可以咯～

```
sed -i 's#http://127.0.0.1:18000#http://10.10.10.10:80#g' ./dist/js/app.48c176d1.js && sed -i 's#http://127.0.0.1:18000#http://10.10.10.10:80#g' ./dist/js/app.48c176d1.js.map
```


### 致谢

https://github.com/ysrc/xunfeng



