package model

import "github.com/jinzhu/gorm"

func AutoMigrate(db *gorm.DB) error {

	// Check if the custom enum type exists
	var typeStatusExists bool
	if err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum')").Row().Scan(&typeStatusExists); err != nil {
		return err
	}

	var typeRoleExists bool
	if err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum')").Row().Scan(&typeRoleExists); err != nil {
		return err
	}

	// If the type does not exist, create it
	if !typeStatusExists {
		if err := db.Exec("CREATE TYPE status_enum AS ENUM ('pending', 'active')").Error; err != nil {
			return err
		}
	}

	if !typeRoleExists {
		if err := db.Exec("CREATE TYPE role_enum AS ENUM ('admin', 'guru')").Error; err != nil {
			return err
		}
		if err := db.Exec("ALTER TABLE gurus ALTER COLUMN role SET DEFAULT 'guru'").Error; err != nil {
			return err
		}
	}

	// AutoMigrate other models
	if err := db.AutoMigrate(&Guru{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Kelas{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Tugas{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&UploadTugas{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Jawaban{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Mapel{}).Error; err != nil {
		return err
	}

	var fkUsersJawaban bool
	if err := db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'jawabans_siswa_nis_users_nis_foreign' )").Row().Scan(&fkUsersJawaban); err != nil {
		return err
	}

	if !fkUsersJawaban {
		if err := db.Model(&Jawaban{}).AddForeignKey("siswa_nis", "users(NIS)", "SET NULL", "CASCADE").Error; err != nil {
			return err
		}
	}

	var fkGuruTugas bool
	if err := db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'tugas_walas_nis_gurus_nis_foreign' )").Row().Scan(&fkGuruTugas); err != nil {
		return err
	}
	if !fkGuruTugas {
		if err := db.Model(&Tugas{}).AddForeignKey("walas_nis", "gurus(NIS)", "SET NULL", "CASCADE").Error; err != nil {
			return err
		}
	}
	var fkJawabanTugas bool
	if err := db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'tugas_id_jawaban_jawabans_id_jawaban_foreign' AND table_name = 'tugas' )").Row().Scan(&fkJawabanTugas); err != nil {
		return err
	}

	if !fkJawabanTugas {
		if err := db.Model(&Tugas{}).AddForeignKey("id_jawaban", "jawabans(Id_jawaban)", "SET NULL", "CASCADE").Error; err != nil {
			return err
		}
	}

	var constraintExists bool
	if err := db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'kelas_walas_nis_gurus_nis_foreign' AND table_name = 'kelas')").Row().Scan(&constraintExists); err != nil {
		return err
	}

	if !constraintExists {
		if err := db.Model(&Kelas{}).AddForeignKey("walas_nis", "gurus(NIS)", "CASCADE", "CASCADE").Error; err != nil {
			return err
		}
	}

	var fkKelasUser bool
	if err := db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'users_kelas_kelas_id_kelas_foreign' AND table_name = 'users')").Row().Scan(&fkKelasUser); err != nil {
		return err
	}

	if !fkKelasUser {
		if err := db.Model(&User{}).AddForeignKey("kelas", "kelas(Id_kelas)", "SET NULL", "CASCADE").Error; err != nil {
			return err
		}

	}

	// Add index for Guru model
	db.Model(&Guru{}).AddIndex("idx_gurus_role", "role")

	return nil
}
