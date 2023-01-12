package main

type IHouse interface {
    Print()
    SetLivingRoom(int)
    SetBedRoom(int)
    Clone() IHouse
}