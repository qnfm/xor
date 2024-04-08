package main

import (
	"log"
	"os"

	"github.com/lukechampine/fastxor"
)

func main() {
	in1, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}
	in2, err := os.Open(os.Args[2])
	if err != nil {
		log.Println(err)
		return
	}
	fi, err := in1.Stat()
	if err != nil {
		log.Fatal(err)
	}
	batch := fi.Size() / 1024

	o, err := os.Create(os.Args[3])
	if err != nil {
		log.Println(err)
		return
	}
	defer in1.Close()
	defer in2.Close()
	defer o.Close()
	buf1 := make([]byte, 1024)
	buf2 := make([]byte, 1024)
	buf3 := make([]byte, 1024)
	if fi.Size() < 1024 {
		sz := int(fi.Size())
		in1.Read(buf1)
		in2.Read(buf2)
		fastxor.Bytes(buf3[0:sz], buf1[0:sz], buf2[0:sz])
		_, err = o.Write(buf3[0:sz])
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	var i int64
	for i = 0; i < batch; i++ {
		in1.Read(buf1)
		in2.Read(buf2)
		fastxor.Block(buf3[0:16], buf1[0:16], buf2[0:16])
		fastxor.Block(buf3[16:32], buf1[16:32], buf2[16:32])
		fastxor.Block(buf3[32:48], buf1[32:48], buf2[32:48])
		fastxor.Block(buf3[48:64], buf1[48:64], buf2[48:64])
		fastxor.Block(buf3[64:80], buf1[64:80], buf2[64:80])
		fastxor.Block(buf3[80:96], buf1[80:96], buf2[80:96])
		fastxor.Block(buf3[96:112], buf1[96:112], buf2[96:112])
		fastxor.Block(buf3[112:128], buf1[112:128], buf2[112:128])
		fastxor.Block(buf3[128:144], buf1[128:144], buf2[128:144])
		fastxor.Block(buf3[144:160], buf1[144:160], buf2[144:160])
		fastxor.Block(buf3[160:176], buf1[160:176], buf2[160:176])
		fastxor.Block(buf3[176:192], buf1[176:192], buf2[176:192])
		fastxor.Block(buf3[192:208], buf1[192:208], buf2[192:208])
		fastxor.Block(buf3[208:224], buf1[208:224], buf2[208:224])
		fastxor.Block(buf3[224:240], buf1[224:240], buf2[224:240])
		fastxor.Block(buf3[240:256], buf1[240:256], buf2[240:256])
		fastxor.Block(buf3[256:272], buf1[256:272], buf2[256:272])
		fastxor.Block(buf3[272:288], buf1[272:288], buf2[272:288])
		fastxor.Block(buf3[288:304], buf1[288:304], buf2[288:304])
		fastxor.Block(buf3[304:320], buf1[304:320], buf2[304:320])
		fastxor.Block(buf3[320:336], buf1[320:336], buf2[320:336])
		fastxor.Block(buf3[336:352], buf1[336:352], buf2[336:352])
		fastxor.Block(buf3[352:368], buf1[352:368], buf2[352:368])
		fastxor.Block(buf3[368:384], buf1[368:384], buf2[368:384])
		fastxor.Block(buf3[384:400], buf1[384:400], buf2[384:400])
		fastxor.Block(buf3[400:416], buf1[400:416], buf2[400:416])
		fastxor.Block(buf3[416:432], buf1[416:432], buf2[416:432])
		fastxor.Block(buf3[432:448], buf1[432:448], buf2[432:448])
		fastxor.Block(buf3[448:464], buf1[448:464], buf2[448:464])
		fastxor.Block(buf3[464:480], buf1[464:480], buf2[464:480])
		fastxor.Block(buf3[480:496], buf1[480:496], buf2[480:496])
		fastxor.Block(buf3[496:512], buf1[496:512], buf2[496:512])
		fastxor.Block(buf3[512:528], buf1[512:528], buf2[512:528])
		fastxor.Block(buf3[528:544], buf1[528:544], buf2[528:544])
		fastxor.Block(buf3[544:560], buf1[544:560], buf2[544:560])
		fastxor.Block(buf3[560:576], buf1[560:576], buf2[560:576])
		fastxor.Block(buf3[576:592], buf1[576:592], buf2[576:592])
		fastxor.Block(buf3[592:608], buf1[592:608], buf2[592:608])
		fastxor.Block(buf3[608:624], buf1[608:624], buf2[608:624])
		fastxor.Block(buf3[624:640], buf1[624:640], buf2[624:640])
		fastxor.Block(buf3[640:656], buf1[640:656], buf2[640:656])
		fastxor.Block(buf3[656:672], buf1[656:672], buf2[656:672])
		fastxor.Block(buf3[672:688], buf1[672:688], buf2[672:688])
		fastxor.Block(buf3[688:704], buf1[688:704], buf2[688:704])
		fastxor.Block(buf3[704:720], buf1[704:720], buf2[704:720])
		fastxor.Block(buf3[720:736], buf1[720:736], buf2[720:736])
		fastxor.Block(buf3[736:752], buf1[736:752], buf2[736:752])
		fastxor.Block(buf3[752:768], buf1[752:768], buf2[752:768])
		fastxor.Block(buf3[768:784], buf1[768:784], buf2[768:784])
		fastxor.Block(buf3[784:800], buf1[784:800], buf2[784:800])
		fastxor.Block(buf3[800:816], buf1[800:816], buf2[800:816])
		fastxor.Block(buf3[816:832], buf1[816:832], buf2[816:832])
		fastxor.Block(buf3[832:848], buf1[832:848], buf2[832:848])
		fastxor.Block(buf3[848:864], buf1[848:864], buf2[848:864])
		fastxor.Block(buf3[864:880], buf1[864:880], buf2[864:880])
		fastxor.Block(buf3[880:896], buf1[880:896], buf2[880:896])
		fastxor.Block(buf3[896:912], buf1[896:912], buf2[896:912])
		fastxor.Block(buf3[912:928], buf1[912:928], buf2[912:928])
		fastxor.Block(buf3[928:944], buf1[928:944], buf2[928:944])
		fastxor.Block(buf3[944:960], buf1[944:960], buf2[944:960])
		fastxor.Block(buf3[960:976], buf1[960:976], buf2[960:976])
		fastxor.Block(buf3[976:992], buf1[976:992], buf2[976:992])
		fastxor.Block(buf3[992:1008], buf1[992:1008], buf2[992:1008])
		fastxor.Block(buf3[1008:1024], buf1[1008:1024], buf2[1008:1024])
		_, err = o.Write(buf3)
		if err != nil {
			log.Println(err)
			return
		}
	}
	if v := int(fi.Size() - batch*1024); v != 0 {
		buf1 = buf1[0:v]
		buf2 = buf2[0:v]
		buf3 = buf3[0:v]

		in1.Read(buf1)
		in2.Read(buf2)

		fastxor.Bytes(buf3, buf1, buf2)
		_, err = o.Write(buf3)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
