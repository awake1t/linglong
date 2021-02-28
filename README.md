
# linglong

  一款资产巡航扫描系统。系统定位是通过masscan+nmap无限循环去发现新增资产，自动进行端口弱口令爆破/、指纹识别、XrayPoc扫描。主要功能包括: `资产探测`、`端口爆破`、`Poc扫描`、`指纹识别`、`定时任务`、`管理后台识别`、`报表展示`。 使用场景是Docker搭建好之后，设置好你要扫描的网段跟爆破任务。就不要管他了，没事过来收漏洞就行了



### 功能清单

- [x] masscan+namp巡航扫描资产
- [x] 创建定时爆破任务(FTP/SSH/SMB/MSSQL/MYSQL/POSTGRESQL/MONGOD)
- [x] 管理后台识别
- [x] 结果导出
- [x] 报表展示
- [x] docker一键部署 [21-02-08] 
- [x] CMS识别 - 结合威胁情报、如果某个CMS爆出漏洞，可以快速定位企业内部有多少资产 [21-02-20]
- [x] poc扫描 - 调用xray的Poc,对新发现的资产自动扫描poc [21-02-20]



## 预览
tip:如果图片加载不出来,[点我去gitee看图片](https://gitee.com/awake1t/linglong)

**首页**
![image](https://github.com/awake1t/linglong/blob/master/img/index.gif)

**资产列表**
![image](https://gitee.com/awake1t/linglong/raw/master/img/ip.png)

**敏感后台**
![image](https://gitee.com/awake1t/linglong/raw/master/img/login.png)

**指纹管理**
![image](https://github.com/awake1t/linglong/blob/master/img/finger.gif)

**任务列表**
![image](https://github.com/awake1t/linglong/blob/master/img/task.gif)

**任务详情**
![image](https://gitee.com/awake1t/linglong/raw/master/img/task-de.png)

**xray**
![image](https://gitee.com/awake1t/linglong/raw/master/img/xray.png)
![image](https://gitee.com/awake1t/linglong/raw/master/img/xray-poc.png)


**设置**
![image](https://github.com/awake1t/linglong/blob/master/img/setting.gif)





**管理后台识别**

  不论甲方乙方。大家在渗透一个网站的时候，很多时候都想尽快找到后台地址。linglong对自己的资产库进行Title识别。然后根据title关键字、url关键字、body关键字(比如url中包含login、body中包含username/password)进行简单区分后台。帮助我们渗透中尽快锁定后台。 



**指纹识别**

  系统会对新发现的资产进行一遍指纹识别, 也可以手动新增指纹。比如某个CMS爆出漏洞，新增指纹扫描一遍系统中存在的资产。可以快速定位到漏洞资产，对于渗透打点或者甲方应急都是极好的



**POC扫描**

  对于任何一个扫描系统，poc扫描都是必不可少的。但是poc的更新一直是所有开源项目面临的一个问题。综合考虑用Xray的poc,系统集成的XRAY版本是1.7.0,感谢Xray对安全圈做出的贡献！ linglong会对每次新发现的资产进行一遍Xray的Poc扫描。如果发现漏洞会自动入库，可以可视化查看漏洞结果



**资产巡航更新**

  masscan可以无限扫描，但是对于失效资产我们也不能一直保存。linglong通过动态设置资产扫描周期，对于N个扫描周期都没有扫描到的资产会进行删除。保存资产的时效性



## 安装

### Docker安装

#### 如果部署在本地体验

如果是**本地体验**下，直接运行如下命令

```bash
git clone https://github.com/awake1t/linglong
cd linglong
docker-compose up -d
```

运行结束后,运行`docker container ls -a`看下是否运行正常

![image](https://github.com/awake1t/linglong/blob/master/img/docker.png)

web访问 http://ip:8001
登录账号:linglong
登录密码:linglong5s


| Web账号        | linglong                                 | linglong5s|
| -------------- | ------------------------------------- | ---------- |
| 类型           | 用户名                                | 密码       |
| mysql数据库 | root                                  | linglong8s  |

**注: 首次运行在设置里修改扫描的网段范围,点击保存后就行了**




#### 如果部署在服务器上

```bash
git clone https://github.com/awake1t/linglong
cd linglong/web

# 把 YourServerIp 换成你的IP地址
sed -i 's#http://127.0.0.1:18000#http://YourServerIp:18000#g' ./dist/js/app.4dccb236.js && sed -i 's#http://127.0.0.1:18000#http://YourServerIp:18000#g' ./dist/js/app.4dccb236.js.map


# 如果之前安装过，使用如下命令删除所有名字包含linglong的历史镜像
docker rmi $(docker images | grep "linglong" | awk '{print $3}') 


# 返回到 linglong的目录下
cd ../
docker-compose up -d
```
![image](https://github.com/awake1t/linglong/blob/master/img/docker2.png)




## 未来

  我觉得一个好的工具就是 **拿来就用、用完既走**。后期会加入漏洞的机器人通知，发现漏洞自动通知到机器人，连你登录系统的步骤都省略。 或者看有没有方式把Goby集成。加油，干饭人！

## 更新日志

- [x] [2021-0210] 指纹资产管理、增删改查
- [x] [2021-0210] 优化资产的查询方式
- [x] [2021-0213] 发现资产POC自动扫描、扫描结果界面查看、删除
- [x] [2021-0214] 密码修改功能，关闭Xray-server-error
- [x] [2021-0215] Docker折腾了好久
- [x] [2021-0225] 更新docker-compose的部署方式

## 致谢

https://github.com/ysrc/xunfeng

https://github.com/chaitin/xray




