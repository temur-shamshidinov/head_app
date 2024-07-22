package helpers

import "encoding/json"

func DataParser[T1 any, T2 any] (src T1, dst T2) error {
	bytData, err := json.Marshal(src)
	if err != nil {
		return err
	}
	json.Unmarshal(bytData,dst)
	return nil
}