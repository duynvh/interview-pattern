package main

func main() {
    normalHouse := NormalHouse{
        LivingRoom: 1,
        BedRoom: 1,
    }
    normalHouse.Print()

    normalHouseClone := normalHouse.Clone()
    normalHouseClone.Print()

    normalHouseClone.SetBedRoom(2)
    normalHouseClone.SetLivingRoom(3)
    normalHouseClone.Print()

    normalHouse.Print()
}