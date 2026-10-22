#!/bin/bash
# Day 20 Mini Project (Task 15): Complete AWS deployment script
# Tasks 4-14 combined into one executable script.
# Prerequisites: AWS CLI configured (aws configure)
# Region: ap-south-1 (Mumbai) — change as needed

set -e

REGION="ap-south-1"
VPC_CIDR="10.0.0.0/16"
PUBLIC_SUBNET_1="10.0.1.0/24"
PUBLIC_SUBNET_2="10.0.2.0/24"
PRIVATE_SUBNET_1="10.0.3.0/24"
PRIVATE_SUBNET_2="10.0.4.0/24"
KEY_NAME="day20-piyush-key"
DB_USER="postgres"
DB_PASS="Piyush2024Secure"
DB_NAME="day20db"
S3_BUCKET="day20-piyush-employee-reports"
APP_NAME="employee-api"

echo "===== Day 20: AWS Deployment ====="

# ---- Task 7: Create VPC ----
echo "[Task 7] Creating VPC..."
VPC_ID=$(aws ec2 create-vpc --cidr-block $VPC_CIDR --region $REGION \
  --tag-specifications "ResourceType=vpc,Tags=[{Key=Name,Value=day20-vpc}]" \
  --query 'Vpc.VpcId' --output text)
echo "  VPC: $VPC_ID"

aws ec2 modify-vpc-attribute --vpc-id $VPC_ID --enable-dns-hostnames '{"Value":true}' --region $REGION

# Internet Gateway
IGW_ID=$(aws ec2 create-internet-gateway --region $REGION \
  --query 'InternetGateway.InternetGatewayId' --output text)
aws ec2 attach-internet-gateway --vpc-id $VPC_ID --internet-gateway-id $IGW_ID --region $REGION

# Subnets
echo "[Task 7] Creating subnets..."
PUB_SUB1=$(aws ec2 create-subnet --vpc-id $VPC_ID --cidr-block $PUBLIC_SUBNET_1 \
  --availability-zone ${REGION}a --region $REGION --query 'Subnet.SubnetId' --output text)
PUB_SUB2=$(aws ec2 create-subnet --vpc-id $VPC_ID --cidr-block $PUBLIC_SUBNET_2 \
  --availability-zone ${REGION}b --region $REGION --query 'Subnet.SubnetId' --output text)
PRV_SUB1=$(aws ec2 create-subnet --vpc-id $VPC_ID --cidr-block $PRIVATE_SUBNET_1 \
  --availability-zone ${REGION}a --region $REGION --query 'Subnet.SubnetId' --output text)
PRV_SUB2=$(aws ec2 create-subnet --vpc-id $VPC_ID --cidr-block $PRIVATE_SUBNET_2 \
  --availability-zone ${REGION}b --region $REGION --query 'Subnet.SubnetId' --output text)

# Route table for public subnets
RT_ID=$(aws ec2 create-route-table --vpc-id $VPC_ID --region $REGION --query 'RouteTable.RouteTableId' --output text)
aws ec2 create-route --route-table-id $RT_ID --destination-cidr-block 0.0.0.0/0 --gateway-id $IGW_ID --region $REGION
aws ec2 associate-route-table --route-table-id $RT_ID --subnet-id $PUB_SUB1 --region $REGION
aws ec2 associate-route-table --route-table-id $RT_ID --subnet-id $PUB_SUB2 --region $REGION

# ---- Task 8: Security Groups ----
echo "[Task 8] Creating Security Groups..."
ALB_SG=$(aws ec2 create-security-group --group-name day20-alb-sg --description "ALB SG" \
  --vpc-id $VPC_ID --region $REGION --query 'GroupId' --output text)
aws ec2 authorize-security-group-ingress --group-id $ALB_SG --protocol tcp --port 80 --cidr 0.0.0.0/0 --region $REGION

EC2_SG=$(aws ec2 create-security-group --group-name day20-ec2-sg --description "EC2 SG" \
  --vpc-id $VPC_ID --region $REGION --query 'GroupId' --output text)
aws ec2 authorize-security-group-ingress --group-id $EC2_SG --protocol tcp --port 8080 \
  --source-group $ALB_SG --region $REGION
aws ec2 authorize-security-group-ingress --group-id $EC2_SG --protocol tcp --port 22 \
  --cidr 0.0.0.0/0 --region $REGION

RDS_SG=$(aws ec2 create-security-group --group-name day20-rds-sg --description "RDS SG" \
  --vpc-id $VPC_ID --region $REGION --query 'GroupId' --output text)
