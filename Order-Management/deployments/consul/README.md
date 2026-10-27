# Native Consul on Windows (No Docker)

1. Install the Consul Windows binary and make sure `consul.exe` is on PATH.
2. From the project root, run `powershell -ExecutionPolicy Bypass -File .\scripts\start-consul.ps1`.
3. Keep the terminal open. Consul UI/API is at `http://localhost:8500`.
4. The `.env` setting `SERVICE_ADVERTISE_ADDRESS=127.0.0.1` is intended for services running on the same Windows machine.
5. Start the Go services after Consul. They register themselves and expose `/health/ready` checks.

The script uses Consul development mode, suitable only for local development/testing. For production, use a configured Consul agent/server deployment.
