package helpers


import (
	"bufio"
	"os"
)


func ReadLines(path string)([]string, error) {
  
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, nil
}

func Sum(slice []int)(int) {
  sum := 0
  for _, val := range slice {
    sum += val
  }
  return sum 
}
