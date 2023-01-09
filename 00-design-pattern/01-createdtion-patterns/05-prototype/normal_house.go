package main

import "fmt"

type NormalHouse struct {
	LivingRoom int
	BedRoom    int
}

func (h *NormalHouse) Print() {
	fmt.Printf("%d living room - %d bed room\n", h.LivingRoom, h.BedRoom)
}

func (h *NormalHouse) SetLivingRoom(room int) {
	h.LivingRoom = room
}

func (h *NormalHouse) SetBedRoom(room int) {
	h.BedRoom = room
}

func (h *NormalHouse) Clone() IHouse {
    return &NormalHouse{
        LivingRoom: h.LivingRoom,
        BedRoom: h.BedRoom,
    }
}