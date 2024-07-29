package main

import "fmt"

// 抽象层
type ComputerFactory interface {
	ProduceCPU() CPU
	ProduceMemory() Memory
	ProduceCard() Card
}

type CPU interface {
	Computer()
}

type Memory interface {
	Store()
}

type Card interface {
	Display()
}

// 实现层（面向抽象编程）
type KinstonCPU struct {
}

func (kc KinstonCPU) Computer() {
	fmt.Println("使用了金士顿的CPU, 开始计算...")
}

type KinstonMemory struct {
}

func (kc KinstonMemory) Store() {
	fmt.Println("使用了金士顿的内存, 开始存储数据...")
}

type KinstonCard struct {
}

func (kc KinstonCard) Display() {
	fmt.Println("使用了金士顿的显卡, 开始显示...")
}

type NvdiaCPU struct {
}

func (n NvdiaCPU) Computer() {
	fmt.Println("使用了英伟达的CPU, 开始计算...")
}

type NvdiaMemory struct {
}

func (n NvdiaMemory) Store() {
	fmt.Println("使用了英伟达的内存, 开始存储数据...")
}

type NvdiaCard struct {
}

func (n NvdiaCard) Display() {
	fmt.Println("使用了英伟达的显卡, 开始显示...")
}

type IntelCPU struct {
}

func (i IntelCPU) Computer() {
	fmt.Println("使用了英特尔的CPU, 开始计算...")
}

type IntelMemory struct {
}

func (i IntelMemory) Store() {
	fmt.Println("使用了英特尔的内存, 开始存储数据...")
}

type IntelCard struct {
}

func (i IntelCard) Display() {
	fmt.Println("使用了英特尔的显卡, 开始显示...")
}

type KinstonFactory struct {
}

func (kf KinstonFactory) ProduceCPU() CPU {
	var cpu CPU
	cpu = new(KinstonCPU)
	return cpu
}

func (kf KinstonFactory) ProduceMemory() Memory {
	var memory Memory
	memory = new(KinstonMemory)
	return memory
}

func (kf KinstonFactory) ProduceCard() Card {
	var card Card
	card = new(KinstonCard)
	return card
}

type NvdiaFactory struct {
}

func (kf NvdiaFactory) ProduceCPU() CPU {
	var cpu CPU
	cpu = new(NvdiaCPU)
	return cpu
}

func (kf NvdiaFactory) ProduceMemory() Memory {
	var memory Memory
	memory = new(NvdiaMemory)
	return memory
}

func (kf NvdiaFactory) ProduceCard() Card {
	var card Card
	card = new(NvdiaCard)
	return card
}

type IntelFactory struct {
}

func (kf IntelFactory) ProduceCPU() CPU {
	var cpu CPU
	cpu = new(IntelCPU)
	return cpu
}

func (kf IntelFactory) ProduceMemory() Memory {
	var memory Memory
	memory = new(IntelMemory)
	return memory
}

func (kf IntelFactory) ProduceCard() Card {
	var card Card
	card = new(IntelCard)
	return card
}

// 业务逻辑层
func main() {
	// 组装一台由金士顿的内存、英伟达的显卡和英特尔的 CPU 组成的电脑
	var factory ComputerFactory
	var cpu CPU
	var memory Memory
	var card Card
	factory = new(IntelFactory)
	cpu = factory.ProduceCPU()
	factory = new(KinstonFactory)
	memory = factory.ProduceMemory()
	factory = new(NvdiaFactory)
	card = factory.ProduceCard()
	cpu.Computer()
	memory.Store()
	card.Display()
}
