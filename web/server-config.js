const isProd = process.env.NODE_ENV === 'production'
// const localhost = 'http://127.0.0.1:9090/'
const baseUrl = process.env.VUE_APP_API_URL
const api = baseUrl
export default {
    isProd,
    api
}
