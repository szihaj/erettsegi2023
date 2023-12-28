package main

import (
    "io"
    "os"
    "fmt"
    "bytes"
    "bufio"
    "slices"
    "strconv"
    "strings"
)

var lanes []byte
var numberOfLanes int
var throws []byte
var players int
var lane int
var fields []string
var selectedLane string

func main() {
    task1()
    task2()
    task3()
    task4()
    task5()
    task6()
    task7()
    task8()
}

func task1() {
    laneFile := "Informatika_forras_Egy2313/Forras/4_Tarsas/osvenyek.txt"
    lanes, numberOfLanes = readFile(laneFile)

    throwsFile := "Informatika_forras_Egy2313/Forras/4_Tarsas/dobasok.txt"
    throws, _ = readFile(throwsFile)
}

func task2() {
    fmt.Println("2. feladat")

    fmt.Print("A dobások száma: ")
    fmt.Println(len(bytes.Fields(throws)))

    fmt.Print("Az ösvények száma: ")
    fmt.Println(numberOfLanes)
}

func task3() {
    lineSeparator := []byte{'\n'}

    max := 0
    iMax := 0
    tmp := bytes.Split(lanes, lineSeparator)
    for i := 0; i < len(tmp); i++ {
        length := len(tmp[i])
        if length > max {
            max = length
            iMax = i
        }
    }

    fmt.Println("")
    fmt.Println("3. feladat")
    fmt.Println(fmt.Sprintf("Az egyik (első) leghosszabb a(z) %d. ösvény, hossza: %d", iMax + 1, max - 1))
}

func task4() {
    lane = readInt("Adja meg az ösvény sorszámát: ");
    for {
        players = readInt("Adja meg a játékosok számát: ");
        if (players >= 2 && players <= 5) {
            break
        }
    }

    lineSeparator := []byte{'\n'}
    tmp := bytes.Split(lanes, lineSeparator)
    selectedLane = string(tmp[lane-1])
    fields = strings.Split(selectedLane, "")
}

func task5() {
    type TypeCounter struct {
        M int
        V int
        E int
    }

    counter := TypeCounter{0, 0, 0}

    // for _, char := range selectedLane {
    //     chr := string(char)
    for _, chr := range fields {
        switch {
        case chr == "M":
            counter.M += 1
        case chr == "V":
            counter.V += 1
        case chr == "E":
            counter.E += 1
        }
    }

    fmt.Println("")
    fmt.Println("5. feladat")
    if counter.M > 0 {
        fmt.Println(fmt.Sprintf("M: %d", counter.M))
    }
    if counter.V > 0 {
        fmt.Println(fmt.Sprintf("V: %d", counter.V))
    }
    if counter.E > 0 {
        fmt.Println(fmt.Sprintf("E: %d", counter.E))
    }
}

func task6() {
    f, _ := os.Create("./kulonleges.txt")
    // io.WriteString(f, "malac")

    // for i, char := range selectedLane {
    //     chr := string(char)
    for i, chr := range fields {
        if (chr == "V" || chr == "E") {
            io.WriteString(f, fmt.Sprintf("%d\t%s\n", i, chr));
        }
    }

    f.Close()
}

func task7() {
    turn := 1
    pos := make([]int, players)
    curPlayer := 1

    for _, c := range bytes.Fields(throws) {
        tmp := string(c)
        throw, _ := strconv.Atoi(tmp)

        if (curPlayer > players) {
            if (slices.Max(pos) >= len(selectedLane)-1) {
                break
            }
            curPlayer = 1
            turn += 1
        }

        pos[curPlayer-1] += throw
        curPlayer += 1
    }

    maxI := 0;
    for i := 1; i < len(pos); i++ {
        if (pos[i] > pos[maxI]) {
            maxI = i
        }
    }

    fmt.Println("")
    fmt.Println("7. feladat")
    fmt.Println(fmt.Sprintf("A játék a(z) %d. körben fejeződött be. A legtávolabb jutó(k) sorszáma: %d", turn, maxI+1))
}

func task8() {
    turn := 1
    curPlayer := 1
    pos := make([]int, players)

    maxFields := len(fields)

    for _, c := range bytes.Fields(throws) {
        tmp := string(c)
        throw, _ := strconv.Atoi(tmp)

        if (curPlayer > players) {
            if (slices.Max(pos) >= len(selectedLane)-1) {
                break
            }
            curPlayer = 1
            turn += 1
        }

        field := "M"
        tmpPos := pos[curPlayer-1] + throw

        if (tmpPos < maxFields) {
            field = fields[tmpPos-1]
        }

        if (field == "E") {
            pos[curPlayer-1] += throw * 2
        } else if (field == "M") {
            pos[curPlayer-1] += throw
        }

        curPlayer += 1
    }


    fmt.Println("")
    fmt.Println("8. feladat")

    var winners []int
    var others []int
    for i := 0; i < len(pos); i++ {
        if (pos[i] > maxFields) {
            winners = append(winners, i+1)
            continue
        }

        others = append(others, i)
    }

    fmt.Print("A nyertes(ek): ")
    for _, i := range winners {
        fmt.Print(i, " ")
    }
    fmt.Println("")
    fmt.Println("A többiek pozíciója:")
    for _, i := range others {
        fmt.Println(fmt.Sprintf("%d. játékos, %d. mező", i+1, pos[i]))
    }
}

func readInt(label string) int {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(label)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")

    intInput, err := strconv.Atoi(input)

    if err != nil {
        fmt.Println("Hiba az érték beolvasésa közben.")
    }

    return intInput
}

func readFile(filename string) ([]byte, int) {
    f, _ := os.Open(filename)
    r := bufio.NewReader(f)

    var contents []byte
    fPos := 0
    lines := 0

    for i := 1; ; i++ {
        line, err := r.ReadBytes('\n')

        if err != nil {
            break
        }
        fPos += len(line)
        lines += 1

        contents = append(contents, line...)
    }

    return contents, lines
}

