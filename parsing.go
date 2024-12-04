package main

import (
	"bufio"	
	"fmt"
	"os"
	"strconv"
	"strings"
)
func ValidateRoomName(name string) error {
	if len(name) == 0 || len(name) > 100 {
		return fmt.Errorf(" Error: invalid data format")
	}
	if name[0] == 'L' || name[0] == '#' || strings.Contains(name, " ") {
		return fmt.Errorf(" Error: invalid data format, room name must not start with 'L', '#' or contain spaces")
	}
	return nil
}

func ParseAnts(line string, colony *Colony) error {
	ants, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf(" Error: invalid data format, number of ants must be an integer")
	}
	if ants <= 0 || ants > 100000 {
		return fmt.Errorf(" Error: invalid data format, number of ants must be between 1 and 100000")
	}
	colony.ants = ants
	return nil
}

func ParseRoom(line string, colony *Colony, roomType string) error {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return fmt.Errorf(" Error: invalid data format, room must be 'name x y'")
	}
	if err := ValidateRoomName(parts[0]); err != nil {
		return err
	}
	x, y := parts[1], parts[2]
	
	// Validate coordinates
	for _, existing := range colony.rooms {
		if existing.x == x && existing.y == y {
			return fmt.Errorf(" Error: invalid data format, room with coordinates (%s, %s) already exists", x, y)
		}
	}

	room := &Room{
		name:    parts[0],
		x:       x,
		y:       y,
		links:   make([]*Room, 0, 4),
		isStart: roomType == "start",
		isEnd:   roomType == "end",
	}

	if _, exists := colony.rooms[room.name]; exists {
		return fmt.Errorf(" Error: invalid data format ,room with name %s already exists", room.name)
	}
	colony.rooms[room.name] = room

	if room.isStart {
		if colony.start != nil {
			return fmt.Errorf(" Error: invalid data format, multiple start rooms defined")
		}
		colony.start = room
	}
	if room.isEnd {
		if colony.end != nil {
			return fmt.Errorf(" Error: invalid data format,multiple end rooms defined")
		}
		colony.end = room
	}
	return nil
}

func parseLink(line string, colony *Colony) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf(" Error: invalid data format, link must be like 'room1-room2'")
	}

	room1, ok1 := colony.rooms[parts[0]]
	room2, ok2 := colony.rooms[parts[1]]
	if !ok1 || !ok2 {
		return fmt.Errorf(" Error: invalid data format, unknown rooms in link")
	}

	if parts[0] == parts[1] {
		return fmt.Errorf(" Error: invalid data format, room cannot link to itself")
	}

	for _, link := range room1.links {
		if link == room2 {
			return fmt.Errorf(" Error: invalid data format, link between %s and %s already exists", parts[0], parts[1])
		}
	}

	room1.links = append(room1.links, room2)
	room2.links = append(room2.links, room1)
	return nil
}
//the function ReadMap reads the file into a struct a.k.a colony that we can use later on to get the params we need like ants rooms..
func ReadMap(fileName string) (*Colony, error) { 
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf(" Error: file %v not found", fileName)
	}
	defer file.Close()

	colony := &Colony{
		rooms: make(map[string]*Room),
	}
	scanner := bufio.NewScanner(file)
	step := 0
	var nextRoomType string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		switch {
		case step == 0:
			if err := ParseAnts(line, colony); err != nil {
				return nil, err
			}
			step = 1
		case line == "##start":
			if nextRoomType != "" {
				return nil, fmt.Errorf(" Error: invalid data format, multiple special room markers")
			}
			nextRoomType = "start"
		case line == "##end":
			if nextRoomType != "" {
				return nil, fmt.Errorf(" Error: invalid data format, multiple special room markers")
			}
			nextRoomType = "end"
		case step == 1:
			if strings.Contains(line, "-") {
				step = 2
				if err := parseLink(line, colony); err != nil {
					return nil, err
				}
				continue
			}
			if err := ParseRoom(line, colony, nextRoomType); err != nil {
				return nil, err
			}
			nextRoomType = ""
		case step == 2:
			if err := parseLink(line, colony); err != nil {
				return nil, err
			}
		}
	}

	if colony.start == nil {
		return nil, fmt.Errorf(" Error: invalid data format, no start room defined")
	}
	if colony.end == nil {
		return nil, fmt.Errorf(" Error: invalid data format, no end room defined")
	}
	return colony, nil
}