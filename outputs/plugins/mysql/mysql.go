package mysql

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"hippo-data-acquisition/commons/logger"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/outputs/output_collection"
)

type MySql struct {
	host          string
	tableName     string
	dataJsonField string
	db            *sql.DB
}

func (m *MySql) testDbConn() error {
	err := m.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (m *MySql) initDb() error {
	var err error
	m.db, err = sql.Open("mysql", m.host)
	if err != nil {
		return err
	}
	return nil
}

func (m *MySql) getDb() error {
	if m.db != nil {
		err := m.testDbConn()
		if err != nil {
			return m.initDb()
		}
	} else {
		return m.initDb()
	}
	return nil
}
func (m *MySql) Insert(sql string, args ...interface{}) (int64, error) {
	stmt, err := m.db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// InitPlugin 初始化参数
func (m *MySql) InitPlugin(config config.OutputConfig) {
	host, ok := config.Params["host"]
	if ok {
		m.host = host.(string)
	} else {
		logger.LogInfo("MySql", "mysql输出插件缺少参数：host")
	}

	tableName, ok := config.Params["tableName"]
	if ok {
		m.tableName = tableName.(string)
	} else {
		logger.LogInfo("MySql", "mysql输出插件缺少参数：tableName")
	}

}

// BeforeExeOutput  执行输出前
func (m *MySql) BeforeExeOutput() {

}

// ExeOutput  执行输出
func (m *MySql) ExeOutput(dataInfo queue.DataInfo) {
	strByte, err := json.Marshal(&dataInfo)
	if err != nil {
		logger.LogInfo("mySql", "输出数据转换成json字符串失败！")
	}

	err = m.getDb()
	if err != nil {
		logger.LogInfo("mySql", "初始化mysql连接失败："+err.Error())
	} else {
		result, insertErr := m.Insert("INSERT INTO "+m.tableName+"(json_data) values (?)", strByte)
		if insertErr != nil {
			logger.LogInfo("mySql", "mysql插入数据失败："+insertErr.Error())
		} else {
			logger.LogInfo("mySql", "mysql插入数据成功："+string(result)+"！")
		}
	}

}

func init() {
	output_collection.Add("mySql", &MySql{
		host: "",
		db:   nil,
	})
}
