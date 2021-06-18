package share_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/cdvelop/margo/storage/share"
)

func Test_GetNewID(t *testing.T) {
	idRequired := 100
	wg := sync.WaitGroup{}
	wg.Add(idRequired)

	idObtained := make(map[string]int)
	var esperar sync.Mutex

	for i := 0; i < idRequired; i++ {
		go func() {
			defer wg.Done()
			id := share.GetNewID()
			esperar.Lock()
			if cantId, exist := idObtained[id]; exist {
				idObtained[id] = cantId + 1
			} else {
				idObtained[id] = 1
			}
			esperar.Unlock()

		}()
	}
	wg.Wait()

	fmt.Printf("total id requeridos: %v obtenidos: %v\n", idRequired, len(idObtained))
	fmt.Printf("%v", idObtained)
	if idRequired != len(idObtained) {
		t.Fail()
	}
}
