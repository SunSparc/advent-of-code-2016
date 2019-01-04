package main

import "fmt"

func main() {
	twoFers := 0
	threeFers := 0
	for _, boxID := range input {
		letterCountMap := getIdMap(boxID)
		twoFerFound := false
		threeFerFound := false
		for _, letterCount := range letterCountMap {
			if letterCount == 2 {
				twoFerFound = true
			}
			if letterCount == 3 {
				threeFerFound = true
			}
		}
		if twoFerFound {
			twoFers++
		}
		if threeFerFound {
			threeFers++
		}
	}
	fmt.Printf("%d * %d\n", twoFers, threeFers)
	checksum := twoFers * threeFers
	fmt.Println("checksum:", checksum)
}

func getIdMap(boxID string) map[rune]int {
	idMap := map[rune]int{}
	for _, letter := range boxID {
		if _, ok := idMap[letter]; ok {
			idMap[letter] = idMap[letter] + 1
		} else {
			idMap[letter] = 1
		}
	}
	return idMap
}

var input = []string{
	"oeambtcgjqnzhgkdylfapoiusr",
	"oewmbtcxjqnzhgvdyltapvqusr",
	"oewmbtchjqnzigkdylfapviuse",
	"oeimbucxjqnzhgkdyyfapviusr",
	"fewmbtcxjqndhgcdylfapviusr",
	"oevgbtccjqnzhgkdylfapviusr",
	"oewmbtcxjqnzhnkdylmapvpusr",
	"oewmbtcxjqnzhxkdyldapvirsr",
	"oewmutccjqnzngkdylfapviusr",
	"oewmbtcxbqnzhgkdsliapviusr",
	"ozwmbtfxjqnzcgkdylfapviusr",
	"oewmbtdxjqnzhgkdypfapsiusr",
	"oeylbtcxjqnzhgyyylfapviusr",
	"oewmbtcxjqnzhgkdrlfakuiusr",
	"oewmbtcujqnxhgkdylfadviusr",
	"oewmbtcxlqpzhgkdylfaaviusr",
	"oewmztcxjqnzhgkdylfqpliusr",
	"oeembtcxjqnzhgkdtlmapviusr",
	"onwmbtcxjqnqhgkdylfapdiusr",
	"oewmbtcxnqnzhgkdylfapsbusr",
	"oeoibtjxjqnzhgkdylfapviusr",
	"oxwmbtcxjqnzhgkdylfapcipsr",
	"oewmbtoxbqnzhgzdylfapviusr",
	"okwubtcxjqnzhgkdylfapiiusr",
	"oewmbtcxjqnzhgodylfapnicsr",
	"oewmitcxjqnzhgkdylfaphlusr",
	"oewmbtaxjqnzhgkhylfapveusr",
	"oewmftcbjqnzhgkdylfapviurr",
	"oewmbtcujqnzbgkdylfapliusr",
	"oeevbtcxjqnzhgkdylfapniusr",
	"oewmbtcxjqnvhgkdylfapnpusr",
	"oewabtcxjqnzhgddylfapviust",
	"oewmbtyxjqnzhgkdvlfapvinsr",
	"jewmbtcxjonzhzkdylfapviusr",
	"oewmbrcxjqnzxgkdylfapoiusr",
	"dewmbtmxjqnzhgkdyvfapviusr",
	"oewmbtctjqnzhgkdmlfapvihsr",
	"oewmbjcxjqnzhgvdylfapviurr",
	"oewmbtcxjqnzhgcdxlfapvfusr",
	"oewmbucxjqnzhgkdyltapvifsr",
	"gewmbtcejqezhgkdylfapviusr",
	"oewebtcxjznzhgkdylfapvhusr",
	"oewmjtcxjqnzhgkdycfakviusr",
	"oewmbtcxjtnvhgkdylfabviusr",
	"oewmbtcxjqnthgkgclfapviusr",
	"hewmbtcxjqnzhgkdwlfapziusr",
	"oewmbtcxjqnzhgkdylfqpviysf",
	"oewmbtcxjvnzhgmdylfapviuse",
	"oewmbtcxjqnphgkdymzapviusr",
	"oewmbtcxjqnzwmkdylfapvbusr",
	"oewmbthxjqnzhgkdylfatvilsr",
	"oewmbtcxaznzhgkhylfapviusr",
	"zewmbscxjqnzhgkdylfatviusr",
	"oewmbecyaqnzhgkdylfapviusr",
	"oewmbtnxjqnzhekxylfapviusr",
	"oewmbtcxjqczhgkdyltnpviusr",
	"yewmbecxjnnzhgkdylfapviusr",
	"oewmbocxjqnzhgkyylfapviusv",
	"oewmxtcxjqnzhgkkylfspviusr",
	"oiwmbtcxjqnzhgkdydfapvgusr",
	"oewmbtcxjqnzngydylfwpviusr",
	"oewmctcxjqnzhgkdelfapviasr",
	"oewmbtcxjqnzhgxdwmfapviusr",
	"oewmntcxjqnzhgkdylfamviusw",
	"oewmatcxjqbzhgkdylfapvhusr",
	"oewmbtcxjqnqhmkdyluapviusr",
	"opwmbtcxjqnzhgkdywfapvilsr",
	"omwmbtcxjqnlhgkdylyapviusr",
	"oewmltcxoqnzhgkdylfapvfusr",
	"oewmbtcxjqtzhgkdyyoapviusr",
	"oewmbtcxjqnzhrkdzlffpviusr",
	"oewmbtqxyqnzhgkdylfalviusr",
	"oeuzbtcxlqnzhgkdylfapviusr",
	"oewmbtcxjqnzhtxdylflpviusr",
	"oewmdtfxjqnzhgkdylfapviufr",
	"ojwmbtcxjqnzhgkdylfypviqsr",
	"oewmbtcxjqnzhgkdylfapvivuf",
	"oewmjtcsjqnzxgkdylfapviusr",
	"ohembtcxjqnzhnkdylfapviusr",
	"oewmptcajqnzhgkdylfapviusd",
	"oewmbtcxjcnwhgkbylfapviusr",
	"oewmbtcxjqnzhgddnlfapvqusr",
	"oewmbtcfjqnzhgkdypfapvzusr",
	"oewdbtccjqnzhgfdylfapviusr",
	"oewmbtcxjpnzhgkdplfaqviusr",
	"oepmbhcxjqnzhgkdylfaaviusr",
	"oewmbtcwjqxzhgkwylfapviusr",
	"oewmatcxjqnchgkdylfapvifsr",
	"omwmbncxjqnzhgkdylfapviuyr",
	"sewmbsckjqnzhgkdylfapviusr",
	"oewubtcxjqnzhgkdyluapvausr",
	"ohwmbtcxqqhzhgkdylfapviusr",
	"oewmbtcxjqnzhgkpylfapnissr",
	"eewmbccxjqnzhgkdylbapviusr",
	"oewmitcyjqnzhgkdylkapviusr",
	"oewmbtcxjvnzhgkdyjfvpviusr",
	"oewmbtcxjqmzhgkdyefagviusr",
	"oewmbtcvjqnzhgkdylpapviufr",
	"oewmbtcxjrnkhgkdylfapsiusr",
	"oewmbtcxjqnzygkdylfaxvipsr",
	"oexmbtcxjqczhgkdyloapviusr",
	"oewmbtcxjqnlhtkdylfapvmusr",
	"oewmbtcxdqjzdgkdylfapviusr",
	"oewmbtclgqnzhgkdylfabviusr",
	"oewmbtvfjqnzhgkdylfapviulr",
	"oewmbtcxjqnzhgkdyllarvijsr",
	"oewmbtyxjqnzhgpdylxapviusr",
	"oeylbtcxjqnzhgkyylfapviusr",
	"oewmbtctjqnzhjkdylfapviulr",
	"oermatcxjqnzhgkdylzapviusr",
	"oewmbtcxjqnztgkdzlfapviutr",
	"oewlbtcxjqnztgkvylfapviusr",
	"oewmbtcxjqzvhgkdylfapviusk",
	"oewmbtcxjqnzmgkdyldapvilsr",
	"felmbtcxjqnzhgkdylfapviusl",
	"oewmbtcxjgnzhgkjylfaeviusr",
	"ovwmbtcxjqpzhgkdylfapvizsr",
	"eewmbtcpjqnzhgkdylfapvijsr",
	"oewmbzcxjqnzhgkdylfaeviutr",
	"tewmbtcljqhzhgkdylfapviusr",
	"oewmbtcujqnzhgkdnliapviusr",
	"oewmbtcljqnzhskdylfapvgusr",
	"oewmbtchjqnzhgkdylmapviuse",
	"oewmbtcxjqnzbgkdylfaiviurr",
	"oewmbtcxjqnzhkkdyloapsiusr",
	"oewjbtcxjqnhhgkdylfapjiusr",
	"odwmbtcnjqnzhgkdylfapvicsr",
	"oewmbccxjqrzwgkdylfapviusr",
	"kewmbtcvjqnzhgkdylaapviusr",
	"okwmbtcxjqnzhgkdylfspvausr",
	"oewmbtcxjynzhgkdyafapviusw",
	"oewmbtcxjqnzhgwdyleayviusr",
	"oewmbtcxjqnzhgkdylfapviicl",
	"oewmbtcxjqnzhgkdyltaeziusr",
	"oewmbtcxrqnzhgkdylftpvizsr",
	"oewsrtcgjqnzhgkdylfapviusr",
	"oewmbtsxgqnzhgxdylfapviusr",
	"oewmbtcxjanzhgtdylfapeiusr",
	"oewybtcgjqnzhgkdylfapviust",
	"ojwmbncxjqnzhgkdylfapgiusr",
	"ocgebtcxjqnzhgkdylfapviusr",
	"oejcbtcxjqnzhgkvylfapviusr",
	"oswmbtcxjqnkhgkdylfapviusb",
	"oewdbtcxjqnzdgkdylfypviusr",
	"oawmutcxjqnzhgkddlfapviusr",
	"oewzbtcxyqnzhgkdylfapviusy",
	"zewmbtcxjqnzkgkdylwapviusr",
	"aewmbtkxjqnzhgkdylfapviuer",
	"oewmbtcxwqnzhgkdyofapviuur",
	"oewmbtcxjqnzggkdylfapanusr",
	"oewmstcxuqnzhgkdylzapviusr",
	"zewmbtcxjqozhgkdelfapviusr",
	"oewzbtcxjqnahgkdyllapviusr",
	"fewmatcxjqnghgkdylfapviusr",
	"oewmbtcxjqnzhgkdylfapviyqb",
	"oewwbtcxjqnzhgkdyljapviqsr",
	"oewmbtbxjqnzhgkxylfapviesr",
	"oewmbtcbjqnphgkdylfapviysr",
	"oewabtcxyqnzhgkdylfabviusr",
	"oewmbtcxhknzhgkdylfapviusd",
	"ozwmbtcljqnzhgkdylfapviksr",
	"tewmbtcxjqnzhgkdylfaxvqusr",
	"oewmbtcxrqnzhgkdytfapvrusr",
	"ohwmbtcxjcnzhgkdyifapviusr",
	"oewmbpcxjqnzhwkdylfaphiusr",
	"oedmbtcxjqnzhgnbylfapviusr",
	"oewmbocxjqnehgkdylfapvbusr",
	"oeymbtcxjqezegkdylfapviusr",
	"oewmbtcxjqnzhgkdllferviusr",
	"oewmbtcxjqnzhgkwmlfawviusr",
	"oewmbtcxienzhgkdyzfapviusr",
	"mewmbtcxjqnzhqkdylfapviwsr",
	"oewmbtcxjqnztgkmylfapvdusr",
	"ouwmbtcxjqnzhokdylpapviusr",
	"oewmctcxjqhzhgmdylfapviusr",
	"oewmbtcyjqnzhmkdylfarviusr",
	"oewmbtcxjqnzhgkdpnfzpviusr",
	"oewmptcxjqnzhgkdylkapviulr",
	"nefmbtcxsqnzhgkdylfapviusr",
	"oewmbtcxwqnzhgkdilfapvizsr",
	"eewmbtcxjqwzhghdylfapviusr",
	"oewmbtixmqnzhgkjylfapviusr",
	"okwmbtcdzqnzhgkdylfapviusr",
	"oewmbtxxjrnzigkdylfapviusr",
	"oewmdycxjqnzhekdylfapviusr",
	"oewmutcxjqnzhgkdylfapsiuqr",
	"oewmbacxjqnzrgkdrlfapviusr",
	"oewmbtcxpqnzhmkdylfapciusr",
	"oewabtcxjqnzhgkdyrcapviusr",
	"oswmbtcxjqnzhgkdrxfapviusr",
	"gewmbtcnjqnzhgkdylvapviusr",
	"newmbtcxjwnzfgkdylfapviusr",
	"lewmbtcxjqnzhgkdylfaptiujr",
	"oewwbtcxjqndhgkdylfapiiusr",
	"oewmbtcxjqnzhggdylfapvqmsr",
	"lewmbtcxjqnzhgkhllfapviusr",
	"oewmbtocjqnzhgkdylfapvhusr",
	"oedmbtcxjqnzhgkdyhfapviusb",
	"oewmbtcxjqnzhgkdylfajvaosr",
	"zewmbtcxjqnzhgkdylfapvsssr",
	"oewmbthxjqnzhskdylfapviuqr",
	"yewmrtcvjqnzhgkdylfapviusr",
	"oewmbtctjqnzhgkdylfabvhusr",
	"oesmstcxjqnzhgkdylfapqiusr",
	"oewmbtcxjqnzzgkdylfopiiusr",
	"otwmbtzxjqnzhgkdylfaxviusr",
	"ouwmbxcxjqnzhgkdylfapvnusr",
	"oewmbtcxjqezhgedylfapvsusr",
	"oesmhtcsjqnzhgkdylfapviusr",
	"oewdbtcxjqnzhgkdilfapvifsr",
	"oewmbtcxjqnzhgudynfamviusr",
	"qewhbtcxjqnzhgkdyxfapviusr",
	"oewmbzcxjqtzhgkdylfapvitsr",
	"oewmbtccjqzzhgkaylfapviusr",
	"jewmbtcxmqnzhlkdylfapviusr",
	"oewmbtcxjqbzhgkdnlfapviusp",
	"oeimbtcdjqnzhgkdylfapviuer",
	"oewtbtcxjqnihgkdylfahviusr",
	"oewmbtcxhqnzhgkdylfapdiudr",
	"oefmbtcxjqyshgkdylfapviusr",
	"oewmbtcxjqnzhgkfglfapviusx",
	"oecmbocxjqnzhgkdmlfapviusr",
	"oewmbtcxjqnzhghdylfavviuhr",
	"oewmbmcxiqnzhgkpylfapviusr",
	"oewmbtcnjqnzhgkrylfanviusr",
	"oewmbocxjqnzhzkdllfapviusr",
	"eewmbtckjqnzhgkdylfapviusg",
	"oewmbtcrjqqzhgkdylfapvigsr",
	"oewmbtcxjqazhgfdylfapjiusr",
	"oetmbtcxjqnzhgldylfapviqsr",
	"oewbbtcxjqnzlgkdylfapviuse",
	"oewmbtcxjqnzhglbolfapviusr",
	"oewmbtcxjqnzcgkdylfapviuhy",
	"oelmbtcxjqfzhgkdylaapviusr",
	"oojmbycxjqnzhgkdylfapviusr",
	"oewmbtrxjqnrhgkdylfapniusr",
	"oewmbtcxjqnzhgkyyhfapviuso",
	"oewabtcxjqnzhskdylfapviusx",
	"oewmbtcrjqnmhgkdylfapvnusr",
	"oewmbtcxjqrzhgkdylfapvpuss",
	"oewmbtcxhqnzwgkddlfapviusr",
	"kewmbtcxjqnzhgkyylfajviusr",
	"oswmbtcxjqnzhgkdjlfapviuss",
	"onwmbtcxjqnchgkdylfapvpusr",
	"oeymbtcxjqnxhikdylfapviusr",
	"oewmblcdjqnzhgkdylfapviysr",
	"oewmbtcxeqczhgudylfapviusr",
	"oewmbpgxjqnzhgkdylfapfiusr",
	"ohwmwtcxjqnzhgkdylftpviusr",
	"zebmbtuxjqnzhgkdylfapviusr",
}
