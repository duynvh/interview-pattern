function getEarliest(room) {
    room.sort((a, b) => a[1] - b[1])
    return room[0]
}

function minMeetingRooms(meetings) {
    if (!meetings) return 0
    if (meetings.length <= 1) return meetings.length;

    meetings.sort((a, b) => a[0] - b[0])
    let rooms = [meetings[0]]

    for (let i = 1; i < meetings.length; i++) {
        let earliestTime = getEarliest(rooms);
        let currentTime = meetings[i]

        if (earliestTime[1] <= currentTime[0]) {
            earliestTime[1] = currentTime[1]
        } else {
            rooms.push(currentTime)
        }
    }

    console.log(rooms.length);
}

minMeetingRooms()
minMeetingRooms([[1,4]])
minMeetingRooms([[1,4], [2,5], [7,9]])//2, Since [1,4] and [2,5] overlap, we need two rooms to hold these two meetings. [7,9] can occur in any of the two rooms later.
minMeetingRooms([[6,7], [2,4], [8,12]])//1, None of the meetings overlap, therefore we only need one room to hold all meetings.
minMeetingRooms([[1,4], [2,3], [3,6]])//2, Since [1,4] overlaps with the other two meetings [2,3] and [3,6], we need two rooms to hold all the meetings.
minMeetingRooms([[4,5], [2,3], [2,4], [3,5]])//2, We will need one room for [2,3] and [3,5], and another room for [2,4] and [4,5].