package util

import (
	"net/http"
	"strings"
)

func TranslateErrorCategories(err error) (errCode int, errMessage string) {
	switch {
	case strings.Contains(err.Error(), "duplicate key value violates unique constraint"):
		return http.StatusBadRequest, "anda sudah pernah mendaftar %s dengan nik/ no hp yang sama, silahkan coba kembali dengan nik/ no hp yang berbeda"
	default:
		return http.StatusInternalServerError, "Can't insert to db"
	}
}