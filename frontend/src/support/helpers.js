export const helpers = {
    getEvent(eventName) {
        const obj = eventName.split(':')

        return {
            type: obj[0],
            channel: obj[1],
            name: obj[2]
        }
    },
    // formatAllTemps(arr) {
    //     let out = {
    //         labels: [],
    //         datasets: []
    //     }
    //
    //     arr.forEach(item => {
    //         out.labels.push(item.name)
    //
    //         let data = []
    //         item.data.forEach(d => {
    //             data.unshift(d.temp)
    //         })
    //
    //         out.datasets.push({
    //             label: item.name,
    //             backgroundColor: '#f87979',
    //             data
    //         })
    //     })
    //
    //     return out
    // }
}