package main

import (
  "fmt"
  "strconv"
  "strings"
  "os"
  "bufio"
  "regexp"
)

// Get explorer direction from moving commands
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

//get explorer final position by start position, moving commands
func runExplorer(input_start_pos string, input_commands string, x_max int, y_max int){

  start_pos := strings.Fields(input_start_pos)
  commands := strings.ReplaceAll(input_commands, " ","")

  x, err_x := strconv.Atoi(string(start_pos[0]))
  y, err_y := strconv.Atoi(string(start_pos[1]))
  direction := strings.ToUpper(string(start_pos[2]))

  if err_x != nil || err_y !=nil {
    fmt.Println("Wrong input on landing location! Please try again.")
    os.Exit(0)
  }

  for i := range commands{
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
    fmt.Println("This Explorer is out of border!")
  }else{
    fmt.Printf("%d %d %s\n",x,y,direction)
  }
}


func main(){
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Please enter your input:")
  input_count := 0
  var x_max, y_max int
  var err_x, err_y error
  var exp_start_pos, exp_commands string
  explorer_commands_map := make(map[string]string)
  re_num :=regexp.MustCompile("[0-9]")
  re_lrm :=regexp.MustCompile("[LRM]")
  re_direction :=regexp.MustCompile("[NSEW]")

  for scanner.Scan(){
    input := strings.ToUpper(scanner.Text())
    // if the user input "Exit", the program will be terminated immediately
    if input == "EXIT" {
      os.Exit(0)
    }

    //"Done" or empty input indicate that the input is finished
    if input == "DONE" || len(input) == 0 {
      if input_count % 2 == 0 {
        fmt.Println("Plase enter the moving commands for your latest Explorer!")
        continue
      }
      break
    }

    //first input must be two number seperated by space indicate the exploring range
    if input_count == 0 {
      if len(re_num.FindAllString(input, -1)) == 0{
          fmt.Println("Wrong input on landing area initialisation! Please enter your landing range again.")
          continue
      }

      input := strings.Fields(input)

      if len(input) != 2 {
        fmt.Println("Wrong input on landing area initialisation! Please enter your landing range again.")
        continue
      }

      x_max, err_x = strconv.Atoi(string(input[0]))
      y_max, err_y = strconv.Atoi(string(input[1]))

      if err_x != nil || err_y != nil || x_max <= 0 || y_max <=0 {
        fmt.Println("Wrong input on landing area initialisation! Please enter your landing range again.")
        continue
      }

      input_count += 1
      continue
    }

    //explorer start position and moving commands must be input one by one
    if input_count % 2 == 1 {
      input_start_pos := strings.Fields(input)

      if len(input_start_pos) != 3 || len(re_num.FindAllString(input_start_pos[0], -1)) == 0 || len(re_num.FindAllString(input_start_pos[1], -1)) == 0 || len(re_direction.FindAllString(input_start_pos[2],-1)) == 0{
        fmt.Println("Wrong input on Explorer landing location! Please enter your Explorer landing location again.")
        continue
      }
      exp_start_pos = input
    }else if input_count % 2 == 0{
      if len(re_lrm.FindAllString(input,-1)) != len(input) {
        fmt.Println("Wrong input on Explorer moving commands! Please enter your Explorer moving commands again.")
        continue
      }
      exp_commands = input
      explorer_commands_map[exp_start_pos] = exp_commands
    }
    input_count += 1
  }

  fmt.Println("\nOutput:")
  explorer_count := 1
  for k, v := range explorer_commands_map{
    fmt.Println("Explorer",explorer_count, ":")
    runExplorer(k,v,x_max,y_max)
    explorer_count += 1
  }
}
