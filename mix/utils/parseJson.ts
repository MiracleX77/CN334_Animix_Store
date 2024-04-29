export function parseJson(jsonString: string) {
    let json = {};
    try {
        json = JSON.parse(jsonString);
    } catch (ex) {
        json = { statusCode: 999, result: null, message: 'ERROR_PARSE_JSON' };
    }
    return json;
}