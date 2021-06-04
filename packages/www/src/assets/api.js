import axios from 'axios'

export const api = axios.create()

api.interceptors.response.use(undefined, async (r) => {
    if ([401, 403].includes(r.statusCode)) {
        location.href = '/'
    }

    throw r
})

export async function initAPI() {
    const u = new URL(location.href)
    const secret = u.searchParams.get('secret')
    if (secret) {
        api.defaults.headers = api.defaults.headers || {}
        api.defaults.headers['X-Secret'] = secret
    }

    await api.post('/server/login')
    return api
}
