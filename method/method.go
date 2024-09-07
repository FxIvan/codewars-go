package method

import(
	_"fmt"
	"strings"
	"github.com/codewars/library"
)

type ProxyValue struct {
    Delete func(prop []string, valueDeleted string) ([]string, error)
    Set    func(prop []string, valueAdd string) ([]string, error)
}

type StringSlice []string

func LastElement(arr []string) string{
	return arr[len(arr)-1]
}

func PopElement(arr []string) []string{
	if len(arr) == 0 {
		return arr
	}
	return arr[:len(arr)-1]
}

func Unshift(arr []string, valueUnshift string) []string{
	return append([]string{valueUnshift}, arr...)
}

func EncodeURI(url string) string{
	uri := strings.Split(url, "/")

	uriCompleted := uri[len(uri)-(len(uri) - 3):] //len(uri) - 3 -> no toma en cuenta el dominio pero si el path

	var urlEncode string

	for _,value := range uriCompleted{
		word := strings.Split(value, "")
		for i , letter := range word{
            word[i] = library.GetValueNumber(letter)
		}

		urlEncode += strings.Join(word, "") + "/"
	}

	return urlEncode
}

func (sl StringSlice) ConcatArray(arr StringSlice) StringSlice{
	return append(sl, arr...)
}

func NewProxyValue() *ProxyValue {
    return &ProxyValue{
        Delete: func(prop []string, valueDeleted string) ([]string, error) {
            var newArray []string
            for _, value := range prop {
                if value != valueDeleted {
                    newArray = append(newArray, value)
                }
            }
            return newArray, nil
        },
        Set: func(prop []string, valueAdd string) ([]string, error) {
            return append(prop, valueAdd), nil
        },
    }
}