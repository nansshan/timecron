package static

import (
	"embed"
	"fmt"
	"log"
	"testing"
)

//go:embed assets
var assets1 embed.FS

// 打开文件
func OpenFile1(filePath string) ([]byte, error) {
	// "assets/sql/task.sql"   文件名要以内嵌的文件夹名称开始
	entries, err := Assets.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		log.Println("打开内嵌文件失败" + filePath)
		return nil, err
	}
	return entries, nil
}

// 获取文件列表
func Test_FileList(t *testing.T) {
	// "assets/sql/task.sql"   文件名要以内嵌的文件夹名称开始
	dirPath := "assets/sql"
	entries, err := assets1.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		log.Println("打开内嵌文件失败" + dirPath)
	}
	var arr []map[string]string
	for _, entry := range entries {
		// fs.DirEntry的Info接口会返回fs.FileInfo，这东西被从os移动到了io/fs，接口本身没有变化
		// info, _ := entry.Info()
		item := make(map[string]string)
		item["name"] = entry.Name()
		item["path"] = dirPath + "/" + entry.Name()
		// fmt.Println("file name:", entry.Name(), "\tisDir:", entry.IsDir(), "\tsize:", info.Size())
		arr = append(arr, item)
	}
	fmt.Println(arr)
}
func FileList1(dirPath string) ([]map[string]string, error) {
	// "assets/sql/task.sql"   文件名要以内嵌的文件夹名称开始
	entries, err := Assets.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		log.Println("打开内嵌文件夹失败" + dirPath)
		return nil, err
	}
	var arr []map[string]string
	for _, entry := range entries {
		// fs.DirEntry的Info接口会返回fs.FileInfo，这东西被从os移动到了io/fs，接口本身没有变化
		item := make(map[string]string)
		item["name"] = entry.Name()
		item["path"] = dirPath + "/" + entry.Name()
		arr = append(arr, item)
	}
	return arr, nil
}

// 遍历读出的文件数组
func Test_for(t *testing.T) {
	fileArr, err := FileList1("assets/sql")
	if err != nil {
		return
	}
	for _, entry := range fileArr {
		fmt.Println(entry["name"])
	}
}
