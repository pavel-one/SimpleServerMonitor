export const appConfig = {
    baseAddress: window.location.hostname,
    wsAddress: getWsAddress(),
    baseURL: window.location.protocol + '//' + window.location.hostname + ':' + window.location.port + '/'
}

function getWsAddress() {
    let prefix = "wss://"

    if (window.location.protocol === 'http:') {
        prefix = "ws://"
    }

    return prefix + window.location.hostname + ':5000'
}