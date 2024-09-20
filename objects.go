package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BackgroundPicture struct {
	vector2                        rl.Vector2
	color                          rl.Color
	speed, dirX, dirY, siz, fd, ro float32
}

type WeaponCrate struct {
	tile      Tile
	weaponNum int
	isOn      bool
}

type Base struct {
	imageRectangle, rectangle, collisionRectangle, caRectangle []rl.Rectangle
	name                                                       []string
	isUnlocked                                                 []bool
	price, level                                               []int
	description                                                []string
	fade                                                       []float32
	pigeon                                                     Tile
	pigeonT                                                    int32
	pigeonoff                                                  bool
	mouse                                                      Enemy
}

type BackgroundInformation struct {
	imageRectangle, rectangle rl.Rectangle
	fade, rotation            float32
	isSmall                   bool
}

type Enemy struct {
	cnt                                                                      rl.Vector2
	name, name2                                                              string
	hp, hpMax, numMax, xp, bleed, count                                      int
	off, lr, idl, idlon, stat, nodmg, fast, fly, follow, xl                  bool
	rotation, rotation2, speed, dirX, dirY, hpY, hpY2, stunF, stunSiz, stunY float32
	image, rectangle, hitImgR, hitImgL, collisionRectangle, carec                                  rl.Rectangle

	idleT, moveChangeT, animFrameT, hpT, atkT, burnT, bleedT, poisonT, stunT, atkT2, freezeT, oilT int32
	//BOSS
	walkLeftAnimation, walkRightAnimation, attackAnimation, idleAnimation Animation
	state                                                                 int
	idleTime, walkTime                                                    int32
	ud, onoff                                                             bool
}

type Item struct {
	rectangle, image         rl.Rectangle
	center                   rl.Vector2
	dropTimer, cooldownTimer int32
	potionPrice, scrollprice []int

	numof, mana, price, listNum int

	isOff, art, notquick, isInvisible, isUnlocked, isNoCrate, isNoChestMove, unique bool

	name string

	des, des2, des3, des4, des5, des6, des7 string
}

type Chest struct {
	rectangle, collisionRectangle, image rl.Rectangle
	slots                                int
	center                               rl.Vector2
	isOpen, audioPlayed                  bool
	item                                 []Item
}

type Weapon struct {
	center                              rl.Vector2
	crec, rec, image, carec, rec2, rec3 rl.Rectangle
	col                                 rl.Color
	T, moveT, collisionTimer, dropTimer            int32

	isOff, lr, direction, isOn, isBelow, isLightning, isFirework, isTurret, isBomb, isPotion, isXp, isRingOfFire, isSludgeGeyser bool

	name                                                                                           string
	description, description2, description3, description4, description5, descrition6, description7 string

	speed, dirX, dirY, rotation, fade, rotationSpeed, WOrig, angle, angle2 float32

	damage, originalDamage, num, bounce, weapListNum, level, special, special2 int
}

type Player struct {
	center                                                                               rl.Vector2
	rectangle, collisionRectangle, caRectangle, perceptionRectangle, bossEffectRectangle rl.Rectangle
	anim                                                                                 []Animation
	weapon, weapon2                                                                      Weapon
	inventory, quik, art                                                                 []Item
	direction, push, gameover                                                            bool
	speed, speed2, speedMax                                                              float32
	rotation                                                                             float32
	angle, angle2, angle3, angle4                                                        float32
	v1, v2, v3, v4                                                                       rl.Vector2

	hp, hpMax, animationNum, xp, mana, manaMax, weaponMax int
	weapons                                               []Weapon

	hpTimer, dampTimer, burnTimer, poisonTimer, attackTimer, pushTimer, resistFireTimer, resistPoisonTimer, invisibleTimer, armorTimer, enemyCollisionTimer, freezeTimer, collectTimer int32
}

type Message struct {
	txt string
	col rl.Color
}

type Effect struct {
	name                                          string
	imageRectangle, rectangle, collisionRectangle rl.Rectangle
	color                                         rl.Color
	isOff, isBelow                                bool
	timer, clearTimer                             int32
	center                                        rl.Vector2
	fade, dirX, dirY, speed, rotation             float32
	circles                                       []Circle
	rectangles                                    []Rectangle
	rectangles2                                   []rl.Rectangle
	collisionRectangles                           []rl.Rectangle
	v1                                            []rl.Vector2
	v2                                            []rl.Vector2
	isOn                                          []bool
}

type Circle struct {
	center            rl.Vector2
	color             rl.Color
	fade, radius      float32
	dirX, dirY, speed float32
}

type Animation struct {
	frames              float32
	height, width, x, y float32
	rectangle           rl.Rectangle
}

type x1scr struct {
	walls, floors, otherTiles                              []Tile
	rectangles, doors                                      []rl.Rectangle
	doorNums                                               []int
	inf                                                    []xinf
	dbg                                                    []xdbg
	chests                                                 []Chest
	items                                                  []Item
	enemies                                                []Enemy
	backgroundInformation, backgroundInformation2, flowers []BackgroundInformation
	weaponCrate                                            WeaponCrate
	weapons                                                []Weapon
	enemyNum                                               int
}

type xdbg struct {
	conecRoom, side int
}

type xinf struct {
	num, numW, numH       int
	width, height         float32
	wallsAdded, isVisible bool
	center                rl.Vector2
}

type Tile struct {
	image, image2, rectangle, rectangle2, cRectangle rl.Rectangle
	color, color2                                    rl.Color
	dirX, dirY, speed, fade, rotation, angle         float32
	isOff, spikes, onoff, solid, firework            bool
	name                                             string
	spikeRectangles                                  []rl.Rectangle
	hp, numType, room                                int
	hpTimer, moveTimer, timer                        int32
	center, vector                                   rl.Vector2
}

type Rectangle struct {
	rectangle         rl.Rectangle
	color             rl.Color
	fade              float32
	dirX, dirY, speed float32
}

type Stats struct {
	strength, luck, dexterity, intelligence, perception      int
	strength2, luck2, dexterity2, intelligence2, perception2 int
	critical, critical2                                      int
}
