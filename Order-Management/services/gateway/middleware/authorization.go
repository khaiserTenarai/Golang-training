package middleware

import authmiddleware "order-management/services/auth/middleware"

// RequireAuth validates the bearer token.
var RequireAuth = authmiddleware.RequireAuth

// RequireRole checks whether the authenticated user has the required role.
var RequireRole = authmiddleware.RequireRole
