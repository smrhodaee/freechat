package configs

import "fmt"

const (
	chaptaKey = "chapta:%s" // "chapta:<uuid>"
)

func GetChaptaKey(uuid string) string {
	return fmt.Sprintf(chaptaKey, uuid)
}