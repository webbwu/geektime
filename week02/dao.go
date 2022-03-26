package main

import (
	"database/sql"

	"github.com/pkg/errors"
)

/* 不应该把  sql.ErrNoRows  抛给业务层，查询不到数据严格不属于错误，返回接口为空即可 */

/* 获取学生名字为xxx的人数 */
func getCountForName(name string, db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT SUM(number) FROM student WHERE name = $1", name).Scan(&count)

	if err == sql.ErrNoRows {
		err = nil
	} else {
		err = errors.Warp(err, "getCountForName error")
	}
	return count, err
}
