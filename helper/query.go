package helper

import (
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func QueryFiltering(db *gorm.DB, query url.Values) *gorm.DB {
	for key, value := range query {
		if key != "sort" && key != "limit" && key != "page" {
			db = db.Where(key + " LIKE " + "'%" + value[0] + "%'")
		}
	}
	return db
}

func QuerySorting(db *gorm.DB, query url.Values) *gorm.DB {
	if !query.Has("sort") {
		return db
	}
	field, method := func() (string, string) {
		str := strings.Split(query.Get("sort"), ":")
		if len(str) == 0 {
			return "", ""
		}
		return str[0], str[1]
	}()
	if field != "" && method != "" {
		db = db.Order(field + " " + method)
	}
	return db
}

func QueryPagination(db *gorm.DB, query url.Values) *gorm.DB {
	var limit, offset int
	if query.Has("limit") {
		res, err := strconv.Atoi(query.Get("limit"))
		if err != nil {
			log.Error("helper:", err.Error())
		} else {
			db = db.Limit(res)
			limit = res
		}
	}
	if query.Has("page") {
		page, err := strconv.Atoi(query.Get("page"))
		if err != nil {
			log.Error("helper:", err.Error())
		} else {
			offset = (page - 1) * limit
			db.Offset(offset)
		}
	}
	return db
}
