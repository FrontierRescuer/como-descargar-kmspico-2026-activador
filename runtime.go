package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.7" )

func 7GjAQmaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X1Zlm9eiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rUHNK1iPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7I4OEq6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pjGJRbg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xICoTZf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1oknA2NnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FWIFaSOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dv4JdH74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 35SEGbTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lYU0bPFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xCSzZbvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vBvPWrGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ThXYRXpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nu5QbKdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YQN8QCUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cU8qzZEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l3rAYBJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yhgMMULUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cSe6r8hjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L99v4szzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g3Gl4uVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 13kGlxj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 104dSMNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DDdpH7buWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xhj7NqMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dKFpM5kuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TkJWgLonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dzRU464iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HSwB2Yk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8xjJTTc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SkUd9TAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NT1WsbG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lPj89jQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9gtPabbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BjvVEmJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZWFndwVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j2FolVjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OGhj2kFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ar2H3PnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIT0KJA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1pPXSgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12aFzKkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mEC2F0n1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ahfePyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 92wEureUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qpKdI6cvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vNo6oQUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rRDrO1KuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lSgJIIxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L9ccKhrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qeguCrVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rNV8itQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UKRUzCrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func avlbx8f1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CxgLogrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vqoD1z2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xgbA1A7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kDmQszvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ry4DnfWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H9dzuPDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6sSsOb9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func twpIJsRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nIl4UoTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qtzEccvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THUgwLzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func STZZCJldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z0fhY6NxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ktIhgFCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uajQHufBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sq5xOAbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VyZvvvxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func avYdJEphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PJecRDZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pFdo0h1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func llbYojbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pu8NMGIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xotVUSbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9siEFPSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YAOMniJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OrrFiEv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7OksTCRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLDn38IEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LFdvcebeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qxUjhb84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7dNTqziXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uw664mD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Po8cOpbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eHDY3YPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dFE1mxtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wWTFkVyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lOmaCqNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rN7ZNgqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GN869jwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6h7c6IknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kwQvo8PFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BnY35sgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8OMi2R39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dsymwkXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N0rTlxMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 191pYbZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aJ0xzX9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SS3NMTIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FzSeOh3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F3jIOfkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bXk2r9ptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HiCkedi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ymiVyUXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7sZyAwCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJELseYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jcyhX2rWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 153DtnqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NTAaxnZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GFfQhGN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CyoW3djbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8kpwDAoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qdkfiWFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L1vHeZ0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zRpnGWwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nblXdf99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rlzECBrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func se3v4p34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f5iwps81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ERqw30X9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KE6qDtItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5vS8M32nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dqYS8OD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 89BZ7tnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nS84VMGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qarWcqXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bgdbh2uHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xv52LzVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V185FVusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qXbJaVdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mB8Gu7PPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NnEZsWDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m7nxHwyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NljrLvIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cksGHNJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mPezbhYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pbomghdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mo7L8q6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UIv89N6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NKKo3O9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ltNzgBDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XxvzbF0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tnKNKFH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j9qjx62gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8d3wvTUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gAZpveS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oUFSaOuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GiNPpLdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOHWoPCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7SMaFFNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sEzXBlSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yle3e4MmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gj4hH16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ko1gZ4oEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qBtdqmD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n3bxLcwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FwDjC4GhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zes6K0cpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RPCYGpumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mvd93Mm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ew41fZcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VRtgOUxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G1mK80MVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VNQElFxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gB3nH8jeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NwO3KI5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pxDLcvZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nSPYfJ8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ToqL2CdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KzxhhfnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zX8qfzXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 00Tj0hppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func akuuhFfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sGxPh05AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hR3bhSYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vfpg08y6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oNirDdJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ALMKVACzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t2eX4soOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xiRruWzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YKrvKXWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I20K1izvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oFPiQQo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Q97yC7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2MX0AX2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WoBDUzvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TgnKyJmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CsmtKlp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vASPIOdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qRbekzyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGafvqQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GI6ubjtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F4VLpxj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S4x3vvvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uCDco9g5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nHPCVOwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MRCpRWzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Du9C3OZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hDkiL1NYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ltZAAmAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lDmqrr5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hJ8l0I3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FmOGsbYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PfQDvB7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jjyb4bD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mKa6Tf51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yoXvYHU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLFABNuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rZwmus6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gt6Gz20WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func quEhKs0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9r5m167HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AKTCyx8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dRPybsd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HPNmgbpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l1S1SGz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FF42XpFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g9rnbEUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MzhWWl3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DKAaOTxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7HhErEBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6YCLfb7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UrarxPbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yVfjAf0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UPHVlaubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p73JW91jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hkmOBfyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func InK1SrufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w6gWlCtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OUgXLvMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BPcb3XDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eYxwdw97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vxix8AYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JXtAIegzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWginSGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0s72tH5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMwohUoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vNfKz2N1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CV8Qy1xrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sPvs6Q3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lgwoCdRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AzPBCGVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kpWY68lEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RqiGNJtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BrYn9PLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func npw3QNA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZHhYweQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uFEmdL3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1wLYL1UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OwlcFdjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b1jxF53aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5imol3wMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func viTBWyrgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GJwSiRGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3eC3LAgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iCypY3zAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 23Lyc2wsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cKZQ4uaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zY6vJeYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MlfV7xgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fVu7hxpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qENYSTooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k94BbeGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zHhTMfxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JQbBrs21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CF9fIn73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xp5yNQLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvenpwVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kMn5zvo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XJQYtR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VFkehmrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQSzkONuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VESiqEaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wEr5V9ZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CMNG50sAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EkjQLYapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P0gcFuqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfEOuOLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NT8AsHigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AtS2apcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dHHQBYe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WlLq9dunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RecAYnf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func djnAdRV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H4lrphJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oXsB56YVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OttQr6LyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2sbnvgO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d2pULTpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ZLBLOy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ktK3d9coWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4v21VTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iTe4Yb9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8lrFMV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ScUR7XjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GSbbzE0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rKmRMbxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OzIBxb9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iBqCKdRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4S0otgV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aZBNLSTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IljllqrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4JoPTXB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vRVd2Wb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AoOdJcjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PtHlSBFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VF0sSzenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IY2nmaztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QNgGqvkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I9I2tkMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AIQfD3DJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qfSvZfbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BcNxa6uQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QHKdSGraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHEiJHG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4dZ2GjugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bwowx3htWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6zsZuv9TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyRuzh9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gwTx1nEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mtfMMWt5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GynRZdEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ia5pfdwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ujNqj0SEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4BOMLHlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5qUTaG6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fyudj8ieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YNHVpQzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FCzFFiFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxahf5qvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZKYWTnqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GKWr3nRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hSQya9SOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0tGhlrEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n7PzosCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func igBR9KZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func brEJBQH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func anmcNWqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SlkjzJSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YPwoXfJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dpKtxMqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Newblf0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 95Iqfh5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NbTemO0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GEWAQOtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sLzQnxm8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4xU2JwnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzyJqoQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHhEvBlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hcuQE90AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfmHCcYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kQHO8KeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8DHBexWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jYQsA2ErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K0wG0t4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mdmPaz9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QD5PkkVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5cWfjbpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UwyYMN9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8sIAIY9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TzR75moIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DZCxU1CYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6m4zAkPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uhYIHitUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ULflT5zNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 76s7wpU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5PKGj7o2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mjQTTsGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2hKqgNNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5GtQi7cLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8LNTI2D1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K0UIc0tyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func slroBzjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mb0GUxi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gtqa7UDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B0LBXll3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func seE9baAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z5FrwKFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9EkoTzvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tYajf7VHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rBjANB0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZ9Ew6ZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SOlzwiACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Apuz1f7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func unh1UnfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4qY51E0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P3bzqrjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8x24f0XZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hdx6Gh7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BFiv25jbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cSD7ubyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ycYsuF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9mMs2pRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5jNG1jbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ka02fnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ijFPtrfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jvoy0H3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6yQT94UpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b18ARBuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4EOreiJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zVFj1e4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func piGYX8pGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nVkTnsLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YM13zj9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7LHvKEIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lo3ggFrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JoAd5y7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ESyqTtvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9rM5COBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LeON5fobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8uwzG7jSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rEi0iYaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rLYEkrC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oreAxnVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xuTrHk5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 42uDubYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d5O5zbzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 201Nt91jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jbS2IRTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yM4QI4nXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7SxAqpjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yDw3y2YlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pqIvsMheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0A62DlLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CPjC8iO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ra2CEKNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vg4KITthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aAR5btbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UfhG5OBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5TUZ8qZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xcs1vuGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aRyBlvClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8z9oSJLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLKvepSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tVjgE0OfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DiQjTGqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 58WYtUNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1WwmwxRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DgjMwv3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b03vBfRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z0t1KADiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PXA7QpJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mAPHbfqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bC6fCYLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RjPdDDChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ueoVQ6DdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OEQRKjukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z2E5yC12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gm7qFgZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xSgMwuWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1MNs785TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7FSjIfxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jEFkw1qAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 17XuTtjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yDDTBXDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2v0UtSmTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxhoFx3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xl6M8vPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OmGnUzg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hHiFfntcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uYkbckhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FszFpBHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lUzuFzyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7j35NQlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ygTJ0AemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mAXKfz6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4em1AibaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YDS8RSgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bPswyT0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func itTvGrU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mq3SQRqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ey4f63obWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GncyIUGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XtmvBiGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0F9ckl6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aN1y83lyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wPyk7zkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CYd3sMljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RQgDuQoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3V5KYehCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5l16LE0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0gCCRZoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2EKYakyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sLfg9ZTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mu6Skj8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TpXyLeOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vvrnDyg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ar298hh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qAO6lay0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qp2W2N9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B48rL18VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XNafwAFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SUDtEhHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GWorWkYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func asgWwx5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5SlSBZrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 84FeTOsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L2RyHfI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dVWJUihlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wVBGLtMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0CVrJCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AVCkul7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GGRx8IJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EG6um28HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3r7tsC1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pNCDI5XrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ThtA7q11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TrmLMpo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PzjzNlVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nb1x4W67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zqTwoXfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gjjVcnZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sbjM5wNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 08V6PIr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VXFdOsyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func koBQORQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ExmisUkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DXsGfaNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uKTrrbHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d1clenCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 099VjkLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E5cvZB6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Myd6uqnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bnfaQ9dTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zhzp0WmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FM30UwJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7jJp5L1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Zbh4R6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xDnBi8mgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LdzN3oKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8WK88UqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gv4PUrojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SkGaJWTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rVVQLawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 61HfZlS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kL6KhaC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func obYZna6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s1hn8UumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nv5BiLNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0t4FoUEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c86Fom4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1r3hiJIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Avr8K9cgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nMeqUm6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3cdxvGAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Uds1TzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 51bqwKmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gXBrBjclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxplSBJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0GV1ke5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zrEdYsaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2mYLAFewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gJYw4oQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NNijPqEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dLSLkZqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 44LYJmOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WDG8ba3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C2AJ93v0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hfTamxcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GohnARgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VTFlo24bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vbfLcrNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DvbusS3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kOECSea8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i99BCmfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FqS9dXUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wy920LPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aVib42hAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C2qS5VoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ms3ko8sIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A5ZwebRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SxpiCMqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z31J6rMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y37V0ZPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BldtdwHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RMtDK2LPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0aO7EqyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1zmHCGw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q7Cw3LXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sRJKW9SWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GlaUifCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JMZbReh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wjudmA55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9g4RhTDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6V3R6MoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P6HTCsFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qm1UtCntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9yLgFLceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zXikpui1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AaaXR57DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func agNOkTroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZfEDMTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1UPdlYkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZSnGJRRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dZTdT7ZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XQhorLMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K8wUPigQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l2R63DbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JCiRRqhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PKTXX3ZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RCF3Ns5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IJCRL0Z7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mKK4sE2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ic33ly7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jHYQb2P1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdNqPSfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eOsvd3GYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FvpBGNF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7p9BTSzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zA7258onWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gHetlVXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Upcu93hFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l7OgDC3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ptNwL3B4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XsSlcZGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3adtbtuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yCxI6qhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kHZJJmi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r2oiGKyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LzccrTC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func caxTcUQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RmZ7PbX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hh2FaiYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XTCXJubbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xGIAsmWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wHBYRJiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WezttIGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opiiZilaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NHMXtaJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZDcbKnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j9POXR2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func guJxH10qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DfkZgmTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bD2ZmqJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1sCCv0fUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AYzMWdOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Pom1zi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ispZlsvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V7sPjQtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBGoZzJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MlfIEsC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UzoPwBdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zSL1j9qqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RnjZHUbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QzES54yfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jJyqag7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ayDjkwanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vVoGu0O5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WZAWQo1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DUkxO99XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fe3a31I6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CA81nflFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FerZvvtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SPrA4kisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0jv3Q0EsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U6DiaAHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func huDjApKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PDMZfd2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RI1StTyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uK7Ycf4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OkhIAvwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u1iXhsYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F5UvTTU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V3UtHlGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pm1ZiHT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MPjNSiixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gjdk7g4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jeCTezqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wAymUKbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KtE7Zk2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YZfRSPucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DVjTQrhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZoUP4muFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TOk6CfLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8dwbZpXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YximKNUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1YONhWgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Czay9iADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8wp0XdzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EjxeB6hYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8mj9M9ETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dJkVrR9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rm8tebJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZXJIQ9LFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8BmIo5L8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y35zaSWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 90E0Js0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2RqHIBy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tW5tiltEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5PFwqXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l0w5GD3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T33fm4EeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4eLNSrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNSz7umhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QwuLwYvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T0QPpXeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kA1jF4fmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W37X5fuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lx5mmR9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eSKAiGOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qsIuPhbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3kJ3sOrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZNKnrTWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j9QcqflJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QgTJFnRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3qsaKSshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YIKIdYuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ewgxjDPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TEnSw3wbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BWlKqaPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mcFNvADQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NvqCvD1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cW3gjNy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xO3cBaLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XgjBqEYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fj5qNIStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NQgjyEezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gtbeiGjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6WGxh32CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ptirEvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5bXtBb4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kiI1zPNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jEBYfVkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 66H5OPEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZTHn000pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uTVt3egfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UcQP0eKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

