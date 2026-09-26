package main


import "fmt"


func main() {


    var name string
    var choice int
    var continueChoice string


    for {


        fmt.Println("\n************** MAIN MENU **************")
        fmt.Println("1. Create Profile")
        fmt.Println("2. View Profile")
        fmt.Println("3. Delete Profile")
        fmt.Println("4. Update Profile")
        fmt.Println("5. View All Profiles")
        fmt.Println("6. Exit")


        fmt.Print("Enter your choice: ")
        fmt.Scan(&choice)


        switch choice {


        case 1:
            fmt.Print("Enter your name: ")
            fmt.Scan(&name)


            fmt.Println("Profile created successfully!")


        case 2:
            fmt.Println("Employee Name:", name)


        case 3:
            name = ""
            fmt.Println("Profile deleted successfully!")


        case 4:
            fmt.Print("Enter new name: ")
            fmt.Scan(&name)


            fmt.Println("Profile updated successfully!")


        case 5:
            fmt.Println("Employee Name:", name)


        case 6:
            fmt.Println("Thank you!")
            return


        default:
            fmt.Println("Wrong choice")
        }


        fmt.Print("\nDo you want to continue? (yes/no): ")
        fmt.Scan(&continueChoice)


        
        if continueChoice != "yes" {
            fmt.Println("Go Employee Management Application")
            break
        }
    }
}