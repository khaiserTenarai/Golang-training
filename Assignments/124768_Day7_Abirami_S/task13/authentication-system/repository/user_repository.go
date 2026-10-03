package repository
import("context";"q13-authentication-system/model";"github.com/jackc/pgx/v5";"github.com/jackc/pgx/v5/pgxpool")
type UserRepository interface{CreateUser(model.User)error;GetByUsername(string)(*model.User,error)}
type UserRepositoryImpl struct{db *pgxpool.Pool}
func NewUserRepository(db *pgxpool.Pool)UserRepository{return &UserRepositoryImpl{db}}
func(r *UserRepositoryImpl)CreateUser(u model.User)error{_,e:=r.db.Exec(context.Background(),`INSERT INTO users(username,password_hash,role)VALUES($1,$2,$3)`,u.Username,u.PasswordHash,u.Role);return e}
func(r *UserRepositoryImpl)GetByUsername(n string)(*model.User,error){var u model.User;e:=r.db.QueryRow(context.Background(),`SELECT id,username,password_hash,role FROM users WHERE username=$1`,n).Scan(&u.ID,&u.Username,&u.PasswordHash,&u.Role);if e==pgx.ErrNoRows{return nil,e};return &u,e}
