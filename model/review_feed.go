package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ReviewFeed struct {
	ReviewList `xml:"entry"`
}

type ReviewList []Review

func (f *ReviewFeed) Json() (string, error) {
	b, err := json.Marshal(f)
	return string(b), err
}

func (f *ReviewFeed) Stats() string {
	str := "stats: "
	for i := 5; i > 0; i-- {
		str = str + f.ratio(i)
		if i != 1 {
			str = str + ", "
		}
	}
	return str
}

func (f *ReviewFeed) ratio(rating int) string {
	rat, count, total := f.Ratio(rating)
	r := strconv.Itoa(rating)
	c := strconv.Itoa(count)
	t := strconv.Itoa(total)
	s := fmt.Sprintf("%2.0f", rat)
	return "[" + r + "]" + s + "p(" + c + "/" + t + ")"
}

func (f *ReviewFeed) Ratio(rating int) (ratio float64, count, total int) {
	r := len(f.Select(rating))
	l := len(f.ReviewList)
	return (float64(r) / float64(l)) * 100, r, l
}

func (f *ReviewFeed) Select(rating int) ReviewList {
	rate := strconv.Itoa(rating)
	list := ReviewList{}
	for _, r := range f.ReviewList {
		if r.Rating == rate {
			list = append(list, r)
		}
	}
	return list
}

func NewReviewFeed(b []byte) *ReviewFeed {
	content := string(b)
	for {
		begin := strings.Index(content, "<content type=\"html\">")
		if begin == -1 {
			break
		}
		rem := content[begin+1:]
		end := strings.Index(rem, "</content>")
		content = content[:begin] + rem[end+len("</content>"):]
	}

	var feed ReviewFeed
	err := xml.Unmarshal([]byte(content), &feed)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// ReviewList[0] is app entry(no review entry)
	feed.ReviewList = feed.ReviewList[1:]
	return &feed
}
