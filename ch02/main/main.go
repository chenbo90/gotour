package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gotour/pkg/cmd"
)

const (
	one = iota + 1
	two
	three
	four
)

var(

	rootCmd = &cobra.Command{
		Use:          "pilot-discovery",
		Short:        "Istio Pilot.",
		Long:         "Istio Pilot provides fleet-wide traffic management capabilities in the Istio Service Mesh.",
		SilenceUsage: true,
	}
	discoveryCmd = &cobra.Command{
		Use:   "discovery",
		Short: "Start Istio proxy discovery service.",
		Args:  cobra.ExactArgs(0),
		RunE: func(c *cobra.Command, args []string) error {
			fmt.Println("haha")
			return nil
		},
	}


)

func main() {
	//fmt.Println("---this is ch02 test--")
	//
	//testIota()
	//
	//testTypeChange()
	//
	//testStr()
	//
	//f1.Util()

	fmt.Println(strings.Contains("widuu", "wi")) //true
	fmt.Println(strings.Contains("wi", "widuu")) //false

}

/**
*	测试Iota
 */
func testIota() {
	//测试iota,默认值是0
	fmt.Println(one, two, three, four)
}

/**
*	测试字符类型转换
 */
func testTypeChange() {
	i := 100
	//整型转字符串
	i2s := strconv.Itoa(i)
	//字符串转整形
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)
}

func testStr() {
	s1 := "Hello World,yellow"
	//判断s1的前缀是否是H
	fmt.Println(strings.HasPrefix(s1, "H"))
	//判断s1是否包含o
	fmt.Println(strings.Contains(s1, "o"))
	//把s1全部转换成大写
	fmt.Println(strings.ToUpper(s1))
	//字符串ell首次出现的位置
	i := strings.Index(s1, "ell")
	//字符串ell最后一次出现的位置
	j := strings.LastIndex(s1, "ell")
	fmt.Println(i, j)

	if err := rootCmd.Execute(); err != nil {
		//log.Errora(err)
		os.Exit(-1)
	}
}

func init() {
	fmt.Println("--我是main包内的init1--")
	cmd.AddFlags(rootCmd)
	rootCmd.AddCommand(discoveryCmd)


}

func init() {
	fmt.Println("--我是main包内的init2--")
}
