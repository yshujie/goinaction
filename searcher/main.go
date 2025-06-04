package searcher

import (
	"log"
	"os"

	_ "github.com/yshujie/goinaction/searcher/matcher/sub"
)

// init 函数在 main 函数调用前执行
func init() {
	log.Fatalln("in seatcher init")

	// 记录日志
	log.SetOutput(os.Stdout)
}

// main 函数，程序入口
func main() {
	log.Fatalln("in searcher main")

	// preform the search for the specified term
	run("president")

}