aws ec2 authorize-security-group-ingress --group-id $RDS_SG --protocol tcp --port 5432 \
  --source-group $EC2_SG --region $REGION

# ---- Task 9: Create S3 Bucket ----
echo "[Task 9] Creating S3 bucket..."
aws s3 mb s3://$S3_BUCKET --region $REGION 2>/dev/null || echo "  Bucket may already exist"

# ---- Task 4: Create IAM Role ----
echo "[Task 4] Creating IAM role..."
aws iam create-role --role-name day20-ec2-role \
  --assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Service":"ec2.amazonaws.com"},"Action":"sts:AssumeRole"}]}' 2>/dev/null || true
aws iam put-role-policy --role-name day20-ec2-role --policy-name day20-policy \
  --policy-document file://iam-policy.json
aws iam create-instance-profile --instance-profile-name day20-ec2-profile 2>/dev/null || true
aws iam add-role-to-instance-profile --instance-profile-name day20-ec2-profile --role-name day20-ec2-role 2>/dev/null || true

# ---- Task 11: Create RDS PostgreSQL ----
echo "[Task 11] Creating RDS PostgreSQL..."
aws rds create-db-subnet-group --db-subnet-group-name day20-db-subnets \
  --db-subnet-group-description "Day20 DB subnets" \
  --subnet-ids $PRV_SUB1 $PRV_SUB2 --region $REGION 2>/dev/null || true

aws rds create-db-instance \
  --db-instance-identifier day20-postgres \
  --db-instance-class db.t3.micro \
  --engine postgres --engine-version 15 \
  --master-username $DB_USER --master-user-password $DB_PASS \
  --db-name $DB_NAME --allocated-storage 20 \
  --vpc-security-group-ids $RDS_SG \
  --db-subnet-group-name day20-db-subnets \
  --no-publicly-accessible \
  --region $REGION 2>/dev/null || echo "  RDS may already exist"

echo "[Task 11] Waiting for RDS to be available (this takes 5-10 minutes)..."
aws rds wait db-instance-available --db-instance-identifier day20-postgres --region $REGION

RDS_ENDPOINT=$(aws rds describe-db-instances --db-instance-identifier day20-postgres \
  --region $REGION --query 'DBInstances[0].Endpoint.Address' --output text)
echo "  RDS Endpoint: $RDS_ENDPOINT"

# ---- Tasks 5, 6: Deploy to EC2 ----
echo "[Tasks 5,6] Launching EC2 instance..."
USER_DATA=$(cat <<'USERDATA'
#!/bin/bash
yum update -y
yum install -y golang git
mkdir -p /opt/employee-api
cd /opt/employee-api
# In production, clone from git or copy the binary
cat > /etc/systemd/system/employee-api.service << 'EOF'
[Unit]
Description=Employee API
After=network.target
[Service]
Type=simple
User=ec2-user
Environment=PORT=8080
Environment=DATABASE_URL=postgres://DB_USER:DB_PASS@RDS_ENDPOINT:5432/DB_NAME?sslmode=require
Environment=AWS_REGION=REGION
Environment=S3_BUCKET=S3_BUCKET
WorkingDirectory=/opt/employee-api
ExecStart=/opt/employee-api/employee-api
Restart=always
[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload
systemctl enable employee-api
USERDATA
)

EC2_ID=$(aws ec2 run-instances \
  --image-id ami-0f58b397bc5c1f2e8 \
  --instance-type t2.micro \
  --key-name $KEY_NAME \
  --security-group-ids $EC2_SG \
  --subnet-id $PUB_SUB1 \
  --iam-instance-profile Name=day20-ec2-profile \
  --associate-public-ip-address \
  --user-data "$USER_DATA" \
  --tag-specifications "ResourceType=instance,Tags=[{Key=Name,Value=day20-employee-api}]" \
  --region $REGION \
  --query 'Instances[0].InstanceId' --output text)

echo "  EC2 Instance: $EC2_ID"

# ---- Task 13: CloudWatch Logging ----
echo "[Task 13] Creating CloudWatch log group..."
aws logs create-log-group --log-group-name /employee-api/application --region $REGION 2>/dev/null || true
aws logs create-log-group --log-group-name /employee-api/access --region $REGION 2>/dev/null || true

echo ""
echo "===== Deployment Complete ====="
echo "VPC:        $VPC_ID"
echo "EC2:        $EC2_ID"
echo "RDS:        $RDS_ENDPOINT"
echo "S3 Bucket:  $S3_BUCKET"
echo "DATABASE_URL: postgres://$DB_USER:$DB_PASS@$RDS_ENDPOINT:5432/$DB_NAME?sslmode=require"
