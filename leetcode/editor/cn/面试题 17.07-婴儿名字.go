package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	names := []string{"John(15)", "Jon(12)", "Chris(13)", "Kris(4)", "Christopher(19)"}
	synonyms := []string{"(Jon,John)", "(John,Johnny)", "(Chris,Kris)", "(Chris,Christopher)"}
	// [Acohsf(86) Aeax(646) Alrksy(69) Ard(82) Avcp(41) Avmzs(39) Axwtno(57) Bnea(179) Chhmx(259) Chycnm(253) Csh(238) Dhe(79) Dnsay(72) Drzsnw(87) Dwifi(237) Ebov(96) Fcclu(70) Fpaf(219) Fvkhz(104) Gauuk(75) Gbkq(77) Gnplfi(109) Hcvcgc(97) Hfp(97) Hljt(51) Ibink(32) Ijveol(46) Jfq(26) Jvqg(47) Kasgmw(95) Kgabb(236) Knpuq(61) Koaak(53) Kufa(95) Mcnef(59) Msyr(211) Mtz(72) Mwwvj(70) Naf(3) Nekuam(17) Npilye(25) Nsgbth(26) Ntfr(70) Nuq(61) Nzaz(51) Obcbxb(124) Ofl(72) Okwuq(96) Oltadg(95) Onnev(77) Qbmk(45) Qiy(26) Qyqifg(85) Ucqh(51) Unsb(26) Uvkdpn(71) Weusps(79) Yeekgc(11) Yjc(94)]
	names = []string{"Fcclu(70)", "Ommjh(63)", "Dnsay(60)", "Qbmk(45)", "Unsb(26)", "Gauuk(75)", "Wzyyim(34)", "Bnea(55)", "Kri(71)", "Qnaakk(76)", "Gnplfi(68)", "Hfp(97)", "Qoi(70)", "Ijveol(46)", "Iidh(64)", "Qiy(26)", "Mcnef(59)", "Hvueqc(91)", "Obcbxb(54)", "Dhe(79)", "Jfq(26)", "Uwjsu(41)", "Wfmspz(39)", "Ebov(96)", "Ofl(72)", "Uvkdpn(71)", "Avcp(41)", "Msyr(9)", "Pgfpma(95)", "Vbp(89)", "Koaak(53)", "Qyqifg(85)", "Dwayf(97)", "Oltadg(95)", "Mwwvj(70)", "Uxf(74)", "Qvjp(6)", "Grqrg(81)", "Naf(3)", "Xjjol(62)", "Ibink(32)", "Qxabri(41)", "Ucqh(51)", "Mtz(72)", "Aeax(82)", "Kxutz(5)", "Qweye(15)", "Ard(82)", "Chycnm(4)", "Hcvcgc(97)", "Knpuq(61)", "Yeekgc(11)", "Ntfr(70)", "Lucf(62)", "Uhsg(23)", "Csh(39)", "Txixz(87)", "Kgabb(80)", "Weusps(79)", "Nuq(61)", "Drzsnw(87)", "Xxmsn(98)", "Onnev(77)", "Owh(64)", "Fpaf(46)", "Hvia(6)", "Kufa(95)", "Chhmx(66)", "Avmzs(39)", "Okwuq(96)", "Hrschk(30)", "Ffwni(67)", "Wpagta(25)", "Npilye(14)", "Axwtno(57)", "Qxkjt(31)", "Dwifi(51)", "Kasgmw(95)", "Vgxj(11)", "Nsgbth(26)", "Nzaz(51)", "Owk(87)", "Yjc(94)", "Hljt(21)", "Jvqg(47)", "Alrksy(69)", "Tlv(95)", "Acohsf(86)", "Qejo(60)", "Gbclj(20)", "Nekuam(17)", "Meutux(64)", "Tuvzkd(85)", "Fvkhz(98)", "Rngl(12)", "Gbkq(77)", "Uzgx(65)", "Ghc(15)", "Qsc(48)", "Siv(47)"}
	synonyms = []string{"(Gnplfi,Qxabri)", "(Uzgx,Siv)", "(Bnea,Lucf)", "(Qnaakk,Msyr)", "(Grqrg,Gbclj)", "(Uhsg,Qejo)", "(Csh,Wpagta)", "(Xjjol,Lucf)", "(Qoi,Obcbxb)", "(Npilye,Vgxj)", "(Aeax,Ghc)", "(Txixz,Ffwni)", "(Qweye,Qsc)", "(Kri,Tuvzkd)", "(Ommjh,Vbp)", "(Pgfpma,Xxmsn)", "(Uhsg,Csh)", "(Qvjp,Kxutz)", "(Qxkjt,Tlv)", "(Wfmspz,Owk)", "(Dwayf,Chycnm)", "(Iidh,Qvjp)", "(Dnsay,Rngl)", "(Qweye,Tlv)", "(Wzyyim,Kxutz)", "(Hvueqc,Qejo)", "(Tlv,Ghc)", "(Hvia,Fvkhz)", "(Msyr,Owk)", "(Hrschk,Hljt)", "(Owh,Gbclj)", "(Dwifi,Uzgx)", "(Iidh,Fpaf)", "(Iidh,Meutux)", "(Txixz,Ghc)", "(Gbclj,Qsc)", "(Kgabb,Tuvzkd)", "(Uwjsu,Grqrg)", "(Vbp,Dwayf)", "(Xxmsn,Chhmx)", "(Uxf,Uzgx)"}
	popular := trulyMostPopular(names, synonyms)
	sort.Strings(popular)
	fmt.Println(popular)
}

