
create database linglong default character set utf8mb4 collate utf8mb4_general_ci;
use linglong;

/* ip资产列表 */
DROP TABLE IF EXISTS `iplist`;
CREATE TABLE `iplist` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  DEFAULT NULL COMMENT '资产ip',
  `port` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '开放端口',
  `protocol` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '协议',
  `cms` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'cms',
  `language` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '代码语言',
  `portnum` int(10) unsigned DEFAULT 0 COMMENT '开放端口数量',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'url',
  `loginurl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '跳转的url',
  `title` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'title',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=315 DEFAULT CHARSET=utf8;


/* 后台资产列表 */
DROP TABLE IF EXISTS `webloginlist`;
CREATE TABLE `webloginlist` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  DEFAULT NULL COMMENT '资产ip',
  `port` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '端口',
  `protocol` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '协议',
   `cms` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'cms',
  `language` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '代码语言',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'url',
  `loginurl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '跳转的url',
  `title` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'title',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=315 DEFAULT CHARSET=utf8;

/*定时任务列表*/
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `task_name` varchar(50) NOT NULL DEFAULT '' COMMENT '任务名称',
  `task_type` varchar(50) NOT NULL DEFAULT '' COMMENT '任务类型',
  `task_cycle` varchar(50) NOT NULL DEFAULT '' COMMENT '任务周期 now/执行一次 day每天xx点执行',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '任务描述',
  `cron_spec` varchar(100) NOT NULL DEFAULT '' COMMENT '时间表达式',
  `concurrent` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否只允许一个实例',
  `command` text NOT NULL COMMENT '命令详情',
  `arge` text NOT NULL COMMENT 'url参数',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0停用 1启用',
  `timeout` smallint(6) NOT NULL DEFAULT '0' COMMENT '超时设置',
  `vuln_num` int(11) NOT NULL DEFAULT '0' COMMENT '漏洞数量',
  `execute_times` int(11) NOT NULL DEFAULT '0' COMMENT '累计执行次数',
  `prev_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上次执行时间',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=315 DEFAULT CHARSET=utf8;


/*定时任务执行记录*/
DROP TABLE IF EXISTS `task_log`;
CREATE TABLE `task_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务ID',
  `output` mediumtext NOT NULL COMMENT '任务输出',
  `error` text NOT NULL COMMENT '错误信息',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  `all_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务总数量',
  `succes_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '成功数量',
  `userdict` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户名字典id',
  `passdict` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '密码字典id',
  `run_time` varchar(200) NOT NULL DEFAULT '' COMMENT '运行时间',
  `process_time` int(11) NOT NULL DEFAULT '0' COMMENT '消耗时间/毫秒',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=315 DEFAULT CHARSET=utf8;


/*端口爆破结果库*/
DROP TABLE IF EXISTS `portbruteres`;
CREATE TABLE `portbruteres` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int NOT NULL DEFAULT 0 COMMENT '定时任务id',
  `task_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '定时任务执行时间',
  `ip` varchar(150) NOT NULL DEFAULT '' COMMENT 'ip',
  `protocol` varchar(150) NOT NULL DEFAULT '' COMMENT '爆破协议',
  `port` int(20) NOT NULL DEFAULT 0 COMMENT 'port',
  `vulntype` tinyint(4) NOT NULL DEFAULT '0' COMMENT '漏洞类型 1弱口令 2未授权 3空口令',
  `user` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `pass` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=315 DEFAULT CHARSET=utf8;

/*全平台任务日志记录*/
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务ID',
  `task_name` varchar(200) NOT NULL DEFAULT '' COMMENT '任务名称',
  `task_type` varchar(200) NOT NULL DEFAULT '' COMMENT '任务类型/ masscan/title识别/端口',
  `all_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务总数量',
  `succes_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '成功数量',
  `run_time` varchar(200) NOT NULL DEFAULT '' COMMENT '运行时间',
  `status` tinyint(4) NOT NULL COMMENT '状态 1完成 2错误',
  `error` text NOT NULL COMMENT '错误信息',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=315 DEFAULT CHARSET=utf8;


CREATE TABLE `auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


/*环境配置数据库*/
DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `login_word` text NOT NULL  COMMENT '敏感后台关键字',
  `login_url` text NOT NULL COMMENT '敏感后台url',
  `masscan_deltime` int NOT NULL DEFAULT 7 COMMENT 'mass删除周期',
  `masscan_thred` int NOT NULL DEFAULT 0 COMMENT 'mass线程',
  `masscan_ip` MEDIUMTEXT NOT NULL  COMMENT 'mass要扫描的列表',
  `masscan_port` text NOT NULL  COMMENT 'mass要扫描的端口',
  `masscan_white` text NOT NULL  COMMENT 'mass不要扫描的列表',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=315 DEFAULT CHARSET=utf8;



INSERT INTO `auth` (`id`, `username`, `password`) VALUES (null, 'awake', 'awakehhhh');

INSERT INTO `setting` (`id`, `login_word`,`login_url`,`masscan_thred`,`masscan_ip`,`masscan_port`,`masscan_white`) VALUES
 (null,  '系统', 'admin', 100, '10.10.10.0/24', '80,22,3306', '10.10.10.10');





