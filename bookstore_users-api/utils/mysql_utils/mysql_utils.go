package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/anabeto93/bookstore/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(field string, value string, err error, defaultMsg string) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewNotFoundError("no matching record found")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error passing database response: %s", err.Error()))
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("%s %s already exists", field, value))
	}
	if strings.TrimSpace(defaultMsg) != "" {
		return errors.NewInternalServerError(defaultMsg)
	}
	return errors.NewInternalServerError("error processing database request")
}