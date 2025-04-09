package info

import (
	client2 "Exchange__info/client"
	"Exchange__info/logger"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

var VersiontoCU = map[string]string{
	"15.2.1748": "Exchange Server 2019 CU15",
	"15.2.1544": "Exchange Server 2019 CU14",
	"15.2.1258": "Exchange Server 2019 CU13",
	"15.2.1118": "Exchange Server 2019 CU12",
	"15.2.986":  "Exchange Server 2019 CU11",
	"15.2.922":  "Exchange Server 2019 CU10",
	"15.2.858":  "Exchange Server 2019 CU9",
	"15.2.792":  "Exchange Server 2019 CU8",
	"15.2.721":  "Exchange Server 2019 CU7",
	"15.2.659":  "Exchange Server 2019 CU6",
	"15.2.595":  "Exchange Server 2019 CU5",
	"15.2.529":  "Exchange Server 2019 CU4",
	"15.2.464":  "Exchange Server 2019 CU3",
	"15.2.397":  "Exchange Server 2019 CU2",
	"15.2.330":  "Exchange Server 2019 CU1",
	"15.2.221":  "Exchange Server 2019 RTM",
	"15.2.196":  "Exchange Server 2019 Preview",
	"15.1.2507": "Exchange Server 2016 CU23",
	"15.1.2375": "Exchange Server 2016 CU22",
	"15.1.2308": "Exchange Server 2016 CU21",
	"15.1.2242": "Exchange Server 2016 CU20",
	"15.1.2176": "Exchange Server 2016 CU19",
	"15.1.2106": "Exchange Server 2016 CU18",
	"15.1.2044": "Exchange Server 2016 CU17",
	"15.1.1979": "Exchange Server 2016 CU16",
	"15.1.1913": "Exchange Server 2016 CU15",
	"15.1.1847": "Exchange Server 2016 CU14",
	"15.1.1779": "Exchange Server 2016 CU13",
	"15.1.1713": "Exchange Server 2016 CU12",
	"15.1.1591": "Exchange Server 2016 CU11",
	"15.1.1531": "Exchange Server 2016 CU10",
	"15.1.1466": "Exchange Server 2016 CU9",
	"15.1.1415": "Exchange Server 2016 CU8",
	"15.1.1261": "Exchange Server 2016 CU7",
	"15.1.1034": "Exchange Server 2016 CU6",
	"15.1.845":  "Exchange Server 2016 CU5",
	"15.1.669":  "Exchange Server 2016 CU4",
	"15.1.544":  "Exchange Server 2016 CU3",
	"15.1.466":  "Exchange Server 2016 CU2",
	"15.1.396":  "Exchange Server 2016 CU1",
	"15.1.225":  "Exchange Server 2016 RTM",
	"15.0.1497": "Exchange Server 2013 CU23",
	"15.0.1473": "Exchange Server 2013 CU22",
	"15.0.1395": "Exchange Server 2013 CU21",
	"15.0.1367": "Exchange Server 2013 CU20",
	"15.0.1365": "Exchange Server 2013 CU19",
	"15.0.1347": "Exchange Server 2013 CU18",
	"15.0.1320": "Exchange Server 2013 CU17",
	"15.0.1293": "Exchange Server 2013 CU16",
	"15.0.1263": "Exchange Server 2013 CU15",
	"15.0.1236": "Exchange Server 2013 CU14",
	"15.0.1210": "Exchange Server 2013 CU13",
	"15.0.1178": "Exchange Server 2013 CU12",
	"15.0.1156": "Exchange Server 2013 CU11",
	"15.0.1130": "Exchange Server 2013 CU10",
	"15.0.1104": "Exchange Server 2013 CU9",
	"15.0.1076": "Exchange Server 2013 CU8",
	"15.0.1044": "Exchange Server 2013 CU7",
	"15.0.995":  "Exchange Server 2013 CU6",
	"15.0.913":  "Exchange Server 2013 CU5",
	"15.0.847":  "Exchange Server 2013 SP1",
	"15.0.775":  "Exchange Server 2013 CU3",
	"15.0.712":  "Exchange Server 2013 CU2",
	"15.0.620":  "Exchange Server 2013 CU1",
	"15.0.516":  "Exchange Server 2013 RTM",
	"14.3.513":  "Exchange Server 2010 SP3-32",
	"14.3.509":  "Exchange Server 2010 SP3-31",
	"14.3.496":  "Exchange Server 2010 SP3-30",
	"14.3.468":  "Exchange Server 2010 SP3-29",
	"14.3.461":  "Exchange Server 2010 SP3-28",
	"14.3.452":  "Exchange Server 2010 SP3-27",
	"14.3.442":  "Exchange Server 2010 SP3-26",
	"14.3.435":  "Exchange Server 2010 SP3-25",
	"14.3.419":  "Exchange Server 2010 SP3-24",
	"14.3.417":  "Exchange Server 2010 SP3-23",
	"14.3.411":  "Exchange Server 2010 SP3-22",
	"14.3.399":  "Exchange Server 2010 SP3-21",
	"14.3.389":  "Exchange Server 2010 SP3-20",
	"14.3.382":  "Exchange Server 2010 SP3-19",
	"14.3.361":  "Exchange Server 2010 SP3-18",
	"14.3.352":  "Exchange Server 2010 SP3-17",
	"14.3.336":  "Exchange Server 2010 SP3-16",
	"14.3.319":  "Exchange Server 2010 SP3-15",
	"14.3.301":  "Exchange Server 2010 SP3-14",
	"14.3.294":  "Exchange Server 2010 SP3-13",
	"14.3.279":  "Exchange Server 2010 SP3-12",
	"14.3.266":  "Exchange Server 2010 SP3-11",
	"14.3.248":  "Exchange Server 2010 SP3-10",
	"14.3.235":  "Exchange Server 2010 SP3-9",
	"14.3.224":  "Exchange Server 2010 SP3-8",
	"14.3.210":  "Exchange Server 2010 SP3-7",
	"14.3.195":  "Exchange Server 2010 SP3-6",
	"14.3.181":  "Exchange Server 2010 SP3-5",
	"14.3.174":  "Exchange Server 2010 SP3-",
	"14.3.169":  "Exchange Server 2010 SP3-3",
	"14.3.158":  "Exchange Server 2010 SP3-2",
	"14.3.146":  "Exchange Server 2010 SP3-1",
	"14.3.123":  "Exchange Server 2010 SP3",
	"14.2.513":  "Exchange Server 2010 SP2-32",
	"14.2.509":  "Exchange Server 2010 SP2-31",
	"14.2.496":  "Exchange Server 2010 SP2-30",
	"14.2.468":  "Exchange Server 2010 SP2-29",
	"14.2.461":  "Exchange Server 2010 SP2-28",
	"14.2.452":  "Exchange Server 2010 SP2-27",
	"14.2.390":  "Exchange Server 2010 SP2-8",
	"14.2.375":  "Exchange Server 2010 SP2-7",
	"14.2.342":  "Exchange Server 2010 SP2-6",
	"14.2.328":  "Exchange Server 2010 SP2-5",
	"14.2.318":  "Exchange Server 2010 SP2-",
	"14.2.309":  "Exchange Server 2010 SP2-3",
	"14.2.298":  "Exchange Server 2010 SP2-2",
	"14.2.283":  "Exchange Server 2010 SP2-1",
	"14.2.247":  "Exchange Server 2010 SP2",
	"14.1.513":  "Exchange Server 2010 SP1-32",
	"14.1.509":  "Exchange Server 2010 SP1-31",
	"14.1.496":  "Exchange Server 2010 SP1-30",
	"14.1.468":  "Exchange Server 2010 SP1-29",
	"14.1.461":  "Exchange Server 2010 SP1-28",
	"14.1.452":  "Exchange Server 2010 SP1-27",
	"14.1.438":  "Exchange Server 2010 SP1-8",
	"14.1.421":  "Exchange Server 2010 SP1-7",
	"14.1.355":  "Exchange Server 2010 SP1-6",
	"14.1.339":  "Exchange Server 2010 SP1-5",
	"14.1.323":  "Exchange Server 2010 SP1-",
	"14.1.289":  "Exchange Server 2010 SP1-3",
	"14.1.270":  "Exchange Server 2010 SP1-2",
	"14.1.255":  "Exchange Server 2010 SP1-1",
	"14.1.218":  "Exchange Server 2010 SP1",
	"14.0.726":  "Exchange Server 2010-5",
	"14.0.702":  "Exchange Server 2010-4",
	"14.0.694":  "Exchange Server 2010-3",
	"14.0.689":  "Exchange Server 2010-2",
	"14.0.682":  "Exchange Server 2010-1",
	"14.0.639":  "Exchange Server 2010 RTM",
}

func Get_Version(URL string, proxyAddr string) error {
	if proxyAddr == "" {
		Client := client2.Global()

		request := Client.R()
		URL = "https://" + URL + "/ecp"

		resp, err := request.Get(URL)
		if err != nil {
			return logger.Log.ErrorMsaf("网络请求失败%s", err)
		}
		reader, err2 := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return logger.Log.ErrorMsaf("HTML解析失败%s", err2)
		}
		//提取标签内容
		text, exists := reader.Find("link").First().Attr("href")
		if exists {
			parts := strings.Split(text, "/")
			logger.Log.InfoMsaf("Exchange Server版本为%s----->%s", parts[3], VersiontoCU[parts[3]])
		} else {
			return logger.Log.ErrorMsaf("标签不存在")
		}
		return nil

	} else {
		Client := client2.Global().SetProxyURL("http://" + proxyAddr)

		request := Client.R()
		URL = "https://" + URL + "/ecp"

		resp, err := request.Get(URL)
		if err != nil {
			return logger.Log.ErrorMsaf("网络请求失败%s", err)
		}
		reader, err2 := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return logger.Log.ErrorMsaf("HTML解析失败%s", err2)
		}
		//提取标签内容
		text, exists := reader.Find("link").First().Attr("href")
		if exists {
			parts := strings.Split(text, "/")
			logger.Log.InfoMsaf("Exchange Server版本为%s----->%s", parts[3], VersiontoCU[parts[3]])
		} else {
			return logger.Log.ErrorMsaf("标签不存在")
		}
		return nil

	}
}
