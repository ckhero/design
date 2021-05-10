/**
 *@Description
 *@ClassName template
 *@Date 2021/5/10 上午11:00
 *@Author ckhero
 */

package template_pattern

import "fmt"

type SkuInfo struct {
	Name string
	Count int
}

type Template interface {
	Login()
	Query() SkuInfo
	CreateDetail(sku SkuInfo) string
}

type BaseTemplate struct {

}

func(BaseTemplate) Exec(tmpl Template) string {
	// 登录
	tmpl.Login()
	// 抓取
	sku := tmpl.Query()
	// 生成详情
	return tmpl.CreateDetail(sku)
}


// 天猫

type Taobao struct {
}

func NewTaobao() Template {
	return &Taobao{}
}

func(*Taobao) Login() {}

func(*Taobao) Query() SkuInfo {
	return SkuInfo{
		Name:  "淘宝商品",
		Count: 100,
	}
}

func(*Taobao) CreateDetail(sku SkuInfo) string{

	return fmt.Sprintf("[taobao] - %s - %d", sku.Name, sku.Count)
}


// 天猫

type JD struct {
}

func NewJD() Template {
	return &JD{}
}

func(*JD) Login() {}

func(*JD) Query() SkuInfo {
	return SkuInfo{
		Name:  "京东商品",
		Count: 200,
	}
}

func(*JD) CreateDetail(sku SkuInfo) string{
	return fmt.Sprintf("[JD] - %s - %d", sku.Name, sku.Count)
}