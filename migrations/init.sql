-- CREATE TABLE IF NOT EXISTS users (
--     id SERIAL PRIMARY KEY,
--     email VARCHAR(255) UNIQUE NOT NULL,
--     password_hash VARCHAR(255) NOT NULL,
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
-- );

CREATE TABLE IF NOT EXISTS Users (
    Id SERIAL PRIMARY KEY,
    PasswordHash VARCHAR(255) NOT NULL,
    Name VARCHAR(255),
    Gender VARCHAR(20) NOT NULL,
    Mail VARCHAR(255) UNIQUE NOT NULL,
    Age INTEGER CHECK (Age > 0),
    Height NUMERIC CHECK (Height > 0),
    Weight NUMERIC CHECK (Weight > 0),
    GoalExercise TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS PlanWeek (
    Id SERIAL PRIMARY KEY,
    UserId INTEGER NOT NULL REFERENCES Users(Id) ON DELETE CASCADE,
    Current BOOLEAN -- Один True, остальные False
);

CREATE TABLE IF NOT EXISTS Day (
    Id SERIAL PRIMARY KEY,
    PlanId INTEGER NOT NULL REFERENCES PlanWeek(Id) ON DELETE CASCADE,
    DayWeek INTEGER NOT NULL CHECK (DayWeek BETWEEN 1 AND 7),
    TypeExercise TEXT -- Храним тип тренировки
    ColoriesAll TEXT -- Храним калории на день
);

CREATE TABLE IF NOT EXISTS Diet (
    Id SERIAL PRIMARY KEY,
    DayId INTEGER NOT NULL REFERENCES Day(Id) ON DELETE CASCADE,
    MealType  TEXT,
    Name VARCHAR(255) NOT NULL,
    Structure TEXT, -- Если хранить состав еды
    Colories VARCHAR(200),
);

CREATE TABLE IF NOT EXISTS Exercise (
    Id SERIAL PRIMARY KEY,
    DayId INTEGER NOT NULL REFERENCES Day(Id) ON DELETE CASCADE,
    Name TEXT
    Info TEXT
    Done BOOLEAN  
);