

export function parseToken(token) {
    let payload = token.split('.')[1];
    let userInfo = JSON.parse(decodeURIComponent(escape(window.atob(payload.replace(/-/g, '+').replace(/_/g, '/')))));
    return userInfo;
}