interface RawBeamMessage {
    Payload: string,
    Group: string,
}

interface BeamMessage {
    timestamp: string,
    payload: string,
    group: string,
    colorGroup?: string,
}
