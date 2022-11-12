package fuzz

import "strconv"
import "github.com/ozontech/cute/asserts/json"
import "github.com/ozontech/cute/internal/utils"


func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            var test interface{}
            content := string(bytes)
            json.Contains(content, test)
            return 0

        case 1:
            var test interface{}
            content := string(bytes)
            json.Equal(content, test)
            return 0

        case 2:
            var test interface{}
            content := string(bytes)
            json.NotEqual(content, test)
            return 0

        case 3:
            test := num
            content := string(bytes)
            json.Length(content, test)
            return 0

        case 4:
            content := string(bytes)
            json.Present(content)
            return 0

        case 5:
            content := string(bytes)
            json.NotEmpty(content)
            return 0

        case 6:
            content := string(bytes)
            json.NotPresent(content)
            return 0

        case 7:
            json.GetValueFromJSON(bytes, "mayhem")
            return 0

        case 8:
            utils.ToJSON(bytes)
            return 0

        default:
            utils.PrettyJSON(bytes)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}