package main

import (
	"encoding/json"
    "fmt"
)

// User 用户类
type User struct {
    Name string
    Website string
    Age int
    Male bool
    Skills []string
}

func main()  {
    user:= User{
        "Emily",
        "https://www.baidu.com",
        18,
        true,
        []string{"Golang","java"},
    }

    // json解析
    u, err := json.Marshal(user)
    if err != nil {
        fmt.Printf("user json encoder failed: %v\n", err)
        return
    }

    fmt.Printf("user json encoder result: %s\n", u)

    var user2 User

    // json反序列化
    err = json.Unmarshal(u, &user2)
    if err != nil {
        fmt.Printf("user2 json decode failed: %v\n", err)
        return
    }

    fmt.Printf("user2 json decode result: %#v\n", user2)

    // 未知结构的json数据
    u3 := []byte(`{"name":"Emily","website":"https://www.sina.com.cn","age":5, "skills":["playing","singing"],"male":true}`)
    var user3 interface{}
    err = json.Unmarshal(u3, &user3)
    if err != nil {
        fmt.Printf("user3 json decode failed: %v\n", err)
        return
    }
    fmt.Printf("user3 json decode result: %#v\n", user3)
    // 处理未知格式的JSON对象
    user5, ok := user3.(map[string]interface{})
    if ok {
        for k, v := range user5 {
            // 按类型
            switch v2 := v.(type) {
            case string:
                fmt.Println(k, "is string ", v2)
            case int:
                fmt.Println(k, "is int ", v2)
            case bool:
                fmt.Println(k, "is bool", v2)
            case []interface{}:
                fmt.Println(k, " is an array:")
                for i, iv := range v2 {
                    fmt.Println(i, iv)
                }
            default:
                fmt.Println(k, "is another type not handle yet")
            }

        }
    }
}
