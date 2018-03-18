/************************************************************
** @Description: parser
** @Author: haodaquan
** @Date:   2018-03-17 22:16
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 22:16
*************************************************************/
package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var HeightRe = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([\d]+)CM</span></td>`)
var WeightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d|\-]+)</span></td>`)
var GenderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var MarriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var EducationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var HukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var IncomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var XinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var HouseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var CarRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var OccupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(content []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(content, ageRe))
	if err != nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(content, HeightRe))
	if err != nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(content, WeightRe))
	if err != nil {
		profile.Weight = weight
	}
	profile.Gender = extractString(content, GenderRe)
	profile.Marriage = extractString(content, MarriageRe)
	profile.Income = extractString(content, IncomeRe)
	profile.Car = extractString(content, CarRe)
	profile.House = extractString(content, HouseRe)
	profile.Hukou = extractString(content, HukouRe)
	profile.Occupation = extractString(content, OccupationRe)
	profile.Xingzuo = extractString(content, XinzuoRe)
	profile.Education = extractString(content, EducationRe)

	profile.Name = name
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
