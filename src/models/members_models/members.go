package members_models

import (
	"api_tugas_minggu4/src/config"

	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Member_name string
	Email       string
	Password    string
	Role        string
	Address     string
	Phone       string
}

// ////////////////////////
func SelectAll_member() *gorm.DB {
	items := []Member{}
	return config.DB.Find(&items)
}

func Select_member(id string) *gorm.DB {
	var item Member
	return config.DB.First(&item, "id = ?", id)
}

// ////////////////////////////////////////////
func Create_member(item *Member) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates_member_customer(id string, newCustomer *Member) *gorm.DB {
	var item Member
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newCustomer)
}

func Updates_member_seller(id string, newSeller *Member) *gorm.DB {
	var item Member
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newSeller)
}

// ///////////////////////////////////////
func Deletes_members(id string) *gorm.DB {
	var item Member
	return config.DB.Delete(&item, "id = ?", id)
}

func FindEmail(input *Member) []Member {
	items := []Member{}
	config.DB.Raw("SELECT * FROM members WHERE email = ?", input.Email).Scan(&items)
	return items
}
