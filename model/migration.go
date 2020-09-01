package model

func migration() {
	DB.AutoMigrate(
		User{},
		BillBoard{},
		Knowledge{},
		Community{},
		Refuge{},
		RefugeFacility{},
		Location{},
		DisasterType{},
		Disaster{},
		DisasterCommunity{},
	)
}
