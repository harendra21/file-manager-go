package models

import (
	"beego/lib"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Site struct {
	Id         int       `orm:"auto"`
	Title      string    `orm:"size(100)" json:"title" valid:"Required"`
	Domain     string    `orm:"size(100)" json:"domain"`
	WpUserName string    `orm:"size(20)" json:"wpUsername"`
	WpEmail    string    `orm:"size(50)" json:"wpEmail"`
	WpPassword string    `orm:"size(50)" json:"wpPassword"`
	WpDatabase string    `orm:"size(50)" json:"wpDatabase"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
}

// GetAllSites function gets all sites
func GetAllSites() []*Site {
	o := orm.NewOrm()
	var sites []*Site
	o.QueryTable(new(Site)).All(&sites)

	return sites
}

// InsertOneSite inserts a single new User record
func InsertOneSite(site Site) *Site {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Site))

	// get prepared statement
	i, _ := qs.PrepareInsert()
	var s Site
	// get now datetime
	site.CreatedAt = time.Now()
	site.UpdatedAt = time.Now()

	fmt.Println("***************")
	fmt.Println(site.Domain)
	// out, err := exec.Command("cd .. && ls").Output()

	out, err := lib.RunCommand("ls -la")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(out)
	// Insert
	id, err := i.Insert(&site)

	if err == nil {
		// successfully inserted
		s = Site{Id: int(id)}
		err := o.Read(&s)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		fmt.Println(err)
		return nil
	}

	return &s
}

// UpdateSite updates an existing user
func UpdateSite(site Site) *Site {
	o := orm.NewOrm()
	s := Site{Id: site.Id}
	var updatedSite Site

	// get existing site
	if o.Read(&s) == nil {
		// updated site
		site.UpdatedAt = time.Now()
		s = site
		_, err := o.Update(&s)
		// read updated site
		if err == nil {
			// update successful
			updatedSite = Site{Id: site.Id}
			o.Read(&updatedSite)
		}
	}

	return &updatedSite
}

// DeleteSite deletes a site
func DeleteSite(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Site{Id: id})
	if err == nil {
		// successfull
		return true
	}

	return false
}

// GetSiteById gets a site with the given id
func GetSiteById(id int) *Site {
	o := orm.NewOrm()
	site := Site{Id: id}
	o.Read(&site)
	return &site
}
