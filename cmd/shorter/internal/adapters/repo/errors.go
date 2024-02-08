package repo

//
//import (
//	"errors"
//	"fmt"
//	"github.com/lib/pq"
//	"strings"
//)
//
//const duplicateURL = "url_table_url_key"
//
//func errconv(err error) error {
//	var myerr *pq.Error
//	errors.As(err, &myerr)
//
//	if strings.HasSuffix(myerr.Message, fmt.Sprintf("unique constraint \"%s\"", duplicateURL)) {
//		return nil
//	}
//	return err
//}
