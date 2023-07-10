export const appConfig = {
    baseAddress: window.location.hostname,
    wsAddress: getWsAddress(),
    baseURL: window.location.protocol + '//' + window.location.hostname + ':' + window.location.port + '/'
}

function getWsAddress() {
    //TODO: move to env
    if (window.location.protocol === 'https:') {
        return "wss://ws.pavel.one"
    }

    return "ws://" + window.location.hostname + ':5000'
}