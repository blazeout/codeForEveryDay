package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 需要保存职位名称、职位类别、工作城市

type JobCateGory struct {
	Name string `json:"name"`
}

type CityList struct {
	Name string `json:"name"`
}

type JobInfo struct {
	Title       string      `json:"title"`
	JobCateGory JobCateGory `json:"job_category"`
	CityList    CityList    `json:"city_list"`
}

type Data struct {
	JobPostList []JobInfo `json:"job_post_list"`
}

type Response struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Err     string `json:"err"`
	Message string `json:"message"`
}

func main() {
	url := "https://xskydata.jobs.feishu.cn/api/v1/search/job/posts?keyword=&limit=36&offset=0&job_category_id_list=&location_code_list=&subject_id_list=&recruitment_id_list=&portal_type=6&job_function_id_list=&portal_entrance=1&_signature=T8-HMgAgEAZvzzmiHKdycU.PhyAAC1C"
	header := map[string]string{
		"authority":       "xskydata.jobs.feishu.cn",
		"accept":          "application/json, text/plain, */*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "zh-CN",
		"content-type":    "application/json",
		"cookie":          "_ga=GA1.2.477680562.1644486932; Hm_lvt_e78c0cb1b97ef970304b53d2097845fd=1644486934; passport_web_did=7064763723940036612; locale=zh-CN; trust_browser_id=fdb11a22-7ebd-48ec-b6bb-d4826e0ebb4c; fid=5deee259-aa7d-47b1-b7ec-c2201ba73acb; Hm_lpvt_e78c0cb1b97ef970304b53d2097845fd=1644983788; X-Risk-Browser-Id=f84f90b18e61a1d4254883d3fca926595f068fc8f242ce9b124593b254a12f31; admin-csrf-token=9ObFeSEeB6TSYwEwiMHsgCQUqu3f+KHrVNj91XZf626IPCUGT1kifWJHJNpbNzpgkMCl6REZ2LsLfsbxNKsYWZDvJLSpAOPjVJNuuIGJU4UIU5Q5i16vM/6HMtrr343iluweBg==; QXV0aHpDb250ZXh0=182fbe3c47ce4c5b96262daec2e7131b; lang=zh; __tea__ug__uid=6973980409471223308; session=XN0YXJ0-9e7m8440-3f00-4f18-8577-d306037ab0d8-WVuZA; session_list=XN0YXJ0-9e7m8440-3f00-4f18-8577-d306037ab0d8-WVuZA_XN0YXJ0-11bha14d-d78f-4c6a-88f8-16e9517bb46c-WVuZA_XN0YXJ0-473t70c6-e8cb-4900-bbd7-18493a504518-WVuZA; channel=saas-career; platform=pc; s_v_web_id=verify_l6kkxen9_jJwrnFtK_k3Rv_4UIb_AN6A_7d1WBIe5lC4l; device-id=7129442156544394755; atsx-csrf-token=QdwsBs-kBuh4KhQnBf1JKIDbRvl9pBQUwE9aiBc2lzw^%^3D",
		"origin":          "https://xskydata.jobs.feishu.cn",
		"referer":         "https://xskydata.jobs.feishu.cn/school",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		"x-csrf-token":    "QdwsBs-kBuh4KhQnBf1JKIDbRvl9pBQUwE9aiBc2lzw=",
	}
	download(url, header)
}

func download(url string, header map[string]string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http GET error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	// 转换成结构体
	responseFromHire := Response{}
	err = json.Unmarshal(body, &responseFromHire)
	if err != nil {
		fmt.Println(err, "转换错误")
	}
	marshal, _ := json.Marshal(responseFromHire.Data)
	if err = ioutil.WriteFile("./info.json", marshal, 0666); err != nil {
		fmt.Println("Writefile Error =", err)
		return
	}
}
