package main

import (
  "fmt"
  "strconv"
  "strings"
  "os"
  "bufio"
)

func getDirection(current_direction string, action string) string{
  action = strings.ToUpper(action)
  current_direction = strings.ToUpper(current_direction)

  var cur_direction_index int
  clockwise :=  []string{0:"N",1:"E",2:"S",3:"W"}

  for i := range clockwise{
    if clockwise[i] == current_direction {
      cur_direction_index = i
      break
    }
  }

  // if cur_direction_index == nil {
  //   fmt.Println("WRONG INPUT IN START DIRECTION! PLEASE TRY AGAIN.")
  //   os.Exit(0)
  // }

  if action == "L" {
    cur_direction_index -= 1
    if cur_direction_index < 0 {
      cur_direction_index += 4
    }
  }else if action == "R"{
    cur_direction_index += 1
    if cur_direction_index > 3 {
      cur_direction_index -= 4
    }
  }

  return clockwise[cur_direction_index]
}

func runExplorer(start_pos string, commands string, x_max int, y_max int){

  default_commands := []string{"L","R","M"}

  start_pos = strings.ReplaceAll(start_pos, " ","")
  commands = strings.ReplaceAll(commands, " ","")

  x, err_x := strconv.Atoi(string(start_pos[0]))
  y, err_y := strconv.Atoi(string(start_pos[1]))
  direction := strings.ToUpper(string(start_pos[2]))

  if err_x != nil || err_y !=nil {
    fmt.Println("WRONG INPUT ON START POSITION! PLEASE TRY AGAIN.")
    os.Exit(0)
  }

  for i := range commands{
    if !inList(string(commands[i]),default_commands) {
      fmt.Println("WRONG INPUT ON ACTION COMMANDS! PLEASE TRY AGAIN.")
      os.Exit(0)
    }
    switch strings.ToUpper(string(commands[i])){
    case "L":
      direction = getDirection(direction, "L")
    case "R":
      direction = getDirection(direction, "R")
    case "M":
      switch strings.ToUpper(direction){
      case "N":
        y += 1
      case "S":
        y -= 1
      case "E":
        x += 1
      case "W":
        x -= 1
      }
    }
  }

  if x > x_max || x < 0 || y > y_max || y < 0 {
    fmt.Println("EXPLORER OUT OF BORDER! PLEASE TRY AGIAN.")
    os.Exit(0)
  }

  fmt.Printf("%d %d %s\n",x,y,direction)
}

func inList(needle string, array []string) bool {
  for i := range array{
    if strings.ToUpper(needle) == strings.ToUpper(array[i]) {
      return true
    }
  }
  return false
}


func main(){
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("PLEASE ENTER YOUR COMMANDS:")
  input_count := 0
  var x_max, y_max int
  var err_x, err_y error
  var exp_start_pos, exp_commands string
  explorer_commands_map := make(map[string]string)

  for scanner.Scan(){
    input := strings.ToUpper(strings.ReplaceAll(scanner.Text()," ",""))
    if input == "DONE" {
      break;
    }
    if input_count == 0 {
      x_max, err_x = strconv.Atoi(string(input[0]))
      y_max, err_y = strconv.Atoi(string(input[1]))

      if err_x != nil || err_y != nil || x_max <= 0 || y_max <=0 {
        fmt.Println("WRONG INPUT ON BORDER POSITIONS! PLEASE TRY AGAIN.")
        os.Exit(0)
      }
      input_count += 1
      continue
    }

    if input_count % 2 == 1 {
      exp_start_pos = input
    }else if input_count % 2 == 0{
      exp_commands = input
      explorer_commands_map[exp_start_pos] = exp_commands
    }

    input_count += 1
  }

  if input_count % 2 == 0 {
    fmt.Println("WRONG INPUT ON YOUR COMMANDS INPUT! PLEASE TRY AGAIN.")
    os.Exit(0)
  }

  fmt.Println("\nOUTPUT:")
  for k, v := range explorer_commands_map{
    runExplorer(k,v,x_max,y_max)
  }
}
