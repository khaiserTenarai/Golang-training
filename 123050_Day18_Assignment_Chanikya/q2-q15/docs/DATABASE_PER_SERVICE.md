# Database per Service

- Employee Service owns employee_db.
- Department Service owns department_db.
- Authentication Service owns auth_db.

Each service connects through GORM.

The employee table stores department_id as a reference.
It does not create a foreign key into department_db.

Each service performs migrations for its own tables.