package plugins

type ScanFunc func(ip string, port string, username string, password string) (err error, result bool)

var (
	ScanFuncMap map[string]ScanFunc
)

func init() {
	ScanFuncMap = make(map[string]ScanFunc)
	ScanFuncMap["FTP"] = ScanFtp
	ScanFuncMap["SSH"] = ScanSsh
	ScanFuncMap["SMB"] = ScanSmb		// 139,445
	ScanFuncMap["MSSQL"] = ScanMssql
	ScanFuncMap["MYSQL"] = ScanMysql
	ScanFuncMap["POSTGRESQL"] = ScanPostgres		// postgres 5432
	ScanFuncMap["REDIS"] = ScanRedis
	//ScanFuncMap["ELASTICSEARCH"] = ScanElastic
	ScanFuncMap["MONGOD"] = ScanMongodb			// 27017
	//ScanFuncMap["JAVADEBUG"] = JavaDebug			// 9091
	//ScanFuncMap["ORACLE"] = ScanOracle			// 1521

	//161:   "SNMP",					snmp的主要作用是对网络设备和设备中的应用程序进行管理，因此，获得了snmp口令后，主要的作用就是查询系统信息
	//
	//4043	rsync
	//rfp
	//ZooKeeper     2181
	//zookeeper是分布式协同管理工具，常用来管理系统配置信息，攻击者能够执行所有只允许由管理员运行的命令。
	//Atlassian Crowd       8095
	//Elasticsearch               8080
	//Jupyter Notebook          8888
}
