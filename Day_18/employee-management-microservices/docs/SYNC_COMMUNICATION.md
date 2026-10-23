# Synchronous REST Communication

When creating an employee:

1. Client sends a request to the API Gateway.
2. Gateway forwards it to Employee Service.
3. Employee Service validates the request.
4. Employee Service calls Department Service.
5. Department Service checks the department ID.
6. Employee Service saves the employee if the department exists.
7. Gateway returns the response to the client.

HTTP calls use timeouts.