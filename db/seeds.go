package db

import "fmt"

func seedDatabase() {
	runMigrate()
	runSeed()
}

func migrateDatabase() {
	runMigrate()
}

func stringPointer(s string) *string {
	return &s
}

func runSeed() {
	sounds := []Sound{
		Sound{
			Name:   stringPointer("air-horn"),
			Path:   stringPointer("/sounds/airhorn"),
			Letter: stringPointer("a"),
			Groups: []Group{
				Group{
					Name: stringPointer("mf"),
				},
			},
		},
		Sound{
			Name:   stringPointer("this-is-america"),
			Path:   stringPointer("/sounds/this-is-america"),
			Letter: stringPointer("b"),
			Groups: []Group{
				Group{
					Name: stringPointer("pop-culture"),
				},
			},
		},
		Sound{
			Name:   stringPointer("alex-laugh"),
			Path:   stringPointer("/sounds/alex-laugh"),
			Letter: stringPointer("c"),
			Groups: []Group{
				Group{
					Name: stringPointer("sound-efects"),
				},
			},
		},
		Sound{
			Name:   stringPointer("emilie-laugh"),
			Path:   stringPointer("/sounds/emilie-laugh"),
			Letter: stringPointer("d"),
		},
		Sound{
			Name:   stringPointer("tech-startup"),
			Path:   stringPointer("/sounds/tech-startup"),
			Letter: stringPointer("e"),
		},
	}

	for _, s := range sounds {
		if err := db.Save(&s).Error; err != nil {
			fmt.Println("Error saving sound to DB ", err)
		}
	}
}

func runMigrate() {
	db.AutoMigrate(&Sound{}, &Group{})
}
