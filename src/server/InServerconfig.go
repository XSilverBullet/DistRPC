package server

//读取配置文件中服务器的IP PORT 验证用户名Username
import(
	"os"
	"strings"
	"io"
	"bufio"
	"fmt"
)
var SERVERIP string
var SERVERPORT string
var USERNAME string
var PASSWORD string
func GetServerAddr(){
	f, err := os.Open("server/server.conf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, ":") == 0 {
			continue
		}
		vals := strings.Split(s, ":")
		fmt.Println(vals[0], vals[1])
		if(strings.Compare(vals[0] , "SERVERIP") == 0){
			SERVERIP = vals[1]
		}
		if(strings.Compare(vals[0] , "SERVERPORT") == 0){
			SERVERPORT = vals[1]
		}
		if(strings.Compare(vals[0], "USERNAME")==0){
			USERNAME = vals[1]
		}
	}
}

