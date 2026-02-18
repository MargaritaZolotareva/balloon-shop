import axios from 'axios'
import router from '@/router'

const api = axios.create({
    baseURL: process.env.VUE_APP_API_URL,
    timeout: 10000,
})

api.interceptors.response.use(
    response => response,
    error => {
        if (!error.response) {
            router.push({ name: 'ServerError' })
            return Promise.reject(error)
        }

        const status = error.response.status

        if (status === 404) {
            router.push({ name: 'NotFound' })
        }

        if (status >= 500) {
            router.push({ name: 'ServerError' })
        }

        return Promise.reject(error)
    }
)

export default api
