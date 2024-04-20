import Cookies from 'js-cookie';

//Works only for client side
export default async function getUserUUID() {
    const uuid = Cookies.get('session_token')?.split('.')[1]
    return uuid
}

//Non async
export function getUUID() {
    const uuid = Cookies.get('session_token')?.split('.')[1]
    return uuid  
}