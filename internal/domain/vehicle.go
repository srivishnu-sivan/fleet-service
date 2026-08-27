package domain

import "time"

type Vehicle struct {
	ID        int    `json:"id"`
	VIN       string `json:"vin"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
	CreatedAt time.Time `json:"created_at"`
}



type CreateVehicleRequest struct {
    VIN   string `json:"vin" binding:"required"`
    Model string `json:"model" binding:"required"`
    Year  int    `json:"year" binding:"required"`
}


/*

this is what I'm gonna do:
create a struct in domain folder named : login.go.
inside it, i'll be creating 
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
and
type LoginResponse struct {
	Token string `json:"token"`
}
inside repository folder
create a struct for db *pgpool.Poolconst names LoginRepository
and a constructor db 
and a method for LoginRepository with name : CreateUser which gonna have role_id, from "job_roles" table and since we already have CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role_id INT NOT NULL REFERENCES job_roles(role_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
I'll be using it to insert data of new users.
this is the phase-1... review my flow and tell me what i missed.. so i'll explain phase-2 which is handler and service layer

*/