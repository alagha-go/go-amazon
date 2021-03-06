package departments

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

const (
	DepartmentsPageUrl = "https://www.amazon.com/gp/navigation/ajax/generic.html?ajaxTemplate=hamburgerMainContent&pageType=Gateway&hmDataAjaxHint=1&navDeviceType=desktop&isSmile=0&isPrime=0&isBackup=false&hashCustomerAndSessionId=ffbb1810709292f57a2507d0669c66868eea6018&isExportMode=true&languageCode=en_US&environmentVFI=AmazonNavigationCards%2Fdevelopment%40B6079054483-AL2_x86_64&secondLayerTreeName=prm_digital_music_hawkfire%2Bkindle%2Bandroid_appstore%2Belectronics_exports%2Bcomputers_exports%2Bsbd_alexa_smart_home%2Barts_and_crafts_exports%2Bautomotive_exports%2Bbaby_exports%2Bbeauty_and_personal_care_exports%2Bwomens_fashion_exports%2Bmens_fashion_exports%2Bgirls_fashion_exports%2Bboys_fashion_exports%2Bhealth_and_household_exports%2Bhome_and_kitchen_exports%2Bindustrial_and_scientific_exports%2Bluggage_exports%2Bmovies_and_television_exports%2Bpet_supplies_exports%2Bsoftware_exports%2Bsports_and_outdoors_exports%2Btools_home_improvement_exports%2Btoys_games_exports%2Bvideo_games_exports%2Bgiftcards%2Bamazon_live%2BAmazon_Global"
	ProxyUrl           = "socks5://127.0.0.1:9050"
)

// collect main departments and their children
func CollectDepartments() Departments {
	collector := colly.NewCollector()
	err := collector.SetProxy(ProxyUrl)
	if err != nil {
		log.Println(err.Error())
		return Departments{}
	}
	var Departments Departments

	collector.OnHTML("body", func(element *colly.HTMLElement) {
		element.ForEach(".hmenu.hmenu-translateX-right", func(index int, element *colly.HTMLElement) {
			Departments = append(Departments, CollectDepartment(element))
		})
	})

	collector.Visit(DepartmentsPageUrl)
	return Departments
}

// collect one main department and its children
func CollectDepartment(element *colly.HTMLElement) Department {
	var department Department
	element.ForEach(".hmenu-item", func(index int, element *colly.HTMLElement) {
		if index >= 2 {
			var dep Department
			dep.Title = strings.TrimSpace(element.Text)
			url := element.Attr("href")
			if url != "" {
				if !strings.Contains(url, "https://") {
					dep.Url = "https://amazon.com" + url
				} else {
					dep.Url = url
				}
				department.Departments = append(department.Departments, dep)
			}
		} else if index == 1 {
			department.Title = strings.TrimSpace(element.Text)
		}
	})

	return department
}

// collect children of any department
func (dep *Department) CollectChildren() {
	collector := colly.NewCollector()
	err := collector.SetProxy(ProxyUrl)
	if err != nil {
		log.Println(err.Error())
		return
	}
	
	collector.OnError(func(r *colly.Response, err error){
		fmt.Println(r.StatusCode)
	})
	
	collector.OnHTML("#s-refinements", func(element *colly.HTMLElement) {
		element.ForEach("li", func(index int, element *colly.HTMLElement) {
			class := strings.TrimSpace(element.Attr("class"))
			if class == "a-spacing-micro apb-browse-refinements-indent-2" || class == "a-spacing-micro s-navigation-indent-2" {
				var department Department
				department.Url = element.ChildAttr("a", "href")
				if !strings.Contains(department.Url, "https://") {
					department.Url = "https://amazon.com" + department.Url
				}
				department.Title = strings.TrimSpace(element.ChildText("a > span"))
				dep.Departments = append(dep.Departments, department)
			}
		})
	})

	collector.Visit(dep.Url)
}