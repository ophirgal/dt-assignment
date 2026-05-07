package migration

import (
	"embed"
	"fmt"
	"io/fs"
	"sort"

	"github.com/ophirgal/dt-assignment/backend/internal/model"

	"gorm.io/gorm"
)

//go:embed seeds/*.sql
var seedFiles embed.FS

func Run(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.Chain{},
		&model.Store{},
		&model.Product{},
		&model.Sale{},
		&model.Forecast{},
	); err != nil {
		return fmt.Errorf("automigrate: %w", err)
	}

	entries, err := fs.ReadDir(seedFiles, "seeds")
	if err != nil {
		return fmt.Errorf("read seeds dir: %w", err)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	for _, entry := range entries {
		data, err := seedFiles.ReadFile("seeds/" + entry.Name())
		if err != nil {
			return fmt.Errorf("read seed %s: %w", entry.Name(), err)
		}
		if err := db.Exec(string(data)).Error; err != nil {
			return fmt.Errorf("exec seed %s: %w", entry.Name(), err)
		}
	}
	return nil
}
