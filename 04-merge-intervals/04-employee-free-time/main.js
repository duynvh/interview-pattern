function findEmployeeFreeTime (schedules) {
    let freeTimes = [];
    let allTimes = [];

    for (let i = 0; i < schedules.length; i++) {
        for (let j = 0; j < schedules[i].length; j++) {
            allTimes.push(schedules[i][j]);
        }
    }

    allTimes.sort((a, b) => a[0] - b[0]);

    //Merge overlap
    for (let i = 1; i < allTimes.length; i++) {
        let current = allTimes[i];
        let previous = allTimes[i - 1];

        if (current[0] <= previous[1]) {
            allTimes[i] = [previous[0], current[1]]
            allTimes.splice(i - 1, 1)
            i--
        }
    }

    for (let i = 1; i < allTimes.length; i++) {
        freeTimes.push([allTimes[i - 1][1], allTimes[i][0]]);
    }

    console.log(freeTimes);
}

findEmployeeFreeTime ([[[1,3], [5,6]], [[2,3], [6,8]]])//[3,5], Both the employees are free between [3,5].
findEmployeeFreeTime ([[[1,3], [9,12]], [[2,4]], [[6,8]]])//[4,6], [8,9], All employees are free between [4,6] and [8,9].
findEmployeeFreeTime ([[[1,3]], [[2,4]], [[3,5], [7,9]]])//[5,7], ll employees are free between [5,7].