package combine_movie

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Main() {
	var path string
	fmt.Print("请输入m3u8文件所在文件夹的路径：")
	_, err := fmt.Scan(&path)
	if err != nil {
		return
	}
	path = strings.TrimRight(path, "\\/")
	fmt.Println(path)
	content, err := ioutil.ReadFile(filepath.Join(path, "local.m3u8"))
	if err != nil {
		fmt.Println("读取m3u8文件失败：", err)
		return
	}
	reg := regexp.MustCompile("#EXTINF:\\d*\\.?\\d*?,\\n(.*?)\\n")
	result := reg.FindAllSubmatch(content, -1)
	if len(result) == 0 {
		fmt.Println("m3u8文件读取切片数为0")
		return
	}
	fp, err := os.OpenFile(filepath.Join(path, "movie.mp4"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("创建合并视频文件失败：", err)
		return
	}
	bufFp := bufio.NewWriter(fp)
	for _, v := range result {
		sliceContent, err := ioutil.ReadFile(filepath.Join(path, string(v[1])))
		if err != nil {
			fmt.Printf("读取切片%s：%s", v[1], err)
			return
		}
		_, err = bufFp.Write(sliceContent)
		if err != nil {
			return
		}
	}
	err = bufFp.Flush()
	if err != nil {
		fmt.Printf("冲刷缓存失败：%s", err)
		return
	}
	fp.Close()
}
