# Architecture Decisions — Day 20 AWS Fundamentals

---

## Decision 1: Dual-Mode App (In-Memory + PostgreSQL)

**What:** The Go app runs in two modes — in-memory (no DB) or PostgreSQL — controlled by the `DATABASE_URL` environment variable.

**Why:** Allows the app to run locally with zero setup for testing and development, while being production-ready for AWS RDS deployment. The handler checks `if h.db != nil` to choose the storage path.

**Trade-off:** Slight code duplication in handlers (two code paths). Worth it for the flexibility — the student can demo the app immediately without PostgreSQL, then connect to RDS when deploying.

---

## Decision 2: ap-south-1 (Mumbai) Region

**What:** All AWS resources default to the Mumbai region.

**Why:** Closest region to India, where the training is likely based. Lower latency for development and testing. Can be changed by modifying the REGION variable in scripts.

---

## Decision 3: VPC with 4 Subnets (2 Public + 2 Private)

**What:** VPC with /16 CIDR, split into 2 public subnets (for EC2/ALB) and 2 private subnets (for RDS).

**Why:**
- Public subnets have internet access via Internet Gateway → needed for EC2 and ALB
- Private subnets have no direct internet access → RDS is only reachable from EC2
- 2 subnets per type across 2 AZs → required for RDS Multi-AZ and ALB

---

## Decision 4: Three Security Groups (ALB, EC2, RDS)

**What:** Separate security groups for each tier.

**Why:** Defense in depth:
- ALB SG: allows port 80 from internet (0.0.0.0/0)
- EC2 SG: allows port 8080 only from ALB SG, SSH (22) for management
- RDS SG: allows port 5432 only from EC2 SG
This means RDS is never directly accessible from the internet.

---

## Decision 5: IAM Role Instead of Access Keys

**What:** EC2 instance uses an IAM Instance Profile/Role for AWS permissions, not hardcoded access keys.

**Why:** Security best practice. Access keys in code or environment variables can leak. IAM roles are automatically rotated by AWS and never exposed in the application.

**Permissions granted:** S3 read/write (for reports), CloudWatch log writing, RDS connect.

---

## Decision 6: RDS db.t3.micro with No Public Access

**What:** Smallest RDS instance, in private subnet, not publicly accessible.

**Why:**
- db.t3.micro is free-tier eligible — no cost for training
- Private subnet + no public access = database is only reachable from within the VPC
- The Go app on EC2 connects via the private RDS endpoint

---

## Decision 7: S3 for Report Storage

**What:** Employee CSV reports are uploaded to S3 via the AWS SDK.

**Why:** S3 is durable (99.999999999%), cheap, and scalable. Better than storing files on EC2 (ephemeral). The S3 handler gracefully degrades — if AWS credentials aren't configured (local mode), it returns a simulated response showing what the command would do.

---

## Decision 8: systemd for EC2 Process Management

**What:** The Go binary runs as a systemd service on EC2.

**Why:** systemd automatically:
- Starts the app on boot
- Restarts on crash (Restart=always)
- Manages environment variables
- Provides log output via `journalctl`
This is the standard way to run services on Linux, not `nohup` or `screen`.

---

## Decision 9: CloudWatch for Logging

**What:** Two CloudWatch log groups: `/employee-api/application` and `/employee-api/access`.

**Why:** EC2 instance logs are lost when the instance terminates. CloudWatch provides persistent, searchable, alertable log storage. In production, the CloudWatch agent on EC2 ships logs automatically.

---

## Decision 10: AWS CLI Scripts Instead of Console

**What:** All AWS resources are created via CLI scripts, not the AWS Console.

**Why:**
- Reproducible — run the script to recreate everything
- Version-controllable — scripts are committed to git
- Auditable — every resource creation is documented
- Faster than clicking through the console

**Alternative:** Terraform (also included conceptually in the architecture). For training, CLI scripts are more transparent about what each AWS API call does.

---

## Summary: AWS Service Mapping

| AWS Service | Task | Purpose |
|-------------|------|---------|
| EC2 | 5, 6 | Run the Go application |
| VPC | 7 | Network isolation |
| Security Groups | 8 | Firewall rules per tier |
| S3 | 9, 10 | Store employee reports |
| RDS PostgreSQL | 11, 12 | Persistent data storage |
| CloudWatch | 13 | Centralized logging |
| IAM | 4 | Permissions for EC2 |
| ALB | 14 | Load balancing + HA |
