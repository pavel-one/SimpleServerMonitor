import {appConfig as config} from "../config/config.js";

export default {
    Connection: new WebSocket(config.wsAddress)
}