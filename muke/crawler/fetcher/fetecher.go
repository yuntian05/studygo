package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)
func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	log.Printf("Fetching url: %s", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	//request.AddCookie(&http.Cookie{
	//	Name:  "sid",
	//	Value: "79992cee-42ff-462c-8281-80770edee22d",
	//})
	//request.AddCookie(&http.Cookie{
	//	Name:  "ec",
	//	Value: "1OJdJVBU-1652896572038-24fc12ec8a8be-1537031358",
	//})
	//request.AddCookie(&http.Cookie{
	//	Name:  "Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2",
	//	Value: "1652896579",
	//})
	//request.AddCookie(&http.Cookie{
	//	Name:  "FSSBBIl1UgzbN7NO",
	//	Value: "5I3tL0pYAY3eqs7_iqGKAIpkHSYD3rxjYxuNH.dwS7u3FtiVmFr6CCi59cLMN_pcKNARZNZKeWoHkivFowvKKwq",
	//})
	//request.AddCookie(&http.Cookie{
	//	Name:  "_exid",
	//	Value: "a23sCGIpVOvk2bpzqz2rjrROv6x6QzAhQAmCN5mvayVEj+oxbG50U1aIpydn2/odj3YkoRzs+GNywgBJWfXeKg==",
	//})
	//request.AddCookie(&http.Cookie{
	//	Name:  "_efmdata",
	//	Value: "f3SUpcjcqp6Qjn27bXPxhibZcFq+EtFNPLxbmJcGruy8J9WjLcycJUHSjCcQnrVzdFzpc2c3I53IxIjkdmn8jjAdo3+uzj/2uR4YTcvitHE=",
	//})
	//request.AddCookie(&http.Cookie{
	//	Name:  "_efmdata",
	//	Value: "f3SUpcjcqp6Qjn27bXPxhibZcFq+EtFNPLxbmJcGruy8J9WjLcycJUHSjCcQnrVzdFzpc2c3I53IxIjkdmn8jjAdo3+uzj/2uR4YTcvitHE=",
	//})
	//request.AddCookie(&http.Cookie{
	//	Name:  "Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2",
	//	Value: "1653386941",
	//})
	//request.AddCookie(&http.Cookie{
	//	Name:  "FSSBBIl1UgzbN7NP",
	//	Value: "53LnXaCtW..aqqqDqeRHsnqLlgEnkQgn2IQEUcYnH99p72mhqP9NRe5Pq9gq4bKdamIpntIataYuR6OvNf1LplDTZIODTJ0.ktaXMkxnyHLR7joqdda55OnU7_7HMUc8ZEyBqO3p8.8cy3AmjXtLWrVWC8aNT3_9UGkFiTaxNlY0XjFNOJbt5v.mhf3Gme0o1.R4OfvWmMEjN9yKudC.wPVXnGs1N653I1WmE.rB07fld3mmW285pJ1eg3z.FqUcrTCBg4n5lQNhPpacBwcmz_hO6B6Ki9G3o4s8JmYhYkO",
	//})
	//options := cookiejar.Options{
	//	PublicSuffixList: publicsuffix.List,
	//}
	//jar, err := cookiejar.New(&options)
	//if err != nil {
	//	return nil, err
	//}
	client := http.Client{
		CheckRedirect: func(req *http.Request,
			via []*http.Request) error {
			//fmt.Println(req)
			return nil
		},
		//Jar: jar,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d\n", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
