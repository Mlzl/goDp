package chainOfResponsibilities

import "fmt"

func Test(){
	nextHandle := new(UpperFilterHandler)

	handle := new(MultiByteFilterHandle)
	handle.SetNextHandler(nextHandle)
	fmt.Println(handle.Handle("过H滤e中l文l留o下 英w文o并r且l转d为答!谢"))
}
