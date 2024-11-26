package main

import (
	"github.com/wmm1996528/cronet-go"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

// URL wraps the net/url.URL type to enable unmarshalling from text
type URL struct {
	url.URL
}

func (u *URL) UnmarshalText(text []byte) error {
	parsed, err := url.Parse(string(text))
	if err != nil {
		return err
	}
	u.URL = *parsed
	return nil
}

func (u *URL) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

//func ConfigureClientCertificate(e *cronet.Engine, certPath string, keyPath string, hostPort []string) {
//	if certPath == "" || keyPath == "" {
//		return
//	}
//	clientCertData, err := os.ReadFile(certPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//	privateKeyData, err := os.ReadFile(keyPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, hostPortPair := range hostPort {
//		e.SetClientCertificate(hostPortPair, clientCertData, privateKeyData)
//	}
//}

func main() {

	engineParams := cronet.NewEngineParams()
	//engineParams.SetProxyServer("http://127.0.0.1:7890")
	engineParams.SetProxyServer("http://unfflcc:IJRBASA-0FHO38C-WZK3EDY-XXDYPAR-REXYTIU-LU3L7ZS-IRXU6RE@usa.rotating.proxyrack.net:9000")
	//engineParams.SetProxyServer("http://user-uni003-region-de-sessid-1125-sesstime-5-keep-true:q39CEBTs5A5YQXor@pr.roxlabs.cn:4600")
	//engineParams.SetProxyServer("http://user-uni003-region-de-sessid-1325-sesstime-5-keep-true:q39CEBTs5A5YQXor@pr.roxlabs.cn:460")
	engineParams.SetEnableHTTP2(false)
	engineParams.SetEnableQuic(false)
	engineParams.SetEnableBrotli(true)
	engineParams.SetEnablePublicKeyPinningBypassForLocalTrustAnchors(true)
	t := cronet.NewCronetTransport(engineParams, true)
	t.Engine.StartNetLogToFile("./1.log", true)
	//ConfigureClientCertificate(&t.Engine, certPath, keyPath, []string{urlArg.Host, proxyArg.Host})
	jar, _ := cookiejar.New(nil)
	//u, err := url.Parse("https://be.wizzair.com")
	//if err != nil {
	//	panic(err)
	//}
	//var cks []*http.Cookie
	////cookieStr := "bm_s=YAAQm/AQAjZdwKWOAQAAB+kDzAFIxULKMNEpBltlqFhoo+H54BiE4kGcryVC2wTXgzUPlSOuBKV816fV4cdhaGA5V3mSU/C5WbPJl0YM/mDyefsMphUUnbxcVkoNCWx+/NGvHnsqKHYbKnJ55Itfx/nmJ6XUcjpXURmHvJFWZiA9br4tFROIUPCQnxRxQrCqhr7IquOa+WwdWaGRdTWEu+ZQD+bc5chkGkOFcMuhEMF0/TcHQC1fL6pUtPsx7ag6Xo2pZlmPOej3YQLdu8J8VLsP4DVuAB4tKjIKt4WKivgxK2YAdNLnpXNXDP5XJE7ztN8ZsOlyNG6idRGO2SmxbTSrpwQR; bm_ss=ab8e18ef4e; akacd_dc1=3890272591~rv=76~id=991b5021cfc7d97f0262771b7d4b8e1f; showbooking=true; bm_sz=0B28FD6630EB9FCCF9C9B22EEF329916~YAAQm/AQAqFewKWOAQAA7O0DzBfV+drlEgeaYiBTvxOWw2ryZ+ow6nY7tzwRNucAkb2z0o+b1j/iea8JBQqt8atqJfp837MWYCCcV/lqiBoti4GdChT6qagGRJxjJ9+R6gvB/CyM7Q6NLpo+M8nalHAkJnz0fhap7Bvnwg3BhqGIenk3Q5vde5/BPXEmt3thkD1ukWPAF7he/pJskNSLId3/5hN++S0scr9etQAQ867ev1QJwecVdDHr/3LL9l5nniskjqqQc5LNr5oqx8nxIAZTp/D2uUrgbvxpXVgALn0JKIB2HYIbGaGQ8fRPmMeUQGfeUQR1l5iJQ6/IbOfhsRk/SpQPN0dbHzCn1j/en28MLY3dj8yPyUWHa/w8vD5uGHGJ83V9hY8zhWnL1WFKZNkyB3Jl~4276802~4473138; _abck=82CE19569DCCF76A01E55836EED6DABA~-1~YAAQm/AQAupmwKWOAQAA1wYEzAsIwM2Mcn5BDnPHlPCNHWXXlqHlsHnPpKkiT7ITHIhgHxx6sdmkD5LvvRCsKHjgSkXo7cXZQasjBSAyxjM1BuHfQDmWGm8fjl2an/STIp3tE8O2z6NJMdSGp/uSpI3hVQQLB2mf5V+svqr5zcAjxpum+EK2PEJh588NEUJmu1RwxK3r4QtfxLfP0Qhj4o5aRp3BS8ud1f/YJoqWxsa4WwplmNh8m1xULzOCMeO8b+J61XkYdMSN8FdOZYbr66blrk6QWOr+Z82a35Cw0C2V9aZcGsplq7rzhO7xLA0YcHCxQNqyLHHkLh3W3BfOyVgUBbzRJ1mE+SlLflW4/sESW7lzxPaKUAdtompabFHDtVPZay92rYKbAvVD9IoixHq4ACTR2VaI~-1~-1~-1"
	//ccc := map[string]string{
	//	"ak_bmsc":                             "D752AAC487F474E3C92CF0F55C041988~000000000000000000000000000000~YAAQCfAgF18XqM2OAQAAr71v0Rcj9Ju7TQM70c2s5U8RquHvEF9qH3GkgriTcxlZ0l7zLLbMJmjqiGTNX7TCcFDQt8xs40+NwGj00FsbM+D2Q2YgxM1WRx/NOSH5UttZqZr+y/YiLY+ck6RpsWu8HHTQojbj3rm4gsZTGIsl5zjkEGjqq7s4CkiCTNtTowyVJTBQTGi9cQSY3zlhDMAPDSyMamCOYMRUzdr3AjGfp1mpKsvFnNf317II6OPpinMt5X/2AXwsyq5+Txix9jIV9Xyus9Xe22ohA3wzb1s+CyjtyHKEfdM4SkV7/boKmqOGMEeaMmbbTrMlF/dMLvFwNQeVLrsVGqipprGE4pUasRzFlOVgBya+UYFwQqTgpFlupvu6SnP3O68lrdEqoB838VRtW5XtN2UB+biOHBAil7cbBJ+SGLcYzhrAJqNcIEE3Ult3ChnBDQM/7DSNTQA=",
	//	"akacd_onewizz_AB_backend_production": "3890363545~rv=1~id=73aaf6f401fe529e36cd8c00e6b8437e",
	//	"_fbp":                                "fb.1.1712910748436.2007491519",
	//	"_ga":                                 "GA1.2.1355254444.1712910749",
	//	"_gid":                                "GA1.2.1635686520.1712910749",
	//	"ASP.NET_SessionId":                   "pjdv4hw2wze5hnz2ncbn3i1d",
	//	"RequestVerificationToken":            "4fa3c21fb01042ebbf6936306e5f1f57",
	//	"_tt_enable_cookie":                   "1",
	//	"_ttp":                                "Q_ijf852-_jmbvIQG8IOzTbNQd9",
	//	"_pin_unauth":                         "dWlkPU56aGtZMlZsWkRrdE1USmtOQzAwT1dNM0xUaGxPR010TmpZeE9EZzJZMlUyTlRWag",
	//	"KP_UIDz-ssn":                         "0G1a582EDr8oLs3zcp9jDhEBSZcGHPkr8zoUn5U6I4lqAUYZhECy4ptsQPzBW5nslCmgF5FCtdiGAPYJ4j2RByHKOWFk74mlIK65uA1obEGqmmWDfNbmxYL90YOT6VIRfemfTBQlzNkWMIivTq54ChvmYRGe6L95fQi7gJVq",
	//	"KP_UIDz":                             "0G1a582EDr8oLs3zcp9jDhEBSZcGHPkr8zoUn5U6I4lqAUYZhECy4ptsQPzBW5nslCmgF5FCtdiGAPYJ4j2RByHKOWFk74mlIK65uA1obEGqmmWDfNbmxYL90YOT6VIRfemfTBQlzNkWMIivTq54ChvmYRGe6L95fQi7gJVq",
	//	"bm_sz":                               "60BAEF4B5774B3817506772B67E0D2B0~YAAQLfAgF/M3rKaOAQAAkiWH0Rccm683bkJ5vAO+Lp3i2mpyxaJ81DwD8AnSjHs9vtFo0NkfECPJE5Cnt8yq+xjNPkP80jGhSbLAsmDjucVQg8RnQUyEIZUY1BtCNul1UPx9yzzmOm/UA4zNnSLmD7DwDQxpr3+zkaxlM9eKLGUJHFZ7ifJpzieM8239cLAaBnNw8UK18fdduu6cHLmfZ8ZpqvY4/zV+AnoSnquJyb9VK7+iBO9+L74YeT+hEb1CT/le86WNMaM+5zRVE+l7DHFjbyBluQtVtssTKI0g/IbrEf9zKnVOLXCmApBKyVhg6J5UfNpOh2Tn7hD6fgSj0JUYBIefnsYRwC3J7SNVjpASPdd5U2TszOp7p4ZjEpQTYn2DjLxur87tnAGXqAC4B2eO5QRhkECYP/BXuhtK0I4S58TnNusBVm1juxBPt+xnAA==~4342580~3162933",
	//	"bm_sv":                               "D8E0B251186AAFFDC0070BA172206CFC~YAAQLfAgF0Q6rKaOAQAAk0WH0Rcbsk1Bgby78JC3z9U89RaM6G52lYBm7ssTP48WHut0lf/ffplOdJ4NowOBqF/zMdaff7YDd5b6qOou2kAuK96uas3kBHwBr4ssSkowWjLIBPQrMBACNit5pGWmkRstc+ON3zPcHAyGLCe1xbFP3X5ImTRjRLY1GlmLqUYRh0RYuYQ0lL7bohGhafpR4KlkSoxKNLmvvwliSxlyCAPahsfZNb7EcxcVz6oGgmiTLWU=~1",
	//	"_abck":                               "AA08EC473771CADDD0E682CBAB37F814~0~YAAQLfAgF4NKrKaOAQAACymI0QvWxYYO59FW4Srd96L+kF8V6qNM8NxZzNAjGY6+GcqQMBA7hMFHH3Mj0MWsgSwwEX34L5bek9xT7QnOh1izqtE0lNKB9RWBSOm3HcUkPMP8REK84RvJd//ax3NT2WufscsRBtoRwTNruSxO7rvwZnVxdlIaVM0L9zQ++RDsM8pfiEPkAMLbnSeVipp8KIHWmwCrw8HJahkjS9Qj4Sr59MnCEnmGTgMB8j6nqXROCSr4G8rgo2wDCT4kmpLMUmzuO82Po2xcq7arDd3Fk0gDL/lKA3q3YJ7n4AzWtrdDAnmMf6WT2txqlYi0G4IlhrxiXEr5pbDYtZMwWo419i9Q3ZIYyMoXfxLOqli0OjwVBRFqXSO7zCSQxK2yWQUx/X7n+10XFMsywib7MIi6DPxxAhlqrbLKcFo2irkSyQ==~-1~-1~-1",
	//	"sec_cpt":                             "CD203D14553DE4208A522222A944429D~3~YAAQLfAgF4RKrKaOAQAACymI0QoYoeslTAd6vfTEcjPMVg1NCkEc/AdOV7IEFE9h6UOhrcd9QzVETv3gmTTiSFjCipOssvciRJM6sv+TeDvSGuYfJZwvTKmApDZQLlFmlC7Xk788C7v60fZ1yOnjhjZ05BIIf/EWRXNgFwLn4VZ42EUo/TsijCebqjQOFmtaqZGHwm3fZagbOwG+x9EwPRlGsYtjjMkj5M2zFQdBXkwlcAbTAlFelvEcX98iSKBsUEFSBGSRXz9UVP/mxA7eqYtXNTCiyajC/TLfV5EMOfZzcNQKXFeK1Z1IL3oVaALgQzzDaxON6OEcMR8ii6byDFoHXuveUdKRXSR8tdNRNXpNqcVvXUcR2qfiDBqoMmBYOf6rcuifNLVWS4O17sVTGYvAxjFKOrBom3U/3JFDpN7AdAVwgBLjfLcqG6cbhyM25d058SMEofJxR+NQ5/NMb4OwiE3XdFTPRzsMF7RdKVDN2aTfKlxl7VJ448cA2sHL+/aIXJGIrtbW4tZllsrprv0Q5AT3OzmbQx02yS/9FPzk44CgUuYK2dQHlJlF8i6179RoRxxRfEBU7vQrbG57WOd3jRnPvcl0A0EhU+YZBI0laqD/E8Q2rf735cHgLShod5n6TtFvvvtr/EG8wdM6SZx5CMMo+f6aYePkp8/k8kq4t1SshfLmFpJMEy5Po5OIEpddHMnBorfukaOjdjj5qhW6cB0JaMR9PO43meM8TrrluPRDZvBcANYcyrjVjE56GS/iQfQgVg8cf1DQZWsTakc1tDc8hoWbtuKnQ8K1ZzfPwAtrPhPcytlqYosZMVTFvLz2HVtBBM69viGTX99wzPLmF4A=",
	//}
	//for k, v := range ccc {
	//	//t := strings.Split(v, "=")
	//	cks = append(cks, &http.Cookie{
	//		Name:  k,
	//		Value: v,
	//	})
	//}
	//fmt.Println(u, cks)
	//jar.SetCookies(u, cks)
	client := &http.Client{
		Transport: t,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: jar,
	}
	//data := "{\"isFlightChange\":false,\"flightList\":[{\"departureStation\":\"BOJ\",\"arrivalStation\":\"VIE\",\"departureDate\":\"2024-07-21\"}],\"adultCount\":1,\"childCount\":0,\"infantCount\":0,\"wdc\":true}"
	//req, _ := http.NewRequest("GET", "https://m.vueling.com", nil)
	data := "{\"DeviceType\":\"WEB\",\"UserAgent\":\"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Mobile Safari/537.36\",\"Udid\":\"2355-6b05-cc33-6e6f-75f5-5682\",\"IP\":\"235.231.34.152\",\"AppVersion\":\"17.32.0\",\"Domain\":\"https://m.vueling.com\",\"DiscountType\":0,\"Paxs\":[{\"PaxType\":\"ADT\",\"Quantity\":\"1\"},{\"PaxType\":\"CHD\",\"Quantity\":\"0\"},{\"PaxType\":\"INF\",\"Quantity\":\"0\"}],\"CurrencyCode\":\"EUR\",\"AirportDateTimeList\":[{\"ArrivalStation\":\"BCN\",\"DepartureStation\":\"LCG\",\"MarketDateDeparture\":\"2024-04-28\"}],\"Language\":\"en-GB\"}"
	//req, _ := http.NewRequest("POST", "https://apimobile.vueling.com/Vueling.Mobile.AvailabilityService.WebAPI/api/V2/AvailabilityController/DoAirPriceSB", strings.NewReader(data))
	req, _ := http.NewRequest("GET", "https://httpbin.org/ip", strings.NewReader(data))
	//req, _ := http.NewRequest("GET", "https://www.jetstar.com/au/en/booking/search-flights?s=true&adults=1&children=0&infants=0&selectedclass1=economy&currency=AUD&mon=true&channel=DESKTOP&origin1=ADL&destination1=BNE&departuredate1=2024-03-13", nil)
	headers := map[string]string{
		"Accept":             "application/json, text/plain, */*",
		"Accept-Language":    "en",
		"Cache-Control":      "no-cache",
		"Connection":         "keep-alive",
		"Content-Type":       "application/json",
		"Origin":             "https://m.vueling.com",
		"Pragma":             "no-cache",
		"Referer":            "https://m.vueling.com/",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-site",
		"User-Agent":         "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Mobile Safari/537.36",
		"sec-ch-ua":          "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"",
		"sec-ch-ua-mobile":   "?1",
		"sec-ch-ua-platform": "\"Android\"",
		"Cookie":             "AKA_A2=A; _abck=7DC49CCB10C0F987AD5287FB8B6CCF97~-1~YAAQhpcwFwX40iWPAQAA0yWjTAuZYNpy+C5yf44Om2GvG15ifV8QnEXrvOE3hHYOHYU1wbuaXE07/80BH2QrVhscipeLawb4VBWjlKOY7u9pi+gZzBp1q25ZjgvGPD3WKFT73y/EOOJ3/Run/X+jNvLi7bSC/fHWIGItP90nmyg8vEsa33V3BSYc21kY2cXiHMs8oAJ6y4w8Ob3szmJfAI7uluVmPb1xdPxSIZwRzsyT8Ztph1UtxABkABR/Vm8nq3Rrza5mGdzb3f/rOMJz3rd/B5Be+b0RZc7JBbv2OiTJRCCo6wq0BKnWK3uIgO09w/98XlgNVJJvLAvZIM4vj/7pdvZsVWj+6lTTEjN6wTNeIv0n8bXFIbGnGf9pY8J4rilj1ZnByS1uv9FZi8uf64VA5YF0oyCB~0~-1~-1; ak_bmsc=BAFA803394F9A21027995AB4EF2811DC~000000000000000000000000000000~YAAQhpcwFwb40iWPAQAA0yWjTBcHdU1xniHmmVnT/R/aIty9ne4GjAaUbFJlW2XbYcjjBWRsxC4H+iL7tN+dSumqg3EDapbhw/xzNfn+DJdPEXjbjk+MCq8m0xwDZ2S3Xb4Do9rgbfimD1ylmeH3WuFP3bxdYWowTusDR7yzxMdUrbspXbqYZZ7UR+HinSSZi0JJNomu54NxW5NnOpm/5VCVtpNpFjyb+4ZIHzlI4hzdjbfs39ZPxtYPHcuhfA3PiX+YZJYZVb03Mp+Gi/kvA719/ektFTnwi8NxO2OYD+2xlqwLclhjTiplwUV/BuOg6/i29n0ZhncOBAr3MC/IgZFOGdQoi/SWARYqcGWOXXwkGrBk58hTL2423a6Iu0qsL+Ue6ECX5lRi2GI+W4Gb6A==; bm_s=YAAQdcMTAqD4hOqOAQAAnB2jTAG+TI/CzgSv3helQ8m8Ulw8pQ6NlMpRMBb/t6XeGEbqzZdfnFutDbzy3NM8jqfRwandpY7HjoJXDjjAd06ivvGnfAEEpOtc5GUdwRE6dj4JP0SwHksk3YScrI0+1E6YHLEMXjdu61z+h1p72p/eYwAozkNRN801/n7oIACSFWQvgBO2vwQLoNTs+P9EpTtN6YCC7FsIZUni8a2tpp5SHyyyX9rHh6bplMzz7+ediB21vYssO9DlxJ21MnqZs7Vo+Oa8E4p/AfahxDEi7WlcVo7eW3+Gy4x9Zf9CnL52CBiJ7YaSh2xPClHRNt7ICW98rVSGFF4PVHeE9V0W2rA=; bm_ss=ab8e18ef4e; bm_sv=0CE92A4C0475517D455E6305FFB7E4C9~YAAQhpcwFwf40iWPAQAA0yWjTBdsoW2C4DbMEEDfHOy7yf9I+4qkVGF9U9nCjDRoZO4JZ649fjuGUyPMtZa2TPTLakFAbtRl/jIqMhWZ3sfPQ0jpuSiyHWzQVH7zd3MTMokhDuqhbs5R3kW4N4TqTj8tlDApc2qajIJJN4nerJqIUP30hiu8E73Fs2/B3SGZMeGaklhTXSXhnr1DbmKkP0CHSS9gOMwG8RxeB/xRkFk5Bi2NLj4dDv9cGDtQyxrufw==~1; bm_sz=7E5F623E20D591A20973DD624C5463BB~YAAQdcMTArn4hOqOAQAALh+jTBd5Y2O2OSTGoidyBknWzVWZDkX6IABE8Qh2YXe3urJdm1kbTVaRMbViQ9hAeawbfVzswXFpNErY7wJ/pAOo2QZYse616xX4xiyLofPwbSSTDuWhHVTb+ot5oZ2fidbtnl4FFZjTzRqQf0eieGwFcMG3y3ugsrYNhdjtjxYD2h3TGoaHsi9ii3F+9GW6r4hSYRqtVAafBtn8W0HF4aelO2oksl2uFEKt1vT3xTTc/4CeLFBdQr78qPv3vGUuBqhuIWadWjupPBDTTaxNZub0vkwaVBpX0d+FMEqIc+Hnn6bcHrvm6BmZxhwsPRa98fES2qKf6BAjX+3oerXMgYBliHWZce+qlFe7kxjTV9k15QkAwe745TJr4rUnoKNBXXeJtQ==~3683907~3224641; sec_cpt=4FE66FEE3C575CD1D1822E6A0B5BE420~3~YAAQhpcwFxz80iWPAQAAraGjTAto2fbPfFqEeXsbHGVA0sAjh6xQuYpEcFE2OV7gY6+mOPxa7n4arEwR43ukhEQOB7YL5nHJNf3VEdaEEVK5wnJrclx7+kvursMGKL/m/9uQ9y+fVpHEI3rpSG/DolZA18h58LhzZoJEBZqLVfu6X5s++c/5YR+shKlPsk/+sLsToGEaUbvNxduLRrLZwiP3FyxNqady5zm6AyP2k9MUssi6Ilt3Ln+4rJMKkRCo86dKF5sm2yG2ZGT+5ETksbIBAXaavZFkqUgnmLnWwy6m+PL2gC0aQWu7xNP0v0Oiay8k8ou+3lJXRLnizvx1qyLrJDBK8cdHzV+ntybvQcwyDDg7u8DAx8UHAs3Ubjm3l7wrlM4yxFYIa4ku2hcyu0Y9pXPcQ2/GDCa+eGSnw/N6hgtoXOeWPuyzmM+uL4g0z3VH02X+OISMtBufrK48O4uPnzKroq4Uko19sRciksirwwltnMEq9Fi/sSLOKDrtdHkUa1ODYxCwr3kQSF+vpfen5ZeK1EK0oMMzHwNrzka//X+2flTQu8Sztvlt9nbDvlY5xcuvJ0CfBtyH0bTDj/SDEZTLhh0XO5GwhU9AF8aLehVi2WfdfPwAsnqvRIbZ1KY11Y4ZNvgm2aWo13nn2X9nWKEPe+qyAb9xGzwFwdoCj9FI02cGUuM2Nk6960rMF/x/aOKMx8skTl3ujtRQuNxnBV7nl9we8ET2u4thi5FQ3UHm/HKUKWMZU0Q7GKw7BEbigQHfeW5gFvuuP4oj5WGV4lHm/7JqcevYr4/idsGfat1v0USoxVZo6bFi3ixYFsd+ruw0q62Fx6GGnA==",
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	//req.Header.Set("cookie", "bm_s=YAAQm/AQAlWcvaWOAQAAFdv1ywErwYCLXjaa74pODGJyllhKS8kUrh6vJdbITuqW7skmCZ4e1JsYNri6EYa5D1YC013zlFtKoCg+KJ7b0nXJoO83godP4mL9xH3l8Z3s/iRyd908D7IM4kBSfhITyLudswOiQskazMM4DYRMmjWJEJPxPaoCsOE2qBQikV1bzhJ9Nj6hQ8PlrusBiytV8MW4W2Rcc09QpGQ489tLxFp1mgUB9LwhhwCKho9yq4dWNHKTqwpKL1wzBLYjlKgfkL3uXWBrowXJ/boZ4bwBj9v8UQOqodB3IB+THFWrCz9k/O5xVWEtgVSYNSoIaWpLkmW6F0UY; bm_ss=ab8e18ef4e; akacd_dc1=3890271670~rv=84~id=a44e2d980875c2973edc0148af4687d9; showbooking=true; bm_sz=A416F2267AD9EBB122C8A05BC133A18E~YAAQm/AQAi2dvaWOAQAAeeD1yxeONPp4iGGiCxEptJzTcD2PKdPUpf7HawF9tqGA3/2Ov2iMWgLBshFz/1YsxuhUjJzr9O/CL6AcOhmIi12MgtO11NUspP39V01CiSkuze6dCiCr7e+DoPyLfAKPtz7wyNdCSfTmR/AknIou7Q4psqDfuUxiMGLKj42+3SLTbQqxwcHYOZTanz+1eYHnZmXHxGRkUnpC9uc5GH1g40zxcQsHKgkv8uERV9FRudi+giTfrdq0Acna0VrHojeWM9Eo5OLFSJNVElPFZxQPBLrjRfNXIRoBqfaV9bYScsnGryf7HtDj37VsccBfNRHupBFyNF9Y/h8eqHVmqxzx0Q9CUStaBmcbbCM0rZpJky0tGfPbxSoEcqQqTpIRuBugVZ7NJnCX~3360051~3490371; _abck=6DA799D5C824B617650ABE27AC944135~-1~YAAQm/AQAvKhvaWOAQAAzPb1ywvyA7qx0v6uOtBDV+5thJZ0QnUBlEZ5R7dBaKolCisVObM867WKqa5UCExH1mGyMmK4hLVPSdTtVj8uEcfc6yTnk+CV10dNB2cegfDNpGqUXJrD9MGEOggYkYGAaqK1pp34UNM7/XPg76iKnbDYbzfK/irjHptmGLYboOiKilaNM1ycJ9oCg+ik0Y104l4f8CbOmBlLWeKNbVI+N8mz05W09/hlOQYjzQxMREroUllrpIzOqhUtksyRiX/0c+sFjPGwVhXjGageneecOHJLyl71DCpKc+uaH/JMpCITO9BFjzCmQ51eCNh/ke75fXfgaLTvUKwmFpvWSztJqJEZojFGYXYxWs2swYElMFMlpBQwemDBHESyGccGZ7H70Tc/JqP3OCPi~-1~-1~-1; device=mobile")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		return
	}
	//resp, err := client.Get(urlArg.String())
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//for k, v := range jar.Cookies(u) {
	//	fmt.Println(k, v)
	//}

	//req2, _ := http.NewRequest("POST", "https://tickets.vueling.com//ScheduleSelectNew.aspx", nil)
	//req2.Header.Set("Host", "tickets.vueling.com")
	//req2.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	//req2.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	//req2.Header.Set("referer", "https://www.vueling.com/de")
	//req2.Header.Set("accept-language", "en")
	////req.Header.Set("cookie", "bm_s=YAAQm/AQAlWcvaWOAQAAFdv1ywErwYCLXjaa74pODGJyllhKS8kUrh6vJdbITuqW7skmCZ4e1JsYNri6EYa5D1YC013zlFtKoCg+KJ7b0nXJoO83godP4mL9xH3l8Z3s/iRyd908D7IM4kBSfhITyLudswOiQskazMM4DYRMmjWJEJPxPaoCsOE2qBQikV1bzhJ9Nj6hQ8PlrusBiytV8MW4W2Rcc09QpGQ489tLxFp1mgUB9LwhhwCKho9yq4dWNHKTqwpKL1wzBLYjlKgfkL3uXWBrowXJ/boZ4bwBj9v8UQOqodB3IB+THFWrCz9k/O5xVWEtgVSYNSoIaWpLkmW6F0UY; bm_ss=ab8e18ef4e; akacd_dc1=3890271670~rv=84~id=a44e2d980875c2973edc0148af4687d9; showbooking=true; bm_sz=A416F2267AD9EBB122C8A05BC133A18E~YAAQm/AQAi2dvaWOAQAAeeD1yxeONPp4iGGiCxEptJzTcD2PKdPUpf7HawF9tqGA3/2Ov2iMWgLBshFz/1YsxuhUjJzr9O/CL6AcOhmIi12MgtO11NUspP39V01CiSkuze6dCiCr7e+DoPyLfAKPtz7wyNdCSfTmR/AknIou7Q4psqDfuUxiMGLKj42+3SLTbQqxwcHYOZTanz+1eYHnZmXHxGRkUnpC9uc5GH1g40zxcQsHKgkv8uERV9FRudi+giTfrdq0Acna0VrHojeWM9Eo5OLFSJNVElPFZxQPBLrjRfNXIRoBqfaV9bYScsnGryf7HtDj37VsccBfNRHupBFyNF9Y/h8eqHVmqxzx0Q9CUStaBmcbbCM0rZpJky0tGfPbxSoEcqQqTpIRuBugVZ7NJnCX~3360051~3490371; _abck=6DA799D5C824B617650ABE27AC944135~-1~YAAQm/AQAvKhvaWOAQAAzPb1ywvyA7qx0v6uOtBDV+5thJZ0QnUBlEZ5R7dBaKolCisVObM867WKqa5UCExH1mGyMmK4hLVPSdTtVj8uEcfc6yTnk+CV10dNB2cegfDNpGqUXJrD9MGEOggYkYGAaqK1pp34UNM7/XPg76iKnbDYbzfK/irjHptmGLYboOiKilaNM1ycJ9oCg+ik0Y104l4f8CbOmBlLWeKNbVI+N8mz05W09/hlOQYjzQxMREroUllrpIzOqhUtksyRiX/0c+sFjPGwVhXjGageneecOHJLyl71DCpKc+uaH/JMpCITO9BFjzCmQ51eCNh/ke75fXfgaLTvUKwmFpvWSztJqJEZojFGYXYxWs2swYElMFMlpBQwemDBHESyGccGZ7H70Tc/JqP3OCPi~-1~-1~-1; device=mobile")
	//resp2, err := client.Do(req2)
	//if err != nil {
	//	return
	//}
	////resp, err := client.Get(urlArg.String())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer resp2.Body.Close()
	//bytess, err := io.ReadAll(resp2.Body)
	//fmt.Println(string(bytess))
}
