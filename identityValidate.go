package main

import (
	"fmt"
	"strconv"
)

var tc int
var sabit int = 29999
var count int = 0
var istenenTCsayisi int = 0

func main() {
	println("Istenen Tc sayisini giriniz:")
	fmt.Scan(&istenenTCsayisi)
	tc = 10000000100
	dogrulamaYap(tc)
}

func dogrulamaYap(tc int) {
	var sorgu1 int = 0
	var durum bool
	tcBasamaklari := make([]int, 11)
	for i := 10; i >= 0; i-- {
		tcBasamaklari[i] = tc % 10
		tc = tc / 10
	}
	ilkSorgu := tcBasamaklari[:10]
	for _, v := range ilkSorgu {
		sorgu1 += v
	}
	durum = sorgu1%10 == tcBasamaklari[10]
	if durum {
		dogrulamaAsama2(tcBasamaklari)
	} else {
		uretimAsamasi(tcBasamaklari)
	}
}
func dogrulamaAsama2(tcBasamaklari []int) {
	var sorgu2 int = 0
	var durum bool
	sorgu2 = tcBasamaklari[0] + tcBasamaklari[2] + tcBasamaklari[4] + tcBasamaklari[6] + tcBasamaklari[8]
	sorgu2 *= 7
	sorgu2 -= tcBasamaklari[1] + tcBasamaklari[3] + tcBasamaklari[5] + tcBasamaklari[7]
	durum = sorgu2%10 == tcBasamaklari[9]
	if durum {
		count++
		println(tc)
		if count == istenenTCsayisi {
			fmt.Printf("%v adet tc dogrulandi", count)
		} else {
			uretimAsamasi(tcBasamaklari)
		}
	} else {
		fmt.Println("TC Asama2 Doğrulanamadı")
		uretimAsamasi(tcBasamaklari)
	}
}
func uretimAsamasi(tcBasamaklari []int) {
	var ilkDokuzBasamaklari []int = make([]int, 9)
	var strTC string = ""
	var icTC int
	var islemTC int
	var islemYapilanDokuzBasamak []int = make([]int, 9)
	var onuncuBasamak int
	var onbirinciBasamak int
	ilkDokuzBasamaklari = tcBasamaklari[:9]

	for i := 0; i < len(ilkDokuzBasamaklari); i++ {
		strTC += strconv.Itoa(ilkDokuzBasamaklari[i])
	}
	icTC, _ = strconv.Atoi(strTC)
	icTC += sabit
	//basamak sayisi 9 olacak eger gecerse tekrar basamak sayisi 9 olacak
	if icTC > 999999999 {
		icTC = icTC % 1000000000
	}
	islemTC = icTC
	for i := 0; i <= 8; i++ {
		islemYapilanDokuzBasamak[i] = islemTC % 10
		islemTC = islemTC / 10
	}
	//10.Basamak bu olmali bu gelen veriye gore
	onuncuBasamak = islemYapilanDokuzBasamak[0] + islemYapilanDokuzBasamak[2] + islemYapilanDokuzBasamak[4] + islemYapilanDokuzBasamak[6] + islemYapilanDokuzBasamak[8]
	onuncuBasamak *= 7
	onuncuBasamak -= islemYapilanDokuzBasamak[1] + islemYapilanDokuzBasamak[3] + islemYapilanDokuzBasamak[5] + islemYapilanDokuzBasamak[7]
	onuncuBasamak %= 10
	//11.Basamak bu olmali bu gelen veriye gore
	onbirinciBasamak = islemYapilanDokuzBasamak[0] + islemYapilanDokuzBasamak[1] + islemYapilanDokuzBasamak[2] + islemYapilanDokuzBasamak[3] + islemYapilanDokuzBasamak[4] + islemYapilanDokuzBasamak[5] + islemYapilanDokuzBasamak[6] + islemYapilanDokuzBasamak[7] + islemYapilanDokuzBasamak[8]
	onbirinciBasamak = (onuncuBasamak + onbirinciBasamak) % 10
	var strTCNUMBER string = strconv.Itoa(icTC)
	strTCNUMBER += strconv.Itoa(onuncuBasamak)
	strTCNUMBER += strconv.Itoa(onbirinciBasamak)
	//println(strTCNUMBER)
	tc, _ = strconv.Atoi(strTCNUMBER)
	dogrulamaYap(tc)
}
