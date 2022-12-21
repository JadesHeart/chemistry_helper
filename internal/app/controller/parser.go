package controller

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"vitalic_project/internal/app/model"
	"vitalic_project/internal/app/repository"
)

func ParseConst() [109]*model.BalanceConst {
	var modelArray [109]*model.BalanceConst
	URL := "https://chemhelp.ru/handbook/tables/dissociation_constant/"
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s ", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	for i := 2; i < 111; i++ {
		modelArray[i-2] = &model.BalanceConst{
			ElName: doc.Find("#tablepress-1 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-1").Text(),
		}
	}
	for i := 2; i < 111; i++ {
		modelArray[i-2].Formula = doc.Find("#tablepress-1 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-2").Text()
	}
	for i := 2; i < 111; i++ {
		modelArray[i-2].FirstParam = doc.Find("#tablepress-1 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-3").Text()
	}
	for i := 2; i < 111; i++ {
		modelArray[i-2].SecondParam = doc.Find("#tablepress-1 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-4").Text()
	}
	return modelArray
}

func ParsePotential() [1051]*model.Potentials {
	var modelArray [1051]*model.Potentials
	URL := "https://chemhelp.ru/handbook/tables/redox_potentials/"
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s ", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	for i := 2; i < 1053; i++ {
		modelArray[i-2] = &model.Potentials{
			Number: doc.Find("#tablepress-6 > tbody > tr.row-" + strconv.Itoa(i) + "> td.column-1").Text(),
		}
	}
	for i := 2; i < 1053; i++ {
		modelArray[i-2].Symbol = doc.Find("#tablepress-6 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-2").Text()
	}
	for i := 2; i < 1053; i++ {

		modelArray[i-2].ElName = doc.Find("#tablepress-6 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-3").Text()
	}
	for i := 2; i < 1053; i++ {
		modelArray[i-2].HalfReactions = doc.Find("#tablepress-6 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-4").Text()
	}
	for i := 2; i < 1053; i++ {
		modelArray[i-2].LastParam = doc.Find("#tablepress-6 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-5").Text()
	}
	return modelArray
}

func ParseThermo() [1005]*model.TermodProp {
	var modelArray [1005]*model.TermodProp
	URL := "https://chemhelp.ru/handbook/tables/thermodynamic_data/"
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s ", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	for i := 2; i < 1007; i++ {
		modelArray[i-2] = &model.TermodProp{
			ElName: doc.Find("#tablepress-3 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-1").Text(),
		}
	}
	for i := 2; i < 1007; i++ {
		modelArray[i-2].Formula = doc.Find("#tablepress-3 > tbody > tr.row-" + strconv.Itoa(i) + "> td.column-2").Text()
	}
	for i := 2; i < 1007; i++ {
		modelArray[i-2].FirstParam = doc.Find("#tablepress-3 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-3").Text()
	}
	for i := 2; i < 1007; i++ {
		modelArray[i-2].SecondParam = doc.Find("#tablepress-3 > tbody > tr.row-" + strconv.Itoa(i) + "> td.column-4").Text()
	}
	for i := 2; i < 1007; i++ {
		modelArray[i-2].ThirdParam = doc.Find("	#tablepress-3 > tbody > tr.row-" + strconv.Itoa(i) + "> td.column-5").Text()
	}
	for i := 2; i < 1007; i++ {
		modelArray[i-2].FourthParam = doc.Find("		#tablepress-3 > tbody > tr.row-" + strconv.Itoa(i) + " > td.column-6").Text()
	}
	return modelArray
}

func ParseInstability() [74]*model.Instability {
	var modelArray [74]*model.Instability
	URL := "https://chemhelp.ru/handbook/tables/complex_ion_dissociation_constants/"
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s ", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	for i := 2; i < 76; i++ {
		modelArray[i-2] = &model.Instability{
			ElName: doc.Find("#tablepress-8 > tbody > tr.row-" + strconv.Itoa(i) + "> td.column-1").Text(),
		}
	}
	for i := 2; i < 76; i++ {

		modelArray[i-2].Ligand = doc.Find("#tablepress-8 > tbody > tr.row-" + strconv.Itoa(i) + "> td.column-2").Text()
	}
	for i := 2; i < 76; i++ {
		modelArray[i-2].Complex = doc.Find("#tablepress-8 > tbody > tr.row-" + strconv.Itoa(i) + "> td.column-3").Text()
	}
	for i := 2; i < 76; i++ {
		modelArray[i-2].LastParam = doc.Find("#tablepress-8 > tbody > tr.row-" + strconv.Itoa(i) + "> td.column-4").Text()
	}
	return modelArray
}

func FillingDatabase(repo *repository.MainRepo) {

	balanceArray := ParseConst()
	for _, balanceConst := range balanceArray {
		repo.BalanceRepo.Create(balanceConst)
	}

	potentialArray := ParsePotential()
	for _, potential := range potentialArray {
		repo.PotentialsRepo.Create(potential)
	}

	thermoArray := ParseThermo()
	for _, potential := range thermoArray {
		repo.TermodPropRepo.Create(potential)
	}

	instabilityArray := ParseInstability()
	for _, instability := range instabilityArray {
		repo.InstabilityRepo.CreateEl(instability)
	}

}
