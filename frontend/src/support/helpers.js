export const helpers = {
    getEvent(eventName) {
        const obj = eventName.split(':')

        return {
            type: obj[0],
            channel: obj[1],
            name: obj[2]
        }
    },
}