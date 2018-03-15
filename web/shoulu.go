package main

//golang 版本 百度,360,sogou,bing,google  网站收录量批量 查询
import (
	//"Public_file"
	"fmt"
	"regexp"
	//"strings"
	//"http_get_post"
	"net/url"
	"strings"
	"time"

	"github.com/levigross/grequests"
)

//===============================
func Get_url(url string) (bool, string) {
	//defer Public_file.Panic_Err()                                  //异常处理
	ro := &grequests.RequestOptions{DialTimeout: 10 * time.Second} //超时设置
	resp, err := grequests.Get(url, ro)
	if err != nil {
		return false, ""
	} else {
		return true, resp.String()
	}
	return false, ""
}

func url_data_re(urlx string) string {
	//defer Public_file.Panic_Err() //异常处理
	bool_err, resp := Get_url(urlx)
	if bool_err != true {
		return "?"
	}
	//fmt.Printf("urlx:%v\n", urlx)
	//=========================
	reg := `该网站共有.*?<b style=\"color:#333\">(.*?)</b>\s+.*?个网页被百度收录`
	sarr := regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("百度11:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}

	reg = `百度为您找到相关结果约(.*?)个` //找出一条百度
	sarr = regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("百度22:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}

	reg = `百度为您找到相关结果(.*?)个` //找出一条百度
	sarr = regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("百度33:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}
	//=========================
	//=========================
	reg = `找到约.*?<span id="scd_num">(.*?)</span>` //找出一条sogou
	sarr = regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("sogou11:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}

	reg = `找到约.*?<em>(.*?)</em>` //找出一条sogou
	sarr = regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("sogou22:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}

	reg = `找到约.*?<resnum id="scd_num">(.*?)</resnum>` //找出一条sogou
	sarr = regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("sogou33:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}

	reg = `搜狗已为您找到约(.*?)条相关结果` //找出一条sogou
	sarr = regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("sogou33:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}

	//=========================
	reg = `<span class="sb_count">(.*?)条结果</span>` //找出一条bing
	sarr = regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("bing11:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}
	//==============================
	reg = `找到相关结果约(.*?)个` //找出一条so360
	sarr = regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("so36011:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}
	//==============================
	reg = `About(.*?)results` //找出一条google.com      About 1,590 results       找到约(.*?)条结果
	sarr = regexp.MustCompile(reg).FindAllSubmatch([]byte(resp), -1)
	if len(sarr) >= 1 {
		//fmt.Printf("so36011:%v\n", string(sarr[0][1]))
		return string(sarr[0][1])
	}
	//==============================
	//fmt.Printf("======%v\n", resp)
	return "!"
}

//===========
func del_tab(value string) string { //去除空格
	if strings.Contains(value, " ") { //去除首尾空格
		value = strings.Trim(value, " ")            //头部
		value = strings.TrimSpace(value)            //尾部
		value = strings.Replace(value, " ", "", -1) //替换
		//fmt.Printf("yyy:%v\n", href)
	}
	value = strings.Replace(value, "，", "", -1) //替换
	value = strings.Replace(value, ",", "", -1) //替换
	if value == "?" || value == "!" {
		return "0"
	}
	return value
}

func baidu_sl(value string) string { //百度收录数量
	//defer Public_file.Panic_Err() //异常处理
	urlx := fmt.Sprintf("http://www.baidu.com/s?wd=site:%s", url.QueryEscape(value))
	return del_tab(url_data_re(urlx))
}

func s360_sl(value string) string { //360收录数量
	urlx := fmt.Sprintf("http://www.haosou.com/s?q=site:%s", url.QueryEscape(value))
	return del_tab(url_data_re(urlx))
}

func sogou_sl(value string) string { //sogou收录数量  web?query=
	urlx := fmt.Sprintf("http://www.sogou.com/sogou?query=site:%s", url.QueryEscape(value))
	return del_tab(url_data_re(urlx))
}

func bing_sl(value string) string { //bing收录数量
	//urlx := fmt.Sprintf("http://www.bing.com/search?q=site:%s&rdrig=B12E61D97E944B7B9740B0392D6197E9", url.QueryEscape(value))
	urlx := fmt.Sprintf("https://cn.bing.com/search?q=site:%s", url.QueryEscape(value))
	return del_tab(url_data_re(urlx))
}

func google_com_sl(value string) string { //google.com收录数量
	///       About 1,590 results
	urlx := fmt.Sprintf("https://www.google.com/search?source=hp&q=site:%s&oq=site:%s", url.QueryEscape(value), url.QueryEscape(value))
	return del_tab(url_data_re(urlx))
}

//===============================
func main() {
	fmt.Printf("=baidu===%v====\n", baidu_sl("www.iphpt.com"))           //百度收录数量
	fmt.Printf("=s360===%v====\n", s360_sl("www.iphpt.com"))             //360收录数量
	fmt.Printf("=sogou===%v====\n", sogou_sl("www.iphpt.com"))           //sogou收录数量
	fmt.Printf("=bing===%v====\n", bing_sl("iphpt.com"))                 //bing收录数量
	fmt.Printf("=google_com===%v====\n", google_com_sl("www.iphpt.com")) //google.com收录数量
}
