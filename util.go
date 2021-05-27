package xxl

import "fmt"

func idSilceToString(ids []uint) (rte string) {

	len := len(ids)

	for i, id := range ids {
		rte = rte + fmt.Sprint(id)
		if i != len-1 {
			rte = rte + ", "
		}
	}

	return rte
}
