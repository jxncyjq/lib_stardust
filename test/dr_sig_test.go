package test

import (
	"fmt"
	sc "github.com/jxncyjq/lib_stardust/helper/utils/crypto"
	"sort"
	"testing"
)

func TestSignInfo(t *testing.T) {
	key := `2d9d60da9726671b3d72ea0508df2d4b`
	t.Log("TestSignInfo")
	chargeMoney := "79"
	channel := "rustore"
	attach := "1,1,399712059115175941,2,22"
	signMap := make(map[string]string)
	signMap["chargeMoney"] = chargeMoney
	signMap["channel"] = channel
	signMap["attach"] = attach

	// 提取 map 的键并排序
	keys := make([]string, 0, len(signMap))
	for key := range signMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	sortMap := make(map[string]string)
	// 按排序后的键输出 map 的值
	for _, key := range keys {
		sortMap[key] = signMap[key]
		fmt.Printf("%s: %s\n", key, sortMap[key])
	}

	paramsString := ""
	for k, v := range sortMap {
		t.Logf("k=%s,v=%s", k, v)
		paramsString += k + "=" + v + "&"
	}

	paramsString = paramsString[:len(paramsString)-1]
	paramsString += "&key=" + key
	t.Logf("paramsString:%s", paramsString)
	t.Log("---------------------------------------------")
	sign := sc.GenerateMD5(paramsString)
	t.Logf("sign:%s", sign)
	paramsString += "&sign=" + sign
	t.Logf(paramsString)
	t.Logf("TestSignInfo end ,request body :%v", signMap)
}
