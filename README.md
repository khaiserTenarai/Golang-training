<div align="center">

<a href="https://github.com/khaiserTenarai/Golang-training">
  <img src="https://readme-typing-svg.demolab.com?font=Fira+Code&size=28&pause=1000&color=00ADD8&center=true&vCenter=true&width=600&lines=Golang+Training+%F0%9F%9A%80;Day+1+%E2%80%A2+Day+2+%E2%80%A2+Day+3+%E2%80%A2+Day+4;Naman+Kumar" alt="Typing SVG" />
</a>

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-in--progress-yellow?style=for-the-badge)
![Days](https://img.shields.io/badge/days--completed-4%2F4-brightgreen?style=for-the-badge)

</div>

---

## 📖 About

This repository contains my **Go training assignments**, organized day by day. Each day builds on core Go concepts — from setup and tooling to language fundamentals and beyond.

## 📁 Project Structure

```text
Golang-training/
├── Day1NamanKumar/
│   └── go-day1/
│       ├── 01-go-installation/
│       ├── 02-hello-cloud-native/
│       ├── 03-cli-calculator/
│       ├── 04-go-toolchain/
│       ├── 05-go-module/
│       ├── 06-external-package/
│       ├── 07-package-design/
│       ├── 08-go-documentation/
│       ├── 09-versioning/
│       ├── 10-environment-variables/
│       ├── 11-cli-employee-search/
│       ├── 12-project-structure/
│       ├── 13-build-automation/
│       ├── 14-troubleshooting/
│       └── 15-mini-project-employee-management/
├── Day2NamanKumar/
│   └── day2/
│       ├── cmd/          # Runnable examples: variables, datatypes, loops, etc.
│       └── internal/     # Supporting packages: employee, loops, salary, textstats
├── Day3NamanKumar/
│   └── employee-project/
│       ├── employee.go
│       ├── report.go
│       ├── main.go
│       └── go.mod
└── Day4/
    └── employee-management/
        └── employee-management/
            ├── main.go
            ├── go.mod
            ├── model/
            │   └── employee.go
            ├── service/
            │   └── employee.go
            └── utility/
                ├── string.go
                └── validation.go
```

## 🗓️ Daily Breakdown

<details>
<summary><strong>📌 Day 1 — Go Fundamentals & Tooling</strong></summary>
<br>

Covers installation, modules, packages, documentation, versioning, environment variables, project structure, build automation, troubleshooting, and a mini employee-management project.

</details>

<details>
<summary><strong>📌 Day 2 — Language Basics</strong></summary>
<br>

Covers variables, data types, zero values, type conversion, operators, strings, runes, slices, maps, switch statements, loops, and scope — plus a small employee CLI exercise.

</details>

<details>
<summary><strong>📌 Day 3 — Employee Management (Intro)</strong></summary>
<br>

A small Go project (`employee-project`) that models employees and generates simple reports, introducing struct-based data modeling and basic package organization.

</details>

<details>
<summary><strong>📌 Day 4 — Employee Management (Structured)</strong></summary>
<br>

A more structured employee management project, split into layers: `model` for data structures, `service` for business logic, and `utility` for helper functions like string handling and validation.

</details>

## 🚀 Running the Code

Each exercise folder has its own `go.mod`. To run an example:

```bash
cd Day1NamanKumar/go-day1/<folder-name>
go run main.go
```

For Day 4:

```bash
cd Day4/employee-management/employee-management
go run main.go
```

To run tests in a folder that has them:

```bash
go test ./...
```

## 🛠️ Tech Stack

![Go](https://img.shields.io/badge/-Go-00ADD8?style=flat-square&logo=go&logoColor=white)
![Git](https://img.shields.io/badge/-Git-F05032?style=flat-square&logo=git&logoColor=white)
![GitHub](https://img.shields.io/badge/-GitHub-181717?style=flat-square&logo=github&logoColor=white)

## 📝 Notes

This repo is part of a Go training program. Each day's folder is self-contained with its own module and, where applicable, a `README.md` explaining that day's exercises in more detail.

---

<div align="center">
  <sub>Built with 💙 while learning Go</sub>
</div>
