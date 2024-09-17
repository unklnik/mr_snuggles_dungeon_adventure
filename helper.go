package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
	"sync"
	"time"
)

// =========================================================================================================
// Math functions
// =========================================================================================================

func DistanceBetweenTwoPoints(v1, v2 rl.Vector2) float32 {
	dx := v2.X - v1.X
	dy := v2.Y - v1.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func PointOnLine(v1, v2 rl.Vector2, dist float32) rl.Vector2 {
	directionVector := rl.NewVector2(v2.X-v1.X, v2.Y-v1.Y)
	length := math.Sqrt(float64(directionVector.X*directionVector.X + directionVector.Y*directionVector.Y))
	normalizedVector := rl.NewVector2(directionVector.X/float32(length), directionVector.Y/float32(length))
	v3 := rl.NewVector2(v1.X+normalizedVector.X*dist, v1.Y+normalizedVector.X*dist)
	return v3
}

func MirrorAngle(angle float32) float32 { //ADD +90 +45 ETC TO CORRECT
	mirroredAngle := 180 - angle
	if mirroredAngle < 0 {
		mirroredAngle += 360
	}
	return mirroredAngle
}

func PointOnCircle(radius, angl float32, cntr rl.Vector2) rl.Vector2 {
	angleInRadians := float64(angl * math.Pi / 180.0)
	x := float32(float64(radius)*math.Cos(angleInRadians) + float64(cntr.X))
	y := float32(float64(radius)*math.Sin(angleInRadians) + float64(cntr.Y))
	return rl.NewVector2(x, y)
}

func FindRadius(center, v2 rl.Vector2) float32 {
	x := v2.X - center.X
	y := v2.Y - center.Y
	radius := math.Sqrt(float64(x*x) + float64(y*y))
	return float32(radius)
}

func AngleBetweenTwoPoints(v1, v2 rl.Vector2) float32 {
	deltaX := float64(v2.X) - float64(v1.X)
	deltaY := float64(v2.Y) - float64(v1.Y)
	radians := math.Atan2(deltaY, deltaX)
	degrees := (radians * 180) / math.Pi
	for degrees >= 360 {
		degrees -= 360
	}
	for degrees < 0 {
		degrees += 360
	}
	return float32(degrees)
}

func AbsDiff(a, b float32) float32 {
	return float32(math.Abs(float64(a - b)))
}

func Abs(value float32) float32 {
	return float32(math.Abs(float64(value)))
}

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandUint8(min, max int) uint8 {
	return (uint8)(min + rand.Intn(max-min))
}

func RandInt32(min, max int) int32 {
	return int32(min + rand.Intn(max-min))
}

func RandF32(min, max float32) float32 {
	tempMin := float64(min)
	tempMax := float64(max)
	return float32(tempMin + rand.Float64()*(tempMax-tempMin))
}

// =========================================================================================================
// Remove functions
// =========================================================================================================

func RemoveItem(slice []xitm, s int) []xitm {
	return append(slice[:s], slice[s+1:]...)
}
func RemoveWeapon(slice []xweap, s int) []xweap {
	return append(slice[:s], slice[s+1:]...)
}
func RemoveTile(slice []xtile, s int) []xtile {
	return append(slice[:s], slice[s+1:]...)
}
func RemoveFx(slice []xfx, s int) []xfx {
	return append(slice[:s], slice[s+1:]...)
}
func RemoveEnemy(slice []xenm, s int) []xenm {
	return append(slice[:s], slice[s+1:]...)
}
func RemoveRectangle(slice []rl.Rectangle, s int) []rl.Rectangle {
	return append(slice[:s], slice[s+1:]...)
}

//=========================================================================================================
// RNG functions
//=========================================================================================================

// Create a global random number generator
var (
	rng      = rand.New(rand.NewSource(time.Now().UnixNano()))
	rngMutex sync.Mutex // Mutex to ensure thread safety
)

func FlipCoin() bool {
	rngMutex.Lock()         // Lock to ensure safe access to the RNG
	defer rngMutex.Unlock() // Unlock after the function completes

	// Generate a random number (0 or 1)
	if rng.Intn(2) == 0 {
		return true
	}
	return false
}

func Roll6() int {
	return RandInt(1, 7)
}
func Roll12() int {
	return RandInt(1, 13)
}
func Roll18() int {
	return RandInt(1, 19)
}
func Roll24() int {
	return RandInt(1, 25)
}
func Roll30() int {
	return RandInt(1, 31)
}
func Roll36() int {
	return RandInt(1, 37)
}

// =========================================================================================================
// Color functions
// =========================================================================================================

func RandColor() rl.Color {
	return rl.NewColor(RandUint8(0, 256), RandUint8(0, 256), RandUint8(0, 256), 255)
}

func RandDarkGreen() rl.Color {
	return rl.NewColor(RandUint8(0, 30), RandUint8(40, 90), RandUint8(0, 40), 255)
}

func RandGreen() rl.Color {
	return rl.NewColor(RandUint8(0, 60), RandUint8(140, 256), uint8(RandInt(0, 60)), 255)
}

func RandRed() rl.Color {
	return rl.NewColor(RandUint8(140, 256), RandUint8(0, 60), RandUint8(0, 60), 255)
}

func RandPink() rl.Color {
	return rl.NewColor(RandUint8(200, 256), RandUint8(10, 110), RandUint8(130, 180), 255)
}

func RandBlue() rl.Color {
	return rl.NewColor(RandUint8(0, 180), RandUint8(0, 180), RandUint8(140, 256), 255)
}

func RandDarkBlue() rl.Color {
	return rl.NewColor(RandUint8(0, 20), RandUint8(0, 20), RandUint8(100, 160), 255)
}

func RandCyan() rl.Color {
	return rl.NewColor(RandUint8(0, 120), RandUint8(200, 256), RandUint8(150, 256), 255)
}

func RandYellow() rl.Color {
	return rl.NewColor(RandUint8(245, 256), RandUint8(200, 256), RandUint8(0, 100), 255)
}

func RandOrange() rl.Color {
	return rl.NewColor(255, RandUint8(70, 170), RandUint8(0, 50), 255)
}

func RandBrown() rl.Color {
	return rl.NewColor(RandUint8(100, 150), RandUint8(50, 120), RandUint8(30, 90), 255)
}

func RandGrey() rl.Color {
	return rl.NewColor(RandUint8(170, 220), RandUint8(170, 220), RandUint8(170, 220), 255)
}

func RandDarkGrey() rl.Color {
	return rl.NewColor(RandUint8(90, 120), RandUint8(90, 120), RandUint8(90, 120), 255)
}

func BrightPink() rl.Color {
	return rl.NewColor(255, 0, 130, 255)
}

func Red() rl.Color {
	return rl.NewColor(237, 29, 36, 255)
}

func BrightYellow() rl.Color {
	return rl.NewColor(255, 240, 31, 255)
}

func DarkRed() rl.Color {
	return rl.NewColor(139, 0, 0, 255)
}

func DarkRed2() rl.Color {
	return rl.NewColor(105, 0, 0, 255)
}
