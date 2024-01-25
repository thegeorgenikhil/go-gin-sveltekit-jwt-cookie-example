import { redirect } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"

export const load: PageServerLoad = async ({ locals, fetch }) => {
    if (!locals.token) {
        throw redirect(302, '/')
    }

    const token = locals.token

    const res = await fetch("http://localhost:8080/user/get-my-id", {
        headers: {
            authorization: `Bearer ${token}`
        }
    })

    const data = await res.json()

    if (res.status === 401) {
        redirect(302, "/auth/login")
    }

    return {
        id: data.id
    }
}