
export default async function apiFetch(path, callback = null, options = {}) {
    const apiURL = `http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api`;
    const response = await fetch(apiURL+path, {...{credentials: 'include'}, ...options});
    if (!response.ok) {
        throw new Error('Failed to fetch data form path: '+path+' with error: '+response.Error);
    }
    const jsonData = await response.json();
    if (jsonData.status !== 'OK') {
        throw new Error(`Error response form path: ${path} resonse:\n${JSON.stringify(jsonData, null, 2)}`);
    }
    if (callback) {
        callback(jsonData);
    }
    return jsonData;
}