package utils
import (
	"fmt"
	"os"
	"time"
)

func GenerateUniqueID() string {
	pid := os.Getpid()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%d_%d", pid, timestamp)
}