// leetcode submit region begin(Prohibit modification and deletion)
func trulyMostPopular(names []string, synonyms []string) []string {
	// 终版
	uni := make(map[string]string) // 1.并查集
	var find func(string) string
	find = func(x string) string {
		if _, ok := uni[x]; !ok {
			return x
		}
		for uni[x] != x {
			x, uni[x] = uni[x], uni[uni[x]]
		}
		return uni[x]
	}
	for _, s := range synonyms {
		idx := strings.Index(s, ",")
		s1, s2 := s[1:idx], s[idx+1:len(s)-1]
		s1, s2 = find(s1), find(s2) // 核心逻辑
		if s1 > s2 {
			s1, s2 = s2, s1
		}
		uni[s1], uni[s2] = s1, s1 // 映射
	}

	freq := make(map[string]int) // 2.统计频率
	for _, s := range names {
		idx := strings.Index(s, "(")
		k, val := s[:idx], s[idx+1:len(s)-1]
		v, _ := strconv.Atoi(val)
		k = find(k)
		freq[k] += v
	}

	ans := make([]string, 0, len(freq)) // 3.输出结果集
	for k, v := range freq {
		ans = append(ans, k+"("+strconv.Itoa(v)+")")
	}
	return ans

	// lc
	//freq := make(map[string]int)
	//for _, s := range names {
	//	idx := strings.Index(s, "(")
	//	k, val := s[:idx], s[idx+1:len(s)-1]
	//	v, _ := strconv.Atoi(val)
	//	freq[k] += v // 统计 name 的频率
	//}
	//
	//union := make(map[string]string)
	//var find func(string) string
	//find = func(x string) string {
	//	if _, ok := union[x]; !ok { // 不能返回 ""
	//		return x
	//	}
	//	if union[x] != x {
	//		union[x] = find(union[x])
	//	}
	//	return union[x]
	//}
	//for _, s := range synonyms {
	//	idx := strings.Index(s, ",")
	//	s1, s2 := s[1:idx], s[idx+1:len(s)-1]
	//	// 核心逻辑
	//	s1, s2 = find(s1), find(s2) // 讨论 s1 s2 是否出现过
	//	if s1 != s2 {
	//		if s1 > s2 { // 求得最小字符串
	//			s1, s2 = s2, s1
	//		}
	//		freq[s1] += freq[s2] // 合并
	//		delete(freq, s2)     // 删除大字符串
	//	}
	//	union[s2] = s1
	//}
	//
	//ans := make([]string, 0, len(freq))
	//for k, v := range freq {
	//	ans = append(ans, k+"("+strconv.Itoa(v)+")")
	//}
	//return ans

	// 个人
	//n := len(synonyms)
	//uni := make([]int, n)
	//for i := range uni {
	//	uni[i] = i
	//}
	//var find func(int) int
	//find = func(x int) int {
	//	if uni[x] != x {
	//		uni[x] = find(uni[x])
	//	}
	//	return uni[x]
	//}
	//union := func(x, y int) {
	//	xr, yr := find(x), find(y)
	//	if xr != yr {
	//		uni[yr] = xr
	//	}
	//}
	//
	//memo := make(map[string]int)
	//for i, s := range synonyms {
	//	idx := strings.Index(s, ",")
	//	s1, s2 := s[1:idx], s[idx+1:len(s)-1]
	//	i1, ok1 := memo[s1]
	//	i2, ok2 := memo[s2]
	//	if ok1 && ok2 {
	//		union(i1, i2)
	//		union(i1, i)
	//	} else if ok1 {
	//		union(i1, i)
	//	} else if ok2 {
	//		union(i2, i)
	//	}
	//	u := find(i)
	//	memo[s1], memo[s2] = u, u
	//}
	//
	//keys := make(map[int]string)
	//for s, u := range memo {
	//	k := find(u)
	//	if keys[k] == "" {
	//		keys[k] = s
	//	} else {
	//		keys[k] = min(keys[k], s)
	//	}
	//}
	//
	//temp := make(map[string]int)
	//for _, s := range names {
	//	idx := strings.Index(s, "(")
	//	k, val := s[:idx], s[idx+1:len(s)-1]
	//	v, _ := strconv.Atoi(val)
	//	if _, ok := memo[k]; !ok {
	//		temp[k] = v
	//	} else {
	//		temp[keys[find(memo[k])]] += v
	//	}
	//}
	//ans := make([]string, 0, len(temp))
	//for k, v := range temp {
	//	ans = append(ans, k+"("+strconv.Itoa(v)+")")
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
