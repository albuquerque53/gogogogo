package main

import "fmt"

func main() {
    user := map[string]string {
        "name": "Albuquerque",
        "role": "Developer",
    }

    fmt.Println(user)

    moderator := map[string]map[string]bool {
        "permissions": {
            "ban": true,
            "timeout": true,
            "privateMessage": true,
        },
        "attributes": {
            "goldMod": false,
        },
    }

    fmt.Println(moderator)
    fmt.Println(moderator["permissions"])
    fmt.Println(moderator["permissions"]["timeout"])
    delete(moderator, "attributes")
    fmt.Println(moderator)

    moderator["userInfo"] = map[string]bool {
        "hasUser": false,
    }
    fmt.Println(moderator)
}
