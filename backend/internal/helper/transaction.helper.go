package helper

import (
	"errors"
)

func RowAffectedChecker(callFunc func() (int64, error)) error {
	rowsAffected, err := callFunc()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Data tidak ditemukan atau sudah diubah oleh pengguna lain")
	}

	return nil
}
