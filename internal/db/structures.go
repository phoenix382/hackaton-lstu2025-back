package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PasswordHash string `gorm:"type:varchar(255);not null"`
	Name         string `gorm:"type:varchar(255);not null;default:''"`
	Gender       string `gorm:"type:varchar(20);not null;default:'other'"`
	Email         string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Age          *uint  `gorm:"type:NUMERIC;check:age > 0"`
	Height       *uint  `gorm:"type:NUMERIC"` // Указатель для optional значений
	Weight       *uint  `gorm:"type:NUMERIC"`
	GoalExercise string `gorm:"type:text;default:''"`
}

type PlanWeek struct {
	gorm.Model
	UserID  uint `gorm:"not null;index;constraint:OnDelete:CASCADE"`
	Current bool `gorm:"type:boolean;default:false"` // NOT NULL с дефолтом
}

type Day struct {
	gorm.Model
	PlanID       uint   `gorm:"not null;index;constraint:OnDelete:CASCADE"`
	DayWeek      uint   `gorm:"check:day_week BETWEEN 1 AND 7;not null"`
	GoalExercise string `gorm:"type:text;default:''"`
	CaloriesAll  string `gorm:"type:text;default:''"`
}

type Diet struct {
	gorm.Model
	DayID     uint   `gorm:"not null;index;constraint:OnDelete:CASCADE"`
	MealType  string `gorm:"type:text"`
	Name      string `gorm:"type:varchar(255);not null"`
	Structure string `gorm:"type:text"`
	Colories  string `gorm:"type:varchar(200)"`
}

type Exercise struct {
	gorm.Model
	DayID uint   `gorm:"not null;index;constraint:OnDelete:CASCADE"`
	Name  string `gorm:"type:text"`
	Info  string `gorm:"type:text"`
	Done  bool   `gorm:"type:boolean;default:false"`
}
